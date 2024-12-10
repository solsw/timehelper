package timehelper

import (
	"reflect"
	"testing"
	"time"
)

func TestBuildTime(t *testing.T) {
	type args struct {
		buildDateTimeStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{name: "empty buildDateTimeStr 1",
			args:    args{buildDateTimeStr: ""},
			want:    time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local),
			wantErr: true,
		},
		{name: "empty buildDateTimeStr 2",
			args:    args{buildDateTimeStr: " , asdfgh"},
			want:    time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local),
			wantErr: true,
		},
		{name: "wrong buildDateTimeStr",
			args:    args{buildDateTimeStr: "qwerty"},
			want:    time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local),
			wantErr: true,
		},
		{name: "Linux",
			args:    args{buildDateTimeStr: "2006-01-02T15:04:05"},
			want:    time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local),
			wantErr: false,
		},
		{name: "Windows",
			args:    args{buildDateTimeStr: "2.01.2006T15:04:05"},
			want:    time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildTime(tt.args.buildDateTimeStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
