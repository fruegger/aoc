package maze

import (
	"advent/aoc/common"
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
	field               []string
	start               pos.Position
	end                 pos.Position
	bestScore           int
	unvisited           map[pos.Position]Node
	distances           map[pos.Position]Node
	maxDistance         int
	stopAfterFirstFound bool
	paths               []*Node
}

func (m *Maze) Initialize(lines []string, start pos.Position, end pos.Position) {
	m.field = lines
	m.start = start
	m.bestScore = math.MaxInt
	m.maxDistance = math.MaxInt
	m.stopAfterFirstFound = false
	m.end = end
	m.unvisited = map[pos.Position]Node{}
	m.distances = map[pos.Position]Node{}
}

type Node struct {
	P        pos.Position
	dist     int
	cameFrom *Node
}

var neighboursDeltas = []pos.Distance{
	{Dx: -1, Dy: 0},
	{Dx: 1, Dy: 0},
	{Dx: 0, Dy: -1},
	{Dx: 0, Dy: 1},
}

func (m *Maze) Traverse(
	distance func(p1 pos.Position, p2 pos.Position) int,
	pathSymbol func(sym uint8) bool,
) int {
	m.initializeUnvisited(pathSymbol)
	m.paths = []*Node{}
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
						newDist := distance(next.P, p2) + next.dist
						if newDist < n2.dist {
							if newDist >= m.maxDistance {
								return m.maxDistance
							}
							n2.dist = newDist
							n2.cameFrom = &next
							m.unvisited[n2.P] = n2
						}
						m.distances[n2.P] = n2
					}
				}
			}
			delete(m.unvisited, next.P)
			if next.P.X == m.end.X && next.P.Y == m.end.Y {
				m.bestScore = next.dist
				m.distances[next.P] = next
				m.paths = append(m.paths, &next)
				if m.stopAfterFirstFound {
					return m.bestScore
				}
			}
		}
	}
	return m.bestScore
}

func (m *Maze) SetMaxDistance(distance int)           { m.maxDistance = distance }
func (m *Maze) SetStopAfterFirstFound(firstOnly bool) { m.stopAfterFirstFound = firstOnly }

func (m *Maze) DistanceFromStart(p pos.Position) int {
	return m.distances[p].dist
}

func (m *Maze) IsOnShortestPath(p pos.Position) bool {
	for _, n := range m.paths {
		for n2 := n; n2.cameFrom != nil; n2 = n2.cameFrom {
			if n2.P == p {
				return true
			}
		}
	}
	return false

}

func (m *Maze) ShortestPath() map[pos.Position]Node {
	return m.distances
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
	m.unvisited[m.start] = Node{P: m.start, dist: 0}
	m.distances[m.start] = m.unvisited[m.start]
	for y := 0; y < len(m.field); y++ {
		for x := 0; x < len(m.field[y]); x++ {
			if pathSymbol(m.field[y][x]) {
				if x != m.start.X || y != m.start.Y {
					p := pos.Position{X: x, Y: y}
					m.unvisited[p] = Node{P: p, dist: math.MaxInt}
				}
			}
		}
	}
}

func (m *Maze) selectNextNode(result *Node) bool {
	minS := math.MaxInt
	if len(m.unvisited) == 0 {
		return false
	}
	for p := range m.unvisited {
		node := m.unvisited[p]
		if node.dist < minS {
			*result = node
			minS = node.dist
		}
	}
	return minS < math.MaxInt
}

func (m *Maze) PrintMaze() {
	for y, v := range m.field {
		for x, c := range v {
			if m.IsOnShortestPath(pos.Position{X: x, Y: y}) {
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
