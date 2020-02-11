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
	Region string `toml:region`
	Type   string `toml:type`
	Name   string `toml:name`
	ID     string `toml:id`
}

func ReadInstanceFile(path string) ([]Instance, error) {
	// TODO: error handling
	buf, _ := ioutil.ReadFile("./instances.yml")
	data := make([]Instance, 20)

	// TODO: error handling
	err := yaml.Unmarshal(buf, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func isRunnable(t time.Time) bool {
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
	fmt.Println("call Handler")

	t := time.Now()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	current := t.In(jst)

	// TODO: error handling
	instances, _ := ReadInstanceFile("./instances.yml")
	for _, s := range instances {
		fmt.Println("Region: ", s.Region)
		fmt.Println("Type: ", s.Type)
		fmt.Println("Name: ", s.Name)
		fmt.Println("ID: ", s.ID)

		switch s.Type {
		case "ec2":
			if isRunnable(current) {
				ec2controller.StartInstance(s.ID)
				fmt.Println("succeeded: start instance.")
			} else {
				ec2controller.StopInstance(s.ID)
				fmt.Println("succeeded: stop instance.")
			}
		case "rds":
			if isRunnable(current) {
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
