package main

import (
	"testing"

	"github.com/SuperDulli/advent-of-code/util"
)

var example = `>
^>v<
^v^v^v^v^v`

var example2 = `^v
^>v<
^v^v^v^v^v`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example", util.SplitLines(example)[0], 2},
		{"example", util.SplitLines(example)[1], 4},
		{"example", util.SplitLines(example)[2], 2},
		{"actual", util.ReadLines("input.txt")[0], 2565},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example", util.SplitLines(example2)[0], 3},
		{"example", util.SplitLines(example2)[1], 3},
		{"example", util.SplitLines(example2)[2], 11},
		{"actual", util.ReadLines("input.txt")[0], 2639},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
