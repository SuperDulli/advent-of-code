package main

import (
	"container/heap"
	"flag"
	"fmt"
	"regexp"
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
	sum := 0
	for _, line := range input {
		room := parseRoom(line)
		if validateChecksum(room) {
			sum += room.id
		}
	}
	return sum
}

func part2(input []string) int {
	for _, line := range input {
		decrypted := decryptName(strings.Split(line, "[")[0])
		if decrypted == "northpole object storage" {
			return util.ConvertToNumber(regexp.MustCompile(`\d+`).FindString(line))
		}
	}
	return -1
}

type room struct {
	name     string
	id       int
	checksum string
}

func parseRoom(s string) room {
	nameParts := regexp.MustCompile(`[^\d\W]+`).FindAllString(s, -1)
	id := util.ConvertToNumber(regexp.MustCompile(`\d+`).FindString(s))
	checksum := regexp.MustCompile(`\[(\w+)\]`).FindStringSubmatch(s)
	return room{
		name:     strings.Join(nameParts[0:len(nameParts)-1], ""),
		id:       id,
		checksum: checksum[1],
	}
}

func countChars(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, c := range s {
		counts[c] += 1
	}
	return counts
}

func validateChecksum(room room) bool {
	counts := countChars(room.name)
	h := getHeap(counts)
	for i := 0; i < 5; i++ {
		char := heap.Pop(h).(kv).Key
		if char != rune(room.checksum[i]) {
			return false
		}
	}
	return true
}

func decryptName(name string) string {
	decrypted := strings.Builder{}
	groups := regexp.MustCompile(`(\D+)-(\d+)`).FindAllStringSubmatch(name, -1)
	text := groups[0][1]
	id := util.ConvertToNumber(groups[0][2])
	for _, char := range text {		
		decrypted.WriteRune(shiftChar(char, id))
	}
	return decrypted.String()
}

func shiftChar(r rune, amount int) rune {
	if r == '-' {
		return ' '
	}
	newChar := int(r) + (amount % 26)
	if newChar > int('z') {
		newChar = int('a') + newChar - int('z') - 1
	}
	return rune(newChar)
}

func getHeap(m map[rune]int) *KVHeap {
	h := &KVHeap{}
	heap.Init(h)
	for k, v := range m {
		heap.Push(h, kv{k, v})
	}
	return h
}

type kv struct {
	Key   rune
	Value int
}

// See https://golang.org/pkg/container/heap/
type KVHeap []kv

// Note that "Less" is greater-than here so we can pop *larger* items.
// ties broken by alphabetization
func (h KVHeap) Less(i, j int) bool {
	if h[i].Value == h[j].Value {
		return h[i].Key < h[j].Key
	}
	return h[i].Value > h[j].Value
}
func (h KVHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h KVHeap) Len() int      { return len(h) }

func (h *KVHeap) Push(x interface{}) {
	*h = append(*h, x.(kv))
}

func (h *KVHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
