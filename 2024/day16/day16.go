package day16

import (
	"advent/aoc/common"
	"advent/aoc/maze"
	"advent/aoc/pos"
	"fmt"
)

func main() {
	lines := common.StartDay(16, "test")
	m := maze.Maze{}
	initializeMaze(&m, lines)

	m.PrintMaze()

	result := m.Traverse(distanceFn)
	//maze.debugMaze()
	fmt.Println("Part1:", result)
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

func distanceFn(p1 pos.Position, p2 pos.Position) int {
	if p1.X == p2.X || p1.Y == p2.Y {
		return 1
	} else {
		//		if m.field[p1.X][p2.Y] != WALL || m.field[p2.X][p1.Y] != WALL {
		return 1002
		//		} else {
		//			return math.MaxInt
		//		}
	}
}

/*
	func (m *Maze) BFSwalk(node Node) {
		p := move.p
		d := move.d
		s := move.s
		if (m.field)[p.Y][p.X] == END {
			fmt.Print(".")
			if move.s < m.bestScore {
				m.bestScore = move.s
			}
		} else {
			if m.accepts(p) {
				d3 := turnClockwise(d)
				p3 := p.Move(d3)
				p1 := p.Move(d)
				if m.accepts(p1) && !m.visited[p1.Y][p1.X] {
					m.visited[p1.Y][p1.X] = true
					mq.Push(Move{p: p1, d: d, s: s + 1})
				}
				if m.accepts(p3) && !m.visited[p3.Y][p3.X] {
					m.visited[p3.Y][p3.X] = true
					mq.Push(Move{p: p3, d: d3, s: s + 1001})
				}
				d2 := turnCounterclockwise(d)
				p2 := p.Move(d2)
				if m.accepts(p2) && !m.visited[p2.Y][p2.X] {
					m.visited[p2.Y][p2.X] = true
					mq.Push(Move{p: p2, d: d2, s: s + 1001})
				}
			}
		}
	}

	func (m *Maze) DFSwalk(move Move) {
		p := move.p
		d := move.d
		s := move.s
		if (m.field)[p.Y][p.X] == END {
			fmt.Print(".")
			if move.s < m.bestScore {
				m.bestScore = move.s
			}
		} else {
			if m.accepts(p) {
				d3 := turnClockwise(d)
				p3 := p.Move(d3)
				if m.accepts(p3) && !m.visited[p3.Y][p3.X] {
					//m.visited[p3.Y][p3.X] = true
					m.DFSwalk(Move{p: p3, d: d3, s: s + 1001})
					m.visited[p3.Y][p3.X] = false
				}
				p1 := p.Move(d)
				if m.accepts(p1) && !m.visited[p1.Y][p1.X] {
					//m.visited[p1.Y][p1.X] = true
					m.DFSwalk(Move{p: p1, d: d, s: s + 1})
					m.visited[p1.Y][p1.X] = false
				}
				d2 := turnCounterclockwise(d)
				p2 := p.Move(d2)
				if m.accepts(p2) && !m.visited[p2.Y][p2.X] {
					//m.visited[p2.Y][p2.X] = true
					m.DFSwalk(Move{p: p2, d: d2, s: s + 1001})
					m.visited[p2.Y][p2.X] = false
				}
			}
		}
	}

	func (m *Maze) walk(move Move) bool {
		p := move.p
		d := move.d
		s := move.s
		if (m.field)[p.Y][p.X] == END {
			fmt.Print(".")
			if move.s < m.bestScore {
				m.bestScore = move.s
			}
		} else {
			if m.accepts(p) {
				d3 := turnClockwise(d)
				p3 := p.Move(d3)
				if m.accepts(p3) && !m.visited[p3.Y][p3.X] {
					m.visited[p3.Y][p3.X] = true
					mq.queueMove(Move{p: p3, d: d3, s: s + 1001})
				}
				p1 := p.Move(d)
				if m.accepts(p1) && !m.visited[p1.Y][p1.X] {
					m.visited[p1.Y][p1.X] = true
					mq.queueMove(Move{p: p1, d: d, s: s + 1})
				}
				d2 := turnCounterclockwise(d)
				p2 := p.Move(d2)
				if m.accepts(p2) && !m.visited[p2.Y][p2.X] {
					m.visited[p2.Y][p2.X] = true
					mq.queueMove(Move{p: p2, d: d2, s: s + 1001})
				}
			}
		}
		return true
	}
*/
/*
func (m *Maze) accepts(p pos.Position) bool {
	return m.field[p.Y][p.X] == NOTHING || (m.field)[p.Y][p.X] == START || (m.field)[p.Y][p.X] == END

}

func (m *Maze) debugMaze() {
	fmt.Println("Debug")
	//	reader := bufio.NewReader(os.Stdin)
	//	text, _ := reader.ReadString('\n')
	//	if text == "q" {
	//		panic("quit")
	//	}

	for y := 0; y < len(m.field); y++ {
		for x := 0; x < len(m.field); x++ {
			_, ok := unvisited[pos.Position{X: y, Y: y}]
			if ok {
				fmt.Print(term.YELLOW + "*" + term.WHITE)
			} else {
				c := m.field[y][x]
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
