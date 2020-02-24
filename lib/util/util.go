package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	// ec2ctrl "github.com/sakuma/aws-sitter/aws/ec2"
	"github.com/sakuma/aws-sitter/lib/holiday"
)

var Verbose bool

type Instance struct {
	Region        string
	InstanceType  string
	Name          string
	ID            string
	StopOnly      bool
	State         string
	OperationMode string
	RunSchedule   string
}

func (i *Instance) isForceRunnable() bool {
	return i.OperationMode == "start"
}

func (i *Instance) IsRunning() bool {
	switch i.State {
	case "running":
		return true
	case "stopping", "stopped", "shutting", "terminated":
		return false
	case "pending":
		// unknown: maybe false
		return false
	default:
		return false
	}
}

func (i *Instance) isRunnable(t time.Time) bool {
	if i.StopOnly {
		return false
	}
	// TODO: invalid format check
	times := strings.Split(i.RunSchedule, "-")
	from, _ := strconv.Atoi(times[0])
	to, _ := strconv.Atoi(times[1])
	if from <= t.Hour() && t.Hour() <= to {
		return true
	}
	return false
}

func IsActive(i Instance) bool {
	// TODO: force runnning
	t := time.Now()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	current := t.In(jst)

	if i.isForceRunnable() {
		return false
	}

	if holiday.IsHoliday(current) {
		return false
	}

	if i.isRunnable(current) {
		return true
	}
	return false
}

func AwsSession(regionName string) *ec2.EC2 {
	session := ec2.New(session.New(&aws.Config{
		Region: aws.String(regionName),
	}))
	return session
}

func DebugPrint(a ...interface{}) {
	if Verbose {
		fmt.Println(a...)
	}
}
