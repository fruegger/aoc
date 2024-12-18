package main

import (
	"advent/aoc/common"
	"advent/aoc/term"
	"math"

	"advent/aoc/pos"
	"fmt"
)

func main() {
	lines := common.StartDay(16, "test2")
	m := Maze{}
	initializeMaze(&m, lines)

	m.PrintMaze()

	result := m.Traverse(distanceFn)
	//maze.debugMaze()
	fmt.Println("Part1:", result)
}

const START = 'S'
const END = 'E'
const WALL = '#'
const NOTHING = '.'

type Maze struct {
	field     []string
	start     Coord
	end       Coord
	bestScore int
}

func (m *Maze) Initialize(lines []string, start pos.Position, end pos.Position) {
	m.field = lines
	m.start = Coord{p: start, d: pos.RIGHT}
	m.bestScore = math.MaxInt
	m.end = Coord{p: end}
}

type Coord struct {
	p pos.Position
	d pos.Direction
}

type Node struct {
	c Coord
	s int
}

var neighboursDeltas = []pos.Distance{
	{Dx: -1, Dy: 0},
	{Dx: 1, Dy: 0},
	{Dx: 0, Dy: -1},
	{Dx: 0, Dy: 1},
}

var unvisited = map[Coord]Node{}

func (m *Maze) Traverse(distance func(p1 Coord, p2 Coord) int) int {
	m.initializeUnvisited()
	hasMore := true
	for hasMore {
		var next Node
		hasMore = selectNextNode(&next)
		if hasMore {
			// determine unvisited neighbours and their distances
			for _, d3 := range pos.Directions {
				//do directions
				c3 := Coord{p: next.c.p, d: d3}
				n3, found := unvisited[c3]
				newDist := distance(next.c, c3)
				if newDist < math.MaxInt {
					newDist += next.s
				}
				if found && newDist < n3.s {
					n3.s = newDist
					unvisited[c3] = n3
				}

			}
			for _, d := range neighboursDeltas {
				//do x,y
				p2 := next.c.p.Add(d)
				if p2.X >= 0 && p2.Y >= 0 && p2.Y < len(m.field) && p2.X < len(m.field[p2.Y]) && m.field[p2.Y][p2.X] != WALL {
					c2 := Coord{p: p2, d: next.c.d}

					n2, found := unvisited[c2]
					newDist := distance(next.c, c2)
					if newDist < math.MaxInt {
						newDist += next.s
					}
					if found && newDist < n2.s {
						n2.s = newDist
						unvisited[c2] = n2
					}
				}
			}
			delete(unvisited, next.c)

			for _, node := range unvisited {
				if node.s < math.MaxInt {
					fmt.Println(node.c, "-", node.s)
				}
			}
			fmt.Println("------")

			if next.c.p.X == m.end.p.X && next.c.p.Y == m.end.p.Y {
				m.bestScore = next.s
				hasMore = false
			}
		}
	}
	return m.bestScore
}

func (m *Maze) initializeUnvisited() {
	for y := 0; y < len(m.field); y++ {
		for x := 0; x < len(m.field[y]); x++ {
			if m.field[y][x] != WALL {
				if x == m.start.p.X && y == m.start.p.Y {
					c2 := m.start
					unvisited[c2] = Node{c: c2, s: 0}
					c2 = Coord{p: m.start.p, d: pos.UP}
					unvisited[c2] = Node{c: c2, s: math.MaxInt}
					c2 = Coord{p: m.start.p, d: pos.LEFT}
					unvisited[c2] = Node{c: c2, s: math.MaxInt}
					c2 = Coord{p: m.start.p, d: pos.DOWN}
					unvisited[c2] = Node{c: c2, s: math.MaxInt}

				} else {
					c := Coord{p: pos.Position{X: x, Y: y}, d: pos.UP}
					unvisited[c] = Node{c: c, s: math.MaxInt}
					c = Coord{p: pos.Position{X: x, Y: y}, d: pos.LEFT}
					unvisited[c] = Node{c: c, s: math.MaxInt}
					c = Coord{p: pos.Position{X: x, Y: y}, d: pos.RIGHT}
					unvisited[c] = Node{c: c, s: math.MaxInt}
					c = Coord{p: pos.Position{X: x, Y: y}, d: pos.DOWN}
					unvisited[c] = Node{c: c, s: math.MaxInt}
				}
			}
		}
	}
}

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

func initializeMaze(m *Maze, lines []string) bool {
	start, sok := common.FindSymbol(lines, START)
	end, eok := common.FindSymbol(lines, END)
	if sok && eok {
		m.Initialize(lines, start, end)
	} else {
		fmt.Println("error, no S or E found")
	}
	return sok && eok
}

func distanceFn(c1 Coord, c2 Coord) int {
	if c1.p.Move(c1.d) == c2.p {
		return 1
	} else {
		if c1.p.X == c2.p.X && c1.p.Y == c2.p.Y {
			return 1000
		} else {
			return math.MaxInt
		}
	}
}

/*
func turnCounterclockwise(d pos.Direction) pos.Direction {
	var result pos.Direction
	switch d {
	case pos.RIGHT:
		result = pos.UP
	case pos.UP:
		result = pos.LEFT
	case pos.LEFT:
		result = pos.DOWN
	case pos.DOWN:
		result = pos.RIGHT
	}
	return result
}

func turnClockwise(d pos.Direction) pos.Direction {
	var result pos.Direction
	switch d {
	case pos.RIGHT:
		result = pos.DOWN
	case pos.DOWN:
		result = pos.LEFT
	case pos.LEFT:
		result = pos.UP
	case pos.UP:
		result = pos.RIGHT
	}
	return result
}
*/
