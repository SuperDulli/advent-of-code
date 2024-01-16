package main

import (
	"testing"

	"github.com/SuperDulli/advent-of-code/util"
)

const example = `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		seconds int
		want  int
	}{
		{"example", util.SplitLines(example), 1000, 1120},
		{"actual", util.ReadLines("input.txt"), 2503, 2660},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, tt.seconds); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		seconds int
		want  int
	}{
		{"example", util.SplitLines(example), 1000, 689},
		{"actual", util.ReadLines("input.txt"), 2503, 1256},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input, tt.seconds); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
