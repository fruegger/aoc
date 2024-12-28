package main

import (
	"advent/aoc/common"
	"advent/aoc/maze"
	"advent/aoc/pos"
	"fmt"
	"maps"
)

func main() {
	lines := common.StartDay(20, "input")

	var m maze.Maze
	initializeMaze(&m, lines)
	m.PrintMaze()
	originalLen := m.Dijkstra(maze.UnityD, maze.IsNothing)
	m.CalculateBestPaths(&m.EndNode)

	fmt.Println("Without cheating: ", originalLen)
	savings := map[int]int{}

	cheats := findPossibleCheats(&m, 1)
	collectSavings(m, cheats, &savings)
	m.PrintMaze()

	printSavings(originalLen, savings)
	total := 0
	for i := 100; i < originalLen; i++ {
		total += savings[i]
	}
	fmt.Println("Part 1: ", total)

	m.PrintMaze()
	originalLen = m.Dijkstra(maze.UnityD, maze.IsNothing)
	m.CalculateBestPaths(&m.EndNode)
	savings = map[int]int{}

	cheats = findPossibleCheats(&m, 19)
	collectSavings(m, cheats, &savings)
	printSavings(originalLen, savings)
	for i := 100; i < originalLen; i++ {
		total += savings[i]
	}
	fmt.Println("Part 2: ", total)
}

func printSavings(originalLen int, savings map[int]int) {
	for i := 1; i < originalLen; i++ {
		if savings[i] != 0 {
			if savings[i] == 1 {
				fmt.Println("There is one cheat that saves ", i, " picoseconds.")
			} else {
				fmt.Println("There are ", savings[i], "cheats that save ", i, " picoseconds.")
			}
		}
	}
}

type Cheat struct {
	start pos.Position
	end   pos.Position
}

func findPossibleCheats(m *maze.Maze, maxCheatLen int) map[Cheat]int {
	// different approach
	var cheats = map[Cheat]int{}
	j := 0
	for i1 := 0; i1 < len(*m.BestPath.Elements()); i1++ {
		n1 := m.Nodes[(*m.BestPath.Elements())[i1]]
		fmt.Print(j, ":")
		j++
		i := 0
		for i2 := i1 + 1; i2 < len(*m.BestPath.Elements()); i2++ {
			n2 := m.Nodes[(*m.BestPath.Elements())[i2]]
			if i%1000 == 0 {
				fmt.Print(".")
			}
			i++
			if m.BestPath.Contains(n1.P) && m.BestPath.Contains(n2.P) {
				dist := m.DistanceFromStart(n1.P) - m.DistanceFromStart(n2.P)
				if dist < 0 {
					dist = -dist
				}
				// Manhattan
				bestD := n1.P.DistanceTo(n2.P)
				best := 0
				if bestD.Dy < 0 {
					best -= bestD.Dy
				} else {
					best += bestD.Dy
				}
				if bestD.Dx < 0 {
					best -= bestD.Dx
				} else {
					best += bestD.Dx
				}
				if best < dist && best <= maxCheatLen+1 && !findAnyCheat(cheats, n2.P, n1.P) {
					cheats[Cheat{start: n1.P, end: n2.P}] = dist - best
				}
			}
		}
		fmt.Println()
	}
	return cheats
}

func findAnyCheat(cheats map[Cheat]int, p1 pos.Position, p2 pos.Position) bool {
	_, found := cheats[Cheat{start: p1, end: p2}]
	return found
}

func collectSavings(m maze.Maze, cheats map[Cheat]int, savings *map[int]int) {
	for v := range maps.Keys(cheats) {
		if !m.IsOnBestPath(v.end) || !m.IsOnBestPath(v.start) {
			if !m.IsOnBestPath(v.end) {
				fmt.Println("end not on best: ", v.end)
			}
			if !m.IsOnBestPath(v.start) {
				fmt.Println("start not on best: ", v.start)
			}
		} else {
			// fmt.Println(v, " - ", v.saving)
			(*savings)[cheats[v]]++
		}
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
