package timehelper

import (
	"strconv"
	"time"
)

// UnixMsec is a helper type to marshal/unmarshal [time.Time] as a Unix time in milliseconds.
type UnixMsec time.Time

// String returns decimal string representation of 't'.
// String implements the [fmt.Stringer] interface.
func (t UnixMsec) String() string {
	return strconv.FormatInt(time.Time(t).UnixMilli(), 10)
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (t UnixMsec) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// used for testing
func unmarshalMsec(s string) (*UnixMsec, error) {
	msec, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return nil, err
	}
	ut := UnixMsec(time.UnixMilli(msec).In(time.UTC))
	return &ut, nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
// Time zone of the unmarshaled value is [time.UTC].
func (t *UnixMsec) UnmarshalText(text []byte) error {
	ut, err := unmarshalMsec(string(text))
	if err != nil {
		return err
	}
	*t = *ut
	return nil
}

// used for testing
func parseMsec(msec int64) UnixMsec {
	return UnixMsec(time.UnixMilli(msec).In(time.UTC))
}

// ParseMSecs parses milliseconds to UnixMsec.
func (t *UnixMsec) ParseMSecs(msec int64) {
	*t = parseMsec(msec)
}
