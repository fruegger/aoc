package main

import (
	"advent/common"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Day 1")
	fmt.Println("=====")

	file := common.OpenFile("./day1/day1_input.txt")
	defer file.Close()

	var leftElements []int
	var rightElements []int

	scanFile(file, &leftElements, &rightElements)
	sortArrays(leftElements, rightElements)

	total := sumDifferences(leftElements, rightElements)
	fmt.Println("Part 1:", total)

	occurrences := countOccurrences(leftElements, rightElements)
	total2 := 0

	for i := 0; i < len(leftElements); i++ {
		total2 = total2 + leftElements[i]*occurrences[i]
	}
	fmt.Println("Part 2:", total2)
}

func scanFile(file *os.File, leftElements *[]int, rightElements *[]int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var left int
		var right int
		line := scanner.Text()
		left, right = process(line)
		*leftElements = append(*leftElements, left)
		*rightElements = append(*rightElements, right)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func process(line string) (int, int) {
	var pair = strings.Split(line, "   ")
	return common.StringToNum(pair[0]), common.StringToNum(pair[1])
}

func sortArrays(leftElements []int, rightElements []int) {
	sort.Slice(leftElements, func(i, j int) bool {
		return leftElements[i] < leftElements[j]
	})
	sort.Slice(rightElements, func(i, j int) bool {
		return rightElements[i] < rightElements[j]
	})
}

func sumDifferences(leftElements []int, rightElements []int) int {
	total := 0
	for i := 0; i < len(leftElements); i++ {
		diff := leftElements[i] - rightElements[i]
		if diff > 0 {
			total = total + diff
		} else {
			total = total - diff
		}
	}
	return total
}

func countOccurrences(leftElements []int, rightElements []int) []int {
	var result []int
	for _, l := range leftElements {
		tot := 0
		for _, r := range rightElements {
			if l == r {
				tot++
			}
		}
		result = append(result, tot)
	}
	return result
}
