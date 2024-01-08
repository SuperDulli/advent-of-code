package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/SuperDulli/advent-of-code/util"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre" // golang's regex engine does not support back references
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
	count := 0
	for _, line := range input {
		if isNice(line) {
			count++
		}
	}
	return count
}

func part2(input []string) int {
	count := 0
	for _, line := range input {
		if pcre.MustCompile(`(..).*\1`, 0).MatcherString(line, 0).Matches() &&
			pcre.MustCompile(`(.).\1`, 0).MatcherString(line, 0).Matches() {
			count++
		}
	}
	return count
}

func isNice(s string) bool {
	vowelCount := 0
	hasRepetition := false
	bannedSeq := map[string]int{
		"ab": 1,
		"cd": 1,
		"pq": 1,
		"xy": 1,
	}
	previous := ' '
	for n, c := range s {
		if vowelCount < 3 && strings.ContainsRune("aeiou", c) {
			vowelCount++
		}
		if n == 0 {
			previous = c
			continue
		}
		substr := string(previous) + string(c)
		if !hasRepetition && c == previous {
			hasRepetition = true
		}
		if _, ok := bannedSeq[substr]; ok {
			return false
		}
		previous = c
	}
	return vowelCount >= 3 && hasRepetition
}
