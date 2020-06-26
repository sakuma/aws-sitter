package util

import (
	"fmt"
	"time"
	// "github.com/sakuma/aws-sitter/lib/holiday"
)

var Verbose bool

func SetCurrentTime(regionName string) time.Time {
	t := time.Now()
	switch regionName {
	case "ap-northeast-1":
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		return t.In(jst)
	}
	return t
}

func DebugPrint(a ...interface{}) {
	if Verbose {
		fmt.Println(a...)
	}
}
