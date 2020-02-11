package holiday

import "time"

func IsHoliday(t time.Time) bool {
	// TODO: national holiday
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}

func IsRunnable(t time.Time) bool {
	switch {
	case t.Hour() < 10:
		return false
	case t.Hour() < 22:
		return true
	default:
	return false
}
}
