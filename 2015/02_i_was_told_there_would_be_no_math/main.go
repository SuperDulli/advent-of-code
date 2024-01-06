package main

import (
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/SuperDulli/advent-of-code/util"
)

func main() {
	var part int
	var inputFileName string
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.StringVar(&inputFileName, "file", "input.txt", "alternate input file")
	flag.Parse()
	fmt.Println("Running part", part)

	input := util.ReadLines(inputFileName)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		sum += area(getDimensions(line))
	}
	return sum
}

func part2(input []string) int {
	sum := 0
	for _, line := range input {
		dimensions := getDimensions(line)
		sum += ribbon(dimensions) + volume(dimensions)
	}
	return sum
}

func getDimensions(s string) []int {
	fields := strings.FieldsFunc(s, func(r rune) bool {
		return r == 'x'
	})
	return util.ConvertToNumbers(fields)
}

// returns the needed area of wrapping paper including slack in square feet
func area(dimensions []int) int {
	// 2*l*w + 2*w*h + 2*h*l
	// 2 * (l*w + w*h + h*l)
	side1 := dimensions[0] * dimensions[1]
	side2 := dimensions[1] * dimensions[2]
	side3 := dimensions[2] * dimensions[0]
	return 2*(side1+side2+side3) + min(side1, side2, side3)
}

// returns the needed length of ribbon in feet
func ribbon(dimensions []int) int {
	// 2 * (smallest_side + medium_size_side)
	slices.Sort(dimensions)
	return 2 * (dimensions[0] + dimensions[1])
}

// returns the needed length of ribbon in feet which equal the volume of the prism
func volume(dimensions []int) int {
	// l*w*h
	return dimensions[0] * dimensions[1] * dimensions[2]
}
