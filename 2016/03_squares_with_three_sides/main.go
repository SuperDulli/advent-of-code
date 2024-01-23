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
	possibleCount := 0
	for _, line := range input {
		sides := util.ConvertToNumbers(strings.Fields(line))
		if trianglePossible(sides) {
			possibleCount++
		}
	}
	return possibleCount
}

func part2(input []string) int {
	possibleCount := 0

	triangles := make([][]int, 3)
	for n, lines := range input {
		numbers := util.ConvertToNumbers(strings.Fields(lines))
		for i := 0; i < 3; i++ {
			triangles[i] = append(triangles[i], numbers[i])
		}
		if n > 0 && (n+1) % 3 == 0 {
			for i := 0; i < 3; i++ {
				if trianglePossible(triangles[i]) {
					possibleCount++
				}
			}
			triangles = make([][]int, 3)
		}
	}

	return possibleCount
}

func trianglePossible(sides []int) bool {
	slices.Sort(sides)
	return sides[0]+sides[1] > sides[2]
}
