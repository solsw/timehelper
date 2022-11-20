package timehelper

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"time"
)

// UnixTime is a helper type to marshal/unmarshal time.Time as a Unix time in seconds.
type UnixTime time.Time

// String returns decimal representation of 't'.
// String implements the fmt.Stringer interface.
func (t UnixTime) String() string {
	return strconv.FormatInt(time.Time(t).Unix(), 10)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
// MarshalBinary returns nil error.
func (t UnixTime) MarshalBinary() ([]byte, error) {
	bb := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(bb, time.Time(t).Unix())
	return bb[:n], nil
}

// used for testing
func unixTimeUnmarshalBin(bb []byte) (*UnixTime, error) {
	sec, err := binary.ReadVarint(bytes.NewReader(bb))
	if err != nil {
		return nil, err
	}
	ut := UnixTime(time.Unix(sec, 0).In(time.UTC))
	return &ut, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (t *UnixTime) UnmarshalBinary(data []byte) error {
	ut, err := unixTimeUnmarshalBin(data)
	if err != nil {
		return err
	}
	*t = *ut
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t UnixTime) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func unixTimeUnmarshalStr(str string) (*UnixTime, error) {
	secs, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		return nil, err
	}
	ut := UnixTime(time.Unix(secs, 0).In(time.UTC))
	return &ut, nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// Location of the unmarshaled value is time.UTC.
func (t *UnixTime) UnmarshalText(text []byte) error {
	ut, err := unixTimeUnmarshalStr(string(text))
	if err != nil {
		return err
	}
	*t = *ut
	return nil
}
