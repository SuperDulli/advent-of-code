package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
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
	magic := 0
	for {
		hash := mD5Hash(input + strconv.Itoa(magic))
		if strings.HasPrefix(hash, "00000") {
			break
		}
		magic++
	}
	return magic
}

func part2(input string) int {
	magic := 0
	for {
		hash := mD5Hash(input + strconv.Itoa(magic))
		if strings.HasPrefix(hash, "000000") {
			break
		}
		magic++
	}
	return magic
}

func mD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
