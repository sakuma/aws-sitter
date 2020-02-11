package rds

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

func StartInstance(instanceID string) (bool, error) {
	svc := rds.New(session.New(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	input := &rds.StartDBInstanceInput{
		DBInstanceIdentifier: aws.String(instanceID),
	}
	_, err := svc.StartDBInstance(input)
	if err != nil {
		return false, err
	}
	return true, nil
}

func StopInstance(instanceID string) (bool, error) {
	svc := rds.New(session.New(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	input := &rds.StopDBInstanceInput{
		DBInstanceIdentifier: aws.String(instanceID),
	}
	_, err := svc.StopDBInstance(input)
	if err != nil {
		return false, err
	}
	return true, nil
}
