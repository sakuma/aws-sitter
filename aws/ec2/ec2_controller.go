package ec2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/sakuma/aws-sitter/lib/util"
)

func awsSession(regionName string) *ec2.EC2 {
	session := ec2.New(session.New(&aws.Config{
		Region: aws.String(regionName),
	}))
	return session
}

func getInstances(region string) []*ec2.Reservation {
	svc := awsSession(region)
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
func Execute(region string) error {
	res := getInstances(region)
	for _, r := range res {
		for _, i := range r.Instances {
			util.DebugPrint("instance ----------")
			instance := util.Instance{}
			instance.ResourceType = "ec2"
			instance.Region = *i.Placement.AvailabilityZone
			instance.ID = *i.InstanceId
			instance.InstanceType = *i.InstanceType
			instance.State = *i.State.Name
			instance.Name = *i.KeyName

			for _, t := range i.Tags {
				v := strings.TrimSpace(*t.Value)
				switch *t.Key {
				case "API_CONTROLLABLE":
					b, _ := strconv.ParseBool(v)
					instance.Controllable = b
				case "API_AUTO_OPERATION_MODE":
					// TODO: validation: [start,stop,auto]
					instance.OperationMode = v
				case "API_RUN_SCHEDULE":
					instance.RunSchedule = v
				}
			}
			if instance.OperationMode == "stop" {
				instance.StopOnly = true
			}
			fmt.Printf("%+v\n", instance)
			if util.IsActive(instance) {
				if instance.IsRunning() {
					fmt.Println("Already Started : ", instance.ID)
				} else {
					_, err := startInstance(region, instance.ID)
					if err == nil {
						fmt.Println("Start instance: ", instance.ID)
					} else {
						fmt.Println("Error: ", instance.ID, ": ", err)
					}
				}
			} else {
				if instance.IsRunning() {
					_, err := stopInstance(region, instance.ID)
					if err != nil {
						fmt.Println("Error: ", instance.ID, ": ", err)
					}
					fmt.Println("Stop instance: ", instance.ID)
				} else {
					fmt.Println("Already stop instance: ", instance.ID)
				}
			}
		}
	}
	return nil
}

func startInstance(region, instanceID string) (bool, error) {
	svc := awsSession(region)
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
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

func stopInstance(region, instanceID string) (bool, error) {
	svc := awsSession(region)
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
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
