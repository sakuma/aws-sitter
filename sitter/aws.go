package sitter

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/rds"
)

type ServiceType string

const (
	EC2 = "ec2"
	RDS = "rds"
)

// Operator is interface for Operation control
type Operator interface {
	InstanceName() string
	Type() string
	AwsSession() interface{}
	// GetInstances() interface{}
	// StartInstance()
	// StopInstance()
}

type Instance struct {
	InstanceType  string
	Region        string
	Name          string
	ResourceType  string
	ID            string
	Controllable  bool
	StopOnly      bool
	State         string
	OperationMode string
	RunSchedule   string
}

func New(instanceType, regionName string) Instance {
	p := &instance{instanceType, regionName}
	if instanceType == RDS {
		return &rdsInstance{p}
	} else {
		return &ec2Instance{p}
	}
}

type rdsInstance struct {
	*instance
}

// func (r *rdsInstance) AwsSession() *rds.RDS {
func (r *rdsInstance) AwsSession() interface{} {
	session := rds.New(session.New(&aws.Config{
		Region: aws.String(r.Region),
	}))
	return session
}

func (r *rdsInstance) GetInstances() []*rds.DBInstance {
	svc := r.AwsSession()
	filter := &rds.DescribeDBInstancesInput{}
	res, err := svc.DescribeDBInstances(filter)
	if err != nil {
		panic(err)
	}
	// fmt.Println(res.DBInstances)
	return res.DBInstances
}

func (r *rdsInstance) InstanceName() string {
	return r.Name
}

func (r *rdsInstance) Type() string {
	return "RDS: "
}

type ec2Instance struct {
	*instance
}

// func (r *ec2Instance) AwsSession() *ec2.EC2 {
func (r *ec2Instance) AwsSession() interface{} {
	session := ec2.New(session.New(&aws.Config{
		Region: aws.String(r.Region),
	}))
	return session
}

func (e *ec2Instance) InstanceName() string {
	return e.Name
}

func (e *ec2Instance) Type() string {
	return "EC2"
}

func (e *ec2Instance) GetInstances() []*ec2.Reservation {
	svc := e.AwsSession()
	filter := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag-key"),
				Values: []*string{
					aws.String("API_CONTROLLABLE"),
				},
			},
		},
	}

	res, err := svc.DescribeInstances(filter)
	if err != nil {
		panic(err)
	}
	return res.Reservations
}

func printInstanceName(i Instance) {
	fmt.Println(i.Type(), i.InstanceName())
}

func Execute(regionName string) {
	// EC2
	i := New(RDS, regionName)
	sitterSession := i.AwsSession()
	res := sitterSession.DescribeDBInstances()

	// res := i.GetInstances()
	for _, i := range res {
		fmt.Println("Name: ", *i.DBInstanceIdentifier)
	}
	// i := instances{RDS, regionName}
	// session := GetSession(EC2, regionName)
	// session
}

// import (
// 	"fmt"

// 	"github.com/aws/aws-sdk-go/service/rds"

// 	// ec2Ctrl "github.com/sakuma/aws-sitter/aws/ec2"
// 	// rdsCtrl "github.com/sakuma/aws-sitter/aws/rds"
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// )

// type InstanceType string

// const (
// 	EC2 = "ec2"
// 	RDS = "rds"
// )

// type Instance interface {
// 	AwsSession()
// 	GetInstances()
// 	// StartInstance()
// 	// StopInstance()
// }

// type Session interface{}

// type instance struct {
// 	InstanceType  string
// 	Region        string
// 	Name          string
// 	ResourceType  string
// 	ID            string
// 	Controllable  bool
// 	StopOnly      bool
// 	State         string
// 	OperationMode string
// 	RunSchedule   string
// }

// type RDSInstance struct {
// 	*instance
// }

// type EC2Instance struct {
// 	*instance
// }

// func (r *RDSInstance) AwsSession() *rds.RDS {
// 	session := rds.New(session.New(&aws.Config{
// 		Region: aws.String(r.Region),
// 	}))
// 	return session
// }

// func (r *RDSInstance) GetInstances() []*rds.DBInstance {
// 	svc := r.AwsSession()
// 	filter := &rds.DescribeDBInstancesInput{}
// 	res, err := svc.DescribeDBInstances(filter)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Println(res.DBInstances)
// 	return res.DBInstances
// }
// func New(instanceType, reagionName string) Instance {
// 	switch instanceType {
// 	case "rds":
// 		r := &instance{}
// 		r.InstanceType = instanceType
// 		r.Region = reagionName
// 		return &RDSInstance{r}
// 	}
// 	// i := &instance{instanceType, reagionName}
// }

// func Execute(regionName string) {
// 	// EC2
// 	i := New(RDS, regionName)
// 	res := i.GetInstances()
// 	for _, i := range res {
// 		fmt.Println("Name: ", *i.DBInstanceIdentifier)
// 	}
// 	// i := instances{RDS, regionName}
// 	// session := GetSession(EC2, regionName)
// 	// session
// }

// // func GetSession(instanceType, regionName string) {
// // 	switch instanceType {
// // 	case EC2:
// // 		return ec2Ctrl.AwsSession(regionName)
// // 	case RDS:
// // 		return ec2Ctrl.AwsSession(regionName)
// // 	}
// // 	//
// // }

// func Swither() bool {
// 	return true
// }
