package rds

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"

	"github.com/sakuma/aws-sitter/lib/util"
)

func awsSession(regionName string) *rds.RDS {
	session := rds.New(session.New(&aws.Config{
		Region: aws.String(regionName),
	}))
	return session
}

func getInstances(region string) []*rds.DBInstance {
	svc := awsSession(region)
	filter := &rds.DescribeDBInstancesInput{}
	res, err := svc.DescribeDBInstances(filter)
	if err != nil {
		panic(err)
	}
	// fmt.Println(res.DBInstances)
	return res.DBInstances
}

func Execute(region string) error {
	res := getInstances(region)
	svc := awsSession(region)
	for _, i := range res {
		util.DebugPrint("instance ----------")
		// fmt.Printf("%+v\n", i)
		instance := util.Instance{}
		instance.ResourceType = "rds"
		instance.Region = *i.AvailabilityZone
		instance.Name = *i.DBInstanceIdentifier
		instance.ID = *i.DBInstanceIdentifier // same
		instance.InstanceType = *i.DBInstanceClass
		instance.State = *i.DBInstanceStatus
		// Get Tags
		input := &rds.ListTagsForResourceInput{
			ResourceName: aws.String(*i.DBInstanceArn),
		}
		result, _ := svc.ListTagsForResource(input)
		for _, t := range result.TagList {
			switch *t.Key {
			case "API_AUTO_OPERATION_MODE":
				instance.OperationMode = strings.TrimSpace(*t.Value)
			case "API_RUN_SCHEDULE":
				instance.RunSchedule = *t.Value
			}
			if instance.OperationMode == "stop" {
				instance.StopOnly = true
			}
		}
		fmt.Printf("%+v\n", instance)

		if util.IsActive(instance) {
			if instance.IsRunning() {
				fmt.Println("Already Started : ", instance.ID)
			}
			if instance.IsStopped() {
				_, err := startInstance(region, instance.ID)
				if err == nil {
					fmt.Println("Start instance: ", instance.ID)
				} else {
					fmt.Println("Error: ", instance.ID, ": ", err)
				}
			} else {
				fmt.Println("Can not be start (instance processing): ", instance.ID)
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
	return nil
}

func startInstance(region, instanceID string) (bool, error) {
	svc := awsSession(region)
	input := &rds.StartDBInstanceInput{
		DBInstanceIdentifier: aws.String(instanceID),
	}
	_, err := svc.StartDBInstance(input)
	if err != nil {
		return false, err
	}
	return true, nil
}

func stopInstance(region, instanceID string) (bool, error) {
	svc := awsSession(region)
	input := &rds.StopDBInstanceInput{
		DBInstanceIdentifier: aws.String(instanceID),
	}
	_, err := svc.StopDBInstance(input)
	if err != nil {
		return false, err
	}
	return true, nil
}
