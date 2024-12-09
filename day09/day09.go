package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Part1() {
	blocks := readDisk("day09/input.txt")
	checksum := fragment((blocks))

	fmt.Println(checksum)
}

func readDisk(fileName string) []string {
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
