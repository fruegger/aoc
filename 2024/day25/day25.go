package main

import (
	"advent/aoc/common"
	"fmt"
)

type Schematic = [5]int

func main() {
	lines := common.StartDay(25, "input")
	var keys []Schematic
	var locks []Schematic
	initKeysAndLocks(lines, &keys, &locks)
	fmt.Println("Locks")
	printSchematics(locks)
	fmt.Println("Keys")
	printSchematics(keys)
	fits := tryKeys(keys, locks)
	fmt.Println()
	fmt.Println("Part1: ", fits)

}

func initKeysAndLocks(lines []string, keys, locks *[]Schematic) {
	var current Schematic

	processingKey := false
	processingLock := false

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) != 0 {
			if lines[i] == "#####" {
				processingLock = true
				current = Schematic{}
			} else {
				if lines[i] == "....." {
					processingKey = true
					current = Schematic{}
				} else {
					fmt.Println("Error processing line:", lines[i])
				}
			}
		}
		i++
		for rowCnt := 0; rowCnt < 5; rowCnt++ {
			for i, c := range lines[i+rowCnt] {
				if c == '#' {
					current[i]++
				}
			}
		}
		i += 5
		if processingLock && lines[i] != "....." ||
			processingKey && processingLock && lines[i] != "#####" {
			fmt.Println("Error processing line:", lines[i])
		}
		i++
		if processingKey {
			*keys = append(*keys, current)
			processingKey = false
		} else {

			if processingLock {
				*locks = append(*locks, current)
				processingLock = false
			}
		}
	}
}

func printSchematics(schematics []Schematic) {
	for _, s := range schematics {
		fmt.Printf("%d,%d,%d,%d,%d", s[0], s[1], s[2], s[3], s[4])
		fmt.Println()
	}
}

func tryKeys(keys []Schematic, locks []Schematic) int {
	result := 0
	for _, k := range keys {
		for _, l := range locks {
			if fits(k, l) {
				result++
			}
		}
	}
	return result
}

func fits(k Schematic, l Schematic) bool {
	for i := 0; i < 5; i++ {
		if k[i]+l[i] > 5 {
			return false
		}
	}
	return true
}
