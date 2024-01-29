package main

import (
	"container/heap"
	"flag"
	"fmt"
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

func part1(input []string) string {
	matrix := util.ConvertToMatrix(input)
	matrix = util.Transpose(matrix)

	builder := strings.Builder{}
	for _, line := range matrix {
		letter := mostFrequentLetter(line)
		builder.WriteString(letter)
	}

	return builder.String()
}

func part2(input []string) string {
	matrix := util.ConvertToMatrix(input)
	matrix = util.Transpose(matrix)

	builder := strings.Builder{}
	for _, line := range matrix {
		letter := leastFrequentLetter(line)
		builder.WriteString(letter)
	}

	return builder.String()
}

func mostFrequentLetter(line []string) string {
	freq := make(map[string]int)
	for _, letter := range line {
		freq[letter] += 1
	}
	h := util.GetHeap(freq)
	return heap.Pop(h).(util.KV).Key
}

func leastFrequentLetter(line []string) string {
	freq := make(map[string]int)
	for _, letter := range line {
		freq[letter] += 1
	}
	h := util.GetHeap(freq)
	for i := 0; i < len(freq) - 1; i++ {
		heap.Pop(h)
	}
	return heap.Pop(h).(util.KV).Key
}
