package main

import (
	"flag"
	"fmt"
	"regexp"

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
		sum += difference(line)
	}
	return sum
}

func part2(input []string) int {
	sum := 0
	for _, line := range input {
		sum += encodeDifference(line)
	}
	return sum
}

func difference(s string) int {
	return len(s) - len(regexp.MustCompile(`\\"|\\\\|\\x[0-9a-z]{0,2}`).ReplaceAllString(s, "@")) +2
}

func encodeDifference(s string) int {
	difference := 2 // the double quotes at the beginning and end
	difference +=  len(regexp.MustCompile(`\\`).FindAllString(s, -1))
	difference +=  len(regexp.MustCompile(`"`).FindAllString(s, -1))
	return difference
}
