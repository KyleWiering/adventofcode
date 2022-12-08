package main

import (
	"fmt"
	"strconv"
	"strings"
)

type grid struct {
	forest  [][]int
	visible [][]int
}

func (g *grid) Populate(fullInput string) {
	rows := strings.Split(fullInput, "\n")
	g.forest = [][]int{}
	g.visible = [][]int{}
	for i, row := range rows {
		var newRow []int
		var newVisible []int
		fmt.Println("row", row)
		for j := 0; j < len(row); j++ {

			item, _ := strconv.Atoi(row[j : j+1])

			newRow = append(newRow, item)
			if j == 0 || j == len(row)-1 || i == 0 || i == len(rows)-1 {
				newVisible = append(newVisible, 1)
			} else {
				newVisible = append(newVisible, 0)
			}

		}
		g.visible = append(g.visible, newVisible)
		g.forest = append(g.forest, newRow)
		if isDebug {
			fmt.Println(g.forest)
			fmt.Println(g.visible)
		}
	}
}

func (g *grid) Visibility() {
	for row := 1; row < len(g.forest)-1; row++ {
		for col := 1; col < len(g.forest[row])-1; col++ {
			if isDebug {
				fmt.Println("row:", row, "col:", col, " val: ", g.forest[row][col])
			}
			if g.isVisible(row, col, len(g.forest), len(g.forest[row])) {
				g.visible[row][col] = 1
				if isDebug {
					fmt.Println(g.visible)
				}
			}
		}
	}
}

func (g *grid) isVisible(x int, y int, length int, width int) bool {
	visible := true
	fmt.Println("down")
	for i := x + 1; i < length; i++ {
		fmt.Println(g.forest[i][y], ">=", g.forest[x][y])
		if g.forest[i][y] >= g.forest[x][y] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	fmt.Println("up")
	for i := x - 1; i >= 0; i-- {
		fmt.Println(g.forest[i][y], ">=", g.forest[x][y])
		if g.forest[i][y] >= g.forest[x][y] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	fmt.Println("right")
	for i := y + 1; i < width; i++ {
		fmt.Println(g.forest[x][i], ">=", g.forest[x][y])
		if g.forest[x][i] >= g.forest[x][y] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	fmt.Println("left")
	for i := y - 1; i >= 0; i-- {
		fmt.Println(g.forest[x][i], ">=", g.forest[x][y])
		if g.forest[x][i] >= g.forest[x][y] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	return false
}

func (g *grid) Total() {
	total := 0
	for i := 0; i < len(g.visible); i++ {
		for j := 0; j < len(g.visible[i]); j++ {
			total = total + g.visible[i][j]
		}
	}
	fmt.Println(total)
}

const isDebug = false
const dummyPayload = `30373
25512
65332
33549
35390`
const fullPayload = ``

func load1() string {
	return dummyPayload
}
func load2() string {
	return fullPayload
}

func main() {
	var input string
	isDummyPayload := false
	Grid := grid{}

	if isDummyPayload {
		input = load1()
	} else {
		input = load2()
	}

	Grid.Populate(input)

	//objects := strings.Split(input, "\n")

	fmt.Println("First operation each row.")
	FirstCase(Grid)

	fmt.Println("Second operation each row.")
	// SecondCase(objects)
}

func FirstCase(Grid grid) {
	Grid.Visibility()
	Grid.Total()
}
