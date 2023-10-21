package timehelper

import (
	"time"
)

// Durations missing in [time] package.
const (
	Tick       time.Duration = 100 * time.Nanosecond
	Day        time.Duration = 24 * time.Hour
	Week       time.Duration = 7 * Day
	CommonYear time.Duration = 365 * Day
	LeapYear   time.Duration = 366 * Day
)

// DurationHMS returns [time.Duration] corresponding to 'hour' hours + 'min' minutes + 'sec' seconds.
func DurationHMS(hour, min, sec int) time.Duration {
	return time.Duration(hour)*time.Hour +
		time.Duration(min)*time.Minute +
		time.Duration(sec)*time.Second
}

// MulDuration returns [time.Duration] that represents a specified multiple of 'd'.
func MulDuration(mul float64, d time.Duration) time.Duration {
	return time.Duration(mul * float64(d))
}

// Days returns the duration as a floating point number of [Day] s.
func Days(d time.Duration) float64 {
	day := d / Day
	nsec := d % Day
	return float64(day) + float64(nsec)/(24*60*60*1e9)
}

// Weeks returns the duration as a floating point number of [Week] s.
func Weeks(d time.Duration) float64 {
	week := d / Week
	nsec := d % Week
	return float64(week) + float64(nsec)/(7*24*60*60*1e9)
}
