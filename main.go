package main

import (
	// "flag"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/sakuma/aws-sitter/lib/util"
	"github.com/sakuma/aws-sitter/sitter"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler() (string, error) {
	// NOTE: making array if other region
	region := "ap-northeast-1"
	// flag.StringVar(&region, "region", "ap-northeast-1", "Execution AWS region.")
	// flag.BoolVar(&util.Verbose, "v", false, "display verbose log")
	// flag.Parse()

	util.DebugPrint("start...")

	ec2 := sitter.EC2{Region: region}
	ec2.Execute()

	rds := sitter.RDS{Region: region}
	rds.Execute()

	// TODO: error handling
	// sitter.Execute(region)
	return "succeded process", nil
}

func main() {
	lambda.Start(Handler)
}
