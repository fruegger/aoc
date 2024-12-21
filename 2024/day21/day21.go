package main

import (
	"advent/aoc/common"
	"advent/aoc/term"
	"fmt"
	"math"
)

type Keyboard struct {
	field []string
	sym   uint8
	moves map[uint8]map[uint8][]string
}

var doorMoves = map[uint8]map[uint8][]string{
	'A': {
		'A': {""},
		'0': {"<"},
		'1': {"^<<", "<^<"},
		'2': {"^<", "<^"},
		'3': {"^"},
		'4': {"^^<<", "^<^<", "<^^<", "<^<^"},
		'5': {"^^<", "^<^", "<^^"},
		'6': {"^^"},
		'7': {"^^^<<", "^^<^<", "^<^^<", "^<^<^", "^<<^^", "<^<^^", "<^^<^", "<^^^<"},
		'8': {"^^^<", "^^<^", "^<^^", "<^^^"},
		'9': {"^^^"},
	},
	'0': {
		'A': {">"},
		'0': {""},
		'1': {"^<"},
		'2': {"^"},
		'3': {"^>", ">^"},
		'4': {"^^<", "^<^"},
		'5': {"^^"},
		'6': {"^^>", "^>^", ">^>"},
		'7': {"^^^<", "^^<^", "^<^^"},
		'8': {"^^^"},
		'9': {"^^^>", "^^>^", "^>^^", ">^^^"},
	},
	'1': {
		'A': {">>v", ">v>"},
		'0': {">v"},
		'1': {""},
		'2': {">"},
		'3': {">>"},
		'4': {"^"},
		'5': {"^>", ">^"},
		'6': {"^>>", ">^>", ">>^"},
		'7': {"^^"},
		'8': {"^^>", "^>^", ">^^"},
		'9': {"^^>>", "^>^>", "^>>^", ">^>^", ">^>^"},
	},
	'2': {
		'A': {"v>", ">v"},
		'0': {"v"},
		'1': {"<"},
		'2': {""},
		'3': {">"},
		'4': {"^<", "<^"},
		'5': {"^"},
		'6': {"^>", ">^"},
		'7': {"^^<", "^<^", "<^^"},
		'8': {"^^"},
		'9': {"^^>", "^>^", ">^^"},
	},
	'3': {
		'A': {"v"},
		'0': {"v<", "<v"},
		'1': {"<<"},
		'2': {"<"},
		'3': {""},
		'4': {"^<<", "<^<", "<<^"},
		'5': {"^<", "<^"},
		'6': {"^"},
		'7': {"^^<<", "^<^<", "<^<^", "<<^^"},
		'8': {"^^<", "^<^", "<^^"},
		'9': {"^^"},
	},
	'4': {
		'A': {">>vv", ">v>v", ">vv>^", "v>v>"},
		'0': {">vv", "v>v"},
		'1': {"v"},
		'2': {">v", "v>"},
		'3': {">>v", ">v>", "v>>"},
		'4': {""},
		'5': {">"},
		'6': {">>"},
		'7': {"^"},
		'8': {"^>", ">^"},
		'9': {"^>>", "^>^", ">>^"},
	},
	'5': {
		'A': {">vv", "v>v", "vv>"},
		'0': {"vv"},
		'1': {"v<", "<v"},
		'2': {"v"},
		'3': {"v>", ">v"},
		'4': {"<"},
		'5': {""},
		'6': {">"},
		'7': {"^<", "<^"},
		'8': {"^"},
		'9': {"^>", ">^"},
	},
	'6': {
		'A': {"vv"},
		'0': {"vv<", "v<v", "<vv"},
		'1': {"v<<", "<v<", "<<v"},
		'2': {"v<", "<^"},
		'3': {"v"},
		'4': {"<<"},
		'5': {"<"},
		'6': {""},
		'7': {"^<<", "<^<", "<<^"},
		'8': {"^<", "<^"},
		'9': {"^"},
	},
	'7': {
		'A': {">>vvv", ">v>vv", ">vv>v", ">vvv>", "v>>vv", "v>v>v", "vv>v>", "vv>>v"},
		'0': {">vvv", "v>vv", "vv>v"},
		'1': {"vv"},
		'2': {"vv>", "v>v", ">vv"},
		'3': {"vv>>", "v>v>", "v>>v", ">v>v", ">vv>", ">>vv"},
		'4': {"v"},
		'5': {"v>", ">v"},
		'6': {"v>>", ">v>", ">>v"},
		'7': {""},
		'8': {">"},
		'9': {">>"},
	},
	'8': {
		'A': {"vvv>", "vv>v", "v>vv", ">vvv"},
		'0': {"vvv"},
		'1': {"vv<", "v<v", "<vv"},
		'2': {"vv"},
		'3': {"vv>", "v>v", ">vv"},
		'4': {"v<", "<v"},
		'5': {"v"},
		'6': {"v>", ">v"},
		'7': {"<"},
		'8': {""},
		'9': {">"},
	},
	'9': {
		'A': {"vvv"},
		'0': {"vvv<", "vv<v", "v<vv", "<vvv"},
		'1': {"vv<<", "v<v<", "v<<v", "<vv<", "<v<v", "<<vv"},
		'2': {"vv<", "v<v", "<vv"},
		'3': {"vv"},
		'4': {"v<<", "<v<", "<<v"},
		'5': {"v<", "<v"},
		'6': {"v"},
		'7': {"<<"},
		'8': {"<"},
		'9': {""},
	},
}

var robotMoves = map[uint8]map[uint8][]string{
	'A': {
		'A': {""},
		'^': {"<"},
		'<': {"v<<", "<v<"},
		'v': {"v<", "<v"},
		'>': {"v"},
	},
	'^': {
		'A': {">"},
		'^': {""},
		'<': {"v<"},
		'v': {"v"},
		'>': {"v>", ">v"},
	},
	'<': {
		'A': {">>^", ">^>"},
		'^': {">^"},
		'<': {""},
		'v': {">"},
		'>': {">>"},
	},
	'v': {
		'A': {">^", "^>"},
		'^': {"^"},
		'<': {"<"},
		'v': {""},
		'>': {">"},
	},
	'>': {
		'A': {"^"},
		'^': {"<^", "^<"},
		'<': {"<<"},
		'v': {"<"},
		'>': {""},
	},
}

func main() {
	lines := common.StartDay(21, "door_keypad")

	door := Keyboard{
		field: lines,
		sym:   'A',
		moves: doorMoves,
	}

	lines = common.ReadDayFile(21, "robot_keypad")

	var robots []Keyboard
	for i := 0; i < 25; i++ {
		robots = append(robots,
			Keyboard{
				field: lines,
				sym:   'A',
				moves: robotMoves,
			})
	}

	printKeyboards(door, robots)
	fmt.Println()

	lines = common.ReadDayFile(21, "input")
	total := 0
	for _, line := range lines {
		total += codeComplexity(line, door, robots, 2)
	}
	fmt.Println("Part1: ", total)

	total = 0
	for _, line := range lines {
		total += codeComplexity(line, door, robots, 25)
	}
	fmt.Println("Part2: ", total)

}

func codeComplexity(code string, door Keyboard, robots []Keyboard, chainLen int) int {
	codeval := common.StringToNum(code[:len(code)-1])
	result := door.enterCode(code)
	for i := 0; i < len(robots) && i < chainLen; i++ {
		var result2 []string
		for _, c := range result {
			part := robots[i].enterCode(c)
			for _, p := range part {
				result2 = append(result2, p)
			}
		}
		result = result2
	}
	minL := math.MaxInt
	var best string
	for _, r := range result {
		if len(r) < minL {
			minL = len(r)
			best = r
		}
	}

	resultLen := len(best)
	fmt.Println(resultLen, " * ", codeval, " - ", result)
	return codeval * resultLen
}

func (k *Keyboard) enterCode(code string) []string {
	var result []string
	var part []string
	for i := 0; i < len(code); i++ {
		part = k.pushKey(code[i])
		if len(result) == 0 {
			result = part
		} else {
			var result2 []string
			for _, r := range result {
				for _, p := range part {
					result2 = append(result2, r+p)
				}
			}
			result = result2
		}
	}
	return result
}

func (k *Keyboard) pushKey(sym uint8) []string {
	var result []string
	for _, variant := range k.moves[k.sym][sym] {
		result = append(result, variant+"A")
	}
	k.sym = sym
	return result
}

func printKeyboards(door Keyboard, robots []Keyboard) {
	fmt.Println("Door          " + term.YELLOW + "|" + term.WHITE + " Robot1        " + term.YELLOW + "|" + term.WHITE + " Robot2        " + term.YELLOW + "|" + term.WHITE + " Robot3")
	for y := 0; y < len(door.field); y++ {
		for x := 0; x < len(door.field[y]); x++ {
			fmt.Print(coloredSym(door.field[y][x], door.sym == door.field[y][x]))
		}
		for _, r := range robots {
			fmt.Print(term.YELLOW + " | " + term.WHITE)
			if y < len(r.field) {
				for x := 0; x < len(r.field[y]); x++ {
					fmt.Print(coloredSym(r.field[y][x], r.field[y][x] == r.sym))
				}
			} else {
				fmt.Print("             ")
			}
		}
		fmt.Println()
	}

}

func coloredSym(s uint8, highlight bool) string {
	if highlight {
		return term.YELLOW + string(s) + term.WHITE
	}
	if (s >= '0' && s <= '9') || s == '>' || s == '<' || s == '^' || s == 'v' {
		return term.BLUE + string(s) + term.WHITE
	} else if s == 'A' {
		return term.GREEN + string(s) + term.WHITE
	} else {
		return string(s)
	}
}

/*
029A: <vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A
029A- v<A<AA>>^A<AA<^A>Av<<A>>^A<A^Av<<A>>^AAv<A>A^A<A>Av<A<A>>^AAA<A<^A>A

980A: <v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A
980A- v<<A>>^AAA<A^Av<A<AA>>^AA<AA<^A>Av<A<A>>^AAA<A<^A>Av<<A>>^A<A>A

179A: <v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
179A- v<<A>>^Av<A<A>>^AA<AA<^A>Av<<A>>^AA<A^Av<<A>>^AA<A>Av<A<A>>^AAA<A<^A>A

456A: <v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A
456A- v<<A>>^AAv<A<A>>^AA<AA<^A>Av<<A>>^A<A>Av<<A>>^A<A>Av<A<A>>^AA<A<^A>A

379A: <v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
379A- v<<A>>^A<A^Av<<A>>^AAv<A<A>>^AA<AA<^A>Av<<A>>^AA<A>Av<A<A>>^AAA<A<^A>A
*/
