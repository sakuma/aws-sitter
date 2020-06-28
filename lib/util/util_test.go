package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetCurrentTime(t *testing.T) {
	var region, zone string
	var current time.Time

	region = "ap-northeast-1"
	current = SetCurrentTime(region)
	zone, _ = current.Zone()
	assert.Equal(t, "JST", zone)
}

func TestMakeStrings(t *testing.T) {
	var testCases = []struct {
		min      int
		max      int
		expected string
	}{
		{1, 5, "12345"},
		{3, 4, "34"},
		{0, 3, "0123"},
		{0, 0, "0"},
	}
	for _, tt := range testCases {
		assert.Equal(t, tt.expected, MakeWeekStrings(tt.min, tt.max))
	}
}
