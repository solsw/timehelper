package timehelper

import (
	"reflect"
	"testing"
	"time"
)

func TestUnixSec_String(t *testing.T) {
	tests := []struct {
		name string
		tr   UnixSec
		want string
	}{
		{name: "1",
			tr:   UnixSec(TimeYMD(2006, 1, 2)),
			want: "1136160000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.String(); got != tt.want {
				t.Errorf("UnixSec.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unmarshalSec(t *testing.T) {
	ut1 := UnixSec(TimeYMD(2006, 1, 2))
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *UnixSec
		wantErr bool
	}{
		{name: "00",
			args:    args{s: ""},
			wantErr: true,
		},
		{name: "01",
			args:    args{s: "qwerty"},
			wantErr: true,
		},
		{name: "1",
			args: args{s: "1136160000"},
			want: &ut1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unmarshalSec(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalSec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshalSec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSec(t *testing.T) {
	type args struct {
		sec int64
	}
	tests := []struct {
		name string
		args args
		want UnixSec
	}{
		{name: "0",
			args: args{sec: 0},
			want: UnixSec(TimeYMD(1970, 1, 1)),
		},
		{name: "1",
			args: args{sec: 1257894000},
			want: UnixSec(time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSec(tt.args.sec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSec() = %v, want %v", time.Time(got), time.Time(tt.want))
			}
		})
	}
}
