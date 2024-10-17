package main

import (
	"flag"
	"fmt"

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
	numbers := util.ConvertToNumbers(input)
	for i, n := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			m := numbers[j]
			result := n + m
			if result == 2020 {
				return n * m
			}
		}
	}
	return 0
}

func part2(input []string) int {
	numbers := util.ConvertToNumbers(input)
	for i, n := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			m := numbers[j]
			if m+n > 2020 {
				continue
			}
			for k := j + 1; k < len(numbers); k++ {
				o := numbers[k]
				result := n + m + o
				if result == 2020 {
					return n * m * o
				}
			}
		}
	}
	return 0
}
