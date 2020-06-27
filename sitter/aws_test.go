package sitter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsWithinScheduleTime(t *testing.T) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	i := Instance{RunSchedule: "10-19"}

	// out of range(7:00)
	i.CurrentTime = time.Date(2019, 8, 31, 7, 00, 0, 0, jst)
	assert.Equal(t, i.isWithinScheduleTime(), false)

	// in time (10:00)
	i.CurrentTime = time.Date(2019, 8, 31, 10, 00, 0, 0, jst)
	assert.Equal(t, i.isWithinScheduleTime(), true)

	// in time (10:01)
	i.CurrentTime = time.Date(2019, 8, 31, 10, 01, 0, 0, jst)
	assert.Equal(t, i.isWithinScheduleTime(), true)

	// in time (19:59)
	i.CurrentTime = time.Date(2019, 8, 31, 19, 59, 0, 0, jst)
	assert.Equal(t, i.isWithinScheduleTime(), true)

	// out of range (20:00)
	i.CurrentTime = time.Date(2019, 8, 31, 20, 00, 0, 0, jst)
	assert.Equal(t, i.isWithinScheduleTime(), false)

	// out of range (20:01)
	i.CurrentTime = time.Date(2019, 8, 31, 20, 01, 0, 0, jst)
	assert.Equal(t, i.isWithinScheduleTime(), false)
}

func TestExecuteMode(t *testing.T) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	i := Instance{RunSchedule: "10-19"}

	////////////////
	i.OperationMode = "start"

	/////////
	i.State = "stopped"

	// out of range time
	i.CurrentTime = time.Date(2019, 8, 31, 7, 00, 0, 0, jst)
	assert.Equal(t, "none", i.executeMode())
	// in time
	i.CurrentTime = time.Date(2019, 8, 31, 10, 01, 0, 0, jst)
	assert.Equal(t, "start", i.executeMode())

	/////////
	i.State = "running"

	// out of range time
	i.CurrentTime = time.Date(2019, 8, 31, 7, 00, 0, 0, jst) // out of range(7:00)
	assert.Equal(t, "none", i.executeMode())
	// in time
	i.CurrentTime = time.Date(2019, 8, 31, 10, 01, 0, 0, jst)
	assert.Equal(t, "none", i.executeMode())

	////////////////
	i.OperationMode = "stop"

	/////////
	i.State = "stopped"

	// out of range time
	i.CurrentTime = time.Date(2019, 8, 31, 7, 00, 0, 0, jst)
	assert.Equal(t, "none", i.executeMode())
	// in time
	i.CurrentTime = time.Date(2019, 8, 31, 10, 01, 0, 0, jst)
	assert.Equal(t, "none", i.executeMode())

	/////////
	i.State = "running"

	// out of range time
	i.CurrentTime = time.Date(2019, 8, 31, 7, 00, 0, 0, jst) // out of range(7:00)
	assert.Equal(t, "stop", i.executeMode())
	// in time
	i.CurrentTime = time.Date(2019, 8, 31, 10, 01, 0, 0, jst)
	assert.Equal(t, "none", i.executeMode())

	////////////////
	i.OperationMode = "auto"

	/////////
	i.State = "stopped"

	// out of range time
	i.CurrentTime = time.Date(2019, 8, 31, 7, 00, 0, 0, jst)
	assert.Equal(t, "none", i.executeMode())
	// in time
	i.CurrentTime = time.Date(2019, 8, 31, 10, 01, 0, 0, jst)
	assert.Equal(t, "start", i.executeMode())

	/////////
	i.State = "running"

	// out of range time
	i.CurrentTime = time.Date(2019, 8, 31, 7, 00, 0, 0, jst) // out of range(7:00)
	assert.Equal(t, "stop", i.executeMode())
	// in time
	i.CurrentTime = time.Date(2019, 8, 31, 10, 01, 0, 0, jst)
	assert.Equal(t, "none", i.executeMode())
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

func TestSetControllable(t *testing.T) {
	i := Instance{}
	var input string

	input = "true"
	i.setControllable(input)
	assert.Equal(t, i.Controllable, true)

	input = "True"
	i.setControllable(input)
	assert.Equal(t, i.Controllable, true)

	input = "TRUE"
	i.setControllable(input)
	assert.Equal(t, i.Controllable, true)

	input = "ｔｒｕｅ"
	i.setControllable(input)
	assert.Equal(t, i.Controllable, false)

	input = "1"
	i.setControllable(input)
	assert.Equal(t, i.Controllable, true)

	input = "false"
	i.setControllable(input)
	assert.Equal(t, i.Controllable, false)

	input = "False"
	i.setControllable(input)
	assert.Equal(t, i.Controllable, false)

	input = "FALSE"
	i.setControllable(input)
	assert.Equal(t, i.Controllable, false)

	input = "no"
	i.setControllable(input)
	assert.Equal(t, i.Controllable, false)

	input = "0"
	i.setControllable(input)
	assert.Equal(t, i.Controllable, false)
}

func TestSetOperationMode(t *testing.T) {
	var input string
	i := Instance{}

	input = "auto"
	i.setOperationMode(input)

	input = "AUTO"
	i.setOperationMode(input)
	assert.Equal(t, "auto", i.OperationMode)

	input = " aUtO"
	i.setOperationMode(input)
	assert.Equal(t, "auto", i.OperationMode)

	input = "   auTo    "
	i.setOperationMode(input)
	assert.Equal(t, "auto", i.OperationMode)

	input = "start"
	i.setOperationMode(input)
	assert.Equal(t, "start", i.OperationMode)

	input = "START"
	i.setOperationMode(input)
	assert.Equal(t, "start", i.OperationMode)

	input = " Start"
	i.setOperationMode(input)
	assert.Equal(t, "start", i.OperationMode)

	input = "   starT    "
	i.setOperationMode(input)
	assert.Equal(t, "start", i.OperationMode)

	input = "stop"
	i.setOperationMode(input)
	assert.Equal(t, "stop", i.OperationMode)

	input = "STOP"
	i.setOperationMode(input)
	assert.Equal(t, "stop", i.OperationMode)

	input = " Stop"
	i.setOperationMode(input)
	assert.Equal(t, "stop", i.OperationMode)

	input = "   stoP    "
	i.setOperationMode(input)
	assert.Equal(t, "stop", i.OperationMode)

	input = "hoge"
	i.setOperationMode(input)
	assert.Equal(t, "", i.OperationMode)
}
