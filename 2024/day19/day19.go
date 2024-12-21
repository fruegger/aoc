package main

import (
	"advent/aoc/common"
	"fmt"
	"strings"
)

var towels []string

func main() {
	lines := common.StartDay(19, "test")
	towels = strings.Split(lines[0], ", ")
	var patterns []string
	for i := 2; i < len(lines); i++ {
		patterns = append(patterns, lines[i])
	}

	var ok bool
	total := 0
	for _, p := range patterns {
		var result []int
		fmt.Println(p, "---")
		result, ok = sequence(p+"*", 0, result)
		if ok {
			total++
		}
	}

	fmt.Println("Part1 : ", total)
}

func sequence(pattern string, pos int, accepted []int) ([]int, bool) {
	matches := resetMatchers()
	matchersLeft := len(matches)

	for i := 0; i+pos < len(pattern) && (matchersLeft > 0); i++ {
		for j := 0; j < len(towels) && (matchersLeft > 0); j++ {
			towel := towels[j]
			if matches[j] {
				if i < len(towel) {
					if towel[i] != pattern[i+pos] {
						matches[j] = false
						matchersLeft--
					}
				} else {
					wasaccept := accepted
					accepted = append(accepted, j)
					if pattern[pos+i] != '*' {
						accepted, matches[j] = sequence(pattern, pos+i, accepted)
						if !matches[j] && len(accepted) > 0 {
							accepted = wasaccept
						} else {
							if pattern[pos+i] != '*' {
								return wasaccept, true
							}
						}
					} else {
						return accepted, true
					}
					//todo check?
					matches[j] = false
					matchersLeft--
				}
			}
		}
	}
	return accepted, matchersLeft > 0
}

func resetMatchers() []bool {
	var result []bool
	for i := 0; i < len(towels); i++ {
		result = append(result)
	}
	return result
}
