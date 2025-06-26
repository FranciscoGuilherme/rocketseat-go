package condicionais

import "time"

func isWeekend(y time.Time) bool {
	switch {
		case y.Weekday() > time.Sunday && y.Weekday() < time.Friday:
			return false
		default:
			return true
	}
}

func isWeekend2(y time.Time) bool {
	switch y.Weekday() {
		case time.Saturday, time.Sunday:
			return true
		default:
			return false
	}
}
