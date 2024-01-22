package main

import (
	"testing"

	"github.com/SuperDulli/advent-of-code/util"
)

const example = `ULL
RRDDD
LURDL
UUUUD`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"example", util.SplitLines(example), 1985},
		{"actual", util.ReadLines("input.txt"), 44558},
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
		want  string
	}{
		{"example", util.SplitLines(example), "5DB3"},
		{"actual", util.ReadLines("input.txt"), "6BBAD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
