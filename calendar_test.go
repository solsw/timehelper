package timehelper

import (
	"reflect"
	"testing"
	"time"
)

func TestJulianToGregorian(t *testing.T) {
	type args struct {
		julian time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{name: "1e",
			args:    args{julian: DateYMD(1000, time.June, 23)},
			want:    time.Time{},
			wantErr: true,
		},
		{name: "2e",
			args:    args{julian: DateYMD(3000, time.February, 8)},
			want:    time.Time{},
			wantErr: true,
		},
		{name: "1",
			args: args{julian: DateYMD(1960, time.June, 23)},
			want: DateYMD(1960, time.June, 23).AddDate(0, 0, 13),
		},
		{name: "2",
			args: args{julian: DateYMD(1959, time.February, 8)},
			want: DateYMD(1959, time.February, 8).AddDate(0, 0, 13),
		},
		{name: "3",
			args: args{julian: DateYMD(1821, time.October, 30)},
			want: DateYMD(1821, time.November, 11),
		},
		{name: "4",
			args: args{julian: DateYMD(1672, time.May, 30)},
			want: DateYMD(1672, time.June, 9),
		},
		{name: "5",
			args: args{julian: DateYMD(1584, time.March, 18)},
			want: DateYMD(1584, time.March, 28),
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
