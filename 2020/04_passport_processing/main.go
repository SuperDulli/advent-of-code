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
	valid := 0
	rawPassport := ""
	for _, line := range input {
		if line == "" {
			passport := parsePassport(rawPassport)
			if containsRequiredFields(passport) {
				valid++
			}
			rawPassport = ""
			continue
		}
		rawPassport = rawPassport + " " + line
	}
	passport := parsePassport(rawPassport)
	if containsRequiredFields(passport) {
		valid++
	}
	return valid
}

func part2(input []string) int {
	return 0
}

func parsePassport(rawPassport string) map[string]string {
	passport := make(map[string]string)

	matches := regexp.MustCompile(`(\w+):(\S+)(\s|$)`).FindAllStringSubmatch(rawPassport, -1)
	for _, m := range matches {
		passport[m[1]] = m[2]
	}

	return passport
}

func containsRequiredFields(passport map[string]string) bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	if len(passport) < len(required) {
		return false
	}
	for _, r := range required {
		_, exists := passport[r]
		if !exists {
			return false
		}
	}
	return true
}
