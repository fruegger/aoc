package main

import (
	"advent/aoc/common"
	"advent/aoc/pos"
	"fmt"
)

func main() {

	lines := common.StartDay(12, "input")

	regions := floodFill(lines)
	total := 0
	for _, r := range regions {
		total += len(r.plots) * r.edges

	}
	fmt.Println("Part1: ", total)

	total = 0
	marks := makeMarks(lines)
	for _, r := range regions {
		sides := r.countSides(marks)
		total += len(r.plots) * sides
	}
	fmt.Println("Part2: ", total)
}

type Plot struct {
	p     pos.Position
	edgeL bool
	edgeR bool
	edgeT bool
	edgeB bool
}

type Region struct {
	plantType uint8
	plots     []Plot
	edges     int
}

func (r Region) find(x int, y int) *Plot {
	for i := 0; i < len(r.plots); i++ {
		if r.plots[i].p.X == x && r.plots[i].p.Y == y {
			return &r.plots[i]
		}
	}
	return nil
}

func (r Region) countSides(marks [][]bool) int {
	result := r.edges
	for _, p1 := range r.plots {
		for _, p2 := range r.plots {
			if !(marks[p1.p.Y][p1.p.X] || marks[p2.p.Y][p2.p.X]) {
				if (p1.p.X-p2.p.X == 1 || p1.p.X-p2.p.X == -1) && p1.p.Y == p2.p.Y {
					if p1.edgeB && p2.edgeB {
						result--
					}
					if p1.edgeT && p2.edgeT {
						result--
					}
				}
				if (p1.p.Y-p2.p.Y == 1 || p1.p.Y-p2.p.Y == -1) && p1.p.X == p2.p.X {
					if p1.edgeL && p2.edgeL {
						result--
					}
					if p1.edgeR && p2.edgeR {
						result--
					}
				}
			}
		}
		marks[p1.p.Y][p1.p.X] = true
	}
	return result
}
func floodFill(lines []string) []Region {
	var result []Region

	marks := makeMarks(lines)
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len((lines)[y]); x++ {
			if !marks[y][x] {
				newRegion := Region{
					plantType: lines[y][x],
					plots: []Plot{{
						p:     pos.Position{x, y},
						edgeL: false,
						edgeB: false,
						edgeR: false,
						edgeT: false,
					}},
					edges: 0,
				}
				marks[y][x] = true
				fillNeighbours(lines, pos.Position{x, y}, &newRegion, &marks)
				result = append(result, newRegion)
			}
		}
	}
	return result
}

func makeMarks(lines []string) [][]bool {
	marks := make([][]bool, len(lines))
	for y := 0; y < len(lines); y++ {
		marks[y] = make([]bool, len((lines)[y]))

		for x := 0; x < len((lines)[y]); x++ {
			marks[y][x] = false
		}
	}
	return marks
}

func fillNeighbours(lines []string, p pos.Position, newRegion *Region, marks *[][]bool) {
	fillAndMark(lines, p, pos.UP, newRegion, marks)
	fillAndMark(lines, p, pos.RIGHT, newRegion, marks)
	fillAndMark(lines, p, pos.DOWN, newRegion, marks)
	fillAndMark(lines, p, pos.LEFT, newRegion, marks)
}

func fillAndMark(lines []string, p pos.Position, d pos.Direction, newRegion *Region, marks *[][]bool) {
	p2 := p.Move(d)
	if p2.Y >= 0 && p2.Y < len(lines) && p2.X >= 0 && p2.X < len(lines[p2.Y]) &&
		!(*marks)[p2.Y][p2.X] && newRegion.plantType == lines[p2.Y][p2.X] {
		newRegion.plots = append(newRegion.plots, Plot{p2, false, false, false, false})
		(*marks)[p2.Y][p2.X] = true
		fillNeighbours(lines, p2, newRegion, marks)
	} else {
		if p2.Y < 0 || p2.Y >= len(lines) || p2.X < 0 || p2.X >= len(lines[p2.Y]) ||
			newRegion.plantType != lines[p2.Y][p2.X] {
			newRegion.edges++
			plot := newRegion.find(p.X, p.Y)
			if d == pos.UP {
				plot.edgeT = true
			}
			if d == pos.DOWN {
				plot.edgeB = true
			}
			if d == pos.LEFT {
				plot.edgeL = true
			}
			if d == pos.RIGHT {
				plot.edgeR = true
			}
		}
	}
}
