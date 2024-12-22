package main

import (
	"advent/aoc/common"
	"fmt"
	"strings"
)

var towels []string

func main() {
	lines := common.StartDay(19, "input")
	towels = strings.Split(lines[0], ", ")
	var patterns []string
	for i := 2; i < len(lines); i++ {
		patterns = append(patterns, lines[i])
	}

	var result [][]int
	total1 := 0
	total2 := 0
	for _, p := range patterns {
		fmt.Println(p)
		var part []int
		result = sequence(p+"*", 0, &part)
		//		fmt.Println(result)
		if len(result) > 0 {
			total1++
		}
		total2 += len(result)
	}

	fmt.Println("Part1 : ", total1)
	fmt.Println("Part2 : ", total2)
}

func sequence(pattern string, pos int, accepted *[]int) [][]int {
	matches := resetMatchers()
	matchersLeft := len(matches)
	var result [][]int
	for i := 0; i+pos < len(pattern) && (matchersLeft > 0); i++ {
		for j := 0; j < len(towels) && (matchersLeft > 0); j++ {
			towel := towels[j] + "*"
			var accepted2 []int
			if matches[j] {
				if i < len(towel) {
					if towel[i] != '*' {
						if towel[i] != pattern[i+pos] {
							matches[j] = false
							matchersLeft--
						}
					} else {
						accepted2 = append(*accepted, j)
						if pattern[pos+i] != '*' {
							result2 := sequence(pattern, pos+i, &accepted2)
							for _, v := range result2 {
								result = append(result, v)
							}
							*accepted = accepted2
						} else {
							*accepted = []int{}
							result = append(result, accepted2)
							//matches[j] = false
							//matchersLeft--
							//							return result, true
						}
					}
				} else {
					matches[j] = false
					matchersLeft--
				}
			}
		}
	}
	return result
}

func resetMatchers() []bool {
	var result []bool
	for i := 0; i < len(towels); i++ {
		result = append(result, true)
	}
	return result
}
