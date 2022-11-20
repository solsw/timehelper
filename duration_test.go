package timehelper

import (
	"testing"
	"time"
)

func TestMulWeek(t *testing.T) {
	type args struct {
		mul float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{mul: 1.0}, want: "168h0m0s"},
		{name: "2", args: args{mul: 1.0 / 7}, want: "24h0m0s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MulWeek(tt.args.mul).String(); got != tt.want {
				t.Errorf("MulWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulDay(t *testing.T) {
	type args struct {
		mul float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{mul: 1.0}, want: "24h0m0s"},
		{name: "2", args: args{mul: 2.2}, want: "52h48m0s"},
		{name: "3", args: args{mul: 1.0 / 3}, want: "8h0m0s"},
		{name: "4", args: args{mul: 0.01}, want: "14m24s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MulDay(tt.args.mul).String(); got != tt.want {
				t.Errorf("MulDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulHour(t *testing.T) {
	type args struct {
		mul float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{mul: 1.0}, want: "1h0m0s"},
		{name: "2", args: args{mul: 1.5}, want: "1h30m0s"},
		{name: "3", args: args{mul: 1.0 / 60}, want: "1m0s"},
		{name: "4", args: args{mul: 1.0 / 3600}, want: "1s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MulHour(tt.args.mul).String(); got != tt.want {
				t.Errorf("MulHour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulMinute(t *testing.T) {
	type args struct {
		mul float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{mul: 1.0}, want: "1m0s"},
		{name: "2", args: args{mul: 1441 + 1.0/60}, want: "24h1m1s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MulMinute(tt.args.mul).String(); got != tt.want {
				t.Errorf("MulMinute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulSecond(t *testing.T) {
	type args struct {
		mul float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{mul: 1.0}, want: "1s"},
		{name: "2", args: args{mul: 1.0 / 11}, want: "90.90909ms"},
		{name: "3", args: args{mul: 86401}, want: "24h0m1s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MulSecond(tt.args.mul).String(); got != tt.want {
				t.Errorf("MulSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulMillisecond(t *testing.T) {
	type args struct {
		mul float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{mul: 1.0}, want: "1ms"},
		{name: "2", args: args{mul: 1.0 / 8}, want: "125µs"},
		{name: "3", args: args{mul: 1.0 / 9}, want: "111.111µs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MulMillisecond(tt.args.mul).String(); got != tt.want {
				t.Errorf("MulMillisecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulMicrosecond(t *testing.T) {
	type args struct {
		mul float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{mul: 1.0}, want: "1µs"},
		{name: "2", args: args{mul: 10000}, want: "10ms"},
		{name: "3", args: args{mul: 1.0 / 6}, want: "166ns"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MulMicrosecond(tt.args.mul).String(); got != tt.want {
				t.Errorf("MulMicrosecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulTick(t *testing.T) {
	type args struct {
		mul float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{mul: 1.0}, want: "100ns"},
		{name: "2", args: args{mul: 0.25}, want: "25ns"},
		{name: "3", args: args{mul: 1.0 / 3}, want: "33ns"},
		{name: "4", args: args{mul: 10}, want: "1µs"},
		{name: "5", args: args{mul: 10000}, want: "1ms"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MulTick(tt.args.mul).String(); got != tt.want {
				t.Errorf("MulTick() = %v, want %v", got, tt.want)
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
		{name: "1", args: args{d: MulDay(2.5)}, want: 2.5},
		{name: "2", args: args{d: MulHour(84)}, want: 3.5},
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
		{name: "1", args: args{d: MulDay(14)}, want: 2},
		{name: "2", args: args{d: MulHour(84)}, want: 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Weeks(tt.args.d); got != tt.want {
				t.Errorf("Weeks() = %v, want %v", got, tt.want)
			}
		})
	}
}
