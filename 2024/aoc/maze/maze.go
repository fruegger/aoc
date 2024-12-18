package maze

import (
	"advent/aoc/pos"
	"advent/aoc/term"
	"fmt"
	"math"
)

const START = 'S'
const END = 'E'
const WALL = '#'
const NOTHING = '.'

type Maze struct {
	field     []string
	start     pos.Position
	end       pos.Position
	bestScore int
}

func (m *Maze) Initialize(lines []string, start pos.Position, end pos.Position) {
	m.field = lines
	m.start = start
	m.bestScore = math.MaxInt
	m.end = end
}

type Node struct {
	p pos.Position
	s int
}

var neighboursDeltas = []pos.Distance{
	{Dx: -1, Dy: 0},
	{Dx: 1, Dy: 0},
	{Dx: 0, Dy: -1},
	{Dx: 0, Dy: 1},
}

var unvisited = map[pos.Position]Node{}

func selectNextNode(result *Node) bool {
	minS := math.MaxInt
	if len(unvisited) == 0 {
		return false
	}
	for p := range unvisited {
		node := unvisited[p]
		if node.s < minS {
			*result = node
			minS = node.s
		}
	}
	return minS < math.MaxInt
}

func (m *Maze) Traverse(distance func(p1 pos.Position, p2 pos.Position) int) int {
	for y := 0; y < len(m.field); y++ {
		for x := 0; x < len(m.field[y]); x++ {
			if m.field[y][x] != WALL {
				if x == m.start.X && y == m.start.Y {
					unvisited[m.start] = Node{p: m.start, s: 0}
				} else {
					p := pos.Position{X: x, Y: y}
					unvisited[p] = Node{p: p, s: math.MaxInt}
				}
			}
		}
	}

	hasMore := true
	for hasMore {
		var next Node
		hasMore = selectNextNode(&next)
		if hasMore {
			// determine unvisited neighbours and their distances
			for _, d := range neighboursDeltas {
				p2 := next.p.Add(d)
				if p2.X >= 0 && p2.Y >= 0 && p2.Y < len(m.field) && p2.X < len(m.field[p2.Y]) && m.field[p2.Y][p2.X] != WALL {
					n2, found := unvisited[p2]
					newDist := distance(next.p, p2) + next.s
					if found {
						if newDist < n2.s {
							n2.s = newDist
							unvisited[n2.p] = n2
						}
					}
				}
			}
			delete(unvisited, next.p)
			/*
				for _, node := range unvisited {
					if node.s < math.MaxInt {
						fmt.Println(node.p, "-", node.s)
					}
				}
				fmt.Println("------")
			*/
			if next.p.X == m.end.X && next.p.Y == m.end.Y {
				m.bestScore = next.s
			}
		}
	}
	return m.bestScore
}

func (m *Maze) PrintMaze() {
	for _, v := range m.field {
		for _, c := range v {
			if c == NOTHING {
				fmt.Print(string(c))
			} else {
				if c == WALL {
					fmt.Print(term.BLUE + string(c) + term.WHITE)
				} else {
					fmt.Print(term.GREEN + string(c) + term.WHITE)

				}
			}
		}
		fmt.Println()
	}
}
