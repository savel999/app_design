package time

import (
	"testing"
	"time"
)

func TestGetDaysDifference(t *testing.T) {
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "from more then to",
			args: args{
				from: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: 0,
		},
		{
			name: "case 1",
			args: args{
				from: time.Date(2021, 1, 1, 13, 0, 0, 0, time.UTC),
				to:   time.Date(2021, 1, 2, 12, 0, 0, 0, time.UTC),
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				from: time.Date(2021, 1, 1, 13, 0, 0, 0, time.UTC),
				to:   time.Date(2021, 1, 2, 14, 0, 0, 0, time.UTC),
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				from: time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC),
				to:   time.Date(2021, 1, 1, 13, 0, 0, 0, time.UTC),
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDaysDifference(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("GetDaysDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
