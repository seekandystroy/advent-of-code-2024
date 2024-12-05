package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	correctUpdates, err := validUpdates(scanner, rules)
	if err != nil {
		fmt.Println(err)
	}

	sumOfMidElements := 0
	for _, update := range correctUpdates {
		sumOfMidElements += update[len(update)/2]
	}

	fmt.Println(sumOfMidElements)
}

func Part2() {
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
	incorrectUpdates, err := invalidUpdates(scanner, rules)
	if err != nil {
		fmt.Println(err)
	}

	correctedUpdates := fixIncorrectUpdates(incorrectUpdates, rules)

	sumOfMidElements := 0
	for _, update := range correctedUpdates {
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

func validUpdates(scanner *bufio.Scanner, rules [][]int) ([][]int, error) {
	updates := [][]int{}

	for scanner.Scan() {
		update, err := getUpdateFromLine(scanner.Text())
		if err != nil {
			return [][]int{}, err
		}

		if isUpdateValid(update, rules) {
			updates = append(updates, update)
		}
	}

	return updates, nil
}

func invalidUpdates(scanner *bufio.Scanner, rules [][]int) ([][]int, error) {
	updates := [][]int{}

	for scanner.Scan() {
		update, err := getUpdateFromLine(scanner.Text())
		if err != nil {
			return [][]int{}, err
		}

		if !isUpdateValid(update, rules) {
			updates = append(updates, update)
		}
	}

	return updates, nil
}

func getUpdateFromLine(pagesString string) ([]int, error) {
	pagesStrings := strings.Split(pagesString, ",")

	update := []int{}
	for _, page := range pagesStrings {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return []int{}, err
		}

		update = append(update, pageInt)
	}

	return update, nil
}

func isUpdateValid(update []int, rules [][]int) bool {
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

	return isUpdateValid
}

func fixIncorrectUpdates(incorrectUpdates [][]int, rules [][]int) [][]int {
	correctedUpdates := [][]int{}
	for _, incorrectUpdate := range incorrectUpdates {
		pageIndicesMap := map[int]int{}
		for i, up := range incorrectUpdate {
			pageIndicesMap[up] = i
		}

		lefts := map[int][]int{}
		for _, rule := range rules {
			left := rule[0]
			right := rule[1]
			_, okLeft := pageIndicesMap[left]
			_, okRight := pageIndicesMap[right]

			if okLeft && okRight {
				lefts[right] = append(lefts[right], left)
			}
		}

		correctedUpdate := []int{}
		slicedMap := [][]int{}
		for k, v := range lefts {
			if len(v) == 1 {
				correctedUpdate = append(correctedUpdate, v[0])
			}

			slicedMap = append(slicedMap, append([]int{k}, v...))
		}
		slices.SortFunc(slicedMap, func(a, b []int) int {
			return len(a) - len(b)
		})

		for _, list := range slicedMap {
			correctedUpdate = append(correctedUpdate, list[0])
		}

		correctedUpdates = append(correctedUpdates, correctedUpdate)
	}

	return correctedUpdates
}
