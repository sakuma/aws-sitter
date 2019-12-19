package main

import (
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sakuma/aws-sitter/lib/holiday"
	ec2controller "github.com/sakuma/aws-sitter/aws/ec2"
)

func isActive(t time.Time) bool {
	return holiday.IsHoliday()
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler() (string, error) {
	instanceID := os.Getenv("INSTANCE_ID")
	t := time.Now()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	current := t.In(jst)
	if isActive(current) {
		ec2controller.StartInstance(instanceID)
		return "succedked: start instance.", nil
	}
	ec2controller.StopInstance(instanceID)
	return "succeeded: stop instance.", nil
}

func main() {
	lambda.Start(Handler)
}
