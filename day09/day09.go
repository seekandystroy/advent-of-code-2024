package day09

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func Part1() {
	blocks := readDiskToFragment("day09/input.txt")
	checksum := fragment(blocks)

	fmt.Println(checksum)
}

func Part2() {
	blocks, emptyBlocksIndices, numbersBySize := readDiskToCompact("day09/input.txt")
	checksum := compact(blocks, emptyBlocksIndices, numbersBySize)

	fmt.Println(checksum)
}

func readDiskToFragment(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	diskMap := scanner.Text()

	disk := []string{}
	for i, numRune := range diskMap {
		positions := int(numRune - '0')

		for p := 0; p < positions; p++ {
			if i%2 == 0 {
				disk = append(disk, strconv.Itoa(i/2))
			} else {
				disk = append(disk, ".")
			}
		}
	}

	return disk
}

func readDiskToCompact(fileName string) ([]string, map[int]int, map[string]int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	diskMap := scanner.Text()

	disk := []string{}
	emptyBlocksIndices := make(map[int]int)
	numbersBySize := make(map[string]int)
	idx := 0

	for i, posRune := range diskMap {
		positions := int(posRune - '0')

		if i%2 == 0 {
			numString := strconv.Itoa(i / 2)
			numbersBySize[numString] = positions

			for p := 0; p < positions; p++ {
				disk = append(disk, numString)
			}
		} else {
			emptyBlocksIndices[idx] = positions
			for p := 0; p < positions; p++ {
				disk = append(disk, ".")
			}
		}

		idx += positions
	}

	return disk, emptyBlocksIndices, numbersBySize
}

func fragment(blocks []string) int {
	leftPtr := 0
	rightPtr := len(blocks) - 1
	checksum := 0

	leftPtr, checksum = moveLeftPtr(leftPtr, blocks)
	rightPtr = moveRightPtr(rightPtr, blocks)
	for leftPtr < rightPtr {
		blocks[leftPtr] = blocks[rightPtr]
		blocks[rightPtr] = "."

		segSum := 0
		leftPtr, segSum = moveLeftPtr(leftPtr, blocks)
		rightPtr = moveRightPtr(rightPtr, blocks)
		checksum += segSum
	}

	return checksum
}

func moveLeftPtr(leftPtr int, blocks []string) (int, int) {
	checksum := 0

	for ; leftPtr < len(blocks) && blocks[leftPtr] != "."; leftPtr++ {
		num, err := strconv.Atoi(blocks[leftPtr])
		if err != nil {
			panic(err)
		}

		checksum += num * leftPtr
	}

	return leftPtr, checksum
}

func moveRightPtr(rightPtr int, blocks []string) int {
	for ; rightPtr > 0 && blocks[rightPtr] == "."; rightPtr-- {
	}

	return rightPtr
}

func compact(blocks []string, emptyBlocksIndices map[int]int, numbersBySize map[string]int) int {
	rightPtr := moveRightPtr(len(blocks)-1, blocks)
	curNumber := blocks[rightPtr]

	for curNumber != "0" {
		numSize := numbersBySize[curNumber]

		for _, idx := range sortedKeys(emptyBlocksIndices) {
			emptyBlocks := emptyBlocksIndices[idx]
			if idx < rightPtr && emptyBlocks >= numSize {
				for i := 0; i < numSize; i++ {
					blocks[idx+i] = curNumber
				}

				for i := rightPtr; i > rightPtr-numSize; i-- {
					blocks[i] = "."
				}

				delete(emptyBlocksIndices, idx)
				emptyBlocksIndices[idx+numSize] = emptyBlocks - numSize
				break
			}
		}

		rightPtr -= numSize
		if blocks[rightPtr] == "." {
			rightPtr = moveRightPtr(rightPtr, blocks)
		}
		curNumber = blocks[rightPtr]
	}

	checksum := 0
	for i, block := range blocks {
		if block != "." {
			num, err := strconv.Atoi(block)
			if err != nil {
				panic(err)
			}

			checksum += i * num
		}

	}
	return checksum
}

func sortedKeys(emptyBlocksIndices map[int]int) []int {
	keys := make([]int, 0, len(emptyBlocksIndices))
	for k := range emptyBlocksIndices {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	return keys
}
