package timehelper

import (
	"time"
)

// PrevWeekday returns [time.Weekday] previous to 'weekday'.
func PrevWeekday(weekday time.Weekday) time.Weekday {
	if weekday == time.Sunday {
		return time.Saturday
	}
	return weekday - 1
}

// NextWeekday returns [time.Weekday] next after 'weekday'.
func NextWeekday(weekday time.Weekday) time.Weekday {
	if weekday == time.Saturday {
		return time.Sunday
	}
	return weekday + 1
}

// WeekdaysInYear returns number of 'weekdays' in the 'year'.
func WeekdaysInYear(year int, weekday time.Weekday) int {
	jan1Weekday := TimeYMD(year, time.January, 1).Weekday()
	if (jan1Weekday == weekday) || (jan1Weekday == PrevWeekday(weekday) && IsLeapYear(year)) {
		return 53
	}
	return 52
}

// PrevClosestWeekday returns date of the 'weekday' closest to 't' and not after 't'.
func PrevClosestWeekday(t time.Time, weekday time.Weekday) time.Time {
	d := int(weekday - t.Weekday())
	if d == 0 {
		return t
	}
	if d > 0 {
		d -= 7
	}
	return t.AddDate(0, 0, d)
}

// NextClosestWeekday returns date of the 'weekday' closest to 't' and not before 't'.
func NextClosestWeekday(t time.Time, weekday time.Weekday) time.Time {
	d := int(weekday - t.Weekday())
	if d == 0 {
		return t
	}
	if d < 0 {
		d += 7
	}
	return t.AddDate(0, 0, d)
}

// FirstWeekdayInYear returns the first occurrence of the 'weekday' in the 'year'.
// The returned value has zero clock and [time.UTC] time zone.
func FirstWeekdayInYear(year int, weekday time.Weekday) time.Time {
	return NextClosestWeekday(TimeYMD(year, time.January, 1), weekday)
}

// LastWeekdayInYear returns the last occurrence of the 'weekday' in the 'year'.
// The returned value has zero clock and [time.UTC] time zone.
func LastWeekdayInYear(year int, weekday time.Weekday) time.Time {
	return PrevClosestWeekday(TimeYMD(year, time.December, 31), weekday)
}
