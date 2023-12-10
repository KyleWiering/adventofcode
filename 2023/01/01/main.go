package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func main() {
	sum :=0
	for _, v := range strings.Split(input, "\n") {
		var letters []string
		var numbers []string
		for _, z := range strings.Split(v, "") {
			if _, err := strconv.Atoi(z); err == nil {
				numbers = append(numbers, z)
			} else {
				letters = append(letters, z)
			}
		}

		number, _ := strconv.Atoi(numbers[0] +""+ numbers[len(numbers)-1])
		sum+=number


	}
	fmt.Println(sum)
}
