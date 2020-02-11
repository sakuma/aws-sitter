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
