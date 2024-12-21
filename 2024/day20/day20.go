package main

import (
	"advent/aoc/common"
	"advent/aoc/maze"
	"advent/aoc/pos"
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := common.StartDay(20, "test")

	var m maze.Maze
	initializeMaze(&m, lines)
	originalLen := m.Traverse(maze.UnityD, maze.IsNothing)
	m.PrintMaze()
	fmt.Println("Without cheating: ", originalLen)
	savings := map[int]int{}

	findPossibleCheatsPart1(&m, lines)
	collectSavings(m, cheatsPart1, &savings)
	m.PrintMaze()

	printSavings(originalLen, savings)
	total := 0
	for i := 101; i < originalLen; i++ {
		total += savings[i]
	}
	fmt.Println("Part 1: ", total)

	savings = map[int]int{}

	findPossibleCheatsPart2(&m, lines)
	collectSavings(m, cheatsPart2, &savings)
	printSavings(originalLen, savings)

}

func printSavings(originalLen int, savings map[int]int) {
	for i := 1; i < originalLen; i++ {
		if savings[i] != 0 {
			if savings[i] == 1 {
				fmt.Println("There is one cheats that saves ", i, " picoseconds.")
			} else {
				fmt.Println("There are ", savings[i], "cheats that save ", i, " picoseconds.")
			}
		}
	}
}

type Cheat struct {
	start  pos.Position
	end    pos.Position
	dir    pos.Direction
	saving int
}

var cheatsPart1 []Cheat

func findPossibleCheatsPart1(m *maze.Maze, lines []string) {
	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			p := pos.Position{X: x, Y: y}
			if lines[p.Y][p.X] == maze.NOTHING || lines[p.Y][p.X] == maze.START {
				var p2 pos.Position
				for _, d := range pos.Directions {
					if canCheatAt(lines, p, d, &p2) && !findCheat(cheatsPart1, p2, d.Opposite()) {
						dist := m.DistanceFromStart(p) - m.DistanceFromStart(p2)
						if dist < 0 {
							dist = -dist
						}
						cheatsPart1 = append(cheatsPart1, Cheat{start: p, dir: d, end: p2, saving: dist - 2})
					}
				}
			}
		}
	}
}

var cheatsPart2 []Cheat

var buf = bufio.NewReader(os.Stdin)

const CHEAT_START = 'A'
const CHEAT_END = 'B'

func findPossibleCheatsPart2(m *maze.Maze, lines []string) {
	// different approach

	r := maze.Maze{}
	for _, n1 := range m.ShortestPath() {
		for _, n2 := range m.ShortestPath() {

			dist := m.DistanceFromStart(n1.P) - m.DistanceFromStart(n2.P)
			if dist < 0 {
				dist = -dist
			}
			if dist > 1 {
				/* debug */
				/*				n1.P.X = 13
								n1.P.Y = 3
								n2.P.X = 12
								n2.P.Y = 5
				*/
				r.Initialize(lines, n1.P, n2.P)
				wasP1 := lines[n1.P.Y][n1.P.X]
				wasP2 := lines[n2.P.Y][n2.P.X]
				r.ChangeSymbol(n1.P, CHEAT_START)
				r.ChangeSymbol(n2.P, CHEAT_END)
				r.SetMaxDistance(21)
				r.SetStopAfterFirstFound(true)
				best := r.Traverse(maze.UnityD, IsWallOrCheat)

				fmt.Print("Saving: ", dist-best)
				if best < dist && best <= 20 && !findAnyCheat(cheatsPart2, n2.P, n1.P) {
					fmt.Println(" accept")
					r.PrintMaze()

					_, err := buf.ReadBytes('\n')
					if err != nil {
					}

				} else {
					fmt.Println(" reject")
				}

				if best < dist && best <= 20 && !findAnyCheat(cheatsPart2, n2.P, n1.P) {
					cheatsPart2 = append(cheatsPart2, Cheat{start: n1.P, end: n2.P, saving: dist - best})
				}
				r.ChangeSymbol(n1.P, wasP1)
				r.ChangeSymbol(n2.P, wasP2)

			}
		}
	}
}

func IsWallOrCheat(sym uint8) bool {
	return sym == maze.WALL || sym == CHEAT_START || sym == CHEAT_END
}

func findCheat(cheats []Cheat, p pos.Position, d pos.Direction) bool {
	for _, cheat := range cheats {
		if cheat.start == p && cheat.dir == d {
			return true
		}
	}
	return false
}

func findAnyCheat(cheats []Cheat, p1 pos.Position, p2 pos.Position) bool {
	for _, cheat := range cheats {
		if cheat.start == p1 && cheat.end == p2 {
			return true
		}
	}
	return false
}

func collectSavings(m maze.Maze, cheats []Cheat, savings *map[int]int) {
	for _, v := range cheats {
		if !m.IsOnShortestPath(v.end) || !m.IsOnShortestPath(v.start) {
			fmt.Println("not on best")
		} else {
			fmt.Println(v, " - ", v.saving)
			(*savings)[v.saving]++
		}
	}
}

func canCheatAt(lines []string, p pos.Position, d pos.Direction, end *pos.Position) bool {
	p2 := p.Move(d)
	*end = p2.Move(d)
	if end.Y > 0 && end.Y < len(lines) && end.X > 0 && end.X < len(lines[end.Y]) {
		return lines[p2.Y][p2.X] == maze.WALL && (lines[end.Y][end.X] == maze.NOTHING || lines[end.Y][end.X] == maze.END)
	} else {
		return false
	}
}

func initializeMaze(m *maze.Maze, lines []string) bool {
	start, sok := common.FindSymbol(lines, maze.START)
	end, eok := common.FindSymbol(lines, maze.END)
	if sok && eok {
		m.Initialize(lines, start, end)
	} else {
		fmt.Println("error, no S or E found")
	}
	return sok && eok
}
