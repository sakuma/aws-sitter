package util

import (
	"fmt"
	"strconv"
	"strings"
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

func SpaceReplaceAll(v string) string {
	conveted := strings.ReplaceAll(v, "　", "")
	return strings.ReplaceAll(conveted, " ", "")
}

func HyphenReplaceAll(v string) string {
	r := strings.NewReplacer(
		// \u**** to \u002D
		"ー", "-", // \u30FC
		"−", "-", // \u2212
		"―", "-", // \u2015
		"－", "-", // \uFF0D
		"﹣", "-", // \uFE63
		"⼀", "-", // \u2F00
		"ー", "-", // \u30FC
		"㆒", "-", // \u3192
	)
	return r.Replace(v)
}

func MakeWeekStrings(min, max int) string {
    a := make([]string, max-min+1)
    for i := range a {
        a[i] = strconv.Itoa(min + i)
    }
    return strings.Join(a, "")
}

func DebugPrint(a ...interface{}) {
	if Verbose {
		fmt.Println(a...)
	}
}
