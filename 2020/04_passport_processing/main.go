package main

import (
	"flag"
	"fmt"
	"regexp"
	"slices"

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
	valid := 0
	rawPassport := ""
	for _, line := range input {
		if line == "" {
			passport := parsePassport(rawPassport)
			if isValid(passport) {
				valid++
			}
			rawPassport = ""
			continue
		}
		rawPassport = rawPassport + " " + line
	}
	passport := parsePassport(rawPassport)
	if isValid(passport) {
		valid++
	}
	return valid
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

func isValid(passport map[string]string) bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	if len(passport) < len(required) {
		return false
	}
	for _, field := range required {
		value, exists := passport[field]
		if !exists {
			return false
		}
		switch field {
		case "byr":
			if !isValidYear(value, 1920, 2002) {
				return false
			}
		case "iyr":
			if !isValidYear(value, 2010, 2020) {
				return false
			}
		case "eyr":
			if !isValidYear(value, 2020, 2030) {
				return false
			}
		case "hgt":
			if !isValidHeight(value) {
				return false
			}
		case "hcl":
			if !isValidHairColor(value) {
				return false
			}
		case "ecl":
			if !isValidEyeColor(value) {
				return false
			}
		case "pid":
			if !isValidPassportID(value) {
				return false
			}
		}
	}
	return true

}

func isValidYear(raw string, min, max int) bool {
	if len(raw) < 4 {
		return false
	}
	year := util.ConvertToNumber(raw)
	if year < min || max < year {
		return false
	}
	return true
}

func isValidHeight(raw string) bool {
	matches := regexp.MustCompile(`^(\d+)(cm|in)$`).FindStringSubmatch(raw)
	if len(matches) < 2 {
		return false
	}
	height := util.ConvertToNumber(matches[1])
	switch matches[2] {
	case "cm":
		if height < 150 || 193 < height {
			return false
		}
	case "in":
		if height < 59 || 76 < height {
			return false
		}
	}
	return true
}

func isValidHairColor(raw string) bool {
	return regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString(raw)
}

func isValidEyeColor(raw string) bool {
	valid := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	return slices.Contains(valid, raw)
}

func isValidPassportID(raw string) bool {
	return regexp.MustCompile(`^\d{9}$`).MatchString(raw)
}
