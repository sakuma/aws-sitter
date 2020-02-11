package holiday

import (
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	var tt time.Time
	var got bool

	// saturday
	tt = time.Date(2019, 8, 31, 7, 34, 0, 0, jst)
	got = IsHoliday(tt)
	if got != true {
		t.Fatal("got: false, should be true")
	}

	// sunday
	tt = time.Date(2019, 9, 1, 5, 59, 59, 0, jst)
	got = IsHoliday(tt)
	if got != true {
		t.Fatal("got: false, should be true")
	}

	// monday
	tt = time.Date(2019, 9, 2, 8, 59, 59, 0, jst)
	got = IsHoliday(tt)
	if got != false {
		t.Fatal("got: true, should be false")
	}
}

func TestIsRunnable(t *testing.T) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	var tt time.Time
	var got bool

	tt = time.Date(2018, 5, 19, 9, 59, 59, 0, jst)
	got = IsRunnable(tt)
	if got != false {
		t.Fatal("got: true, should be false")
	}

	tt = time.Date(2018, 5, 19, 10, 00, 00, 0, jst)
	got = IsRunnable(tt)
	if got != true {
		t.Fatal("got: false, should be true")
	}

	tt = time.Date(2018, 5, 19, 15, 00, 00, 0, jst)
	got = IsRunnable(tt)
	if got != true {
		t.Fatal("got: false, should be true")
	}

	tt = time.Date(2018, 5, 19, 21, 59, 59, 0, jst)
	got = IsRunnable(tt)
	if got != true {
		t.Fatal("got: false, should be true")
	}

	tt = time.Date(2018, 5, 19, 22, 00, 00, 0, jst)
	got = IsRunnable(tt)
	if got != false {
		t.Fatal("got: false, should be true")
	}
}
