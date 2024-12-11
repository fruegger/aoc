package main

import (
	"advent/aoc/common"
	"advent/aoc/pos"
	"fmt"
)

func main() {
	lines := common.StartDay(10, "input")
	paths := findPaths(lines, false)
	total := countScores(paths)
	fmt.Println("Part 1: ", total)

	paths = findPaths(lines, true)
	total = countScores(paths)
	fmt.Println("Part 2: ", total)
}

func countScores(paths []Path) int {
	total := 0
	for _, p := range paths {
		total += p.score
	}
	return total
}

type Path struct {
	head     pos.Position
	current  pos.Position
	score    int
	traveled []pos.Direction
}

func findPaths(lines []string, distinct bool) []Path {
	var result []Path
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			visited := clearVisited(lines)
			if lines[y][x] == '0' {
				start := pos.Position{x, y}
				newPath := Path{
					head:     start,
					current:  start,
					traveled: []pos.Direction{},
				}
				gradientAscent(lines, &newPath, visited, distinct)
				result = append(result, newPath)
			}
		}
	}
	return result
}

func clearVisited(lines []string) [][]bool {
	var result [][]bool
	for y := 0; y < len(lines); y++ {
		var boolLine []bool
		for x := 0; x < len(lines[y]); x++ {
			boolLine = append(boolLine, false)
		}
		result = append(result, boolLine)
	}
	return result
}

func gradientAscent(lines []string, path *Path, visited [][]bool, distinct bool) {
	gradientAscentOneDirection(lines, path, visited, pos.UP, distinct)
	gradientAscentOneDirection(lines, path, visited, pos.RIGHT, distinct)
	gradientAscentOneDirection(lines, path, visited, pos.DOWN, distinct)
	gradientAscentOneDirection(lines, path, visited, pos.LEFT, distinct)
}

func gradientAscentOneDirection(lines []string, path *Path, visited [][]bool, d pos.Direction, distinct bool) {
	var p2 pos.Position
	canMove := false
	canMove, p2 = checkDirection(lines, visited, path.current, d)
	if canMove {
		save := path.current
		path.current = p2
		path.traveled = append(path.traveled, d)
		visited[p2.Y][p2.X] = !distinct
		if pos.AtPosition(lines, p2) == '9' {
			path.score++
		} else {
			gradientAscent(lines, path, visited, distinct)
		}
		path.current = save
	}
}

func checkDirection(lines []string, visited [][]bool, p pos.Position, d pos.Direction) (bool, pos.Position) {
	//all line have the same length; there is at least one line.
	max := pos.Position{X: len(lines[0]), Y: len(lines)}
	p2 := p.Move(d)
	if p2.X >= 0 && p2.X < max.X && p2.Y >= 0 && p2.Y < max.Y && !visited[p2.Y][p2.X] {
		grad := int8(pos.AtPosition(lines, p2)) - int8(pos.AtPosition(lines, p))
		return grad == 1, p2
	}
	return false, p2
}
