package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
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
	seq := input[0]
	for i := 0; i < 40; i++ {
		seq = lookAndSay(seq)
	}
	return len(seq)
}

func part2(input []string) int {
	seq := input[0]
	for i := 0; i < 50; i++ {
		seq = lookAndSay(seq)
	}
	return len(seq)
}

func lookAndSay(s string) string  {
	matches := regexp.MustCompile(`1{1,3}|2{1,3}|3{1,3}`).FindAllString(s, -1)
	builder := strings.Builder{}
	for _, match := range matches {
		count := len(match)
		builder.WriteString(strconv.Itoa(count))
		builder.WriteByte(match[0])
	}
	return builder.String()
}
