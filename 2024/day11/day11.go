package main

import (
	"advent/aoc/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := common.StartDay(11, "input")
	input := scanInput(lines[0])

	var total uint64 = 0
	for _, v := range input {
		total += blink(v, 0, 25)
	}
	fmt.Println("Part1 : ", total)
	//should be 175006
	clear(cache)
	total = 0
	for _, v := range input {
		total += blink(v, 0, 75)
	}
	fmt.Println("Part2 : ", total)
}

type Pair struct {
	value uint64
	level uint
}

var cache = map[Pair]uint64{}

func blink(val uint64, currentblink, maxblinks uint) uint64 {
	result, ok := cache[Pair{val, currentblink}]
	if ok {
		return result
	} else {
		if currentblink == maxblinks {
			return 1
		}
		if val == 0 {
			stones := blink(1, currentblink+1, maxblinks)
			return stones
		} else {
			st := strconv.Itoa(int(val))
			if len(st)%2 == 0 {
				stl := st[:len(st)/2]
				str := st[len(st)/2:]
				left, _ := strconv.Atoi(stl)
				right, _ := strconv.Atoi(str)
				stones := blink(uint64(left), currentblink+1, maxblinks) + blink(uint64(right), currentblink+1, maxblinks)
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

func scanInput(line string) []uint64 {
	vals := strings.Split(line, " ")
	result := make([]uint64, len(vals))
	for i := 0; i < len(vals); i++ {
		result[i] = uint64(common.StringToNum(vals[i]))
	}
	return result
}
