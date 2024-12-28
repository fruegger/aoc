package main

import (
	"advent/aoc/common"
	"advent/aoc/ds"
	"advent/aoc/pos"
	"advent/aoc/term"
	"fmt"
	"math"
)

func main() {
	lines := common.StartDay(16, "input")
	m := Maze{}
	initializeMaze(&m, lines)

	m.PrintMaze()

	result := m.Dijkstra(distanceFn)
	m.calculateBestPaths(&m.endNode)
	best := m.bestPath.Size()
	m.PrintMaze()

	//maze.debugMaze()
	fmt.Println("Part1:", result)
	fmt.Println("Part2:", best)
}

const START = 'S'
const END = 'E'
const WALL = '#'
const NOTHING = '.'

type Maze struct {
	field     []string
	start     Coord
	end       Coord
	endNode   Node
	bestPath  ds.Set[pos.Position]
	bestScore int
}

func (m *Maze) Initialize(lines []string, start pos.Position, end pos.Position) {
	m.field = lines
	m.start = Coord{p: start, d: pos.RIGHT}
	m.bestScore = math.MaxInt
	m.bestPath = ds.Set[pos.Position]{}
	m.end = Coord{p: end}
}

type Coord struct {
	p pos.Position
	d pos.Direction
}

type Node struct {
	c    Coord
	dist int
	//	best bool
	camefrom []*Node
}

var neighboursDeltas = []pos.Distance{
	{Dx: -1, Dy: 0},
	{Dx: 1, Dy: 0},
	{Dx: 0, Dy: -1},
	{Dx: 0, Dy: 1},
}

var unvisited = map[Coord]Node{}

func (m *Maze) Dijkstra(distance func(m *Maze, p1 Coord, p2 Coord) int) int {
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
				newDist := distance(m, next.c, c3)
				if newDist < math.MaxInt {
					newDist += next.dist
				}
				if found {
					if newDist <= n3.dist {
						n3.dist = newDist
						n3.camefrom = append(n3.camefrom, &next)
						unvisited[c3] = n3
					}
				}
			}
			for _, d := range neighboursDeltas {
				//do x,y
				p2 := next.c.p.Add(d)
				if p2.X >= 0 && p2.Y >= 0 && p2.Y < len(m.field) && p2.X < len(m.field[p2.Y]) && m.field[p2.Y][p2.X] != WALL {
					c2 := Coord{p: p2, d: next.c.d}

					n2, found := unvisited[c2]
					newDist := distance(m, next.c, c2)
					if newDist < math.MaxInt {
						newDist += next.dist
					}
					if found {
						if newDist <= n2.dist {
							n2.dist = newDist
							n2.camefrom = append(n2.camefrom, &next)
							unvisited[c2] = n2
						}
					}
				}
			}
			delete(unvisited, next.c)

			if next.c.p == m.end.p {
				if next.dist <= m.bestScore {
					m.bestScore = next.dist
					m.endNode.camefrom = append(m.endNode.camefrom, &next)
				}
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
					unvisited[c2] = Node{c: c2, dist: 0}
					c2 = Coord{p: m.start.p, d: pos.UP}
					unvisited[c2] = Node{c: c2, dist: math.MaxInt}
					c2 = Coord{p: m.start.p, d: pos.LEFT}
					unvisited[c2] = Node{c: c2, dist: math.MaxInt}
					c2 = Coord{p: m.start.p, d: pos.DOWN}
					unvisited[c2] = Node{c: c2, dist: math.MaxInt}

				} else {
					if x == m.end.p.X && y == m.end.p.Y {
						c := Coord{p: pos.Position{X: x, Y: y}, d: pos.UP}
						unvisited[c] = Node{c: c, dist: math.MaxInt}
						m.endNode = unvisited[c]
					} else {
						c := Coord{p: pos.Position{X: x, Y: y}, d: pos.UP}
						unvisited[c] = Node{c: c, dist: math.MaxInt}
						c = Coord{p: pos.Position{X: x, Y: y}, d: pos.LEFT}
						unvisited[c] = Node{c: c, dist: math.MaxInt}
						c = Coord{p: pos.Position{X: x, Y: y}, d: pos.RIGHT}
						unvisited[c] = Node{c: c, dist: math.MaxInt}
						c = Coord{p: pos.Position{X: x, Y: y}, d: pos.DOWN}
						unvisited[c] = Node{c: c, dist: math.MaxInt}
					}
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
		if node.dist < minS {
			*result = node
			minS = node.dist
		}
	}
	return minS < math.MaxInt
}

func (m *Maze) calculateBestPaths(w *Node) {
	// add all nodes reachable form the end back to start whose distance is min.
	minDist := math.MaxInt
	for _, n := range w.camefrom {
		if n.dist < minDist {
			minDist = n.dist
		}
	}
	for _, n := range w.camefrom {
		if n != nil { //&& n.dist <= minDist {
			m.bestPath.Add(n.c.p)
			m.calculateBestPaths(n)
		}
	}
}

func (m *Maze) PrintMaze() {
	for y, v := range m.field {
		for x, c := range v {
			if m.isOnBestPath(pos.Position{X: x, Y: y}) {
				if c == NOTHING {
					fmt.Print(term.YELLOW + "O" + term.WHITE)
				} else {
					fmt.Print(term.YELLOW + string(c) + term.WHITE)
				}

			} else {
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
		}
		fmt.Println()
	}
}

func (m *Maze) isOnBestPath(p pos.Position) bool {
	for _, n := range *m.bestPath.Elements() {
		if n == p {
			return true
		}
	}
	return false
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

func distanceFn(m *Maze, c1 Coord, c2 Coord) int {
	if c1.p.Move(c1.d) == c2.p || c2.p == m.end.p {
		return 1
	} else {
		if c1.p == c2.p {
			if c1.d == c2.d || c1.p == m.end.p {
				return 0
			} else {
				return 1000
			}
		} else {
			return math.MaxInt
		}
	}
}
