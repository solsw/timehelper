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
	if j.Before(TimeYMD(1582, time.October, 5)) || j.After(TimeYMD(2200, time.February, 14)) {
		return Date{}, fmt.Errorf("unsupported Julian date: %d-%d-%d", julian.Year, julian.Month, julian.Day)
	}
	var g time.Time
	switch {
	case j.After(TimeYMD(2100, time.February, 15)):
		g = j.AddDate(0, 0, 14)
	case j.After(TimeYMD(1900, time.February, 16)):
		g = j.AddDate(0, 0, 13)
	case j.After(TimeYMD(1800, time.February, 17)):
		g = j.AddDate(0, 0, 12)
	case j.After(TimeYMD(1700, time.February, 18)):
		g = j.AddDate(0, 0, 11)
	default:
		g = j.AddDate(0, 0, 10)
	}
	return Date{g.Year(), g.Month(), g.Day()}, nil
}

// GregorianToJulian converts [Gregorian] date to [Julian] date.
//
// [Gregorian]: https://en.wikipedia.org/wiki/Gregorian_calendar
// [Julian]: https://en.wikipedia.org/wiki/Julian_calendar
func GregorianToJulian(gregorian Date) (Date, error) {
	if err := gregorian.Valid(); err != nil {
		return Date{}, err
	}
	g := gregorian.Time()
	if g.Before(TimeYMD(1582, time.October, 15)) || g.After(TimeYMD(2200, time.February, 28)) {
		return Date{}, fmt.Errorf("unsupported Gregorian date: %d-%d-%d", gregorian.Year, gregorian.Month, gregorian.Day)
	}
	var j time.Time
	switch {
	case g.After(TimeYMD(2100, time.February, 28)):
		j = g.AddDate(0, 0, -14)
	case g.After(TimeYMD(1900, time.February, 28)):
		j = g.AddDate(0, 0, -13)
	case g.After(TimeYMD(1800, time.February, 28)):
		j = g.AddDate(0, 0, -12)
	case g.After(TimeYMD(1700, time.February, 28)):
		j = g.AddDate(0, 0, -11)
	default:
		j = g.AddDate(0, 0, -10)
	}
	return Date{j.Year(), j.Month(), j.Day()}, nil
}
