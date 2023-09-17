package timehelper

import (
	"fmt"
	"time"
)

// ValidMonth reports whether 'month' represents a valid month.
func ValidMonth(month time.Month) error {
	if !(time.January <= month && month <= time.December) {
		return fmt.Errorf("invalid month '%d'", month)
	}
	return nil
}

// ValidWeekday reports whether 'weekday' is valid.
func ValidWeekday(weekday time.Weekday) error {
	if !(time.Sunday <= weekday && weekday <= time.Saturday) {
		return fmt.Errorf("invalid weekday '%d'", weekday)
	}
	return nil
}

// ValidWeek reports whether 'week' is valid for 'year'.
func ValidWeek(year, week int) error {
	_, lastWeek := TimeYMD(year, time.December, 31).ISOWeek()
	if lastWeek == 1 {
		lastWeek = 52
	}
	if !(1 <= week && week <= lastWeek) {
		return fmt.Errorf("invalid week '%d' for year '%d'", week, year)
	}
	return nil
}

// ValidYMD reports whether 'year', 'month', 'day' constitute a valid date.
func ValidYMD(year int, month time.Month, day int) error {
	if err := ValidMonth(month); err != nil {
		return err
	}
	if !(1 <= day && day <= DaysInMonth(year, month)) {
		return fmt.Errorf("invalid date: %d-%d-%d", year, month, day)
	}
	return nil
}

// ValidHMS reports whether 'hour', 'min', 'sec' constitute a valid clock time.
func ValidHMS(hour, min, sec int) error {
	if !(0 <= hour && hour <= 23 && 0 <= min && min <= 59 && 0 <= sec && sec <= 59) {
		return fmt.Errorf("invalid clock time: %d:%d:%d", hour, min, sec)
	}
	return nil
}
