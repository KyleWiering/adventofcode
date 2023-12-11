package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

type game struct {
	Value int
	Green int
	Blue int
	Red int
}

func (g *game) parseGame(s string) {
	s = strings.ReplaceAll(s, "Game ", "")
	g.Value, _ = strconv.Atoi(strings.Split(s, ":")[0])
	Draws := strings.Split(s, ":")[1]
	rounds := strings.Split(Draws, ";")

	gameOver := false
	
	for _, v := range rounds {
		if gameOver {
			continue
		}
		round := strings.Split(v, ",")
		for _, z := range round {
			if gameOver {
				continue
			}
			splitter := strings.Split(strings.Trim(z, " "), " ")
			value, _ := strconv.Atoi(strings.Trim(splitter[0], " "))
			if splitter[1] == "red" && value > 12 {
				gameOver = true
				g.Value = 0
			} else if splitter[1] == "green" && value > 13 {
				gameOver = true
				g.Value = 0
			} else if splitter[1] == "blue" && value > 14 {
				gameOver = true
				g.Value = 0
			}
		}


	}
}


func (g *game) getGameValues(maxR, maxG, maxB int) int {
	return g.Value
}

func main() {
	finalCount := 0
	for _, v := range strings.Split(input, "\n") {
		var game game
		game.parseGame(v)
		finalCount += game.getGameValues(12, 13, 14)
	}
	fmt.Println(finalCount)
}
