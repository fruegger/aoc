package main

import (
	"advent/aoc/common"
	"fmt"
	"sort"
	"strings"
)

// order[x][y] >0  -> x follows y
// order[x][y] <0  -> x precedes y
// order[x][y] =0  -> order is unknown
var order [100][100]int

func main() {
	fmt.Println("Day 5")
	fmt.Println("=====")

	file := common.OpenFile("./day5/day5_input_h.txt")
	defer file.Close()

	var lines []string
	common.ScanLines(file, &lines)

	cnt := readOrderRules(lines)
	updates := readUpdates(lines, cnt+1)

	alreadyOrdered := checkOrder(updates)
	total := sumCenters(updates, func(i int) bool { return alreadyOrdered[i] })
	fmt.Println("Part 1: ", total)

	for _, v := range updates {
		sort.Slice(v, func(i, j int) bool {
			return order[v[i]][v[j]] == -1
		})
	}
	total = sumCenters(updates, func(i int) bool { return !alreadyOrdered[i] })
	fmt.Println("Part 2: ", total)

}

func readOrderRules(lines []string) int {
	i := 0
	for ; i < len(lines) && len(lines[i]) > 0; i++ {
		parts := strings.Split(lines[i], "|")
		first := common.StringToNum(parts[0])
		second := common.StringToNum(parts[1])
		order[second][first] = 1
		order[first][second] = -1
	}
	return i
}

func readUpdates(lines []string, start int) [][]int {
	var updates [][]int
	for i := 0; i < len(lines)-start; i++ {
		var update []int
		values := strings.Split(lines[i+start], ",")
		for _, v := range values {
			update = append(update, common.StringToNum(v))
		}
		updates = append(updates, update)
	}
	return updates
}
func checkOrder(updates [][]int) []bool {
	result := make([]bool, len(updates))
	for i, update := range updates {
		ordered := true
		for i := 0; i+1 < len(update) && ordered; i++ {
			compare := order[update[i+1]][update[i]]
			ordered = compare > 0
		}
		result[i] = ordered
	}
	return result
}

func sumCenters(updates [][]int, isOrdered func(int) bool) int {
	total := 0
	for i, update := range updates {
		if isOrdered(i) {
			total += centerPage(update)
		}
	}
	return total
}

func centerPage(update []int) int {
	return update[len(update)/2]
}
