package main

import (
	"flag"

	ec2ctrl "github.com/sakuma/aws-sitter/aws/ec2"
	"github.com/sakuma/aws-sitter/lib/util"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
// func Handler() (string, error) {

// 	// TODO: error handling
// 	instances, _ := readInstanceFile("./instances.yml")
// 	for _, s := range instances {
// 		switch s.Type {
// 		case "ec2":
// 			if isActive(s) {
// 				ec2controller.StartInstance(s.ID)
// 				fmt.Println("succeeded: start instance.")
// 			} else {
// 				ec2controller.StopInstance(s.ID)
// 				fmt.Println("succeeded: stop instance.")
// 			}
// 		case "rds":
// 			if isActive(s) {
// 				rdscontroller.StartInstance(s.ID)
// 				fmt.Println("succeeded: start instance.")
// 			} else {
// 				rdscontroller.StopInstance(s.ID)
// 				fmt.Println("succeeded: stop instance.")
// 			}
// 		}
// 	}
// 	return "succeded process", nil
// }

func main() {
	region := "ap-northeast-1"
	flag.StringVar(&region, "region", "ap-northeast-1", "Execution AWS region.")
	flag.BoolVar(&util.Verbose, "v", false, "display verbose log")
	flag.Parse()

	util.DebugPrint("start...")

	ec2ctrl.Execute(region)
	// lambda.Start(Handler)
}
