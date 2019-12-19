package main

import (
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func isActive(t time.Time) bool {
	return false
}

func startInstance(instanceID string) (bool, error) {
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

func stopInstance(instanceID string) (bool, error) {
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


// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler() (string, error) {
	instanceID := os.Getenv("INSTANCE_ID")
	t := time.Now()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	current := t.In(jst)
	if isActive(current) {
		startInstance(instanceID)
		return "succeded: start instance.", nil
	}
	stopInstance(instanceID)
	return "succeeded: stop instance.", nil
}

func main() {
	lambda.Start(Handler)
}
