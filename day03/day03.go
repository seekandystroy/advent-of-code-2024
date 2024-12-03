package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1() {
	sum, err := readMuls("day03/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
}

func readMuls(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	sumOfMuls := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
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

			sumOfMuls += frst * scnd
		}
	}

	return sumOfMuls, nil
}
