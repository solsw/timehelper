package timehelper

import (
	"strconv"
	"time"
)

// UnixSec is a helper type to marshal/unmarshal [time.Time] as a Unix time in seconds.
type UnixSec time.Time

// String returns decimal string representation of 't'.
// String implements the [fmt.Stringer] interface.
func (t UnixSec) String() string {
	return strconv.FormatInt(time.Time(t).Unix(), 10)
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (t UnixSec) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// used for testing
func unmarshalSec(s string) (*UnixSec, error) {
	sec, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return nil, err
	}
	ut := UnixSec(time.Unix(sec, 0).In(time.UTC))
	return &ut, nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
// Time zone of the unmarshaled value is [time.UTC].
func (t *UnixSec) UnmarshalText(text []byte) error {
	ut, err := unmarshalSec(string(text))
	if err != nil {
		return err
	}
	*t = *ut
	return nil
}

// used for testing
func parseSec(sec int64) UnixSec {
	return UnixSec(time.Unix(int64(sec), 0).In(time.UTC))
}

// ParseSec parses seconds to UnixSec.
func (t *UnixSec) ParseSec(sec int64) {
	*t = parseSec(sec)
}

// MarshalJSON implements the [encoding/json.Marshaler] interface.
func (t UnixSec) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface.
func (t *UnixSec) UnmarshalJSON(bb []byte) error {
	return t.UnmarshalText(bb)
}
