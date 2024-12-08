package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func StringToNum(part string) int {
	result, _err := strconv.Atoi(part)
	if _err != nil {
		log.Fatal(_err)
	}
	return result
}

func ScanLines(file *os.File, lines *[]string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		*lines = append(*lines, line)
	}
}

func CopyLines(source []string) []string {
	var lines2 []string
	for _, v := range source {
		lines2 = append(lines2, v)
	}
	return lines2
}

func StartDay(day uint8, inputType string) []string {
	fmt.Println("Day ", day)
	fmt.Println("=====")

	filename := fmt.Sprintf("./day%d/day%d_%s.txt", day, day, inputType)
	file := OpenFile(filename)
	defer file.Close()
	var lines []string

	ScanLines(file, &lines)
	return lines
}
