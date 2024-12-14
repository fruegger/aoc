package main

import (
	"advent/aoc/intmath"
	"advent/aoc/pos"
	"fmt"
	"math"
)

type Machine struct {
	aButton pos.Distance
	bButton pos.Distance
	price   pos.Position
}

const IMPOSSIBLE_PRICE = math.MaxInt

func (m Machine) hasSolution(offset int) bool {
	gx := intmath.Gcd(m.aButton.Dx, m.bButton.Dx)
	gy := intmath.Gcd(m.aButton.Dy, m.bButton.Dy)
	return (m.price.X+offset)%gx == 0 && (m.price.Y+offset)%gy == 0
}

func (m Machine) minimumPrice(offset int) int {
	minprice := IMPOSSIBLE_PRICE
	gx, x0x, y0x := intmath.GcdExtended(m.aButton.Dx, m.bButton.Dx)
	cPrimex := (m.price.X + offset) / gx
	aPrimex := m.aButton.Dx / gx
	bPrimex := m.bButton.Dx / gx
	xpx := cPrimex * x0x
	ypx := cPrimex * y0x

	gy, x0y, y0y := intmath.GcdExtended(m.aButton.Dy, m.bButton.Dy)
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
