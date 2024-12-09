package main

import (
	"advent/aoc/common"
	"advent/aoc/pos"
	"advent/aoc/term"
	"fmt"
	"strings"
)

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

	file := common.OpenFile("./day6/day6_input_h.txt")
	defer file.Close()

	var lines []string
	common.ScanLines(file, &lines)
	initialPos := findInitialPosition(lines)

	//save a copy for part 2
	lines2 := common.CopyLines(lines)

	fmt.Print(term.WHITE)
	printLines(lines)

	predictPath(initialPos, pos.UP, lines)

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
			p := pos.Position{X: x, Y: y}
			if p != initialPos && lines2[y][x] != SYM_Osbtacle {
				save := lines2[y][x]
				replaceSymAtPos(p, lines2, SYM_AddedObstruction)
				if testForLoop(initialPos, pos.UP, lines2, totalElements) {
					replaceSymAtPos(p, obstructions, SYM_AddedObstruction)
				}
				replaceSymAtPos(p, lines2, save)
			}
		}
	}

	printLines(obstructions)
	fmt.Println("Part2: ", countSymbols(obstructions, SYM_AddedObstruction))
}

func findInitialPosition(lines []string) pos.Position {
	found := false
	posy := 0
	for i := 0; i < len(lines) && !found; i++ {
		found = strings.Contains(lines[i], pos.UP.Symbol)
		if found {
			posy = i
		}
	}
	posx := strings.Index(lines[posy], pos.UP.Symbol)
	return pos.Position{X: posx, Y: posy}
}

func lookAhead(p pos.Position, d pos.Direction, lines []string) ObjectType {
	obectPos := p.Move(d)
	var result ObjectType
	if obectPos.Y < 0 || obectPos.X < 0 || obectPos.Y >= len(lines) || obectPos.X >= len(lines[obectPos.Y]) {
		result = OBJ_Exit
	} else {
		if lines[obectPos.Y][obectPos.X] == SYM_Osbtacle || lines[obectPos.Y][obectPos.X] == SYM_AddedObstruction {
			result = OBJ_Obstacle
		} else {
			result = OBJ_Nothing
		}
	}
	return result
}

func predictPath(p pos.Position, d pos.Direction, lines []string) {
	obj := lookAhead(p, d, lines)
	if obj == OBJ_Nothing {
		replaceSymAtPos(p, lines, SYM_Breadcrumb)
		predictPath(p.Move(d), d, lines)
	} else {
		if obj == OBJ_Obstacle {
			predictPath(p, d.TurnRight(), lines)
		} else {
			replaceSymAtPos(p, lines, SYM_Breadcrumb)
		}
	}
}

func replaceSymAtPos(p pos.Position, lines []string, newSym uint8) {
	newline := []uint8(lines[p.Y])
	newline[p.X] = newSym
	lines[p.Y] = string(newline)
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
func testForLoop(p pos.Position, d pos.Direction, lines []string, countdown int) bool {
	if countdown <= 0 {
		return true
	}
	obj := lookAhead(p, d, lines)
	if obj == OBJ_Nothing {
		return testForLoop(p.Move(d), d, lines, countdown-1)
	} else {
		if obj == OBJ_Obstacle || obj == OBJ_AddedObstruction {
			return testForLoop(p, d.TurnRight(), lines, countdown-1)
		}
	}
	return false
}

//some goofy term stuff

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
		result = term.GREEN + string(c) + term.WHITE
	case SYM_Breadcrumb:
		result = term.BLUE + string(c) + term.WHITE
	case SYM_Start:
		result = term.RED + string(c) + term.WHITE
	case SYM_AddedObstruction:
		result = term.YELLOW + string(c) + term.WHITE
	default:
		result = string(c)
	}
	return result
}
