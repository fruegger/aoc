package main

import (
	"advent/aoc/common"
	"advent/aoc/pos"
	"advent/aoc/term"
	"fmt"
)

func main() {
	lines := common.StartDay(15, "input")
	warehouse, moves, ok := scanInput(lines, false)
	if !ok {
		return
	}
	helpWith(&warehouse, moves)

	fmt.Println()
	fmt.Println("Part1: ", warehouse.score())

	warehouse2, _, ok2 := scanInput(lines, true)
	if !ok2 {
		return
	}

	tot := 0
	for _, v := range warehouse2.floor {
		for _, c := range v {
			if c == WIDEBOXL {
				tot++
			}
		}
	}
	fmt.Println(tot, "Boxes.")
	helpWith(&warehouse2, moves)
	tot = 0
	for _, v := range warehouse2.floor {
		for _, c := range v {
			if c == WIDEBOXL {
				tot++
			}
		}
	}
	fmt.Println(tot, "Boxes.")

	fmt.Println()
	fmt.Println("Part2: ", warehouse2.score())

}

func helpWith(warehouse *Warehouse, moves []pos.Direction) {
	warehouse.printWarehouse()
	for _, move := range moves {
		warehouse.moveRobot(move)
		//fmt.Println("Move: ", move.Symbol)
		//warehouse.printWarehouse()

	}
	warehouse.printWarehouse()
}

const BOX = 'O'
const ROBOT = '@'
const WALL = '#'
const WIDEBOXL = '['
const WIDEBOXR = ']'
const NOTHING = '.'

type Warehouse struct {
	floor []string
	robot pos.Position
	wide  bool
}

func (w *Warehouse) printWarehouse() {
	for _, line := range w.floor {
		for _, c := range line {
			switch c {
			case BOX:
				fmt.Print(term.BLUE, string(BOX), term.WHITE)
			case WIDEBOXL:
				fmt.Print(term.BLUE, string(WIDEBOXL), term.WHITE)
			case WIDEBOXR:
				fmt.Print(term.BLUE, string(WIDEBOXR), term.WHITE)
			case ROBOT:
				fmt.Print(term.RED, string(ROBOT), term.WHITE)
			case WALL:
				fmt.Print(term.GREEN, string(WALL), term.WHITE)
			default:
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}
	fmt.Println("Robot: ", w.robot)
}

func (w *Warehouse) score() int {
	total := 0
	for y := 0; y < len(w.floor); y++ {
		for x := 0; x < len(w.floor[y]); x++ {
			if w.wide {
				if w.floor[y][x] == WIDEBOXL {
					total += 100*y + x
				}

			} else {
				if w.floor[y][x] == BOX {
					total += 100*y + x
				}
			}
		}
	}
	return total
}

func (w *Warehouse) findRobot() bool {
	for y := 0; y < len(w.floor); y++ {
		for x := 0; x < len(w.floor[y]); x++ {
			if w.floor[y][x] == ROBOT {
				w.robot = pos.Position{X: x, Y: y}
				return true
			}
		}
	}
	return false
}

func scanInput(lines []string, wide bool) (Warehouse, []pos.Direction, bool) {
	w := Warehouse{}
	var m []pos.Direction
	for _, v := range lines {
		if len(v) > 0 {
			if v[0] == WALL {
				if wide {
					w.floor = append(w.floor, wideTranslation(v))
				} else {
					w.floor = append(w.floor, v)
				}
			} else {
				for _, c := range v {
					m = append(m, pos.Directions[string(c)])
				}
			}
		}
	}
	w.wide = wide
	ok := w.findRobot()
	return w, m, ok
}

func wideTranslation(line string) string {
	result := ""
	for _, c := range line {
		var part string
		switch c {
		case WALL:
			part = string(WALL) + string(WALL)
		case ROBOT:
			part = string(ROBOT) + string(NOTHING)
		case NOTHING:
			part = string(NOTHING) + string(NOTHING)
		case BOX:
			part = string(WIDEBOXL) + string(WIDEBOXR)
		}
		result += part
	}
	return result
}

func (w *Warehouse) moveRobot(move pos.Direction) {
	p2 := w.robot.Move(move)
	switch w.floor[p2.Y][p2.X] {
	case BOX, WIDEBOXR, WIDEBOXL:
		{
			w.pushBoxes(p2, move)
		}
	case WALL:
		{
		}
	default:
		{
			swap(&w.floor, w.robot, p2)
			w.robot = p2
		}
	}
}

func (w *Warehouse) pushBoxes(p2 pos.Position, move pos.Direction) {
	if w.wide {
		switch move {
		case pos.LEFT:
			if w.pushBoxLeft(p2) {
				w.floor[p2.Y] = w.floor[p2.Y][:p2.X] + string(ROBOT) + string(NOTHING) + w.floor[p2.Y][p2.X+2:]
				w.robot = p2
			}
		case pos.RIGHT:
			if w.pushBoxRight(p2) {
				w.floor[p2.Y] = w.floor[p2.Y][:p2.X-1] + string(NOTHING) + string(ROBOT) + w.floor[p2.Y][p2.X+1:]
				w.robot = p2
			}
		case pos.UP:
			if w.pushBoxUp(p2) {
				w.floor[p2.Y] = w.floor[p2.Y][:p2.X] + string(ROBOT) + w.floor[p2.Y][p2.X+1:]
				w.floor[p2.Y+1] = w.floor[p2.Y+1][:p2.X] + string(NOTHING) + w.floor[p2.Y+1][p2.X+1:]
				w.robot = p2
			}

		case pos.DOWN:
			if w.pushBoxDown(p2) {
				w.robot = p2
				w.floor[p2.Y] = w.floor[p2.Y][:p2.X] + string(ROBOT) + w.floor[p2.Y][p2.X+1:]
				w.floor[p2.Y-1] = w.floor[p2.Y-1][:p2.X] + string(NOTHING) + w.floor[p2.Y-1][p2.X+1:]
			}
		}
	} else {
		var p3 = p2.Move(move)
		for w.floor[p3.Y][p3.X] == BOX {
			p3 = p3.Move(move)
		}
		if w.floor[p3.Y][p3.X] != WALL {
			swap(&w.floor, p2, p3)
			swap(&w.floor, p2, w.robot)
			w.robot = p2
		}
	}
}

func (w *Warehouse) pushBoxUp(p2 pos.Position) bool {
	if !w.canPushBoxUp(p2) {
		return false
	}
	p3 := p2.Move(pos.UP)
	var p4 pos.Position
	if w.floor[p2.Y][p2.X] == WIDEBOXR {
		p4 = p3.Move(pos.LEFT)
	} else {
		p4 = p3.Move(pos.RIGHT)
	}
	if w.floor[p3.Y][p3.X] != NOTHING {
		w.pushBoxUp(p3)
	}
	if w.floor[p4.Y][p4.X] != NOTHING {
		w.pushBoxUp(p4)
	}
	w.floor[p3.Y] = w.floor[p3.Y][:p3.X] + string(w.floor[p3.Y+1][p3.X]) + w.floor[p3.Y][p3.X+1:]
	w.floor[p4.Y] = w.floor[p4.Y][:p4.X] + string(w.floor[p4.Y+1][p4.X]) + w.floor[p4.Y][p4.X+1:]
	w.floor[p3.Y+1] = w.floor[p3.Y+1][:p3.X] + string(NOTHING) + w.floor[p3.Y+1][p3.X+1:]
	w.floor[p4.Y+1] = w.floor[p4.Y+1][:p4.X] + string(NOTHING) + w.floor[p4.Y+1][p4.X+1:]

	return true
}

func (w *Warehouse) pushBoxDown(p2 pos.Position) bool {
	if !w.canPushBoxDown(p2) {
		return false
	}
	p3 := p2.Move(pos.DOWN)
	var p4 pos.Position
	if w.floor[p2.Y][p2.X] == WIDEBOXR {
		p4 = p3.Move(pos.LEFT)
	} else {
		p4 = p3.Move(pos.RIGHT)
	}
	if w.floor[p3.Y][p3.X] != NOTHING {
		w.pushBoxDown(p3)
	}
	if w.floor[p4.Y][p4.X] != NOTHING {
		w.pushBoxDown(p4)
	}
	w.floor[p3.Y] = w.floor[p3.Y][:p3.X] + string(w.floor[p3.Y-1][p3.X]) + w.floor[p3.Y][p3.X+1:]
	w.floor[p4.Y] = w.floor[p4.Y][:p4.X] + string(w.floor[p4.Y-1][p4.X]) + w.floor[p4.Y][p4.X+1:]
	w.floor[p3.Y-1] = w.floor[p3.Y-1][:p3.X] + string(NOTHING) + w.floor[p3.Y-1][p3.X+1:]
	w.floor[p4.Y-1] = w.floor[p4.Y-1][:p4.X] + string(NOTHING) + w.floor[p4.Y-1][p4.X+1:]

	return true
}

func (w *Warehouse) pushBoxLeft(p2 pos.Position) bool {
	if !w.canPushBoxLeft(p2) {
		return false
	}
	p3 := p2.Move(pos.LEFT).Move(pos.LEFT)
	if w.floor[p3.Y][p3.X] != NOTHING {
		w.pushBoxLeft(p3)
	}
	w.floor[p3.Y] = w.floor[p3.Y][:p3.X] + string(WIDEBOXL) + string(WIDEBOXR) + string(NOTHING) + w.floor[p3.Y][p2.X+1:]
	return true
}

func (w *Warehouse) pushBoxRight(p2 pos.Position) bool {
	if !w.canPushBoxRight(p2) {
		return false
	}
	p3 := p2.Move(pos.RIGHT).Move(pos.RIGHT)
	if w.floor[p3.Y][p3.X] != NOTHING {
		w.pushBoxRight(p3)
	}
	w.floor[p3.Y] = w.floor[p3.Y][:p2.X] + string(NOTHING) + string(WIDEBOXL) + string(WIDEBOXR) + w.floor[p3.Y][p3.X+1:]

	return true

}

func (w *Warehouse) canPushBoxUp(p2 pos.Position) bool {
	p3 := p2.Move(pos.UP)
	var p4 pos.Position
	if w.floor[p2.Y][p2.X] == WIDEBOXR {
		p4 = p3.Move(pos.LEFT)
	} else {
		p4 = p3.Move(pos.RIGHT)
	}
	if w.floor[p3.Y][p3.X] == WALL || w.floor[p4.Y][p4.X] == WALL {
		return false
	} else {
		if w.floor[p3.Y][p3.X] == NOTHING && w.floor[p4.Y][p4.X] == NOTHING {
			return true
		} else {
			return (w.floor[p3.Y][p3.X] == NOTHING || w.canPushBoxUp(p3)) &&
				(w.floor[p4.Y][p4.X] == NOTHING || w.canPushBoxUp(p4))
		}
	}

}
func (w *Warehouse) canPushBoxDown(p2 pos.Position) bool {
	p3 := p2.Move(pos.DOWN)
	var p4 pos.Position
	if w.floor[p2.Y][p2.X] == WIDEBOXR {
		p4 = p3.Move(pos.LEFT)
	} else {
		p4 = p3.Move(pos.RIGHT)
	}
	if w.floor[p3.Y][p3.X] == WALL || w.floor[p4.Y][p4.X] == WALL {
		return false
	} else {
		if w.floor[p3.Y][p3.X] == NOTHING && w.floor[p4.Y][p4.X] == NOTHING {
			return true
		} else {
			return (w.floor[p3.Y][p3.X] == NOTHING || w.canPushBoxDown(p3)) &&
				(w.floor[p4.Y][p4.X] == NOTHING || w.canPushBoxDown(p4))
		}
	}
}

func (w *Warehouse) canPushBoxLeft(p2 pos.Position) bool {
	p3 := p2.Move(pos.LEFT).Move(pos.LEFT)
	if w.floor[p3.Y][p3.X] == WALL {
		return false
	} else {
		if w.floor[p3.Y][p3.X] == NOTHING {
			return true
		} else {
			return w.canPushBoxLeft(p3)
		}
	}
}

func (w *Warehouse) canPushBoxRight(p2 pos.Position) bool {
	p3 := p2.Move(pos.RIGHT).Move(pos.RIGHT)
	if w.floor[p3.Y][p3.X] == WALL {
		return false
	} else {
		if w.floor[p3.Y][p3.X] == NOTHING {
			return true
		} else {
			return w.canPushBoxRight(p3)
		}
	}
}

func swap(warehouse *[]string, p1 pos.Position, p2 pos.Position) {
	save1 := string((*warehouse)[p1.Y][p1.X])
	save2 := string((*warehouse)[p2.Y][p2.X])
	(*warehouse)[p1.Y] = (*warehouse)[p1.Y][:p1.X] + save2 + (*warehouse)[p1.Y][p1.X+1:]
	(*warehouse)[p2.Y] = (*warehouse)[p2.Y][:p2.X] + save1 + (*warehouse)[p2.Y][p2.X+1:]
}
