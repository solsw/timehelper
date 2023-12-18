package timehelper

import (
	"reflect"
	"testing"
	"time"
)

func TestUnixMsec_String(t *testing.T) {
	tests := []struct {
		name string
		tr   UnixMsec
		want string
	}{
		{name: "1",
			tr:   UnixMsec(TimeYMD(2006, 1, 2)),
			want: "1136160000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.String(); got != tt.want {
				t.Errorf("UnixMsec.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unmarshalMsec(t *testing.T) {
	ut1 := UnixMsec(TimeYMD(2006, 1, 2))
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *UnixMsec
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
			args: args{s: "1136160000000"},
			want: &ut1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unmarshalMsec(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalMsec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshalMsec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseMsec(t *testing.T) {
	type args struct {
		msec int64
	}
	tests := []struct {
		name string
		args args
		want UnixMsec
	}{
		{name: "0",
			args: args{msec: 0},
			want: UnixMsec(TimeYMD(1970, 1, 1)),
		},
		{name: "1",
			args: args{msec: 1702828942634},
			want: UnixMsec(time.Date(2023, 12, 17, 16, 2, 22, 634000000, time.UTC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseMsec(tt.args.msec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMsec() = %v, want %v", time.Time(got), time.Time(tt.want))
			}
		})
	}
}
