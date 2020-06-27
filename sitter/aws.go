package sitter

import (
	"strconv"
	"strings"
	"time"

	"github.com/sakuma/aws-sitter/lib/util"
	"golang.org/x/text/width"
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
	CurrentTime      time.Time
}

func (i *Instance) isWithinScheduleTime() bool {
	// TODO: invalid format check
	runSchedule := i.RunSchedule
	splittedSchedule := strings.Split(i.RunSchedule, ":")
	timeRange := splittedSchedule[0]
	var weekdays string

	// jadge of weekday
	if strings.Contains(runSchedule, ":") {
		weekdays = splittedSchedule[1]
		switch {
		case strings.Contains(weekdays, "-"):
			weekRange := strings.Split(weekdays, "-")
			weekdayString := strconv.Itoa(int(i.CurrentTime.Weekday()))
			min, _ := strconv.Atoi(weekRange[0])
			max, _ := strconv.Atoi(weekRange[1])
			weekStrings := util.MakeWeekStrings(min, max)
			if !strings.Contains(weekStrings, weekdayString) {
				return false
			}
		case strings.Contains(weekdays, ","):
			weekdayString := strconv.Itoa(int(i.CurrentTime.Weekday()))
			if !strings.Contains(weekdays, weekdayString) {
				return false
			}
		}
	}

	// jadge of time range
	times := strings.Split(timeRange, "-")
	from, _ := strconv.Atoi(strings.TrimSpace(times[0]))
	to, _ := strconv.Atoi(strings.TrimSpace(times[1]))
	if from <= i.CurrentTime.Hour() && i.CurrentTime.Hour() <= to {
		return true
	}
	return false
}

func (i *Instance) executeMode() string {
	switch i.OperationMode {
	case "start":
		if i.isStopped() && i.isWithinScheduleTime() {
			return "start"
		}
	case "stop":
		if i.isRunning() && !i.isWithinScheduleTime() {
			return "stop"
		}
	case "auto":
		if i.isWithinScheduleTime() {
			if i.isStopped() {
				return "start"
			}
		} else {
			if i.isRunning() {
				return "stop"
			}
		}
	default:
		return "none"
	}
	return "none"
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

func (i *Instance) setControllable(inputValue string) {
	// NOTE: Full Width Char is false
	// [Maybe][TODO]: convert to Full to Half
	b, _ := strconv.ParseBool(inputValue)
	i.Controllable = b
}

func (i *Instance) setOperationMode(inputValue string) {
	var v string
	v = strings.TrimSpace(inputValue)
	v = strings.ToLower(v)

	if v == "auto" || v == "start" || v == "stop" {
		i.OperationMode = v
	} else {
		i.OperationMode = ""
	}
}

func (i *Instance) setRunSchedule(inputValue string) {
	v := inputValue
	v = util.SpaceReplaceAll(v)
	v = util.HyphenReplaceAll(v)
	v = width.Narrow.String(v)
	i.RunSchedule = v
}
