package timehelper

import (
	"reflect"
	"testing"
	"time"
)

func TestJulianToGregorian(t *testing.T) {
	type args struct {
		julian Date
	}
	tests := []struct {
		name    string
		args    args
		want    Date
		wantErr bool
	}{
		{name: "1e",
			args:    args{julian: Date{1500, time.June, 23}},
			want:    Date{},
			wantErr: true,
		},
		{name: "2e",
			args:    args{julian: Date{2300, time.June, 23}},
			want:    Date{},
			wantErr: true,
		},
		{name: "1",
			args: args{julian: Date{1960, time.June, 10}},
			want: Date{1960, time.June, 23},
		},
		{name: "2",
			args: args{julian: Date{1821, time.October, 30}},
			want: Date{1821, time.November, 11},
		},
		{name: "3",
			args: args{julian: Date{1672, time.May, 30}},
			want: Date{1672, time.June, 9},
		},
		{name: "4",
			args: args{julian: Date{1584, time.March, 18}},
			want: Date{1584, time.March, 28},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JulianToGregorian(tt.args.julian)
			if (err != nil) != tt.wantErr {
				t.Errorf("JulianToGregorian() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JulianToGregorian() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGregorianToJulian(t *testing.T) {
	type args struct {
		gregorian Date
	}
	tests := []struct {
		name    string
		args    args
		want    Date
		wantErr bool
	}{
		{name: "1e",
			args:    args{gregorian: Date{1500, time.June, 23}},
			want:    Date{},
			wantErr: true,
		},
		{name: "2e",
			args:    args{gregorian: Date{2300, time.June, 23}},
			want:    Date{},
			wantErr: true,
		},
		{name: "1",
			args: args{gregorian: Date{1960, time.June, 23}},
			want: Date{1960, time.June, 10},
		},
		{name: "2",
			args: args{gregorian: Date{1821, time.November, 11}},
			want: Date{1821, time.October, 30},
		},
		{name: "3",
			args: args{gregorian: Date{1672, time.June, 9}},
			want: Date{1672, time.May, 30},
		},
		{name: "4",
			args: args{gregorian: Date{1584, time.March, 28}},
			want: Date{1584, time.March, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GregorianToJulian(tt.args.gregorian)
			if (err != nil) != tt.wantErr {
				t.Errorf("GregorianToJulian() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GregorianToJulian() = %v, want %v", got, tt.want)
			}
		})
	}
}
