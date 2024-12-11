package main

import (
	"advent/aoc/common"
	"fmt"
	"strings"
)

func main() {
	lines := common.StartDay(11, "test2")
	input := lines[0]
	var c BlinkStoneCompiler
	c.init()
	c.start(input)
	result := 0
	for i := 0; i < 25; i++ {
		for !c.scanner.Eof {
			c.parse()
		}
		result = len(strings.Split(c.backend.output, " "))
		c.start(c.backend.output)
	}
	fmt.Println("Part1 : ", result-1)
}
