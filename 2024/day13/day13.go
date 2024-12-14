package main

import (
	"advent/aoc/common"
	"advent/aoc/pos"
	"fmt"
	"strings"
)

func main() {
	lines := common.StartDay(13, "input")

	machines := readMachineDefinitions(lines)

	total := 0
	for _, machine := range machines {
		hasSolution := machine.hasSolution(0)
		if hasSolution {
			price := machine.minimumPrice(0)
			if price < IMPOSSIBLE_PRICE {
				total += price
			}
		}
	}
	fmt.Println("Par1 : ", total)

	total = 0
	for _, machine := range machines {
		hasSolution := machine.hasSolution(10000000000000)
		if hasSolution {
			price := machine.minimumPrice(10000000000000)
			if price < IMPOSSIBLE_PRICE {
				total += price
			}
		}
	}
	fmt.Println("Part2 : ", total)

}

func readMachineDefinitions(lines []string) []Machine {
	var result []Machine
	var aB *pos.Distance
	var bB *pos.Distance
	var pr *pos.Position
	for _, line := range lines {
		if strings.Contains(line, "Button A: ") {
			dx := common.StringToNum(line[12:14])
			dy := common.StringToNum(line[18:20])
			aB = &pos.Distance{Dx: dx, Dy: dy}
		} else {
			if strings.Contains(line, "Button B: ") {
				dx := common.StringToNum(line[12:14])
				dy := common.StringToNum(line[18:20])
				bB = &pos.Distance{Dx: dx, Dy: dy}
			} else {
				if strings.Contains(line, "Prize: ") {
					st := strings.Split(line[7:], ", ")
					px := common.StringToNum(st[0][2:])
					py := common.StringToNum(st[1][2:])
					pr = &pos.Position{X: px, Y: py}
				}
			}
		}
		if aB != nil && bB != nil && pr != nil {
			result = append(result, Machine{aButton: *aB, bButton: *bB, price: *pr})
			aB = nil
			bB = nil
			pr = nil
		}
	}
	return result
}
