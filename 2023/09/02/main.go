package main

import (
	"fmt"
	"strconv"
	"strings"
)

type numbList struct {
	Numbers []int
	History *numbList
}

func (n *numbList) populateList(newList []int) {
	n.Numbers = newList
	fmt.Println(newList)
	allZeroes := true
	var history []int
	for i:=1; i < len(n.Numbers); i++ {
		diff := n.Numbers[i] - n.Numbers[i-1]
		if diff != 0 {
			allZeroes = false
		}
		history = append(history, diff)
	}
	if allZeroes {
		history = append(history, 0)

		return
	} else {
		n.History = &numbList{}
		n.History.populateList(history)
		n.Numbers = append(n.Numbers, n.History.Numbers[len(n.History.Numbers)-1] + n.Numbers[len(n.Numbers)-1])
		//fmt.Println(n.Numbers )
	}

}

var testinput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func main() {
	fullList := strings.Split(input, "\n")
	total := 0
	for _, v := range fullList {
		var numberList []int
		numberStrings := strings.Split(v, " ")
		for _, z := range numberStrings {
			value, _ := strconv.Atoi(z)

			numberList = append([]int{value}, numberList...)
		}
		var numbs numbList
		numbs.populateList(numberList)
		fmt.Println("adding", numbs.Numbers[len(numbs.Numbers)-1])
		total += numbs.Numbers[len(numbs.Numbers)-1]
	}
	fmt.Println(total)
}
