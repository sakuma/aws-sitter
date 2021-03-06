package sitter

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/sakuma/aws-sitter/lib/util"
)

type EC2 struct {
	Region string
	Instance
}

func (e EC2) awsSession() *ec2.EC2 {
	session := ec2.New(session.New(&aws.Config{
		Region: aws.String(e.Region),
		// LogLevel: aws.LogLevel(aws.LogDebugWithRequestErrors | aws.LogDebugWithRequestRetries),
	}))
	return session
}

func (e EC2) getInstances() []*ec2.Reservation {
	svc := e.awsSession()
	filter := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag-key"),
				Values: []*string{
					aws.String("API_CONTROLLABLE"),
				},
			},
		},
	}

	res, err := svc.DescribeInstances(filter)
	if err != nil {
		panic(err)
	}
	return res.Reservations
}

//
func (e EC2) Execute() error {
	res := e.getInstances()
	for _, r := range res {
		for _, i := range r.Instances {
			util.DebugPrint("instance ----------")
			instance := Instance{}
			instance.ResourceType = "ec2"
			instance.AvailabilityZone = *i.Placement.AvailabilityZone
			instance.ID = *i.InstanceId
			instance.InstanceType = *i.InstanceType
			instance.State = *i.State.Name
			instance.Name = *i.KeyName

			for _, t := range i.Tags {
				switch *t.Key {
				case "API_CONTROLLABLE":
					instance.setControllable(*t.Value)
				case "API_AUTO_OPERATION_MODE":
					instance.setOperationMode(*t.Value)
				case "API_RUN_SCHEDULE":
					instance.setRunSchedule(*t.Value)
				}
			}
			if instance.OperationMode == "stop" {
				instance.StopOnly = true
			}
			// fmt.Printf("%+v\n", instance)
			instance.CurrentTime = util.SetCurrentTime(e.Region)
			ec2 := EC2{Region: e.Region, Instance: instance}

			if !instance.Controllable {
				continue
			}

			mode := instance.executeMode()
			switch mode {
			case "start":
				_, err := ec2.startInstance()
				if err == nil {
					fmt.Println("Start instance: ", instance.ID)
				} else {
					fmt.Println("Error: ", instance.ID, ": ", err)
				}
			case "stop":
				_, err := ec2.stopInstance()
				if err == nil {
					fmt.Println("Stop instance: ", instance.ID)
				} else {
					fmt.Println("Error: ", instance.ID, ": ", err)
				}
			default:
				continue
			}
		}
	}
	return nil
}

func (e EC2) startInstance() (bool, error) {
	svc := e.awsSession()
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(e.Instance.ID),
		},
	}
	_, err := svc.StartInstances(input)
	if err != nil {
		// TODO: error handling
		// aerr, ok := err.(awserr.Error); ok {
		return false, err
	}
	return true, err
}

func (e EC2) stopInstance() (bool, error) {
	svc := e.awsSession()
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(e.Instance.ID),
		},
	}
	_, err := svc.StopInstances(input)
	if err != nil {
		// TODO: error handling
		// aerr, ok := err.(awserr.Error); ok {
		return false, err
	}
	return true, err
}
