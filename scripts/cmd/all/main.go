package main

import (
	"github.com/SuperDulli/advent-of-code/scripts"
	"github.com/SuperDulli/advent-of-code/scripts/aoc"
)

func main() {
	day, year, cookie := aoc.ParseFlags()
	scripts.CreateSkeleton(day, year)
	aoc.GetInput(day, year, cookie)
	aoc.GetPrompt(day, year, cookie)
}
