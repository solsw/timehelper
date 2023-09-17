package timehelper

import (
	"fmt"
	"time"
)

// JulianToGregorian converts [Julian] date to [Gregorian] date.
//
// [Gregorian]: https://en.wikipedia.org/wiki/Gregorian_calendar
// [Julian]: https://en.wikipedia.org/wiki/Julian_calendar
func JulianToGregorian(julian Date) (Date, error) {
	if err := julian.Valid(); err != nil {
		return Date{}, err
	}
	j := julian.Time()
	if j.Before(TimeYMD(1300, time.March, 1)) || j.After(TimeYMD(2400, time.February, 28)) {
		return Date{}, fmt.Errorf("unsupported Julian date: %d-%d-%d", julian.Year, julian.Month, julian.Day)
	}
	var g time.Time
	switch {
	case julian.Time().After(TimeYMD(2300, time.February, 28)):
		g = julian.Time().AddDate(0, 0, 16)
	case julian.Time().After(TimeYMD(2200, time.February, 28)):
		g = julian.Time().AddDate(0, 0, 15)
	case julian.Time().After(TimeYMD(2100, time.February, 28)):
		g = julian.Time().AddDate(0, 0, 14)
	case julian.Time().After(TimeYMD(1900, time.February, 28)):
		g = julian.Time().AddDate(0, 0, 13)
	case julian.Time().After(TimeYMD(1800, time.February, 28)):
		g = julian.Time().AddDate(0, 0, 12)
	case julian.Time().After(TimeYMD(1700, time.February, 28)):
		g = julian.Time().AddDate(0, 0, 11)
	case julian.Time().After(TimeYMD(1500, time.February, 28)):
		g = julian.Time().AddDate(0, 0, 10)
	case julian.Time().After(TimeYMD(1400, time.February, 28)):
		g = julian.Time().AddDate(0, 0, 9)
	case julian.Time().After(TimeYMD(1300, time.February, 28)):
		g = julian.Time().AddDate(0, 0, 8)
	}
	return Date{g.Year(), g.Month(), g.Day()}, nil
}
