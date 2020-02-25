package util

import (
	"fmt"
	"testing"
	"time"
)

func TestIsRunning(t *testing.T) {
	var got bool
	i := Instance{}

	i.State = "runningg"
	got = i.IsRunning()
	if got != true {
		t.Fatal("got: false, should be true")
	}

	i.State = "available"
	got = i.IsRunning()
	if got != true {
		t.Fatal("got: false, should be true")
	}

	i.State = "stopping"
	got = i.IsRunning()
	if got != false {
		t.Fatal("got: true, should be false")
	}

	i.State = "stopped"
	got = i.IsRunning()
	if got != false {
		t.Fatal("got: true, should be false")
	}

	i.State = "shutting"
	got = i.IsRunning()
	if got != false {
		t.Fatal("got: true, should be false")
	}

	i.State = "terminated"
	got = i.IsRunning()
	if got != false {
		t.Fatal("got: true, should be false")
	}

	i.State = "hoge"
	got = i.IsRunning()
	if got != false {
		t.Fatal("got: true, should be false")
	}
}

func TestIsStopped(t *testing.T) {
	var got bool
	i := Instance{}

	i.State = "stopped"
	got = i.IsStopped()
	if got != true {
		t.Fatal("got: false, should be true")
	}

	i.State = "stopping"
	got = i.IsStopped()
	if got != true {
		t.Fatal("got: false, should be true")
	}

	i.State = "shutting"
	got = i.IsStopped()
	if got != true {
		t.Fatal("got: false, should be true")
	}

	i.State = "terminated"
	got = i.IsStopped()
	if got != true {
		t.Fatal("got: false, should be true")
	}

	i.State = "running"
	got = i.IsStopped()
	if got != false {
		t.Fatal("got: true, should be false")
	}
}


func TestIsWithinScheduleTime(t *testing.T) {

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	var tt time.Time
	var got bool
	i := Instance{RunSchedule: "10-19"}

	// out of range(7:00)
	tt = time.Date(2019, 8, 31, 7, 00, 0, 0, jst)
	got = i.isWithinScheduleTime(tt)
	if got != false {
		t.Fatal("got: true, should be false")
	}

	// in time (10:00)
	tt = time.Date(2019, 8, 31, 10, 00, 0, 0, jst)
	got = i.isWithinScheduleTime(tt)
	if got != true {
		t.Fatal("got: false, should be true")
	}

	// in time (10:01)
	tt = time.Date(2019, 8, 31, 10, 01, 0, 0, jst)
	got = i.isWithinScheduleTime(tt)
	if got != true {
		t.Fatal("got: false, should be true")
	}

	// in time (19:59)
	tt = time.Date(2019, 8, 31, 19, 59, 0, 0, jst)
	got = i.isWithinScheduleTime(tt)
	if got != true {
		t.Fatal("got: false, should be true")
	}

	// out of range (20:00)
	tt = time.Date(2019, 8, 31, 20, 00, 0, 0, jst)
	got = i.isWithinScheduleTime(tt)
	if got != false {
		t.Fatal("got: true, should be false")
	}

	// out of range (20:01)
	tt = time.Date(2019, 8, 31, 20, 01, 0, 0, jst)
	got = i.isWithinScheduleTime(tt)
	fmt.Println(got)
	if got != false {
		t.Fatal("got: true, should be false")
	}
}

func TestIsActive(t *testing.T) {
	// TODO
}
