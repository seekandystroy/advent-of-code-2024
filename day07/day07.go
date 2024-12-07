package day07

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day07/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		testVal, numbersStack := getEquationFromLine(scanner.Text())

		if solvable(testVal, numbersStack) {
			sum += testVal
		}
	}

	fmt.Println(sum)
}

func Part2() {
	file, err := os.Open("day07/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		testVal, numbersStack := getEquationFromLine(scanner.Text())

		if solvableWithConcat(testVal, numbersStack) {
			sum += testVal
		}
	}

	fmt.Println(sum)
}

func getEquationFromLine(line string) (int, []int) {
	before, after, _ := strings.Cut(line, ":")

	result, err := strconv.Atoi(before)
	if err != nil {
		panic(err)
	}

	numbersStrings := strings.Fields(after)
	numbers := []int{}
	for _, str := range numbersStrings {
		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	slices.Reverse(numbers)
	return result, numbers
}

func solvable(testVal int, numbersStack []int) bool {
	length := len(numbersStack)

	if length == 1 {
		return testVal == numbersStack[0]
	} else {
		frst := numbersStack[length-1]
		scnd := numbersStack[length-2]

		mulStack := slices.Delete(slices.Clone(numbersStack), length-2, length)
		mulStack = append(mulStack, frst*scnd)

		sumStack := slices.Delete(slices.Clone(numbersStack), length-2, length)
		sumStack = append(sumStack, frst+scnd)

		return solvable(testVal, mulStack) || solvable(testVal, sumStack)
	}
}

func solvableWithConcat(testVal int, numbersStack []int) bool {
	length := len(numbersStack)

	if length == 1 {
		return testVal == numbersStack[0]
	} else {
		frst := numbersStack[length-1]
		scnd := numbersStack[length-2]

		mulStack := slices.Delete(slices.Clone(numbersStack), length-2, length)
		mulStack = append(mulStack, frst*scnd)

		sumStack := slices.Delete(slices.Clone(numbersStack), length-2, length)
		sumStack = append(sumStack, frst+scnd)

		concatStack := slices.Delete(slices.Clone(numbersStack), length-2, length)
		concatStack = append(concatStack, numberConcat(frst, scnd))

		return solvableWithConcat(testVal, mulStack) || solvableWithConcat(testVal, sumStack) || solvableWithConcat(testVal, concatStack)
	}
}

func numberConcat(frst int, scnd int) int {
	concat, err := strconv.Atoi(strconv.Itoa(frst) + strconv.Itoa(scnd))
	if err != nil {
		panic(err)
	}

	return concat
}
