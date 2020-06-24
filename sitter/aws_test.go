package sitter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsWithinScheduleTime(t *testing.T) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	var tt time.Time

	i := Instance{RunSchedule: "10-19"}

	// out of range(7:00)
	tt = time.Date(2019, 8, 31, 7, 00, 0, 0, jst)
	assert.Equal(t, i.isWithinScheduleTime(tt), false)

	// in time (10:00)
	tt = time.Date(2019, 8, 31, 10, 00, 0, 0, jst)
	assert.Equal(t, i.isWithinScheduleTime(tt), true)

	// in time (10:01)
	tt = time.Date(2019, 8, 31, 10, 01, 0, 0, jst)
	assert.Equal(t, i.isWithinScheduleTime(tt), true)

	// in time (19:59)
	tt = time.Date(2019, 8, 31, 19, 59, 0, 0, jst)
	assert.Equal(t, i.isWithinScheduleTime(tt), true)

	// out of range (20:00)
	tt = time.Date(2019, 8, 31, 20, 00, 0, 0, jst)
	assert.Equal(t, i.isWithinScheduleTime(tt), false)

	// out of range (20:01)
	tt = time.Date(2019, 8, 31, 20, 01, 0, 0, jst)
	assert.Equal(t, i.isWithinScheduleTime(tt), false)
}

func TestIsRunning(t *testing.T) {
	i := Instance{}

	i.State = "running"
	assert.Equal(t, i.isRunning(), true)

	i.State = "available"
	assert.Equal(t, i.isRunning(), true)

	i.State = "stopping"
	assert.Equal(t, i.isRunning(), false)

	i.State = "stopped"
	assert.Equal(t, i.isRunning(), false)

	i.State = "shutting"
	assert.Equal(t, i.isRunning(), false)

	i.State = "terminated"
	assert.Equal(t, i.isRunning(), false)

	i.State = "hoge"
	assert.Equal(t, i.isRunning(), false)
}

func TestIsStopped(t *testing.T) {
	i := Instance{}

	i.State = "stopped"
	assert.Equal(t, i.isStopped(), true)

	i.State = "stopping"
	assert.Equal(t, i.isStopped(), true)

	i.State = "shutting"
	assert.Equal(t, i.isStopped(), true)

	i.State = "terminated"
	assert.Equal(t, i.isStopped(), true)

	i.State = "running"
	assert.Equal(t, i.isStopped(), false)
}
