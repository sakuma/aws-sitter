package sitter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsWithinScheduleTime(t *testing.T) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	var testCases = []struct {
		runSchedule string
		inputTime   time.Time
		expected    bool
	}{
		{"10-20", time.Date(2011, 4, 30, 9, 59, 0, 0, jst), false},
		{"10-20", time.Date(2011, 4, 30, 10, 00, 0, 0, jst), true},
		{"10-20", time.Date(2011, 4, 30, 15, 00, 0, 0, jst), true},
		{"10-20", time.Date(2011, 4, 30, 20, 59, 0, 0, jst), true},
		{"10-20", time.Date(2011, 4, 30, 21, 00, 0, 0, jst), false},

		{"9-23:4-6", time.Date(2011, 4, 28, 16, 45, 0, 0, jst), true},
		{"9-23:3-6", time.Date(2011, 4, 29, 16, 45, 0, 0, jst), true},
		{"9-23:3-6", time.Date(2011, 4, 30, 16, 45, 0, 0, jst), true},
		{"9-23:3-5", time.Date(2011, 4, 30, 16, 45, 0, 0, jst), false},

		{"9-23:3,6", time.Date(2011, 4, 30, 16, 45, 0, 0, jst), true},
		{"9-23:1,5", time.Date(2011, 4, 30, 16, 45, 0, 0, jst), false},
	}
	for _, tt := range testCases {
		i := Instance{RunSchedule: tt.runSchedule, CurrentTime: tt.inputTime}
		assert.Equal(t, tt.expected, i.isWithinScheduleTime())
	}
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
	var testCases = []struct {
		state    string
		expected bool
	}{
		{"running", true},
		{"available", true},
		{"stopping", false},
		{"stopped", false},
		{"shutting", false},
		{"terminated", false},
		{"hoge", false},
	}

	for _, tt := range testCases {
		i := Instance{State: tt.state}
		assert.Equal(t, tt.expected, i.isRunning())
	}
}

func TestIsStopped(t *testing.T) {
	var testCases = []struct {
		state    string
		expected bool
	}{
		{"stopped", true},
		{"stopping", true},
		{"shutting", true},
		{"terminated", true},
		{"running", false},
	}

	for _, tt := range testCases {
		i := Instance{State: tt.state}
		assert.Equal(t, tt.expected, i.isStopped())
	}
}

func TestSetControllable(t *testing.T) {
	i := Instance{}
	var testCases = []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"True", true},
		{"TRUE", true},
		{"ｔｒｕｅ", true},
		{"1", true},
		{"t r u e", true},

		{"false", false},
		{"False", false},
		{"FALSE", false},
		{"　FA　LSE　", false},
		{"no", false},
		{"0", false},
	}

	for _, tt := range testCases {
		i.setControllable(tt.input)
		assert.Equal(t, tt.expected, i.Controllable)
	}
}

func TestSetOperationMode(t *testing.T) {
	i := Instance{}
	var testCases = []struct {
		mode     string
		expected string
	}{
		{"auto", "auto"},
		{"AUTO", "auto"},
		{" aUtO", "auto"},
		{"   auTo    ", "auto"},
		{"ａuT　o", "auto"},

		{"start", "start"},
		{"START", "start"},
		{" Start　", "start"},
		{" starT    ", "start"},
		{" ｓtaR　T", "start"},

		{"stop", "stop"},
		{"STOP", "stop"},
		{"Stop", "stop"},
		{" Stop\n", "stop"},
		{"   stoP    ", "stop"},

		{"hoge", ""},
	}
	for _, tt := range testCases {
		i.setOperationMode(tt.mode)
		assert.Equal(t, tt.expected, i.OperationMode)
	}
}

func TestSetRunSchedule(t *testing.T) {
	var testCases = []struct {
		input    string
		expected string
	}{
		{"10-19", "10-19"},
		{"10 - 19", "10-19"},
		{" 8-22", "8-22"},
		{"8-22　", "8-22"}, // wide-width space
		{"12-23 ", "12-23"},
		{"  7-9  ", "7-9"},
		{"１-２２", "1-22"},
		// several fullwidth haypen
		{"2ー１１", "2-11"},
		{"3−１１", "3-11"},
		{"4―１１", "4-11"},
		{"5－１１", "5-11"},
		{"6﹣11", "6-11"},
		{"7⼀１１", "7-11"},
		{"8ー１１", "8-11"},
		{"9㆒１１", "9-11"},
		{"12-34:1-5　", "12-34:1-5"},
		{"12-34：1-5", "12-34:1-5"},
		{"12-34:１-５", "12-34:1-5"},
		{"12-34:１，３", "12-34:1,3"},
	}

	for _, tt := range testCases {
		i := Instance{}
		i.setRunSchedule(tt.input)
		assert.Equal(t, tt.expected, i.RunSchedule)
	}
}
