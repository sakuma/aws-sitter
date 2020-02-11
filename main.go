package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	ec2controller "github.com/sakuma/aws-sitter/aws/ec2"
	rdscontroller "github.com/sakuma/aws-sitter/aws/rds"
	"github.com/sakuma/aws-sitter/lib/holiday"
	"gopkg.in/yaml.v2"
)

type Instance struct {
	Region   string `yaml:region`
	Type     string `yaml:type`
	Name     string `yaml:name`
	ID       string `yaml:id`
	StopOnly bool   `yaml:stop_only`
}

func readInstanceFile(path string) ([]Instance, error) {
	// TODO: error handling
	buf, _ := ioutil.ReadFile("./instances.yml")
	data := make([]Instance, 50)

	// TODO: error handling
	err := yaml.Unmarshal(buf, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func isActive(s Instance) bool {
	// TODO: force runnning
	if holiday.IsHoliday(t) {
		return false
	}
	if holiday.IsRunnable() {
		return true
	}
	return false
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler() (string, error) {

	// TODO: error handling
	instances, _ := readInstanceFile("./instances.yml")
	for _, s := range instances {
		switch s.Type {
		case "ec2":
			if isActive(s) {
				ec2controller.StartInstance(s.ID)
				fmt.Println("succeeded: start instance.")
			} else {
				ec2controller.StopInstance(s.ID)
				fmt.Println("succeeded: stop instance.")
			}
		case "rds":
			if isActive(s) {
				rdscontroller.StartInstance(s.ID)
				fmt.Println("succeeded: start instance.")
			} else {
				rdscontroller.StopInstance(s.ID)
				fmt.Println("succeeded: stop instance.")
			}
		}
	}
	return "succeded process", nil
}

func main() {
	lambda.Start(Handler)
}
