package main

import (
	"advent/aoc/common"
	"advent/aoc/pos"
	"advent/aoc/term"
	"fmt"
)

func main() {
	lines := common.StartDay(8, "input")
	antiNodes := findAllAntinodes(lines, false)
	printMap(lines, antiNodes)
	fmt.Println("Part 1: ", len(antiNodes))

	antiNodes = findAllAntinodes(lines, true)
	printMap(lines, antiNodes)
	fmt.Println("Part 2: ", len(antiNodes))

}

func findAllAntinodes(lines []string, resonant bool) []pos.Position {
	var result []pos.Position
	for y, v := range lines {
		for x, c := range v {
			if c != '.' {
				findAntinodesForPos(pos.Position{X: x, Y: y}, lines, &result, resonant)
				if resonant {
					findResonantAntennas(pos.Position{X: x, Y: y}, lines, &result)
				}
			}
		}
	}
	return result
}

func findAntinodesForPos(p pos.Position, lines []string, result *[]pos.Position, resonant bool) {
	sym := lines[p.Y][p.X]
	for y, v := range lines {
		for x := range v {
			p2 := pos.Position{X: x, Y: y}
			if v[x] == sym && !p.Equals(p2) {
				findResonantAntinodes(p, p2, result, len(v), len(lines), resonant)
			}
		}
	}
}

func findResonantAntennas(p pos.Position, lines []string, result *[]pos.Position) {
	sym := lines[p.Y][p.X]
	for y, v := range lines {
		for x := range v {
			p2 := pos.Position{X: x, Y: y}
			if v[x] == sym && !p.Equals(p2) {
				if isAntennaResonant(p, p2, lines, len(v), len(lines)) {
					addUnique(result, p)
				}
			}
		}
	}
}

func findResonantAntinodes(p1prime pos.Position, p2prime pos.Position, result *[]pos.Position, sizeX int, sizeY int, resonant bool) {
	diff := p1prime.DistanceTo(p2prime)
	done := false
	for !done {
		p2prime = p2prime.Move(diff)
		var added bool
		inMap2 := isInMap(p2prime, sizeX, sizeY)
		if inMap2 {
			added = addUnique(result, p2prime)
		}
		p1prime = p1prime.Move(pos.Distance{Dx: -diff.Dx, Dy: -diff.Dy})
		inMap1 := isInMap(p1prime, sizeX, sizeY)
		if inMap1 {
			added = added || addUnique(result, p1prime)
		}
		done = !resonant || !(inMap1 || inMap2)
	}
}

func isAntennaResonant(p1prime pos.Position, p2prime pos.Position, lines []string, sizeX int, sizeY int) bool {
	diff := p1prime.DistanceTo(p2prime)
	symbol := lines[p1prime.Y][p1prime.X]
	done := false
	var found bool
	for !done {
		p2second := p2prime.Move(diff)
		inMap2 := isInMap(p2second, sizeX, sizeY)
		if inMap2 {
			found = lines[p2second.Y][p2second.X] == symbol
		}
		p2prime = p2second

		p1second := p1prime.Move(pos.Distance{Dx: -diff.Dx, Dy: -diff.Dy})
		inMap1 := isInMap(p1second, sizeX, sizeY)
		if inMap1 {
			found = lines[p1second.Y][p1second.X] == symbol
		}
		p1prime = p1second
		// if we found a third antenna aligned with the given 2 or if we ran out of the map in both directions, then we're done.
		done = found || !(inMap1 || inMap2)
	}
	return !found
}

func isInMap(p pos.Position, sizeX int, sizeY int) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < sizeX && p.Y < sizeY
}

func addUnique(positions *[]pos.Position, p pos.Position) bool {
	found := contains(*positions, p)
	if !found {
		*positions = append(*positions, p)
	}
	return !found
}

func contains(positions []pos.Position, p pos.Position) bool {
	found := false
	for i := 0; i < len(positions) && !found; i++ {
		found = positions[i].Equals(p)
	}
	return found
}

func printMap(lines []string, antinodes []pos.Position) {
	for y, v := range lines {
		for x, c := range v {
			hasAntinode := contains(antinodes, pos.Position{X: x, Y: y})
			if c != '.' {
				if hasAntinode {
					fmt.Print(term.RED)
				} else {
					fmt.Print(term.GREEN)
				}
				fmt.Print(string(c) + term.WHITE)

			} else {
				if hasAntinode {
					fmt.Print(term.BLUE + "#" + term.WHITE)
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}
