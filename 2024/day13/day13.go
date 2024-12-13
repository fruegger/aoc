package main

import (
	"advent/aoc/common"
	"advent/aoc/pos"
	"fmt"
	"strings"
)

type Machine struct {
	aButton pos.Distance
	bButton pos.Distance
	price   pos.Position
}

const IMPOSSIBLE_PRICE = 401

func (m Machine) minimumPrice(offset int) int {
	minCost := IMPOSSIBLE_PRICE
	costa := 0
	ax := 0
	ay := 0
	tx := m.price.X + offset
	ty := m.price.Y + offset
	fmt.Print("M")

	for aCnt := 0; ax <= tx && ay <= ty; aCnt++ {
		bx := 0
		by := 0
		costb := 0
		for bCnt := 0; ax+bx <= tx && ay+by <= ty; bCnt++ {
			posX := ax + bx
			posY := ay + by
			cost := costa + costb
			if posX == tx && posY == ty && cost < minCost {
				minCost = cost
			}
			bx += m.bButton.Dx
			by += m.bButton.Dy
			costb++
			fmt.Print(".")
		}
		ax += m.aButton.Dx
		ay += m.aButton.Dy
		costa += 3
		fmt.Print(".")
	}
	fmt.Println()
	return minCost
}

func main() {
	lines := common.StartDay(13, "input_h")

	machines := readMachineDefinitions(lines)

	total := 0
	for _, machine := range machines {
		price := machine.minimumPrice(0)
		if price < IMPOSSIBLE_PRICE {
			total += price
		}
	}
	fmt.Println("Part1 : ", total)

	total = 0
	for _, machine := range machines {
		price := machine.minimumPrice(10000000000000)
		if price < IMPOSSIBLE_PRICE {
			total += price
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
