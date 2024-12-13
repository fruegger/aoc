package main

import (
	"advent/aoc/common"
	"advent/aoc/pos"
	"fmt"
	"math"
	"strings"
)

type Machine struct {
	aButton pos.Distance
	bButton pos.Distance
	price   pos.Position
}

const IMPOSSIBLE_PRICE = math.MaxInt

func gcd(a int, b int) int {
	var t int
	for b != 0 {
		t = b
		b = a % b
		a = t
	}
	return a
}

//returns gcd, x0 and y0)
func gcdExtended(a int, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	g, x1, y1 := gcdExtended(b%a, a)
	return g, y1 - (b/a)*x1, x1
}

func (m Machine) hasSolution(offset int) bool {
	gx := gcd(m.aButton.Dx, m.bButton.Dx)
	gy := gcd(m.aButton.Dy, m.bButton.Dy)
	return (m.price.X+offset)%gx == 0 && (m.price.Y+offset)%gy == 0
}

func (m Machine) minimumPrice(offset int) int {
	minprice := IMPOSSIBLE_PRICE
	gx, x0x, y0x := gcdExtended(m.aButton.Dx, m.bButton.Dx)
	cPrimex := (m.price.X + offset) / gx
	aPrimex := m.aButton.Dx / gx
	bPrimex := m.bButton.Dx / gx
	xpx := cPrimex * x0x
	ypx := cPrimex * y0x

	gy, x0y, y0y := gcdExtended(m.aButton.Dy, m.bButton.Dy)
	cPrimey := (m.price.Y + offset) / gy
	aPrimey := m.aButton.Dy / gy
	bPrimey := m.bButton.Dy / gy
	xpy := cPrimey * x0y
	ypy := cPrimey * y0y

	kminx := -xpx / bPrimex
	kmaxx := ypx / aPrimex

	if kmaxx < kminx {
		store := kmaxx
		kmaxx = kminx
		kminx = store
	}

	kminy := -xpy / bPrimey
	kmaxy := ypy / aPrimey
	if kmaxy < kminy {
		store := kmaxy
		kmaxy = kminy
		kminy = store
	}

	dkx := kmaxx - kminx
	dky := kmaxy - kmaxx
	if dkx < dky {
		for k := kminx; k < kmaxx; k++ {
			ax := xpx + k*bPrimex
			bx := ypx - k*aPrimex
			tx := ax*m.aButton.Dx + bx*m.bButton.Dx
			ty := ax*m.aButton.Dy + bx*m.bButton.Dy
			if ty == m.price.Y+offset {
				price := ax*3 + bx
				if minprice > price {
					minprice = price
				}
				fmt.Print(".", tx, ",", ty)
			}

		}
	} else {
		for k := kminy; k < kmaxy; k++ {
			ay := xpy + k*bPrimey
			by := ypy - k*aPrimey
			tx := ay*m.aButton.Dx + by*m.bButton.Dx
			ty := ay*m.aButton.Dy + by*m.bButton.Dy
			if tx == m.price.X+offset {
				price := ay*3 + by
				if minprice > price {
					minprice = price
				}
				fmt.Print(".", tx, ",", ty)
			}
		}
	}
	return minprice
}

/*
func (m Machine) minimumPrice(offset int) int {
	minCost := IMPOSSIBLE_PRICE
	costa := 0
	ax := 0
	ay := 0
	tx := m.price.X + offset
	ty := m.price.Y + offset

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
		}
		ax += m.aButton.Dx
		ay += m.aButton.Dy
		costa += 3
	}
	fmt.Println()
	return minCost
}
*/
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
