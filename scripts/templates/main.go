package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SuperDulli/advent-of-code/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	input := util.ReadLines(os.Args[1])

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input []string) int {
	return 0
}

func part2(input []string) int {
	return 0
}
