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

func part1(doorId string) string {
	password := ""
	index := 0
	for len(password) < 8 {
		hash := mD5Hash(doorId + strconv.Itoa(index))
		if strings.HasPrefix(hash, "00000") {
			password += string(hash[5])
		}
		index++
	}
	return password
}

func part2(doorId string) string {
	password := "--------"
	index := 0
	found := 0
	for found < 8 {
		hash := mD5Hash(doorId + strconv.Itoa(index))
		index++
		if strings.HasPrefix(hash, "00000") {
			posInPassword, err := strconv.Atoi(string(hash[5]))
			if err != nil{
				continue
			}
			if posInPassword < 0 || posInPassword >= 8 || password[posInPassword] != '-' {
				// Use only the first result for each position, and ignore invalid positions.
				continue
			}
			password = replaceAtIndex(password, rune(hash[6]), posInPassword)
			found++
		}
	}
	return password
}

func mD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}
