package function

import (
	"testing"
)

func Test_isPossible(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{111111}, 1},
		{"test2", args{122345}, 1},
		{"test3", args{223450}, 0},
		{"test4", args{123789}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossible(tt.args.x); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPossible2(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{112233}, 1},
		{"test2", args{123444}, 0},
		{"test3", args{111122}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossible2(tt.args.x); got != tt.want {
				t.Errorf("isPossible2(%d) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}

func Test_getPossible(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{111000, 111222}, 46},
		{"test2", args{111111, 111111}, 1},
		{"test3", args{223450, 223455}, 1},
		{"test4", args{223450, 223454}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPossible(tt.args.start, tt.args.end, isPossible); got != tt.want {
				t.Errorf("getPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}