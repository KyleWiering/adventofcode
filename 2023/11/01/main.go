package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var input = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

type Position struct {
	Value string
}

type Galaxy struct {
	X, Y int
}

func distanceFormula(x1, y1, x2, y2 float64) float64 {
	return math.Abs(math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

func distance(grid [][]Position, galaxy1, galaxy2 Galaxy) int {
	spare := make([][]Position, len(grid))
	for i := range spare {
		spare[i] = make([]Position, len(grid[i]))
		copy(spare[i], grid[i])
	}

	max := distanceFormula(float64(galaxy1.X), float64(galaxy1.Y), float64(galaxy2.X), float64(galaxy2.Y))
	depth := findRoute(spare, galaxy1, galaxy2, 0, max)

	val, _ :=strconv.Atoi(depth[galaxy2.Y][galaxy2.X].Value)
	fmt.Println(galaxy1.X, galaxy1.Y, "-->", galaxy2.X, galaxy2.Y, "==", val)
	//for _, row := range depth {
	//	for _, col := range row {
	//		if len(col.Value) == 1 {
	//			fmt.Print("0"+col.Value + " ")
	//		} else {
	//			fmt.Print(col.Value + " ")
	//		}
	//	}
	//	fmt.Println()
	//}
	//fmt.Println()
	return val
}

func findRoute(grid [][]Position, g1, g2 Galaxy, depth int, max float64) [][]Position {

	if g1.X == g2.X && g1.Y == g2.Y {
		if weight, err := strconv.Atoi(grid[g1.Y][g1.X].Value); err != nil {
			grid[g1.Y][g1.X].Value = strconv.Itoa(depth)
		} else if depth <= weight {
			grid[g1.Y][g1.X].Value = strconv.Itoa(depth)
		}
		return grid
	}
	if grid[g2.Y][g2.X].Value != "#" && grid[g2.Y][g2.X].Value != "." {
		weight, _ := strconv.Atoi(grid[g2.Y][g2.X].Value)
		if depth >= weight {
			return grid
		}
	}
	dist := distanceFormula(float64(g1.X), float64(g1.Y), float64(g2.X), float64(g2.Y))
	if g1.X < 0 || g1.Y < 0 || g1.X >= len(grid[0]) || g1.Y >= len(grid) || depth > len(grid) *2 || dist > max {
		return grid
	}
	if grid[g1.Y][g1.X].Value == "#" || grid[g1.Y][g1.X].Value == "." {
		grid[g1.Y][g1.X].Value = strconv.Itoa(depth)
	} else {
		weight, _ := strconv.Atoi(grid[g1.Y][g1.X].Value)
		if depth < weight {
			grid[g1.Y][g1.X].Value = strconv.Itoa(depth)
		} else {
			return grid
		}
	}

	if distanceFormula(float64(g1.X), float64(g1.Y-1), float64(g2.X), float64(g2.Y)) < dist {
		grid = findRoute(grid, Galaxy{g1.X, g1.Y - 1}, g2, depth+1, max)
	}
	
	if distanceFormula(float64(g1.X), float64(g1.Y+1), float64(g2.X), float64(g2.Y)) < dist {
		grid = findRoute(grid, Galaxy{g1.X, g1.Y + 1}, g2, depth+1, max)	
	}
	if distanceFormula(float64(g1.X-1), float64(g1.Y), float64(g2.X), float64(g2.Y)) < dist {
		grid = findRoute(grid, Galaxy{g1.X - 1, g1.Y}, g2, depth+1, max)	
	}
	if distanceFormula(float64(g1.X+1), float64(g1.Y), float64(g2.X), float64(g2.Y)) < dist {
		grid = findRoute(grid, Galaxy{g1.X + 1, g1.Y}, g2, depth+1, max)	
	}
	

	return grid
}

type Map struct {
	Pos [][]Position
	EmptyRows []bool
	EmptyCols []bool
	Galaxies []Galaxy
}

func (m *Map) Print() {
	for _, row := range m.Pos {
		for _, col := range row {
			print(col.Value)
		}
		println()
	}

	fmt.Println(m.Galaxies)
}

func (m *Map) Populate(input string) {
	lines := strings.Split(input, "\n")
	m.Pos = make([][]Position, len(lines))
	m.EmptyRows = make([]bool, len(lines))
	for i := range m.EmptyRows {
		m.EmptyRows[i] = true
	}
	
	for y, line := range lines {
		val := strings.Split(line, "")
		if len(m.EmptyCols) < len(val) {
			m.EmptyCols = make([]bool, len(val))
			for i := range m.EmptyCols {
				m.EmptyCols[i] = true
			}
		}
			m.Pos[y] = make([]Position, len(val))
		
		for x, v := range val {
			if v != "." {
				m.EmptyRows[y] = false
				m.EmptyCols[x] = false
			}

			m.Pos[y][x].Value = v
		}
	}
	fmt.Println(m.EmptyRows)
	fmt.Println(m.EmptyCols)
	m.ExpandGalaxy()
}

func (m *Map) ExpandGalaxy() {
	m.ExpandRows()
	m.ExpandColumn()
}

func (m *Map) ExpandRows() {
	for i := len(m.EmptyRows) - 1; i >= 0; i-- {
		if m.EmptyRows[i] {
			m.Pos = append(m.Pos[:i+1], m.Pos[i:]...)
			m.Pos[i] = make([]Position, len(m.Pos[i-1]))
			for j := range m.Pos[i] {
				m.Pos[i][j].Value = "."
			}
		}
	}
}

func (m *Map) ExpandColumn() {
	for i := len(m.EmptyCols) - 1; i >= 0; i-- {
		if m.EmptyCols[i] {
			for j := range m.Pos {
				m.Pos[j] = append(m.Pos[j][:i+1], m.Pos[j][i:]...)
				m.Pos[j][i].Value = "."
			}
		}
	}
}

func main() {
	universe := Map{}
	universe.Populate(input)
	universe.Print()

	universe.Chart()

}

func (m *Map) Chart() {
	loc := 1
	for y, row := range m.Pos {
		for x, col := range row {
			if col.Value == "#" {
				//m.Pos[y][x].Value = strconv.Itoa(loc)
				loc++
				m.Galaxies = append(m.Galaxies, Galaxy{x, y})
			}
		}
	}

	m.Print()
	total := 0
	for i := 0; i< len(m.Galaxies); i++ {
		for j:= i; j < len(m.Galaxies); j++ {
			if m.Galaxies[i].X == m.Galaxies[j].X && m.Galaxies[i].Y == m.Galaxies[j].Y {
			 //	fmt.Printf("Galaxy %d and %d are the same\n", i, j)
				continue
			} else {
				// fmt.Printf("Galaxy %d and %d are not the same\n", i, j)
				fmt.Print(i+1, "->", j+1, " ")
				total += distance(m.Pos, m.Galaxies[i], m.Galaxies[j])
				fmt.Println()
			}
		}
	}
	fmt.Println(total)
}
