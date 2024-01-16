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
	seconds := 2503

	if part == 1 {
		ans := part1(input, seconds)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input []string, seconds int) int {
	distances := []int{}
	for _, line := range input {
		speed, dashTime, restTime := parseInput(line)
		distance := calcDistance(0, speed, dashTime, restTime, seconds)
		distances = append(distances, distance)
	}
	return slices.Max(distances)
}

func part2(input []string) int {
	return 0
}

func parseInput(line string) (int, int, int) {
	numbers := util.ConvertToNumbers(regexp.MustCompile(`\d+`).FindAllString(line, -1))
	speed, dashTime, restTime := numbers[0], numbers[1], numbers[2]
	return speed, dashTime, restTime
}

func calcDistance(distance, speed, dashTime, restTime, seconds int) int {
	if seconds <= 0 {
		return distance
	}
	if seconds <= dashTime {
		return distance + speed*seconds
	}
	distance += speed * dashTime
	return calcDistance(distance, speed, dashTime, restTime, seconds-dashTime-restTime)
}
