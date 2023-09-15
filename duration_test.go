package timehelper

import (
	"reflect"
	"testing"
	"time"
)

func TestMulDuration(t *testing.T) {
	type args struct {
		mul float64
		d   time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1Week", args: args{mul: 1.0, d: Week}, want: "168h0m0s"},
		{name: "2Week", args: args{mul: 1.0 / 7, d: Week}, want: "24h0m0s"},
		{name: "1Day", args: args{mul: 1.0, d: Day}, want: "24h0m0s"},
		{name: "2Day", args: args{mul: 2.2, d: Day}, want: "52h48m0s"},
		{name: "3Day", args: args{mul: 1.0 / 3, d: Day}, want: "8h0m0s"},
		{name: "4Day", args: args{mul: 0.01, d: Day}, want: "14m24s"},
		{name: "1Hour", args: args{mul: 1.0, d: time.Hour}, want: "1h0m0s"},
		{name: "2Hour", args: args{mul: 1.5, d: time.Hour}, want: "1h30m0s"},
		{name: "3Hour", args: args{mul: 1.0 / 60, d: time.Hour}, want: "1m0s"},
		{name: "4Hour", args: args{mul: 1.0 / 3600, d: time.Hour}, want: "1s"},
		{name: "1Minute", args: args{mul: 1.0, d: time.Minute}, want: "1m0s"},
		{name: "2Minute", args: args{mul: 1441 + 1.0/60, d: time.Minute}, want: "24h1m1s"},
		{name: "1Second", args: args{mul: 1.0, d: time.Second}, want: "1s"},
		{name: "2Second", args: args{mul: 1.0 / 11, d: time.Second}, want: "90.90909ms"},
		{name: "3Second", args: args{mul: 86401, d: time.Second}, want: "24h0m1s"},
		{name: "1Millisecond", args: args{mul: 1.0, d: time.Millisecond}, want: "1ms"},
		{name: "2Millisecond", args: args{mul: 1.0 / 8, d: time.Millisecond}, want: "125µs"},
		{name: "3Millisecond", args: args{mul: 1.0 / 9, d: time.Millisecond}, want: "111.111µs"},
		{name: "1Microsecond", args: args{mul: 1.0, d: time.Microsecond}, want: "1µs"},
		{name: "2Microsecond", args: args{mul: 10000, d: time.Microsecond}, want: "10ms"},
		{name: "3Microsecond", args: args{mul: 1.0 / 6, d: time.Microsecond}, want: "166ns"},
		{name: "1Tick", args: args{mul: 1.0, d: Tick}, want: "100ns"},
		{name: "2Tick", args: args{mul: 0.25, d: Tick}, want: "25ns"},
		{name: "3Tick", args: args{mul: 1.0 / 3, d: Tick}, want: "33ns"},
		{name: "4Tick", args: args{mul: 10, d: Tick}, want: "1µs"},
		{name: "5Tick", args: args{mul: 10000, d: Tick}, want: "1ms"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MulDuration(tt.args.mul, tt.args.d).String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MulDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDays(t *testing.T) {
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{d: MulDuration(2.5, Day)}, want: 2.5},
		{name: "2", args: args{d: MulDuration(84, time.Hour)}, want: 3.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Days(tt.args.d); got != tt.want {
				t.Errorf("Days() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeeks(t *testing.T) {
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{d: MulDuration(14, Day)}, want: 2},
		{name: "2", args: args{d: MulDuration(84, time.Hour)}, want: 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Weeks(tt.args.d); got != tt.want {
				t.Errorf("Weeks() = %v, want %v", got, tt.want)
			}
		})
	}
}
