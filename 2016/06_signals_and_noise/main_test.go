package main

import (
	"testing"

	"github.com/SuperDulli/advent-of-code/util"
)

const example = `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{"example", util.SplitLines(example), "easter"},
		{"actual", util.ReadLines("input.txt"), "ygjzvzib"},
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
		input []string
		want  string
	}{
		{"example", util.SplitLines(example), "advent"},
		{"actual", util.ReadLines("input.txt"), "pdesmnoz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}
