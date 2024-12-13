package main

import (
	"advent/aoc/common"
	"fmt"
)

func main() {
	lines := common.StartDay(11, "test2")
	input := lines[0]

	initProductions()
	initTenPow()

	var c BlinkStoneCompiler
	c.init()
	c.start(input)
	result := c.compile(6)
	fmt.Println("Part1 : ", result)

	c.init()
	c.start(input)
	result = c.compile(75)
	fmt.Println("Part2 : ", result)
}
