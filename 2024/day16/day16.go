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
	lines := common.StartDay(16, "test")
	start, ok := common.FindSymbol(lines, START)
	if !ok {
		fmt.Println("error, no S or E found")
		return
	}
	fmt.Println("Start:", start)
	printMaze(lines)

	scorer := math.MaxInt
	walk(&lines, start, pos.RIGHT, 0, &scorer)

	lines2 := common.CopyLines(lines)
	scoreu := math.MaxInt
	walk(&lines2, start, pos.UP, 0, &scoreu)

	fmt.Println("Score R:", scorer)
	fmt.Println("Score U:", scoreu)

}

func walk(maze *[]string, p pos.Position, d pos.Direction, score int, bestSolution *int) {
	/*
		if (*maze)[p.Y][p.X] == END {
			if score < *bestSolution {
				*bestSolution = score
			}
		} else {
			if (*maze)[p.Y][p.X] == NOTHING || (*maze)[p.Y][p.X] == START {
				walk(maze, p.Move(d), d, score+1, bestSolution)
				d2 := turnCounterclockwise(d)
				walk(maze, p.Move(d2), d2, score+1001, bestSolution)
				//			d3 := turnCounterclockwise(turnCounterclockwise(d2))
				//			walk(maze, p.Move(d3), d3, score+3001, bestSolution)
			}
		}
	*/
	p2 := p.Move(d)
	if (*maze)[p2.Y][p2.X] == END {
		if score < *bestSolution {
			*bestSolution = score
		}
	} else {
		if (*maze)[p2.Y][p2.X] == NOTHING {
			common.ChangeSymbol(maze, p, '*')
			walk(maze, p2, d, score+1, bestSolution)
		} else {
			if d == pos.DOWN {
				// all possible directions used up; backtrack
				common.ChangeSymbol(maze, p, '*')
			} else {
				d2 := turnCounterclockwise(d)
				walk(maze, p, d2, score+1000, bestSolution)
			}
		}
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

func printMaze(lines []string) {
	for _, v := range lines {
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
