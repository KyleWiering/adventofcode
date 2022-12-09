package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type pos struct {
	X    int
	Y    int
	next *pos
}

func (p *pos) Print() {
	if p.next != nil {
		fmt.Print("(", p.X, ",", p.Y, "),")
		p.next.Print()
	} else {
		fmt.Print("(", p.X, ",", p.Y, ")")
		fmt.Println()
	}
}

func (p *pos) Build(count int) *pos {
	if count < 1 {
		return nil
	}
	next := pos{
		X:    0,
		Y:    0,
		next: nil,
	}

	p.next = next.Build(count - 1)
	return p
}

func (p *pos) Move(direction string, count int) {
	switch direction {
	case "U":
		p.Y += count
	case "D":
		p.Y -= count
	case "L":
		p.X -= count
	case "R":
		p.X += count
	}

	if p.next != nil {
		p.next.Follow(*p)
	}
}

func (p *pos) Follow(prev pos) {
	xDist := math.Abs(float64(p.X - prev.X))
	yDist := math.Abs(float64(p.Y - prev.Y))

	if xDist > 1 && yDist > 1 {

	} else if xDist > 1 && yDist > 0 {
		p.Y = prev.Y
	} else if yDist > 1 && xDist > 0 {
		p.X = prev.X
	}
	if xDist > 1 {
		p.ChaseX(prev)
	}
	if yDist > 1 {
		p.ChaseY(prev)
	}

	if p.next != nil {
		p.next.Follow(*p)
	}

	if p.next == nil {
		Visit(*p)
	}
}

func (p *pos) ChaseX(head pos) {
	if p.X > head.X {
		p.X--
	}

	if p.X < head.X {
		p.X++
	}
}

func (p *pos) ChaseY(head pos) {
	if p.Y > head.Y {
		p.Y--
	}
	if p.Y < head.Y {
		p.Y++
	}
}

var visited []pos
var rope pos
var isDummyPayload = true

const isDebug = false
const dummyPayload = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
const dummyPayload2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
const fullPayload = ``

func main() {
	var input string

	if isDummyPayload {
		input = dummyPayload
	} else {
		input = fullPayload
	}

	moves := strings.Split(input, "\n")

	visited = []pos{}
	rope = pos{0, 0, nil}
	rope.Build(2)
	fmt.Println("First operation each row.")
	FirstCase(moves)
	fmt.Println("Count:", len(visited))

	if isDummyPayload {
		input = dummyPayload2
	} else {
		input = fullPayload
	}
	moves = strings.Split(input, "\n")
	visited = []pos{}
	rope = pos{0, 0, nil}
	rope.Build(10)
	fmt.Println("Second operation each row.")
	SecondCase(moves)

	fmt.Println("Count:", len(visited))
}

func FirstCase(moves []string) {
	for _, move := range moves {
		movement := strings.Split(move, " ")
		count, _ := strconv.Atoi(movement[1])
		Move(movement[0], count)
	}
}

func SecondCase(moves []string) {
	for _, move := range moves {
		movement := strings.Split(move, " ")

		count, _ := strconv.Atoi(movement[1])
		Move(movement[0], count)
		if isDebug {
			rope.Print()
		}
	}
}

func Move(direction string, count int) {
	for i := 0; i < count; i++ {
		rope.Move(direction, 1)
	}
}

func Visit(tail pos) {
	found := false
	for _, place := range visited {
		if place.X == tail.X && place.Y == tail.Y {
			found = true
		}
	}
	if !found {
		visited = append(visited, tail)
	}
}
