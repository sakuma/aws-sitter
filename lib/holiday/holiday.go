package holiday

import "time"

func IsHoliday(t time.Time) bool {
	// TODO: national holiday
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}

func IsRunnable() bool {
	// TODO
	return false
}
