package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1() {
	sum, err := sumAllMulsFromFile("day03/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
}

func Part2() {
	sum, err := sumEnabledMulsFromFile("day03/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
}

func sumAllMulsFromFile(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	sumOfMuls := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineSum, err := sumMulsFromLine(line)
		if err != nil {
			return 0, err
		}

		sumOfMuls += lineSum
	}

	return sumOfMuls, nil
}

func sumEnabledMulsFromFile(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	sumOfMuls := 0
	// Assumes input is all in one line - multiple lines break the assumption that the delimiters are ^ or do()/$ or don't()
	re := regexp.MustCompile(`(?:^|do\(\)).*?(?:mul\(\d+,\d+\))*?(?:$|don\'t\(\))`)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	matches := re.FindAllString(line, -1)

	for _, match := range matches {
		matchSum, err := sumMulsFromLine(match)
		if err != nil {
			return 0, err
		}

		sumOfMuls += matchSum
	}

	return sumOfMuls, nil
}

func sumMulsFromLine(line string) (int, error) {
	sum := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		frst, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}
		scnd, err := strconv.Atoi(match[2])
		if err != nil {
			return 0, err
		}

		sum += frst * scnd
	}
	return sum, nil
}
