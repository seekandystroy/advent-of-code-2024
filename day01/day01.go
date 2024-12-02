package day01

import (
	"advent-of-code-2024/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	leftList, rightList, err := getListsFromFile("day01/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	totalDistance := 0

	for i := 0; i < len(leftList); i++ {
		totalDistance += utils.Abs(leftList[i] - rightList[i])
	}

	fmt.Println(totalDistance)
}

func Part2() {
	leftList, rightList, err := getListsFromFile("day01/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	rightOccurrences := make(map[int]int)

	for i := 0; i < len(rightList); i++ {
		val := rightList[i]
		_, exists := rightOccurrences[val]

		if !exists {
			rightOccurrences[val] = 1
		} else {
			rightOccurrences[val] += 1
		}
	}

	similarityScore := 0

	for i := 0; i < len(leftList); i++ {
		leftValue := leftList[i]
		occurences, exists := rightOccurrences[leftValue]

		if exists {
			similarityScore += leftValue * occurences
		}
	}

	fmt.Println(similarityScore)
}

func getListsFromFile(fileName string) ([]int, []int, error) {
	leftList := []int{}
	rightList := []int{}

	file, err := os.Open(fileName)
	if err != nil {
		return leftList, rightList, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Fields(line)
		first, err := strconv.Atoi(split[0])
		if err != nil {
			return leftList, rightList, err
		}

		second, err := strconv.Atoi(split[1])
		if err != nil {
			return leftList, rightList, err
		}

		leftList = append(leftList, first)
		rightList = append(rightList, second)
	}

	if err := scanner.Err(); err != nil {
		return leftList, rightList, err
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	return leftList, rightList, nil
}
