package main

import (
	"advent/aoc/common"
	"fmt"
)

const EMPTY_BLOCK = -1

func main() {

	lines := common.StartDay(9, "input")
	//let's try simply a list if file ids per block (free := -1)
	diskBlocks := decodeBlocks(lines[0])
	fmt.Println(diskBlocks)
	compactBlocks(&diskBlocks)
	fmt.Println(diskBlocks)
	checksum := computeChecksum(diskBlocks)
	fmt.Println("Part 1 : ", checksum)

	diskBlocks = decodeBlocks(lines[0])
	moveBlocks(&diskBlocks)
	fmt.Println(diskBlocks)
	checksum = computeChecksum(diskBlocks)
	fmt.Println("Part 2 : ", checksum)

}

func decodeBlocks(line string) []int {
	var diskBlocks []int
	fileId := 0
	isFile := true
	for _, nrBlocks := range line {
		var value int
		if isFile {
			value = fileId
		} else {
			value = EMPTY_BLOCK
		}
		for j := nrBlocks - '0'; j > 0; j-- {
			diskBlocks = append(diskBlocks, value)
		}
		if isFile {
			fileId++
		}
		isFile = !isFile
	}
	return diskBlocks
}

func compactBlocks(diskBlocks *[]int) {
	readPtr := len(*diskBlocks) - 1
	writePtr := 0
	for writePtr < readPtr {
		// find the next empty block
		for writePtr < readPtr && (*diskBlocks)[writePtr] >= 0 {
			writePtr++
		}
		//swap blocks
		if writePtr < readPtr {
			(*diskBlocks)[writePtr] = (*diskBlocks)[readPtr]
			(*diskBlocks)[readPtr] = EMPTY_BLOCK
		}
		writePtr++

		readPtr--
		// skip empty blocks
		for writePtr < readPtr && (*diskBlocks)[readPtr] < 0 {
			readPtr--
		}

	}
}

func moveBlocks(diskBlocks *[]int) {
	readPtr := len(*diskBlocks) - 1
	for readPtr > 0 {
		// calculate the size of the next file to move
		writePtr := 0
		blockSizeNeeded := 1
		for readPtr > 0 && (*diskBlocks)[readPtr] == (*diskBlocks)[readPtr-1] {
			blockSizeNeeded++
			readPtr--
		}

		// find the next empty block large enough
		blockSizeAvailable := 0
		for blockSizeAvailable < blockSizeNeeded && writePtr < readPtr {
			for writePtr < readPtr && (*diskBlocks)[writePtr] >= 0 {
				writePtr++
			}
			j := 0
			for ; writePtr < readPtr && blockSizeAvailable < blockSizeNeeded && (*diskBlocks)[writePtr+j] < 0; j++ {
				blockSizeAvailable++
			}
			if blockSizeAvailable < blockSizeNeeded {
				writePtr += j
				blockSizeAvailable = 0
			}
		}
		if blockSizeAvailable == blockSizeNeeded {
			//swap blocks
			for i := 0; i < blockSizeNeeded; i++ {
				(*diskBlocks)[writePtr+i] = (*diskBlocks)[readPtr+i]
				(*diskBlocks)[readPtr+i] = EMPTY_BLOCK
			}
		}
		readPtr--
		// skip empty blocks
		for readPtr > 0 && (*diskBlocks)[readPtr] < 0 {
			readPtr--
		}

	}
}

func computeChecksum(diskBlocks []int) int {
	total := 0
	for i, c := range diskBlocks {
		if c > 0 {
			total += i * diskBlocks[i]
		}
	}
	return total
}
