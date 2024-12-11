package timehelper

import (
	"errors"
	"strings"
	"time"
)

// BuildTime parses 'buildDateTimeStr' as local time set by -ldflags during build process.
//
// First, declare 'var buildDateTimeStr string' in the 'main' package.
// Then, add '-ldflags' build flag to 'go build' command.
// Under Windows: -ldflags "-X 'main.buildDateTimeStr=%DATE%T%TIME%'".
// Under Linux: -ldflags "-X 'main.buildDateTimeStr=$(date +%Y-%m-%dT%H:%M:%S)'".
func BuildTime(buildDateTimeStr string) (time.Time, error) {
	if cidx := strings.LastIndex(buildDateTimeStr, ","); cidx >= 0 {
		buildDateTimeStr = buildDateTimeStr[:cidx]
	}
	buildDateTimeStr = strings.ReplaceAll(buildDateTimeStr, " ", "")
	if buildDateTimeStr == "" {
		return time.Time{}, errors.New("empty buildDateTimeStr")
	}
	// try to parse buildDateTimeStr set under Linux: 'buildDateTimeStr=$(date +%Y-%m-%dT%H:%M:%S)'
	t, err := time.ParseInLocation("2006-01-02T15:04:05", buildDateTimeStr, time.Local)
	if err == nil {
		return t, nil
	}
	// try to parse buildDateTimeStr set under Windows: 'buildDateTimeStr=%DATE%T%TIME%'
	t, err = time.ParseInLocation("2.01.2006T15:04:05", buildDateTimeStr, time.Local)
	if err == nil {
		return t, nil
	}
	return time.Time{}, errors.New("wrong buildDateTimeStr")
}
