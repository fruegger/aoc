package main

import (
	"advent/aoc/common"
	"fmt"
	"strings"
)

func main() {
	lines := common.StartDay(7, "input")
	fmt.Println(lines)
	total := sumCorrectEquations(lines, false)
	fmt.Println("Part 1 :", total)

	total = sumCorrectEquations(lines, true)
	fmt.Println("Part 2 :", total)
}

func sumCorrectEquations(lines []string, considerConcat bool) int {
	total := 0
	for _, v := range lines {
		values, expectedResult := scanLine(v)
		if findOpCombination(values, expectedResult, considerConcat) {
			total += expectedResult
		}
	}
	return total
}

func scanLine(line string) ([]int, int) {
	equation := strings.Split(line, ": ")
	result := common.StringToNum(equation[0])
	operands := strings.Split(equation[1], " ")
	var values []int
	for _, operand := range operands {
		values = append(values, common.StringToNum(operand))
	}
	return values, result
}

func findOpCombination(values []int, expectedResult int, considerConcat bool) bool {
	incorrect := true
	for combination := 0; combination < 1<<((len(values)-1)<<1) && incorrect; combination++ {
		sum := values[0]
		for i := 0; i < len(values)-1; i++ {
			sum = applyOp((combination>>(i<<1))&3, sum, values[i+1], considerConcat)
		}
		incorrect = sum != expectedResult
	}
	return !incorrect
}

func applyOp(opId, opd1 int, opd2 int, considerConcat bool) int {
	var result int
	if opId == 1 {
		result = opd1 * opd2
	} else {
		if opId == 2 && considerConcat {
			result = concat(opd1, opd2)
		} else {
			result = opd1 + opd2
		}
	}
	return result
}

func concat(o1 int, o2 int) int {
	for decade := 10; ; decade *= 10 {
		if o2 < decade {
			return o1*decade + o2
		}
	}
}
