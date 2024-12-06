package main

import (
	"advent/common"
	"fmt"
	"strings"
)

type Position struct {
	x int
	y int
}

type Direction struct {
	dx     int
	dy     int
	symbol string
}

var RIGHT = Direction{dx: 1, dy: 0, symbol: ">"}
var DOWN = Direction{dx: 0, dy: 1, symbol: "v"}
var LEFT = Direction{dx: -1, dy: 0, symbol: "<"}
var UP = Direction{dx: 0, dy: -1, symbol: "^"}

func (p Position) move(d Direction) Position {
	return Position{
		x: p.x + d.dx,
		y: p.y + d.dy,
	}
}

func (d Direction) turnRight() Direction {
	var result Direction
	switch d {
	case RIGHT:
		result = DOWN
	case DOWN:
		result = LEFT
	case LEFT:
		result = UP
	case UP:
		result = RIGHT
	}
	return result
}

type ObjectType uint8

const (
	OBJ_Obstacle ObjectType = iota
	OBJ_Exit
	OBJ_Nothing
	OBJ_AddedObstruction
)

const SYM_Osbtacle = '#'
const SYM_Breadcrumb = '*'
const SYM_Start = '^'
const SYM_AddedObstruction = 'O'

func main() {
	fmt.Println("Day 6")
	fmt.Println("=====")

	file := common.OpenFile("./day6_input.txt")
	defer file.Close()

	var lines []string
	common.ScanLines(file, &lines)
	initialPos := findInitialPosition(lines)

	//save a copy for part 2
	lines2 := common.CopyLines(lines)

	fmt.Print(WHITE)
	printLines(lines)

	predictPath(initialPos, UP, lines)

	fmt.Println("Breadcrumbs")

	printLines(lines)

	fmt.Println("Part 1: ", countSymbols(lines, SYM_Breadcrumb))

	fmt.Println("Obstacles")

	obstructions := common.CopyLines(lines2)

	// calculate how large the whole map is
	totalElements := 0
	for _, v := range lines2 {
		totalElements += len(v)
	}

	//place an obstruction at each spot where there is no obstacle and is not the start in turn and test for loops
	for y := 0; y < len(lines2); y++ {
		for x := 0; x < len(lines2[y]); x++ {
			p := Position{x: x, y: y}
			if p != initialPos && lines2[y][x] != SYM_Osbtacle {
				save := lines2[y][x]
				replaceSymAtPos(p, lines2, SYM_AddedObstruction)
				if testForLoop(initialPos, UP, lines2, totalElements) {
					replaceSymAtPos(p, obstructions, SYM_AddedObstruction)
				}
				replaceSymAtPos(p, lines2, save)
			}
		}
	}

	printLines(obstructions)
	fmt.Println("Part2: ", countSymbols(obstructions, SYM_AddedObstruction))
}

func findInitialPosition(lines []string) Position {
	found := false
	posy := 0
	for i := 0; i < len(lines) && !found; i++ {
		found = strings.Contains(lines[i], UP.symbol)
		if found {
			posy = i
		}
	}
	posx := strings.Index(lines[posy], UP.symbol)
	return Position{x: posx, y: posy}
}

func lookAhead(p Position, d Direction, lines []string) ObjectType {
	obectPos := p.move(d)
	var result ObjectType
	if obectPos.y < 0 || obectPos.x < 0 || obectPos.y >= len(lines) || obectPos.x >= len(lines[obectPos.y]) {
		result = OBJ_Exit
	} else {
		if lines[obectPos.y][obectPos.x] == SYM_Osbtacle || lines[obectPos.y][obectPos.x] == SYM_AddedObstruction {
			result = OBJ_Obstacle
		} else {
			result = OBJ_Nothing
		}
	}
	return result
}

func predictPath(p Position, d Direction, lines []string) {
	obj := lookAhead(p, d, lines)
	if obj == OBJ_Nothing {
		replaceSymAtPos(p, lines, SYM_Breadcrumb)
		predictPath(p.move(d), d, lines)
	} else {
		if obj == OBJ_Obstacle {
			predictPath(p, d.turnRight(), lines)
		} else {
			replaceSymAtPos(p, lines, SYM_Breadcrumb)
		}
	}
}

func replaceSymAtPos(p Position, lines []string, newSym uint8) {
	newline := []uint8(lines[p.y])
	newline[p.x] = newSym
	lines[p.y] = string(newline)
}

func countSymbols(lines []string, sym int32) int {
	total := 0
	for _, v := range lines {
		for _, c := range v {
			if c == sym {
				total++
			}
		}
	}
	return total
}

// the idea is to keep on going until an exit is found or an amount of spots sufficient to cover the whole map has been
// visited (this can be improved, given some time)
func testForLoop(p Position, d Direction, lines []string, countdown int) bool {
	if countdown <= 0 {
		return true
	}
	obj := lookAhead(p, d, lines)
	if obj == OBJ_Nothing {
		return testForLoop(p.move(d), d, lines, countdown-1)
	} else {
		if obj == OBJ_Obstacle || obj == OBJ_AddedObstruction {
			return testForLoop(p, d.turnRight(), lines, countdown-1)
		}
	}
	return false
}

//some goofy term stuff

const GREEN = "\x1B[32m"
const BLUE = "\x1B[34m"
const WHITE = "\x1B[97m"
const RED = "\x1B[31m"
const YELLOW = "\x1B[33m"

func printLines(lines []string) {
	for _, v := range lines {
		for _, c := range v {
			fmt.Print(terminalString(c))
		}
		fmt.Println()
	}
}
func terminalString(c rune) string {
	var result string
	switch c {
	case SYM_Osbtacle:
		result = GREEN + string(c) + WHITE
	case SYM_Breadcrumb:
		result = BLUE + string(c) + WHITE
	case SYM_Start:
		result = RED + string(c) + WHITE
	case SYM_AddedObstruction:
		result = YELLOW + string(c) + WHITE
	default:
		result = string(c)
	}
	return result
}
