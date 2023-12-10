package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`


func main() {
	sum :=0
	for _, v := range strings.Split(input, "\n") {
		fmt.Print(v, " ")
		var letters []string
		var numbers []string

				v = strings.ReplaceAll(v, "one", "one1one")
				v = strings.ReplaceAll(v, "two", "two2two")
				v = strings.ReplaceAll(v, "three", "three3three")
				v = strings.ReplaceAll(v, "four", "four4four")
				v = strings.ReplaceAll(v, "five", "five5five")
				v = strings.ReplaceAll(v, "six", "six6six")
				v = strings.ReplaceAll(v, "seven", "seven7seven")
				v = strings.ReplaceAll(v, "eight", "eight8eight")
				v = strings.ReplaceAll(v, "nine", "nine9nine")


		for _, z := range strings.Split(v, "") {
			if _, err := strconv.Atoi(z); err == nil {
				numbers = append(numbers, z)
			} else {
				letters = append(letters, z)
			}
		}



		number, _ := strconv.Atoi(numbers[0] +""+ numbers[len(numbers)-1])
		fmt.Print(number, sum)
		sum+=number
		fmt.Print(" = ",sum)
		fmt.Println()

	}
	fmt.Println(sum)
}
