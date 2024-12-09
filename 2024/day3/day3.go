package main

import (
	"advent/aoc/common"
	"fmt"
)

type Mul struct {
	source string
	opd1   int
	opd2   int
	result int
}

type ScanState uint8

const (
	SCAN_Start ScanState = iota
	SCAN_M
	SCAN_U
	SCAN_L
	SCAN_D
	SCAN_O
	SCAN_N
	SCAN_Quote
	SCAN_T
	SCAN_MulLeftPar
	SCAN_DoLeftPar
	SCAN_DontLeftPar
	SCAN_Comma
)
const PART1 = true
const PART2 = false

func main() {
	fmt.Println("Day 3")
	fmt.Println("=====")

	file := common.OpenFile("./day3/day3_input_h.txt")
	defer file.Close()
	var lines []string
	common.ScanLines(file, &lines)

	var validOps1 []Mul
	mulEnabled := true
	for _, v := range lines {
		mulEnabled = parseMul(v, &validOps1, mulEnabled, PART1)
	}
	total1 := calcTotal(validOps1)

	var validOps2 []Mul
	mulEnabled = true
	for _, v := range lines {
		mulEnabled = parseMul(v, &validOps2, mulEnabled, PART2)
	}
	total2 := calcTotal(validOps2)

	fmt.Println("Part1: ", total1)
	fmt.Println("Part2: ", total2)
}

func parseMul(line string, validOps *[]Mul, mulEnabled bool, part1 bool) bool {
	state := SCAN_Start
	var opd1 int
	var opd2 int
	var opcnt int
	var sourcestart int
	for i := 0; i < len(line); i++ {
		switch state {
		case SCAN_Start:
			sourcestart = i
			if line[i] == 'm' {
				state = SCAN_M
			} else if line[i] == 'd' {
				state = SCAN_D
			}
		case SCAN_M:
			if line[i] == 'u' {
				state = SCAN_U
			} else {
				state = SCAN_Start
			}
		case SCAN_U:
			if line[i] == 'l' {
				state = SCAN_L
			} else {
				state = SCAN_Start
			}
		case SCAN_L:
			if line[i] == '(' {
				state = SCAN_MulLeftPar
				opd1 = 0
				opd2 = 0
				opcnt = 0
			} else {
				state = SCAN_Start
			}
		case SCAN_MulLeftPar:
			if isDigit(line[i]) && opcnt < 3 {
				opd1 = opd1*10 + int(line[i]-'0')
				opcnt++
			} else if line[i] == ',' {
				state = SCAN_Comma
				opcnt = 0
			} else {
				state = SCAN_Start
			}
		case SCAN_Comma:
			if isDigit(line[i]) && opcnt < 3 {
				opd2 = opd2*10 + int(line[i]-'0')
				opcnt++
			} else {
				if line[i] == ')' {
					if mulEnabled || part1 {
						accepted := Mul{
							source: line[sourcestart : i+1],
							opd1:   opd1,
							opd2:   opd2,
							result: opd1 * opd2,
						}
						*validOps = append(*validOps, accepted)
					}
					opcnt = 0
				}
				state = SCAN_Start
			}
		case SCAN_D:
			if line[i] == 'o' {
				state = SCAN_O
			} else {
				state = SCAN_Start
			}
		case SCAN_O:
			switch line[i] {
			case '(':
				state = SCAN_DoLeftPar
			case 'n':
				state = SCAN_N
			default:
				state = SCAN_Start
			}
		case SCAN_N:
			if line[i] == '\'' {
				state = SCAN_Quote
			} else {
				state = SCAN_Start
			}
		case SCAN_Quote:
			if line[i] == 't' {
				state = SCAN_T
			} else {
				state = SCAN_Start
			}
		case SCAN_T:
			if line[i] == '(' {
				state = SCAN_DontLeftPar
			} else {
				state = SCAN_Start
			}
		case SCAN_DoLeftPar:
			if line[i] == ')' {
				mulEnabled = true
			}
			state = SCAN_Start
		case SCAN_DontLeftPar:
			if line[i] == ')' {
				mulEnabled = false
			}
			state = SCAN_Start
		default:
			state = SCAN_Start
		}
	}
	return mulEnabled
}

func isDigit(v uint8) bool { return v >= '0' && v <= '9' }

func calcTotal(validOps []Mul) int {
	total := 0
	for _, op := range validOps {
		total += op.result
		fmt.Println("mul(", op.opd1, ",", op.opd2, ") = ", op.result, " [", op.source, "]")
	}
	return total
}
