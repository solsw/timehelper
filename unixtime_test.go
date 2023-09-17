package timehelper

import (
	"encoding/binary"
	"reflect"
	"testing"
	"time"
)

func TestUnixTime_MarshalBinary(t *testing.T) {
	tests := []struct {
		name string
		ut   UnixTime
		want int64
	}{
		{name: "zero time", ut: UnixTime(time.Time{}), want: -62135596800},
		{name: "1", ut: UnixTime(TimeYMD(2006, 1, 2)), want: 1136160000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotbb, _ := tt.ut.MarshalBinary()
			got, _ := binary.Varint(gotbb)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnixTime.MarshalBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unixTimeUnmarshalBin(t *testing.T) {
	utz := UnixTime(time.Time{})
	bbz, _ := utz.MarshalBinary()
	ut0 := UnixTime(TimeYMD(1970, 1, 1))
	ut1 := UnixTime(TimeYMD(2006, 1, 2))
	bb1, _ := ut1.MarshalBinary()
	type args struct {
		bb []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *UnixTime
		wantErr bool
	}{
		{name: "00", args: args{}, wantErr: true},
		{name: "01", args: args{bb: make([]byte, 0)}, wantErr: true},
		{name: "zero time", args: args{bb: bbz}, want: &utz},
		{name: "0", args: args{bb: []byte{0}}, want: &ut0},
		{name: "1", args: args{bb: bb1}, want: &ut1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unixTimeUnmarshalBin(tt.args.bb)
			if (err != nil) != tt.wantErr {
				t.Errorf("unixTimeUnmarshalBin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unixTimeUnmarshalBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnixTime_String(t *testing.T) {
	tests := []struct {
		name string
		tr   UnixTime
		want string
	}{
		{name: "1", tr: UnixTime(TimeYMD(2006, 1, 2)), want: "1136160000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.String(); got != tt.want {
				t.Errorf("UnixTime.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unixTimeUnmarshalStr(t *testing.T) {
	ut1 := UnixTime(TimeYMD(2006, 1, 2))
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *UnixTime
		wantErr bool
	}{
		{name: "00", args: args{str: ""}, wantErr: true},
		{name: "01", args: args{str: "qwerty"}, wantErr: true},
		{name: "1", args: args{str: "1136160000"}, want: &ut1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unixTimeUnmarshalStr(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("unixTimeUnmarshalStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unixTimeUnmarshalStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
