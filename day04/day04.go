package day04

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rows   int
	cols   int
	matrix [][]rune
)

func Part1() {
	var err error

	matrix, err = readMatrix("day04/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	rows = len(matrix)
	cols = len(matrix[0])
	totalXMAS := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 'X' {
				totalXMAS += countNWMatchXMAS(r, c)
				totalXMAS += countNMatchXMAS(r, c)
				totalXMAS += countNEMatchXMAS(r, c)
				totalXMAS += countEMatchXMAS(r, c)
				totalXMAS += countSEMatchXMAS(r, c)
				totalXMAS += countSMatchXMAS(r, c)
				totalXMAS += countSWMatchXMAS(r, c)
				totalXMAS += countWMatchXMAS(r, c)
			}
		}
	}

	fmt.Println(totalXMAS)
}

func Part2() {
	var err error

	matrix, err = readMatrix("day04/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	rows = len(matrix)
	cols = len(matrix[0])
	totalCrossMAS := 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if matrix[r][c] == 'A' {
				totalCrossMAS += countCrossMASMAS(r, c)
				totalCrossMAS += countCrossMASSAM(r, c)
				totalCrossMAS += countCrossSAMSAM(r, c)
				totalCrossMAS += countCrossSAMMAS(r, c)
			}
		}
	}

	fmt.Println(totalCrossMAS)
}

func readMatrix(fileName string) ([][]rune, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return [][]rune{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowNum := 0
	matrix := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()

		row := []rune(line)
		matrix = append(matrix, row)

		rowNum++
	}

	return matrix, nil
}

func countNWMatchXMAS(r int, c int) int {
	if r >= 3 && c >= 3 &&
		matrix[r-1][c-1] == 'M' &&
		matrix[r-2][c-2] == 'A' &&
		matrix[r-3][c-3] == 'S' {
		return 1
	}

	return 0
}

func countNMatchXMAS(r int, c int) int {
	if r >= 3 &&
		matrix[r-1][c] == 'M' &&
		matrix[r-2][c] == 'A' &&
		matrix[r-3][c] == 'S' {
		return 1
	}

	return 0
}

func countNEMatchXMAS(r int, c int) int {
	if r >= 3 && c <= cols-4 &&
		matrix[r-1][c+1] == 'M' &&
		matrix[r-2][c+2] == 'A' &&
		matrix[r-3][c+3] == 'S' {
		return 1
	}

	return 0
}

func countEMatchXMAS(r int, c int) int {
	if c <= cols-4 &&
		matrix[r][c+1] == 'M' &&
		matrix[r][c+2] == 'A' &&
		matrix[r][c+3] == 'S' {
		return 1
	}

	return 0
}

func countSEMatchXMAS(r int, c int) int {
	if r <= rows-4 && c <= cols-4 &&
		matrix[r+1][c+1] == 'M' &&
		matrix[r+2][c+2] == 'A' &&
		matrix[r+3][c+3] == 'S' {
		return 1
	}

	return 0
}

func countSMatchXMAS(r int, c int) int {
	if r <= rows-4 &&
		matrix[r+1][c] == 'M' &&
		matrix[r+2][c] == 'A' &&
		matrix[r+3][c] == 'S' {
		return 1
	}

	return 0
}

func countSWMatchXMAS(r int, c int) int {
	if r <= rows-4 && c >= 3 &&
		matrix[r+1][c-1] == 'M' &&
		matrix[r+2][c-2] == 'A' &&
		matrix[r+3][c-3] == 'S' {
		return 1
	}

	return 0
}

func countWMatchXMAS(r int, c int) int {
	if c >= 3 &&
		matrix[r][c-1] == 'M' &&
		matrix[r][c-2] == 'A' &&
		matrix[r][c-3] == 'S' {
		return 1
	}

	return 0
}

// M.S
// .A.
// M.S
func countCrossMASMAS(r int, c int) int {
	if matrix[r-1][c-1] == 'M' &&
		matrix[r+1][c-1] == 'M' &&
		matrix[r-1][c+1] == 'S' &&
		matrix[r+1][c+1] == 'S' {
		return 1
	}

	return 0
}

// M.M
// .A.
// S.S
func countCrossMASSAM(r int, c int) int {
	if matrix[r-1][c-1] == 'M' &&
		matrix[r+1][c-1] == 'S' &&
		matrix[r-1][c+1] == 'M' &&
		matrix[r+1][c+1] == 'S' {
		return 1
	}

	return 0
}

// S.M
// .A.
// S.M
func countCrossSAMSAM(r int, c int) int {
	if matrix[r-1][c-1] == 'S' &&
		matrix[r+1][c-1] == 'S' &&
		matrix[r-1][c+1] == 'M' &&
		matrix[r+1][c+1] == 'M' {
		return 1
	}

	return 0
}

// S.S
// .A.
// M.M
func countCrossSAMMAS(r int, c int) int {
	if matrix[r-1][c-1] == 'S' &&
		matrix[r+1][c-1] == 'M' &&
		matrix[r-1][c+1] == 'S' &&
		matrix[r+1][c+1] == 'M' {
		return 1
	}

	return 0
}
