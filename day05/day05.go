package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day05/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules, err := readRules(scanner)
	if err != nil {
		fmt.Println(err)
	}
	correctUpdates, err := validateUpdates(scanner, rules)
	if err != nil {
		fmt.Println(err)
	}

	sumOfMidElements := 0
	for _, update := range correctUpdates {
		sumOfMidElements += update[len(update)/2]
	}

	fmt.Println(sumOfMidElements)
}

func readRules(scanner *bufio.Scanner) ([][]int, error) {
	rules := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		ruleStrings := strings.Split(line, "|")
		left, err := strconv.Atoi(ruleStrings[0])
		if err != nil {
			return rules, err
		}
		right, err := strconv.Atoi(ruleStrings[1])
		if err != nil {
			return rules, err
		}

		rules = append(rules, []int{left, right})
	}

	return rules, nil
}

func validateUpdates(scanner *bufio.Scanner, rules [][]int) ([][]int, error) {
	updates := [][]int{}

	for scanner.Scan() {
		pagesStrings := strings.Split(scanner.Text(), ",")
		update := []int{}
		for _, page := range pagesStrings {
			pageInt, err := strconv.Atoi(page)
			if err != nil {
				return [][]int{}, err
			}

			update = append(update, pageInt)
		}

		pageIndicesMap := map[int]int{}
		for i, page := range update {
			pageIndicesMap[page] = i
		}

		isUpdateValid := true
		for i := 0; i < len(rules) && isUpdateValid; i++ {
			left := rules[i][0]
			right := rules[i][1]
			leftIndex, okLeft := pageIndicesMap[left]
			rightIndex, okRight := pageIndicesMap[right]

			// fmt.Printf("Rule: %d|%d, indices: %d(%t) %d(%t)\n", left, right, leftIndex, okLeft, rightIndex, okRight)

			if okLeft && okRight && leftIndex > rightIndex {
				isUpdateValid = false
			}
		}

		if isUpdateValid {
			updates = append(updates, update)
		}
	}

	return updates, nil
}
