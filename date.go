package timehelper

import (
	"time"
)

// Date combines year, month and day.
type Date struct {
	Year  int
	Month time.Month
	Day   int
}

// DateFromTime extracts Date part from 't'.
func DateFromTime(t time.Time) Date {
	return Date{
		Year:  t.Year(),
		Month: t.Month(),
		Day:   t.Day(),
	}
}

// Valid reports whether 'd' represents a valid date.
func (d Date) Valid() error {
	return ValidYMD(d.Year, d.Month, d.Day)
}

// Time returns [time.Time] corresponding to 'd'.
// The returned value has zero clock and [time.UTC] time zone.
func (d Date) Time() time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, time.UTC)
}

// TimeClock returns [time.Time] corresponding to 'd' + 'c'.
// The returned value has zero nanoseconds and [time.UTC] time zone.
func (d Date) TimeClock(c Clock) time.Time {
	return time.Date(d.Year, d.Month, d.Day, c.Hour, c.Minute, c.Second, 0, time.UTC)
}
