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
	assert.Equal(t, zone, "Asia/Tokyo")
}
