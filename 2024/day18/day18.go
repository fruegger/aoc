package main

import (
	"advent/aoc/common"
	"advent/aoc/maze"
	"advent/aoc/pos"
	"fmt"
	"math"
	"strings"
)

func main() {

	const SIZEX = 71
	const SIZEY = 71

	const BYTESCOMING = 1024

	lines := common.StartDay(18, "input")
	m := maze.Maze{}
	initializeMaze(&m, lines, SIZEX, SIZEY, BYTESCOMING)
	m.PrintMaze()

	result := m.Traverse(distanceFn)
	fmt.Println("Part1 :", result)
	start := BYTESCOMING + 1
	end := len(lines) - 1
	for start != end {
		walk := (start + end) / 2
		initializeMaze(&m, lines, SIZEX, SIZEY, walk)
		result = m.Traverse(distanceFn)
		if result == math.MaxInt {
			end = walk - 1
		} else {
			start = walk + 1
		}
	}
	m.PrintMaze()
	fmt.Println("Part2 :", lines[start])

}

func distanceFn(p1 pos.Position, p2 pos.Position) int {
	return 1
	/*
		dist :=p1.DistanceTo(p2)
		if dist.Dx<0 { dist.Dx = - dist.Dx }
		if dist.Dy<0 { dist.Dy = - dist.Dy }
		return dist.Dx + dist.Dy

	*/
}

func initializeMaze(m *maze.Maze, lines []string, maxx int, maxy int, bytes int) {
	line := ""
	var field []string
	for x := 0; x < maxx; x++ {
		line = line + string(maze.NOTHING)
	}
	for y := 0; y < maxy; y++ {
		field = append(field, line)
	}

	for i := 0; i < len(lines) && i < bytes; i++ {
		parts := strings.Split(lines[i], ",")
		px := common.StringToNum(parts[0])
		py := common.StringToNum(parts[1])
		common.ChangeSymbol(&field, pos.Position{X: px, Y: py}, maze.WALL)
	}
	m.Initialize(field, pos.Position{X: 0, Y: 0}, pos.Position{X: maxx - 1, Y: maxy - 1})
}
