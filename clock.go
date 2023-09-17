package timehelper

import (
	"time"
)

// Clock combines hour, minute and second.
type Clock struct {
	Hour   int
	Minute int
	Second int
}

// Valid reports whether 'c' represents a valid clock time.
func (c Clock) Valid() error {
	return ValidHMS(c.Hour, c.Minute, c.Second)
}

// Duration returns [time.Duration] corresponding to c.Hour hours + c.Minute minutes + c.Second seconds.
func (c Clock) Duration() time.Duration {
	return DurationHMS(c.Hour, c.Minute, c.Second)
}

// TimeDate returns [time.Time] corresponding to 'd' + 'c'.
// The returned value has zero nanoseconds and [time.UTC] time zone.
func (c Clock) TimeDate(d Date) time.Time {
	return time.Date(d.Year, d.Month, d.Day, c.Hour, c.Minute, c.Second, 0, time.UTC)
}
