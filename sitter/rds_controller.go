package sitter

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"

	"github.com/sakuma/aws-sitter/lib/util"
)

type RDS struct {
	Region string
	Instance
}

func (r RDS) awsSession() *rds.RDS {
	session := rds.New(session.New(&aws.Config{
		Region: aws.String(r.Region),
	}))
	return session
}

func (r RDS) getInstances() []*rds.DBInstance {
	svc := r.awsSession()
	filter := &rds.DescribeDBInstancesInput{}
	res, err := svc.DescribeDBInstances(filter)
	if err != nil {
		panic(err)
	}
	// fmt.Println(res.DBInstances)
	return res.DBInstances
}

func (r RDS) Execute() error {
	res := r.getInstances()
	svc := r.awsSession()
	for _, i := range res {
		util.DebugPrint("instance ----------")
		// fmt.Printf("%+v\n", i)
		instance := Instance{}
		instance.ResourceType = "rds"
		instance.AvailabilityZone = *i.AvailabilityZone
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
			case "API_CONTROLLABLE":
				instance.setControllable(*t.Value)
			case "API_AUTO_OPERATION_MODE":
				instance.setOperationMode(*t.Value)
			case "API_RUN_SCHEDULE":
				instance.RunSchedule = *t.Value
			}
			if instance.OperationMode == "stop" {
				instance.StopOnly = true
			}
		}
		// fmt.Printf("%+v\n", instance)
		instance.CurrentTime = util.SetCurrentTime(r.Region)
		rds := RDS{Region: r.Region, Instance: instance}

		if !instance.Controllable {
			continue
		}

		mode := instance.executeMode()
		switch mode {
		case "start":
			_, err := rds.startInstance()
			if err == nil {
				fmt.Println("Start instance: ", instance.ID)
			} else {
				fmt.Println("Error: ", instance.ID, ": ", err)
			}
		case "stop":
			_, err := rds.stopInstance()
			if err == nil {
				fmt.Println("Stop instance: ", instance.ID)
			} else {
				fmt.Println("Error: ", instance.ID, ": ", err)
			}
		default:
			continue
		}
	}
	return nil
}

func (r RDS) startInstance() (bool, error) {
	svc := r.awsSession()
	input := &rds.StartDBInstanceInput{
		DBInstanceIdentifier: aws.String(r.Instance.ID),
	}
	_, err := svc.StartDBInstance(input)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r RDS) stopInstance() (bool, error) {
	svc := r.awsSession()
	input := &rds.StopDBInstanceInput{
		DBInstanceIdentifier: aws.String(r.Instance.ID),
	}
	_, err := svc.StopDBInstance(input)
	if err != nil {
		return false, err
	}
	return true, nil
}
