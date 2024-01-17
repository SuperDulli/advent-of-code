package main

import (
	"testing"

	"github.com/SuperDulli/advent-of-code/util"
)

const example = `Sue 1: goldfish: 6, trees: 9, akitas: 0
Sue 2: goldfish: 7, trees: 1, akitas: 0
Sue 3: cars: 10, akitas: 6, perfumes: 7
Sue 4: perfumes: 2, vizslas: 0, cars: 6`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		// {"example", util.SplitLines(example), 0},
		{"actual", util.ReadLines("input.txt"), 0},
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
