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
	treesHit := 0
	trees := util.ConvertToMatrix(input)
	slope := util.Vector2D{X: 3, Y: 1}
	for x, y := 0, 0; y < len(trees); x, y = (x + slope.X) % len(trees[y]), y+slope.Y {
		if trees[y][x] == "#" {
			treesHit++
		}
	}
	return treesHit
}

func part2(input []string) int {
	return 0
}
