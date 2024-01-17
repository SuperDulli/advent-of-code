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
	for n, sue := range parseInput(input) {
		if matchesResult(sue) {
			return n + 1
		}
	}
	return 0
}

func part2(input []string) int {
	for n, sue := range parseInput(input) {
		if matchesResultPart2(sue) {
			return n + 1
		}
	}
	return 0
}

func parseInput(input []string) []map[string]int {
	sues := []map[string]int{}
	for _, line := range input {
		fields := regexp.MustCompile(`(\w+): (\d+)`).FindAllStringSubmatch(line, -1)
		sue := make(map[string]int, len(fields))
		for i := 0; i < len(fields); i++ {
			sue[fields[i][1]] = util.ConvertToNumber(fields[i][2])
		}
		sues = append(sues, sue)
	}
	return sues
}

func matchesResult(sue map[string]int) bool {
	result := make(map[string]int)
	result["children"] = 3
	result["cats"] = 7
	result["samoyeds"] = 2
	result["pomeranians"] = 3
	result["akitas"] = 0
	result["vizslas"] = 0
	result["goldfish"] = 5
	result["trees"] = 3
	result["cars"] = 2
	result["perfumes"] = 1

	for key, value := range sue {
		if result[key] != value {
			return false
		}
	}
	return true
}

func matchesResultPart2(sue map[string]int) bool {
	result := make(map[string]int)
	result["children"] = 3
	result["cats"] = 7
	result["samoyeds"] = 2
	result["pomeranians"] = 3
	result["akitas"] = 0
	result["vizslas"] = 0
	result["goldfish"] = 5
	result["trees"] = 3
	result["cars"] = 2
	result["perfumes"] = 1

	for key, value := range sue {
		if key == "cats" || key == "trees" {
			if value <= result[key] {
				return false
			} else {
				continue
			}
		}
		if key == "pomeranians" || key == "goldfish" {
			if value >= result[key] {
				return false
			} else {
				continue
			}
		}
		if result[key] != value {
			return false
		}
	}
	return true
}
