package main

import (
	"testing"

	"github.com/SuperDulli/advent-of-code/util"
)

const example = `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"example", util.SplitLines(example), 1514},
		{"actual", util.ReadLines("input.txt"), 185371},
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
		{"example", util.SplitLines(example), 0},
		{"actual", util.ReadLines("input.txt"), 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
