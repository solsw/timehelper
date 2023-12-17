package timehelper

import (
	"math/bits"
	"strconv"
	"time"
)

// UnixTime is a helper type to marshal/unmarshal [time.Time] as a Unix time in seconds.
type UnixTime time.Time

// String returns decimal string representation of 't'.
// String implements the [fmt.Stringer] interface.
func (t UnixTime) String() string {
	return strconv.FormatInt(time.Time(t).Unix(), 10)
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (t UnixTime) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// used for testing
func unmarshalStr(str string) (*UnixTime, error) {
	secs, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		return nil, err
	}
	ut := UnixTime(time.Unix(secs, 0).In(time.UTC))
	return &ut, nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
// Time zone of the unmarshaled value is [time.UTC].
func (t *UnixTime) UnmarshalText(text []byte) error {
	ut, err := unmarshalStr(string(text))
	if err != nil {
		return err
	}
	*t = *ut
	return nil
}

// used for testing
func parseSecs(s uint64) UnixTime {
	return UnixTime(time.Unix(int64(s), 0).In(time.UTC))
}

// ParseSecs parses seconds to UnixTime.
func (t *UnixTime) ParseSecs(s uint64) {
	*t = parseSecs(s)
}

// used for testing
func parseMSecs(ms uint64) UnixTime {
	sec, msec := bits.Div64(0, ms, 1000)
	return UnixTime(time.Unix(int64(sec), int64(msec*1000000)).In(time.UTC))
}

// ParseMSecs parses milliseconds to UnixTime.
func (t *UnixTime) ParseMSecs(ms uint64) {
	*t = parseMSecs(ms)
}
