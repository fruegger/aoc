package main

import (
	"advent/aoc/common"
	"fmt"
)

const size = 140

func main() {
	fmt.Println("Day 4")
	fmt.Println("=====")

	file := common.OpenFile("./day4/day4_input.txt")
	defer file.Close()
	var lines []string
	common.ScanLines(file, &lines)
	total := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			total += countXMas(x, y, lines)
		}
	}
	fmt.Println("Part1: ", total)

	total = 0
	for y := 1; y+1 < len(lines); y++ {
		for x := 1; x+1 < len(lines[y]); x++ {
			total += countCrossMas(x, y, lines)
		}
	}
	fmt.Println("Part2: ", total)

}
func countXMas(x int, y int, lines []string) int {
	total := 0
	if checkRight(x, y, lines) {
		total++
	}
	if checkRightDown(x, y, lines) {
		total++
	}
	if checkDown(x, y, lines) {
		total++
	}
	if checkLeftDown(x, y, lines) {
		total++
	}
	if checkLeft(x, y, lines) {
		total++
	}
	if checkLeftUp(x, y, lines) {
		total++
	}
	if checkUp(x, y, lines) {
		total++
	}
	if checkRightUp(x, y, lines) {
		total++
	}
	return total
}

func checkRight(x int, y int, lines []string) bool {
	if x > size-4 {
		return false
	} else {
		return lines[y][x] == 'X' && lines[y][x+1] == 'M' && lines[y][x+2] == 'A' && lines[y][x+3] == 'S'
	}
}

func checkRightDown(x int, y int, lines []string) bool {
	if y > size-4 || x > size-4 {
		return false
	} else {
		return lines[y][x] == 'X' && lines[y+1][x+1] == 'M' && lines[y+2][x+2] == 'A' && lines[y+3][x+3] == 'S'
	}
}

func checkDown(x int, y int, lines []string) bool {
	if y > size-4 {
		return false
	} else {
		return lines[y][x] == 'X' && lines[y+1][x] == 'M' && lines[y+2][x] == 'A' && lines[y+3][x] == 'S'
	}
}

func checkLeftDown(x int, y int, lines []string) bool {
	if x < 3 || y > size-4 {
		return false
	} else {
		return lines[y][x] == 'X' && lines[y+1][x-1] == 'M' && lines[y+2][x-2] == 'A' && lines[y+3][x-3] == 'S'
	}
}

func checkLeft(x int, y int, lines []string) bool {
	if x < 3 {
		return false
	} else {
		return lines[y][x] == 'X' && lines[y][x-1] == 'M' && lines[y][x-2] == 'A' && lines[y][x-3] == 'S'
	}
}

func checkLeftUp(x int, y int, lines []string) bool {
	if x < 3 || y < 3 {
		return false
	} else {
		return lines[y][x] == 'X' && lines[y-1][x-1] == 'M' && lines[y-2][x-2] == 'A' && lines[y-3][x-3] == 'S'
	}
}

func checkUp(x int, y int, lines []string) bool {
	if y < 3 {
		return false
	} else {
		return lines[y][x] == 'X' && lines[y-1][x] == 'M' && lines[y-2][x] == 'A' && lines[y-3][x] == 'S'
	}
}

func checkRightUp(x int, y int, lines []string) bool {
	if y < 3 || x > size-4 {
		return false
	} else {
		return lines[y][x] == 'X' && lines[y-1][x+1] == 'M' && lines[y-2][x+2] == 'A' && lines[y-3][x+3] == 'S'
	}
}

func countCrossMas(x int, y int, lines []string) int {
	total := 0
	if checkXRight(x, y, lines) {
		total++
	}
	if checkXDown(x, y, lines) {
		total++
	}
	if checkXLeft(x, y, lines) {
		total++
	}
	if checkXUp(x, y, lines) {
		total++
	}
	return total
}

func checkXRight(x int, y int, lines []string) bool {
	return lines[y][x] == 'A' &&
		lines[y-1][x-1] == 'M' && lines[y+1][x-1] == 'M' &&
		lines[y-1][x+1] == 'S' && lines[y+1][x+1] == 'S'
}

func checkXDown(x int, y int, lines []string) bool {
	return lines[y][x] == 'A' &&
		lines[y-1][x-1] == 'M' && lines[y-1][x+1] == 'M' &&
		lines[y+1][x-1] == 'S' && lines[y+1][x+1] == 'S'
}

func checkXLeft(x int, y int, lines []string) bool {
	return lines[y][x] == 'A' &&
		lines[y-1][x+1] == 'M' && lines[y+1][x+1] == 'M' &&
		lines[y-1][x-1] == 'S' && lines[y+1][x-1] == 'S'
}

func checkXUp(x int, y int, lines []string) bool {
	return lines[y][x] == 'A' &&
		lines[y+1][x-1] == 'M' && lines[y+1][x+1] == 'M' &&
		lines[y-1][x-1] == 'S' && lines[y-1][x+1] == 'S'
}
