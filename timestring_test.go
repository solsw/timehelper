package timehelper

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestNewTimeString(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    TimeString
		wantErr bool
	}{
		{name: "0",
			args: args{t: time.Time{}},
			want: "",
		},
		{name: "1",
			args: args{t: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)},
			want: "2006-01-02T15:04:05Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimeString(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTimeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTimeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeString_Time(t *testing.T) {
	tests := []struct {
		name    string
		ts      TimeString
		want    time.Time
		wantErr bool
	}{
		{name: "0",
			ts:   TimeString(""),
			want: time.Time{},
		},
		{name: "1",
			ts:   TimeString("2006-01-02T15:04:05Z"),
			want: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ts.Time()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeString.Time() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeString.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

type times2 struct {
	T  time.Time  `json:"time,omitempty"`
	Ts TimeString `json:"timeString,omitempty"`
}

func TestTimeStringJSON2(t *testing.T) {
	t0 := time.Time{}
	ts0, _ := NewTimeString(t0)
	t1 := TimeYMD(2006, 1, 2)
	ts1, _ := NewTimeString(t1)
	tests := []struct {
		name string
		t2   times2
		want string
	}{
		{name: "0",
			t2:   times2{T: t0, Ts: ts0},
			want: `{"time":"0001-01-01T00:00:00Z"}`,
		},
		{name: "1",
			t2:   times2{T: t1, Ts: ts1},
			want: `{"time":"2006-01-02T00:00:00Z","timeString":"2006-01-02T00:00:00Z"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb, _ := json.Marshal(tt.t2)
			if !reflect.DeepEqual(string(bb), tt.want) {
				t.Errorf("TimeStringJSON = %v, want %v", string(bb), tt.want)
			}
		})
	}
}
