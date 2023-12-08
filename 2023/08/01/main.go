package main

import (
	"fmt"
	"strings"
)

var input = "LLR"
var steps = `AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`
var teststeps2 = "RL"
var testinput2 = `AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

type node struct {
	self string
	left string
	right string
}
func (n *node) Print() {
	println(n.self, n.left, n.right)
}

func (n *node) String() string {
	return n.self
}

func (n *node) Left() string {
	return n.left
}

func (n *node) Right() string {
	return n.right
}

func (n *node) SetNode(self, left, right string) {
	n.self = self
	n.left = left
	n.right = right
}

func (n *node) ParseInputToNode(input string ) {
	split := strings.Split(input, " = ")
	self := strings.Trim(split[0], " ")
	path := strings.Split(strings.Trim(split[1], " ()" ), ", ")

	n.SetNode(self, path[0], path[1])
}



func main() {
	rows := strings.Split(input, "\n")
	stepMap := make(map[string]*node)
	for i, row := range rows {
		fmt.Println(i)
		n := new(node)
		n.ParseInputToNode(row)
		n.Print()

		stepMap[n.self] = n
	}
	for _, step := range stepMap {
		fmt.Println(step.self, step.left, step.right)
	}
	fmt.Println(stepMap)

	count := 0
	position := "AAA"
	for {
		for _, step := range steps {
			if step == 'L' {
				position = stepMap[position].Left()
			} else if step == 'R' {
				position = stepMap[position].Right()
			}
			count++
			if position == "ZZZ" {
			 	fmt.Println(count)
				return
			}
		}
	}
}
