package main

import (
	"flag"
	"fmt"
	"log"
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
	keypad := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	start := util.Vector2D{X: 1, Y: 1} // key 5

	code := 0
	for _, line := range input {
		for _, letter := range line {
			direction := translateDirection(letter)
			start.SafeMove(direction, 3, 3)
		}
		key := keypad[start.Y][start.X]
		code *= 10
		code += key
	}
	return code
}

func part2(input []string) string {
	keypad := [][]string{
		{"#", "#", "1", "#", "#"},
		{"#", "2", "3", "4", "#"},
		{"5", "6", "7", "8", "9"},
		{"#", "A", "B", "C", "#"},
		{"#", "#", "D", "#", "#"},
	}
	pos := util.Vector2D{X: 0, Y: 2} // key 5

	code := strings.Builder{}
	for _, line := range input {
		for _, letter := range line {
			direction := translateDirection(letter)
			newPos := util.SafeMove(pos, direction, 5, 5)
			if keypad[newPos.Y][newPos.X] != "#" {
				pos = newPos
			}
		}
		key := keypad[pos.Y][pos.X]
		code.WriteString(key)
	}

	return code.String()
}

func translateDirection(r rune) util.Vector2D {
	switch r {
	case 'D':
		return util.Up2D()
	case 'U':
		return util.Down2D()
	case 'L':
		return util.Left2D()
	case 'R':
		return util.Right2D()
	default:
		log.Fatal("Unknown direction: ", r)
	}
	return util.Vector2D{}
}
