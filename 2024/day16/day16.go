package main

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

func main() {
	lines := common.StartDay(16, "input")
	maze := Maze{}
	maze.initialize(lines)
	start, ok := common.FindSymbol(lines, START)
	if !ok {
		fmt.Println("error, no S or E found")
		return
	}
	fmt.Println("Start:", start)
	maze.printMaze()

	maze.traverse()
	maze.debugMaze()
	fmt.Println("Part1:", maze.bestScore)
}

type Maze struct {
	field     []string
	start     pos.Position
	visited   [][]bool
	bestScore int
}

func (m *Maze) initialize(lines []string) bool {
	m.field = lines
	var ok bool
	m.start, ok = common.FindSymbol(lines, START)
	if !ok {
		return false
	}
	for y := 0; y < len(m.field); y++ {
		var line []bool
		for y := 0; y < len(m.field); y++ {
			line = append(line, false)
		}
		m.visited = append(m.visited, line)
	}
	m.bestScore = math.MaxInt
	return true
}

var iter = 0

type Move struct {
	p pos.Position
	d pos.Direction
	s int
}

type Queue struct {
	moves []Move
}

var mq = Queue{}

func (q *Queue) queueMove(m Move) {
	q.moves = append(q.moves, m)
}

func (q *Queue) dequeueMove() (Move, bool) {
	if len(q.moves) == 0 {
		return Move{}, false
	}
	m := q.moves[0]
	q.moves = q.moves[1:]
	return m, true
}

func (m *Maze) traverse() {
	//m.DFSwalk(Move{m.start, pos.RIGHT, 0})
	mq.queueMove(Move{m.start, pos.RIGHT, 0})
	m.visited[m.start.Y][m.start.X] = true
	var move Move
	ok := true
	for ok {
		move, ok = mq.dequeueMove()
		if ok {
			m.BFSwalk(move)
		}
	}
}

func (m *Maze) BFSwalk(move Move) {
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
				mq.queueMove(Move{p: p1, d: d, s: s + 1})
			}
			if m.accepts(p3) && !m.visited[p3.Y][p3.X] {
				m.visited[p3.Y][p3.X] = true
				mq.queueMove(Move{p: p3, d: d3, s: s + 1001})
			}
			d2 := turnCounterclockwise(d)
			p2 := p.Move(d2)
			if m.accepts(p2) && !m.visited[p2.Y][p2.X] {
				m.visited[p2.Y][p2.X] = true
				mq.queueMove(Move{p: p2, d: d2, s: s + 1001})
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
				m.visited[p3.Y][p3.X] = true
				m.DFSwalk(Move{p: p3, d: d3, s: s + 1001})
				m.visited[p3.Y][p3.X] = false
			}
			p1 := p.Move(d)
			if m.accepts(p1) && !m.visited[p1.Y][p1.X] {
				m.visited[p1.Y][p1.X] = true
				m.DFSwalk(Move{p: p1, d: d, s: s + 1})
				m.visited[p1.Y][p1.X] = false
			}
			d2 := turnCounterclockwise(d)
			p2 := p.Move(d2)
			if m.accepts(p2) && !m.visited[p2.Y][p2.X] {
				m.visited[p2.Y][p2.X] = true
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
			if m.visited[y][x] {
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

func (m *Maze) printMaze() {
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
