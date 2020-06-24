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

func (i *Instance) executeMode() string {
	switch i.OperationMode {
	case "start":
		if i.isShouldBeStart() {
			return "start"
		}
	case "stop":
		if i.isShouldBeStop() {
			return "stop"
		}
	case "auto":
		t := time.Now()
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		current := t.In(jst)

		if i.isWithinScheduleTime(current) {
			return "start"
		} else {
			return "stop"
		}
	default:
		return "none"
	}
	return "none"
}

func (i *Instance) isShouldBeStart() bool {
	if i.isStopped() {
		return true
	}
	return false
}

func (i *Instance) isShouldBeStop() bool {
	if i.isRunning() {
		return true
	}
	return false
}

func (i *Instance) isRunning() bool {
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

func (i *Instance) isStopped() bool {
	list := []string{"stopping", "stopped", "shutting", "terminated"}
	for _, s := range list {
		if s == i.State {
			return true
		}
	}
	return false
}
