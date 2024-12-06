package day06

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func Part1() {
	matrix, startPos, err := readMatrixAndStartPos("day06/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(patrol(matrix, startPos))
}

func readMatrixAndStartPos(fileName string) ([][]rune, []int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowNum := 0
	matrix := [][]rune{}
	startPos := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		row := []rune(line)
		matrix = append(matrix, row)

		caretPos := slices.Index(row, '^')
		if caretPos != -1 {
			startPos = []int{rowNum, caretPos}
		}

		rowNum++
	}

	return matrix, startPos, nil
}

func patrol(matrix [][]rune, startPos []int) int {
	distinctVisits := 0
	maxRows := len(matrix)
	maxCols := len(matrix[0])
	row := startPos[0]
	col := startPos[1]
	direction := matrix[row][col]

	for row >= 0 && row < maxRows && col >= 0 && col < maxCols {
		if matrix[row][col] != 'X' {
			matrix[row][col] = 'X'
			distinctVisits++
		}

		if direction == '^' {
			if row > 0 && matrix[row-1][col] == '#' {
				direction = '>'
				col++
			} else {
				row--
			}
		} else if direction == '>' {
			if col < maxCols-1 && matrix[row][col+1] == '#' {
				direction = 'v'
				row++
			} else {
				col++
			}
		} else if direction == 'v' {
			if row < maxRows-1 && matrix[row+1][col] == '#' {
				direction = '<'
				col--
			} else {
				row++
			}
		} else if direction == '<' {
			if col > 0 && matrix[row][col-1] == '#' {
				direction = '^'
				row--
			} else {
				col--
			}
		}
	}

	return distinctVisits
}
