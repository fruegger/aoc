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

	var result int
	total1 := 0
	total2 := 0
	for _, p := range patterns {
		fmt.Println(p)
		result = sequence(p + "*")
		if result > 0 {
			total1++
		}
		total2 += result
	}

	fmt.Println("Part1 : ", total1)
	fmt.Println("Part2 : ", total2)
}

var cache = map[string]int{}

func sequence(pattern string) int {
	res, found := cache[pattern]
	if found {
		return res
	}
	matches := resetMatchers()
	matchersLeft := len(matches)
	var result int
	for i := 0; i < len(pattern) && (matchersLeft > 0); i++ {
		for j := 0; j < len(towels) && (matchersLeft > 0); j++ {
			towel := towels[j] + "#"
			if matches[j] {
				if towel[i] != '#' {
					if towel[i] != pattern[i] {
						matches[j] = false
						matchersLeft--
					}
				} else {
					if pattern[i] != '*' {
						result2 := sequence(pattern[i:])
						result += result2
						matches[j] = false
						matchersLeft--

					} else {
						result++
						cache[pattern] = result
						return result
					}
				}
			}
		}
	}
	cache[pattern] = result
	return result
}

func resetMatchers() []bool {
	var result []bool
	for i := 0; i < len(towels); i++ {
		result = append(result, true)
	}
	return result
}
