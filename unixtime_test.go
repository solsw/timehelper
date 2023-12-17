package timehelper

import (
	"reflect"
	"testing"
	"time"
)

func TestUnixTime_String(t *testing.T) {
	tests := []struct {
		name string
		tr   UnixTime
		want string
	}{
		{name: "1",
			tr:   UnixTime(TimeYMD(2006, 1, 2)),
			want: "1136160000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.String(); got != tt.want {
				t.Errorf("UnixTime.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unmarshalStr(t *testing.T) {
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
		{name: "00",
			args:    args{str: ""},
			wantErr: true,
		},
		{name: "01",
			args:    args{str: "qwerty"},
			wantErr: true,
		},
		{name: "1",
			args: args{str: "1136160000"},
			want: &ut1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unmarshalStr(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshalStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSecs(t *testing.T) {
	type args struct {
		s uint64
	}
	tests := []struct {
		name string
		args args
		want UnixTime
	}{
		{name: "0",
			args: args{s: 0},
			want: UnixTime(TimeYMD(1970, 1, 1)),
		},
		{name: "1",
			args: args{s: 1257894000},
			want: UnixTime(time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSecs(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSecs() = %v, want %v", time.Time(got), time.Time(tt.want))
			}
		})
	}
}

func Test_parseMSecs(t *testing.T) {
	type args struct {
		ms uint64
	}
	tests := []struct {
		name string
		args args
		want UnixTime
	}{
		{name: "0",
			args: args{ms: 0},
			want: UnixTime(TimeYMD(1970, 1, 1)),
		},
		{name: "1",
			args: args{ms: 1702828942634},
			want: UnixTime(time.Date(2023, 12, 17, 16, 2, 22, 634000000, time.UTC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseMSecs(tt.args.ms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMSecs() = %v, want %v", time.Time(got), time.Time(tt.want))
			}
		})
	}
}
