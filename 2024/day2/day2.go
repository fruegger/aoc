package main

import (
	"advent/common"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 2")
	fmt.Println("=====")

	file := common.OpenFile("./day2/day2_input.txt")
	defer file.Close()

	var reports [][]int

	scanReports(file, &reports)
	safeCount := countSafe(reports, isSafe)
	fmt.Println("Part 1 - Safe:", safeCount)
	safeCount = countSafe(reports, isSafeRelaxed)
	fmt.Println("Part 2 - Save:", safeCount)
}

func scanReports(file *os.File, reports *[][]int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var levels = strings.Split(line, " ")
		var report []int
		for _, level := range levels {
			report = append(report, common.StringToNum(level))
		}
		*reports = append(*reports, report)
	}
}

type safeFn func(report []int) bool

func countSafe(reports [][]int, safe safeFn) int {
	total := 0
	for _, report := range reports {
		if safe(report) {
			total++
		}
	}
	return total
}

func isSafe(report []int) bool {
	safe := true
	decreasing := report[0] > report[1]
	for i := 0; i+1 < len(report) && safe; i++ {
		diff := report[i+1] - report[i]
		if decreasing {
			diff = -diff
		}
		safe = diff >= 1 && diff <= 3
	}
	return safe
}

func isSafeRelaxed(report []int) bool {
	for j := 0; j < len(report); j++ {
		modified := removeLevel(report, j)
		if isSafe(modified) {
			return true
		}
	}
	return false
}

func removeLevel(report []int, pos int) []int {
	var result []int
	for i, v := range report {
		if i != pos {
			result = append(result, v)
		}
	}
	return result
}
