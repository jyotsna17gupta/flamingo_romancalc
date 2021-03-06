package pinkflamingo

import (
	"testing"
)

func TestIsFibonacci(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   bool
	}{
		{
			name:   "when number is 8",
			number: 8,
			want:   true,
		},
		{
			name:   "when number is 1",
			number: 1,
			want:   true,
		},
		{
			name:   "when number is 16",
			number: 16,
			want:   false,
		},
		{
			name:   "when number is 4",
			number: 4,
			want:   false,
		},
		{
			name:   "when number is 22",
			number: 22,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFibonacci(tt.number); got != tt.want {
				t.Errorf("IsFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPerfectSquare(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "when number is 25",
			n:    25,
			want: true,
		},
		{
			name: "when number is 21",
			n:    21,
			want: false,
		},
		{
			name: "when number is 15",
			n:    15,
			want: false,
		},
		{
			name: "when number is 4",
			n:    4,
			want: true,
		},
		{
			name: "when number is 100",
			n:    100,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.n); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
