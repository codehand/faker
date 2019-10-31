package faker

import (
	"fmt"
	"testing"
)

// TestDate ...
func TestDate(t *testing.T) {
	for i := 0; i < randIntRange(1, 100); i++ {
		t.Run(fmt.Sprintf("TestDate%d", i+1), func(t *testing.T) {
			got := Date()
			if got.String() == "" {
				t.Errorf("TestDate() = %v, want not empty", got)
			}
		})
	}
}

// TestDay ...
func TestDay(t *testing.T) {
	type args struct {
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestDay1",
			args: args{
				m: 1,
			},
			want: 31,
		},
		{
			name: "TestDay2",
			args: args{
				m: 9,
			},
			want: 30,
		},
		{
			name: "TestDay3",
			args: args{
				m: 2,
			},
			want: 28,
		},
	}
	//1, 3, 5, 7, 8, 10, 12
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day(tt.args.m); got < 1 || got > tt.want {
				t.Errorf("Day() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestYear ...
func TestYear(t *testing.T) {
	for i := 0; i < randIntRange(1, 100); i++ {
		t.Run(fmt.Sprintf("TestYear%d", i+1), func(t *testing.T) {
			got := Year()
			if got < 1900 {
				t.Errorf("Year() = %v, want not empty", got)
			}
		})
	}
}

// TestHour ...
func TestHour(t *testing.T) {
	for i := 0; i < randIntRange(1, 100); i++ {
		t.Run(fmt.Sprintf("TestHour%d", i+1), func(t *testing.T) {
			got := Hour()
			if got < 0 || got > 23 {
				t.Errorf("Hour() = %v, want not invalid", got)
			}
		})
	}
}
