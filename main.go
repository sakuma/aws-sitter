package main

import (
	"flag"

	"github.com/aws/aws-lambda-go/lambda"
	// sitter "github.com/sakuma/aws-sitter/aws"
	// ec2Ctrl "github.com/sakuma/aws-sitter/aws/ec2"
	rdsCtrl "github.com/sakuma/aws-sitter/sitter"
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
	// sitter.Execute(region)
	// ec2Ctrl.Execute(region)
	rdsCtrl.Execute(region)
	return "succeded process", nil
}

func main() {
	lambda.Start(Handler)
}
