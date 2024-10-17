package main

import (
	"flag"
	"fmt"
	"regexp"
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
	validPasswords := 0
	for _, line := range input {
		passwordAndPolicy := parseLine(line)
		if passwordAndPolicy.check() {
			validPasswords++
		}
	}

	return validPasswords
}

func part2(input []string) int {
	return 0
}

type PasswordAndPolicy struct {
	min      int
	max      int
	symbol   rune
	password string
}

func parseLine(line string) PasswordAndPolicy {
	matches := regexp.MustCompile(`^(\d+)-(\d+)\s(\w):\s(\w+)$`).FindStringSubmatch(line)
	return PasswordAndPolicy{
		min:      util.ConvertToNumber(matches[1]),
		max:      util.ConvertToNumber(matches[2]),
		symbol:   rune(matches[3][0]),
		password: matches[4],
	}
}

func (p PasswordAndPolicy) check() bool {
	symbolCount := strings.Count(p.password, string(p.symbol))
	return symbolCount >= p.min && symbolCount <= p.max
}
