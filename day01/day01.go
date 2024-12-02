package day01

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var (
		leftList  = []int{}
		rightList = []int{}
	)

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Fields(line)
		first, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		second, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		leftList = append(leftList, first)
		rightList = append(rightList, second)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	var totalDistance = 0

	for i := 0; i < len(leftList); i++ {
		totalDistance += abs(leftList[i] - rightList[i])
	}

	fmt.Println(totalDistance)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
