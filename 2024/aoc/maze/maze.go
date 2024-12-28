package maze

import (
	"advent/aoc/common"
	"advent/aoc/ds"
	"advent/aoc/pos"
	"advent/aoc/term"
	"fmt"
	"math"
)

const START = 'S'
const END = 'E'
const WALL = '#'
const NOTHING = '.'
const OBSTACLE = 'O'

type Maze struct {
	field     []string
	Start     pos.Position
	End       pos.Position
	BestScore int
	//28.12
	unvisited map[pos.Position]*Node
	Nodes     map[pos.Position]*Node
	//	distances           map[pos.Position]*Node
	EndNode             Node
	maxDistance         int
	stopAfterFirstFound bool
	BestPath            ds.Set[pos.Position]
}

func (m *Maze) Initialize(lines []string, start pos.Position, end pos.Position) {
	m.field = lines
	m.Start = start
	m.BestScore = math.MaxInt
	m.maxDistance = math.MaxInt
	m.stopAfterFirstFound = false
	m.End = end
	m.unvisited = map[pos.Position]*Node{}
	m.Nodes = map[pos.Position]*Node{}

	m.BestPath = ds.Set[pos.Position]{}
}

type Node struct {
	P    pos.Position
	Dist int
	//28.12
	cameFrom []*Node
}

var neighboursDeltas = []pos.Distance{
	{Dx: -1, Dy: 0},
	{Dx: 1, Dy: 0},
	{Dx: 0, Dy: -1},
	{Dx: 0, Dy: 1},
}

func (m *Maze) Dijkstra(
	distance func(p1 pos.Position, p2 pos.Position) int,
	pathSymbol func(sym uint8) bool,
) int {
	m.initializeUnvisited(pathSymbol)
	//	m.paths = []*Node{}
	hasMore := true
	for hasMore {
		var next Node
		hasMore = m.selectNextNode(&next)
		if hasMore {
			// determine unvisited neighbours and their distances
			for _, d := range neighboursDeltas {
				p2 := next.P.Add(d)
				if p2.X > 0 && p2.Y > 0 && p2.Y+1 < len(m.field) && p2.X+1 < len(m.field[p2.Y]) && pathSymbol(m.field[p2.Y][p2.X]) {
					n2, found := m.unvisited[p2]
					if found {
						newDist := distance(next.P, p2) + next.Dist
						if newDist <= n2.Dist {
							if newDist > m.maxDistance {
								return m.maxDistance
							}
							n2.Dist = newDist
							n2.cameFrom = append(n2.cameFrom, &next)
							m.unvisited[n2.P] = n2
						}
					}
				}
			}
			delete(m.unvisited, next.P)
			if next.P == m.End {
				m.BestScore = next.Dist
				m.EndNode.cameFrom = append(m.EndNode.cameFrom, &next)

				if m.stopAfterFirstFound {
					return m.BestScore
				}
			}
		}
	}
	return m.BestScore
}

func (m *Maze) SetMaxDistance(distance int)           { m.maxDistance = distance }
func (m *Maze) SetStopAfterFirstFound(firstOnly bool) { m.stopAfterFirstFound = firstOnly }

func (m *Maze) CalculateBestPaths(w *Node) {
	// add all nodes reachable form the end back to start whose distance is min.
	minDist := math.MaxInt
	for _, n := range w.cameFrom {
		if n.Dist < minDist {
			minDist = n.Dist
		}
	}
	for _, n := range w.cameFrom {
		if n != nil { //&& n.dist <= minDist {
			m.BestPath.Add(n.P)
			m.CalculateBestPaths(n)
		}
	}
}

func IsNotWall(sym uint8) bool {
	return sym != WALL
}

func IsNothing(sym uint8) bool {
	return sym == NOTHING || sym == END || sym == START
}

func UnityD(p pos.Position, p2 pos.Position) int {
	if p == p2 {
		return 0
	} else {
		return 1
	}
}

func (m *Maze) ChangeSymbol(p pos.Position, sym uint8) {
	common.ChangeSymbol(&m.field, p, sym)
}

func (m *Maze) initializeUnvisited(pathSymbol func(sym uint8) bool) {
	startNode := Node{P: m.Start, Dist: 0}
	m.unvisited[m.Start] = &startNode
	m.Nodes[m.Start] = &startNode
	//m.paths = append(m.paths, &startNode)
	for y := 0; y < len(m.field); y++ {
		for x := 0; x < len(m.field[y]); x++ {
			if pathSymbol(m.field[y][x]) {
				if x != m.Start.X || y != m.Start.Y {
					p := pos.Position{X: x, Y: y}
					m.unvisited[p] = &Node{P: p, Dist: math.MaxInt}
					m.Nodes[p] = m.unvisited[p]
				}
			}
		}
	}
	m.EndNode = *m.unvisited[m.End]
}

func (m *Maze) selectNextNode(result *Node) bool {
	minS := math.MaxInt
	if len(m.unvisited) == 0 {
		return false
	}
	for p := range m.unvisited {
		node := m.unvisited[p]
		if node.Dist < minS {
			*result = *node
			minS = node.Dist
		}
	}
	return minS < math.MaxInt
}

func (m *Maze) IsOnBestPath(p pos.Position) bool {
	for _, n := range *m.BestPath.Elements() {
		if n == p {
			return true
		}
	}
	return false
}

func (m *Maze) DistanceFromStart(p pos.Position) int {
	return m.Nodes[p].Dist
}

func (m *Maze) PrintMaze() {
	for y, v := range m.field {
		for x, c := range v {
			if m.IsOnBestPath(pos.Position{X: x, Y: y}) {
				fmt.Print(term.YELLOW + string(c) + term.WHITE)
			} else {
				switch c {
				case NOTHING:
					fmt.Print(string(c))
				case WALL:
					fmt.Print(term.BLUE + string(c) + term.WHITE)
				case OBSTACLE:
					fmt.Print(term.YELLOW + string(c) + term.WHITE)
				case 'A':
					fmt.Print(term.YELLOW + string(c) + term.WHITE)
				case 'B':
					fmt.Print(term.YELLOW + string(c) + term.WHITE)
				default:
					fmt.Print(term.GREEN + string(c) + term.WHITE)
				}
			}
		}
		fmt.Println()
	}
}
