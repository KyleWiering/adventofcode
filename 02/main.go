package main

import (
	"fmt"
	"strings"
)

// A for Rock, B for Paper, and C for Scissors.
// X for Rock, Y for Paper, and Z for Scissors.
// (1 for Rock, 2 for Paper, and 3 for Scissors)
// 0 if you lost, 3 if the round was a draw, and 6 if you won
func main() {
	input := `A Y
B X
C Z`
	score := 0
	strategy := strings.Split(input, "\n")
	for _, value := range strategy {

		round := strings.Split(value, " ")
		score += calcType(round[1]) + calcRound(round[0], round[1])
	}

	fmt.Println(score)
}

func calcType(b string) int {
	switch b {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}
	return 0
}

func calcRound(a string, b string) int {
	switch a {
	case "A": // rock
		switch b {
		case "X": // rock
			return 3
		case "Y": // paper
			return 6
		case "Z": // scissors
			return 0
		}
	case "B": // paper
		switch b {
		case "X": // rock
			return 0
		case "Y": // paper
			return 3
		case "Z": // scissors
			return 6
		}
	case "C": // scissors
		switch b {
		case "X": // rock
			return 6
		case "Y": // paper
			return 0
		case "Z": // scissors
			return 3
		}
	}

	return 0
}


// part two:
// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"

func calcRound2(a string, b string) int {
	switch a {
	case "A": // rock
		switch b {
		case "X": // scissors
			return 0 + 3
		case "Y": // Rock
			return 3 + 1
		case "Z": // Paper
			return 6 + 2
		}
	case "B": // paper
		switch b {
		case "X": // rock
			return 0 + 1
		case "Y": // paper
			return 3 + 2
		case "Z": // scissors
			return 6 + 3
		}
	case "C": // scissors
		switch b {
		case "X": // paper
			return 0 + 2
		case "Y": // scissors
			return 3 + 3
		case "Z": // rock
			return 6 + 1
		}
	}

	return 0
}
