package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reportsList = [][]int{}
	for scanner.Scan() {
		line := scanner.Text()

		report, err := lineToNumbersList(line)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		reportsList = append(reportsList, report)
	}

	var safeReportsCount = 0
	for i := 0; i < len(reportsList); i++ {
		safe := isReportSafe(reportsList[i])

		if safe {
			safeReportsCount += 1
		}
	}

	fmt.Println(safeReportsCount)
}

func lineToNumbersList(line string) ([]int, error) {
	split := strings.Fields(line)

	var numbers = []int{}

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

	if previousDiff == 0 || abs(previousDiff) > 3 {
		return false
	}

	for i := 1; i < len(report)-1; i++ {
		var (
			curr = report[i]
			next = report[i+1]
		)

		diff := curr - next

		if diff == 0 ||
			((diff < 0) != (previousDiff < 0)) ||
			abs(diff) > 3 {
			return false
		}

		previousDiff = diff
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
