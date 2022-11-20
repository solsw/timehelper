package timehelper

import (
	"time"
)

// Durations missing in time package.
const (
	Tick       time.Duration = 100 * time.Nanosecond
	Day        time.Duration = 24 * time.Hour
	Week       time.Duration = 7 * Day
	CommonYear time.Duration = 365 * Day
	LeapYear   time.Duration = 366 * Day
)

// DurationHMS returns time.Duration corresponding to 'hour' hours + 'min' minutes + 'sec' seconds.
func DurationHMS(hour, min, sec int) time.Duration {
	return time.Duration(hour)*time.Hour + time.Duration(min)*time.Minute + time.Duration(sec)*time.Second
}

// MulDuration returns time.Duration that represents a specified multiple of 'd'.
func MulDuration(mul float64, d time.Duration) time.Duration {
	return time.Duration(mul * float64(d))
}

// MulWeek returns time.Duration that represents a specified multiple of week.
func MulWeek(mul float64) time.Duration {
	return MulDuration(mul, Week)
}

// MulDay returns time.Duration that represents a specified multiple of day.
func MulDay(mul float64) time.Duration {
	return MulDuration(mul, Day)
}

// MulHour returns time.Duration that represents a specified multiple of hour.
func MulHour(mul float64) time.Duration {
	return MulDuration(mul, time.Hour)
}

// MulMinute returns time.Duration that represents a specified multiple of minute.
func MulMinute(mul float64) time.Duration {
	return MulDuration(mul, time.Minute)
}

// MulSecond returns time.Duration that represents a specified multiple of second.
func MulSecond(mul float64) time.Duration {
	return MulDuration(mul, time.Second)
}

// MulMillisecond returns time.Duration that represents a specified multiple of millisecond.
func MulMillisecond(mul float64) time.Duration {
	return MulDuration(mul, time.Millisecond)
}

// MulMicrosecond returns time.Duration that represents a specified multiple of microsecond.
func MulMicrosecond(mul float64) time.Duration {
	return MulDuration(mul, time.Microsecond)
}

// MulTick returns time.Duration that represents a specified multiple of tick.
func MulTick(mul float64) time.Duration {
	return MulDuration(mul, Tick)
}

// Days returns the duration as a floating point number of days.
func Days(d time.Duration) float64 {
	day := d / Day
	nsec := d % Day
	return float64(day) + float64(nsec)/(24*60*60*1e9)
}

// Weeks returns the duration as a floating point number of weeks.
func Weeks(d time.Duration) float64 {
	week := d / Week
	nsec := d % Week
	return float64(week) + float64(nsec)/(7*24*60*60*1e9)
}
