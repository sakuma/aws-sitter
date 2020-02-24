package ec2

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/sakuma/aws-sitter/lib/util"
)

func getInstances(region string) []*ec2.Reservation {
	svc := util.AwsSession(region)
	filter := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:API_CONTROLL_ON_OR_OFF"),
				Values: []*string{
					aws.String("ON"),
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
			instance := util.Instance{}
			fmt.Println("instance ----------")
			instance.Region = *i.Placement.AvailabilityZone
			instance.ID = *i.InstanceId
			instance.InstanceType = *i.InstanceType
			instance.State = *i.State.Name
			instance.Name = *i.KeyName

			for _, t := range i.Tags {
				switch *t.Key {
				case "API_AUTO_OPERATION_MODE":
					instance.OperationMode = strings.TrimSpace(*t.Value)
				case "API_RUN_SCHEDULE":
					instance.RunSchedule = *t.Value
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
	svc := util.AwsSession(region)
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
	svc := util.AwsSession(region)
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
