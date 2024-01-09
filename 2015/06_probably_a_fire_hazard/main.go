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
	lightMatrix := make([][]int, 1000)
	for i := range lightMatrix {
		lightMatrix[i] = make([]int, 1000)
	}

	for _, line := range input {
		cmd := parse(line)
		for x := cmd.start.X; x <= cmd.end.X; x++ {
			for y := cmd.start.Y; y <= cmd.end.Y; y++ {
				switch cmd.command {
				case "turn on":
					lightMatrix[y][x] = 1
				case "turn off":
					lightMatrix[y][x] = 0
				case "toggle":
					if lightMatrix[y][x] == 1 {
						lightMatrix[y][x] = 0
					} else {
						lightMatrix[y][x] = 1
					}
				}
			}
		}
	}

	count := 0
	for _, row := range lightMatrix {
		for _, light := range row {
			if light == 1 {
				count++
			}
		}
	}
	return count
}

func part2(input []string) int {
	lightMatrix := make([][]int, 1000)
	for i := range lightMatrix {
		lightMatrix[i] = make([]int, 1000)
	}

	for _, line := range input {
		cmd := parse(line)
		for x := cmd.start.X; x <= cmd.end.X; x++ {
			for y := cmd.start.Y; y <= cmd.end.Y; y++ {
				switch cmd.command {
				case "turn on":
					lightMatrix[y][x] += 1
				case "turn off":
					if lightMatrix[y][x] > 0 {
						lightMatrix[y][x] -= 1
					}
				case "toggle":
					lightMatrix[y][x] += 2
				}
			}
		}
	}

	brightness := 0
	for _, row := range lightMatrix {
		for _, light := range row {
				brightness += light
		}
	}
	return brightness
}

type instruction struct {
	command string
	start   util.Vector2D
	end     util.Vector2D
}

func parse(line string) instruction {
	result := instruction{}
	result.command = regexp.MustCompile(`turn on|turn off|toggle`).FindString(line)
	matches := regexp.MustCompile(`(\d+),(\d+) through (\d+),(\d+)`).FindStringSubmatch(line)
	result.start.X = util.ConvertToNumber(matches[1])
	result.start.Y = util.ConvertToNumber(matches[2])
	result.end.X = util.ConvertToNumber(matches[3])
	result.end.Y = util.ConvertToNumber(matches[4])
	return result
}
