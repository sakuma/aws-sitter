package util

import (
	"fmt"
	// "github.com/sakuma/aws-sitter/lib/holiday"
)

var Verbose bool

func DebugPrint(a ...interface{}) {
	if Verbose {
		fmt.Println(a...)
	}
}
