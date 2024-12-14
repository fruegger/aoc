package main

import (
	"advent/aoc/common"
	"fmt"
	"strings"
)

func main() {
	lines := common.StartDay(11, "input")
	input := scanInput(lines[0])
	initTenPow()

	var total uint = 0
	for _, v := range input {
		total += blink(v, 0, 25)
	}
	fmt.Println("Part1 : ", total)
	clear(cache)
	total = 0
	for _, v := range input {
		total += blink(v, 0, 75)
	}
	fmt.Println("Part2 : ", total)
}

type Pair struct {
	value uint
	level uint
}

var cache = map[Pair]uint{}

func blink(val uint, currentblink, maxblinks uint) uint {
	result, ok := cache[Pair{val, currentblink}]
	if ok {
		return result
	} else {
		if currentblink == maxblinks {
			return 1
		}
		if val == 0 {
			return blink(1, currentblink+1, maxblinks)
		} else {
			if isEvenLength(val) {
				d := divisorHalf(val)
				left := val / d
				right := val % d
				stones := blink(left, currentblink+1, maxblinks) + blink(right, currentblink+1, maxblinks)
				cache[Pair{val, currentblink}] = stones
				return stones
			} else {
				stones := blink(val*2024, currentblink+1, maxblinks)
				cache[Pair{val, currentblink}] = stones
				return stones
			}
		}
	}
}

func scanInput(line string) []uint {
	vals := strings.Split(line, " ")
	result := make([]uint, len(vals))
	for i := 0; i < len(vals); i++ {
		result[i] = uint(common.StringToNum(vals[i]))
	}
	return result
}

var tenpow []uint

func divisorHalf(n uint) uint {
	var i uint
	i = 1
	for n > tenpow[i] {
		i++
	}
	return tenpow[i>>1]
}

func isEvenLength(n uint) bool {
	var i uint = 0
	for n > tenpow[i] {
		i++
	}
	return i%2 == 1
}

func initTenPow() {
	var tp uint = 1
	for i := 0; i < 20; i++ {
		tenpow = append(tenpow, tp)
		tp *= 10
	}
}
