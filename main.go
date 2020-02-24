package main

import (
	"flag"

	"github.com/aws/aws-lambda-go/lambda"
	ec2ctrl "github.com/sakuma/aws-sitter/aws/ec2"
	"github.com/sakuma/aws-sitter/lib/util"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler() (string, error) {
	region := "ap-northeast-1"
	flag.StringVar(&region, "region", "ap-northeast-1", "Execution AWS region.")
	flag.BoolVar(&util.Verbose, "v", false, "display verbose log")
	flag.Parse()

	util.DebugPrint("start...")

	// TODO: error handling
	ec2ctrl.Execute(region)
	return "succeded process", nil
}

func main() {
	// ec2ctrl.Execute(region)
	lambda.Start(Handler)
}
