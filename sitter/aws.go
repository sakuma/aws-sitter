package sitter

import (
	"strconv"
	"strings"
	"time"
)

type Instance struct {
	InstanceType     string
	AvailabilityZone string
	Name             string
	ResourceType     string
	ID               string
	Controllable     bool
	StopOnly         bool
	State            string
	OperationMode    string
	RunSchedule      string
}

func (i *Instance) IsRunning() bool {
	// NOTE: rds status
	// "available", "stopping"
	switch i.State {
	case "running", "available":
		return true
	case "stopping", "stopped", "shutting", "terminated":
		return false
	case "pending":
		// unknown: maybe false
		return false
	default:
		return false
	}
}

func (i *Instance) IsStopped() bool {
	list := []string{"stopping", "stopped", "shutting", "terminated"}
	for _, s := range list {
		if s == i.State {
			return true
		}
	}
	return false
}

func (i *Instance) isWithinScheduleTime(t time.Time) bool {
	// TODO: invalid format check
	times := strings.Split(i.RunSchedule, "-")
	from, _ := strconv.Atoi(strings.TrimSpace(times[0]))
	to, _ := strconv.Atoi(strings.TrimSpace(times[1]))
	if from <= t.Hour() && t.Hour() <= to {
		return true
	}
	return false
}

func (i *Instance) IsActive() bool {
	// TODO: force runnning

	if !i.isOperatable() {
		return true
	}

	t := time.Now()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	current := t.In(jst)

	// if holiday.IsHoliday(current) {
	// 	return false
	// }

	if i.isWithinScheduleTime(current) {
		return true
	}

	return false
}

func (i *Instance) isOperatable() bool {
	if !i.Controllable {
		return false
	}

	switch i.OperationMode {
	case "start":
		return true
	case "stop":
		return true
	case "auto":
		return true
	}
	return false
}
