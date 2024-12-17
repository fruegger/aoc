package main

import (
	"advent/aoc/common"
	"advent/aoc/ds"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	lines := common.StartDay(17, "input")
	var m ThreeBitMachine
	m.init(lines)
	m.printAsm()
	m.run(false)
	fmt.Println("Part 1: ")
	for _, m := range m.output {

		fmt.Print(m, ",")
	}
	fmt.Println()
	m.init(lines)
	done := false

	// assume the register is divided by 8 at every cycle and has no effect on the next
	// see analysis
	// search all matching ra s for any step; bsf - analog.
	results := ds.Queue[uint]{}
	results.Push(0)
	for !done {
		equal := false
		var value uint
		for !done && results.Pull(&value) {
			for cnt := 0; cnt < 8; cnt++ {
				m.init(lines)
				m.rA = value*8 + uint(cnt)
				m.run(false)

				equal = len(m.output) <= len(m.code)
				for i := 1; i <= len(m.output); i++ {
					equal = equal && m.code[len(m.code)-i] == m.output[len(m.output)-i]
				}
				if equal {
					results.Push(value*8 + uint(cnt))
				}
				done = len(m.output) > len(m.code)

			}
		}
	}
	var val uint = 0
	results.Min(&val, uintcompare)

	fmt.Println("Part 2: ", results)
	//	v := 216148338630335
	var v = val
	for {
		m.init(lines)
		m.rA = v
		m.run(false)

		equal := len(m.output) == len(m.code)
		for i := 1; i <= len(m.output); i++ {
			equal = equal && m.code[len(m.code)-i] == m.output[len(m.output)-i]
		}
		if equal {
			fmt.Println("Part3 : ", v)
		}
		v--
	}
	// we missed this one in part  2 -- 216148338630253
}

func uintcompare(v1, v2 uint) int {
	if v1 > v2 {
		return 1
	} else {
		if v1 < v2 {
			return -1
		} else {
			return 0
		}
	}
}

type Op struct {
	name          string
	sourceIsCombo bool
	exec          func(bm *ThreeBitMachine, op uint)
}

type ThreeBitMachine struct {
	rA     uint
	rB     uint
	rC     uint
	code   []uint8
	pc     int
	output []uint8
	opSet  map[uint8]Op
}

func (bm *ThreeBitMachine) run(debug bool) {
	buf := bufio.NewReader(os.Stdin)
	for bm.pc < len(bm.code) {
		opc := bm.code[bm.pc]
		var opd uint
		if bm.opSet[opc].sourceIsCombo {
			opd = bm.combo()
		} else {
			opd = bm.literal()
		}
		bm.opSet[opc].exec(bm, opd)
		bm.pc += 2
		if debug {
			_, err := buf.ReadBytes('\n')
			if err != nil {
			}
			bm.printAsm()
		}
	}
}

func (bm *ThreeBitMachine) init(lines []string) {
	bm.pc = 0
	bm.output = []uint8{}
	bm.code = []uint8{}
	bm.initOpSet()
	for _, v := range lines {
		if strings.Contains(v, "Register A:") {
			bm.rA = uint(common.StringToNum(strings.Split(v, ": ")[1]))
		} else {
			if strings.Contains(v, "Register B:") {
				bm.rB = uint(common.StringToNum(strings.Split(v, ": ")[1]))
			} else {
				if strings.Contains(v, "Register C:") {
					bm.rC = uint(common.StringToNum(strings.Split(v, ": ")[1]))
				} else {
					if strings.Contains(v, "Program:") {
						bytes := strings.Split(strings.Split(v, ": ")[1], ",")
						for _, b := range bytes {
							bm.code = append(bm.code, uint8(common.StringToNum(b)))
						}
					}
				}
			}
		}
	}
}

func (bm *ThreeBitMachine) initOpSet() {
	bm.opSet = make(map[uint8]Op)
	bm.opSet[0] = Op{
		name:          "adv",
		sourceIsCombo: true,
		exec: func(bm *ThreeBitMachine, op uint) {
			result := divOp(bm.rA, op)
			bm.rA = result
		},
	}
	bm.opSet[1] = Op{
		name:          "bxl",
		sourceIsCombo: false,
		exec: func(bm *ThreeBitMachine, op uint) {
			bm.rB = bm.rB ^ op
		},
	}
	bm.opSet[2] = Op{
		name:          "bst",
		sourceIsCombo: true,
		exec: func(bm *ThreeBitMachine, op uint) {
			bm.rB = op & 7
		},
	}
	bm.opSet[3] = Op{
		name:          "jnz",
		sourceIsCombo: false,
		exec: func(bm *ThreeBitMachine, op uint) {
			if bm.rA != 0 {
				bm.pc = int(op) - 2 // will be incremented
			}
		},
	}
	bm.opSet[4] = Op{
		name:          "bxc",
		sourceIsCombo: false,
		exec: func(bm *ThreeBitMachine, op uint) {
			bm.rB = bm.rC ^ bm.rB
		},
	}
	bm.opSet[5] = Op{
		name:          "out",
		sourceIsCombo: true,
		exec: func(bm *ThreeBitMachine, op uint) {
			bm.output = append(bm.output, uint8(op&7))
		},
	}
	bm.opSet[6] = Op{
		name:          "bdv",
		sourceIsCombo: true,
		exec: func(bm *ThreeBitMachine, op uint) {
			result := divOp(bm.rA, op)
			bm.rB = result
		},
	}
	bm.opSet[7] = Op{
		name:          "cdv",
		sourceIsCombo: true,
		exec: func(bm *ThreeBitMachine, op uint) {
			result := divOp(bm.rA, op)
			bm.rC = result
		},
	}
}

func divOp(o1 uint, o2 uint) uint {
	return uint(float64(o1) / math.Trunc(math.Pow(2.0, float64(o2))))
}

func (bm *ThreeBitMachine) combo() uint {
	opd := bm.operand()
	switch opd {
	case 0, 1, 2, 3:
		{
			return uint(opd)
		}
	case 4:
		return bm.rA
	case 5:
		return bm.rB
	case 6:
		return bm.rC
	default:
		//error
		return 0
	}
}

func (bm *ThreeBitMachine) literal() uint {
	return uint(bm.operand())
}

func (bm *ThreeBitMachine) operand() uint8 {
	return bm.code[bm.pc+1]
}

func (bm *ThreeBitMachine) printAsm() {
	fmt.Println("A:", bm.rA, " B:", bm.rB, " C:", bm.rC)
	fmt.Println("---------------------")

	for i := 0; i < len(bm.code); i += 2 {
		if bm.pc == i {
			fmt.Print(">")
		} else {
			fmt.Print(" ")
		}
		if i < 10 {
			fmt.Print(" ")
		}
		fmt.Print(i)
		fmt.Print(": ")
		fmt.Print(bm.opSet[bm.code[i]].name, " ")
		if bm.opSet[bm.code[i]].sourceIsCombo && bm.code[i+1] > 3 {
			fmt.Print("R" + string(bm.code[i+1]-4+'A'))
		} else {
			fmt.Print(bm.code[i+1])
		}
		fmt.Println()
	}
	fmt.Println()
}
