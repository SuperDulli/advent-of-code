package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/SuperDulli/advent-of-code/util"
)

func main() {
	var part int
	var inputFileName string
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.StringVar(&inputFileName, "file", "input.txt", "alternate input file")
	flag.Parse()
	fmt.Println("Running part", part)

	input := util.ReadLines(inputFileName)[0]

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	visited := make(map[util.Vector2D]int)
	pos := util.Vector2D{
		X: 0,
		Y: 0,
	}
	visited[pos] = 1

	for _, c := range input {
		pos = move(pos, c)
		visited[pos] = 1
	}
	return len(visited)
}

func part2(input string) int {
	visited := make(map[util.Vector2D]int)
	pos := util.Vector2D{
		X: 0,
		Y: 0,
	}
	posRobo := util.Vector2D{
		X: 0,
		Y: 0,
	}
	visited[pos] = 1

	for n, c := range input {
		if n%2 == 0 {
			pos = move(pos, c)
			visited[pos] = 1
		} else {
			posRobo = move(posRobo, c)
			visited[posRobo] = 1
		}
	}
	return len(visited)
}

func move(start util.Vector2D, direction rune) util.Vector2D {
	switch direction {
	case '^':
		return util.MoveUp(start)
	case 'v':
		return util.MoveDown(start)
	case '<':
		return util.MoveLeft(start)
	case '>':
		return util.MoveRight(start)
	default:
		log.Fatal("unknown direction: ", direction)
	}
	return util.Vector2D{}
}
