package common

import (
	"bufio"
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
