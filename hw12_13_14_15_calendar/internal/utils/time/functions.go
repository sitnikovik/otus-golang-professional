package time

import "time"

// EndOfWeek returns the time of the next Sunday from the given time.
// This function calculates the time up to and including Sunday.
func EndOfWeek(t time.Time) time.Time {
	// calc the days to next Sunday
	daysToSunday := int(time.Sunday - t.Weekday())
	if daysToSunday < 0 {
		daysToSunday += 7
	}
	toTime := t.AddDate(0, 0, daysToSunday)
	// set time to 23:59:59
	toTime = time.Date(
		toTime.Year(),
		toTime.Month(),
		toTime.Day(),
		23, 59, 59, 0,
		toTime.Location(),
	)

	return toTime
}

// EndOfDay returns the time at the end of the given day.
// This function calculates the time up to and including 23:59:59 of the given day.
func EndOfDay(t time.Time) time.Time {
	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		23, 59, 59, 0,
		t.Location(),
	)
}

// EndOfMonth returns the time at the end of the given month.
// This function calculates the time up to and including the last day of the month at 23:59:59.
func EndOfMonth(t time.Time) time.Time {
	// Get the first day of the next month
	firstDayNextMonth := time.Date(
		t.Year(),
		t.Month()+1,
		1,
		0, 0, 0, 0,
		t.Location(),
	)
	// Subtract one second to get the last moment of the current month
	lastMomentCurrentMonth := firstDayNextMonth.Add(-time.Second)
	return lastMomentCurrentMonth
}

// UpToDays returns the time at the end of the given year.
func UpToDays(t time.Time, days int) time.Time {
	toTime := t.AddDate(0, 0, days)
	// set time to 00:00:00
	toTime = time.Date(
		toTime.Year(),
		toTime.Month(),
		toTime.Day(),
		0, 0, 0, 0,
		toTime.Location(),
	)

	return toTime
}
