package main
// https://adventofcode.com/2022/leaderboard/private join_key: 2277564-35497492

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	objects := strings.Split(input, "\n")
	FirstCase(objects)
	SecondCase(objects)
}

func FirstCase(objects []string) {
	total := 0
	for _, obj := range objects {
		pairs := strings.Split(obj, ",")
		total = total + isFullOverlap(pairs[0], pairs[1])
	}
	fmt.Println(total)
}

func splitPair(a string) (int, int) {
	splitter := strings.Split(a, "-")
	aa, _ := strconv.Atoi(splitter[0])
	bb, _ := strconv.Atoi(splitter[1])

	return aa, bb
}

func isFullOverlap(a string, b string) int {
	x, y := splitPair(a)
	c, d := splitPair(b)

	if (x <= c && y >= d) || (x >= c && y <= d) {
		//fmt.Println(x, y, ",", c, d)
		return 1
	}
	return 0
}

func SecondCase(objects []string) {
	total := 0
	for _, obj := range objects {
		pairs := strings.Split(obj, ",")
		total = total + isPartialOverlap(pairs[0], pairs[1])
	}
	fmt.Println(total)
}

func isPartialOverlap(a string, b string) int {
	x, y := splitPair(a)
	c, d := splitPair(b)

	if y >= c && d >= x {
		//fmt.Println(x, y, ",", c, d)
		return 1
	}
	return 0
}
