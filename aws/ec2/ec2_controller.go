package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func StartInstance(instanceID string) (bool, error) {
	svc := ec2.New(session.New(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
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

func StopInstance(instanceID string) (bool, error) {
	svc := ec2.New(session.New(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
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
