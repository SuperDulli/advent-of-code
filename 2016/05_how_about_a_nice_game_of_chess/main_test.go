package main

import (
	"testing"

	"github.com/SuperDulli/advent-of-code/util"
)

const example = `abc`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"example", example, "18f47a30"},
		{"actual", util.ReadLines("input.txt")[0], "2414bc77"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"example", example, "05ace8e3"},
		{"actual", util.ReadLines("input.txt")[0], "437e60fc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}
