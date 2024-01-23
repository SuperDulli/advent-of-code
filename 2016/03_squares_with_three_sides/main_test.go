package main

import (
	"testing"

	"github.com/SuperDulli/advent-of-code/util"
)

const example = `5 10 25`

const example2 = `101 301 501
102 302 502
103 303 503
201 401 601
202 402 602
203 403 603`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"example", util.SplitLines(example), 0},
		{"actual", util.ReadLines("input.txt"), 917},
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
		{"example", util.SplitLines(example2), 6},
		{"actual", util.ReadLines("input.txt"), 1649},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
