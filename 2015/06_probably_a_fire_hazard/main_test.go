package main

import (
	"testing"

	"github.com/SuperDulli/advent-of-code/util"
)

const example = `turn on 0,0 through 999,999
toggle 0,0 through 999,0
turn off 499,499 through 500,500`

const example2 = `turn on 0,0 through 0,0
toggle 0,0 through 999,999`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"example", util.SplitLines(example), 1_000_000 - 1_004},
		{"actual", util.ReadLines("input.txt"), 543903},
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
		input []string
		want  int
	}{
		{"example", util.SplitLines(example2), 2000001},
		{"actual", util.ReadLines("input.txt"), 14687245},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
