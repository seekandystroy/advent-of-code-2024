package day02

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
	reportsList, err := getListOfReports("day02/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	safeReportsCount := 0
	for i := 0; i < len(reportsList); i++ {
		if isReportSafe(reportsList[i]) {
			safeReportsCount += 1
		}
	}

	fmt.Println(safeReportsCount)
}

func Part2() {
	reportsList, err := getListOfReports("day02/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	safeReportsCount := 0
	for i := 0; i < len(reportsList); i++ {
		currReport := reportsList[i]

		if isReportSafe(currReport) {
			safeReportsCount += 1
		} else {
			for j := 0; j < len(currReport); j++ {
				reportWithoutJPosition := slices.Delete(slices.Clone(currReport), j, j+1)

				if isReportSafe(reportWithoutJPosition) {
					safeReportsCount += 1
					break
				}
			}
		}
	}

	fmt.Println(safeReportsCount)
}

func getListOfReports(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	reportsList := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()

		report, err := lineToNumbersList(line)
		if err != nil {
			return nil, err
		}

		reportsList = append(reportsList, report)
	}

	return reportsList, nil
}

func lineToNumbersList(line string) ([]int, error) {
	split := strings.Fields(line)

	numbers := []int{}

	for i := 0; i < len(split); i++ {
		number, err := strconv.Atoi(split[i])
		if err != nil {
			return []int{}, err
		}

		numbers = append(numbers, number)
	}
	return numbers, nil
}

func isReportSafe(report []int) bool {
	previousDiff := report[0] - report[1]

	if previousDiff == 0 || utils.Abs(previousDiff) > 3 {
		return false
	}

	for i := 1; i < len(report)-1; i++ {
		curr := report[i]
		next := report[i+1]

		diff := curr - next

		if diff == 0 ||
			((diff < 0) != (previousDiff < 0)) ||
			utils.Abs(diff) > 3 {
			return false
		}

		previousDiff = diff
	}
	return true
}
