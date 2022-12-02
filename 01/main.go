package main

import (
	"fmt"
	"strconv"
	"strings"
)

// A for Rock, B for Paper, and C for Scissors.
// X for Rock, Y for Paper, and Z for Scissors.
// (1 for Rock, 2 for Paper, and 3 for Scissors)
// 0 if you lost, 3 if the round was a draw, and 6 if you won
func main() {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	elves := strings.Split(input, "\n\n")
	var cals []int
	for _, value := range elves {
		sub := strings.Split(value, "\n")
		total := 0
		for i := 0; i < len(sub); i++ {
			intVal, _ := strconv.Atoi(sub[i])
			total += intVal
		}

		cals = append(cals, total)
	}
	fmt.Println(maxArr(cals))
}

func maxArr(input []int) int {

	second := 0
	third := 0
	max := 0
	for _, val := range input {
		if val > max {
			third = second
			second = max
			max = val
		} else {
			if val > second {
				third = second
				second = val
			} else {
				if val > third {
					third = val
				}
			}
		}
	}
	return max + second + third
}


// func maxArrPart1(input []int) int {
// 	max := 0
// 	for _, val := range input {
// 		if val > max {
// 			max = val
// 		}
// 	}
// 	return max
// }
