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

func Part2() {
	matrix, startPos, err := readMatrixAndStartPos("day06/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bruteForceCycleFinding(matrix, startPos))
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

		row, col, direction = moveOnce(matrix, row, col, direction)
	}

	return distinctVisits
}

func bruteForceCycleFinding(matrix [][]rune, startPos []int) int {
	cycles := 0
	maxRows := len(matrix)
	maxCols := len(matrix[0])

	for r := 0; r < maxRows; r++ {
		for c := 0; c < maxCols; c++ {
			if matrix[r][c] != '#' && !(r == startPos[0] && c == startPos[1]) {
				matrixWithNewO := deepCloneMatrix(matrix)
				matrixWithNewO[r][c] = '#'
				if tortoiseAndHare(matrixWithNewO, startPos) {
					cycles++
				}
			}
		}
	}

	return cycles
}
func tortoiseAndHare(matrix [][]rune, startPos []int) bool {
	maxRows := len(matrix)
	maxCols := len(matrix[0])
	slowRow := startPos[0]
	slowCol := startPos[1]
	slowDirection := '^'
	fastRow := startPos[0]
	fastCol := startPos[1]
	fastDirection := '^'

	for slowRow >= 0 && slowRow < maxRows && slowCol >= 0 && slowCol < maxCols &&
		(fastRow >= 0 && fastRow < maxRows && fastCol >= 0 && fastCol < maxCols) {

		slowRow, slowCol, slowDirection = moveOnce(matrix, slowRow, slowCol, slowDirection)
		fastRow, fastCol, fastDirection = moveOnce(matrix, fastRow, fastCol, fastDirection)
		fastRow, fastCol, fastDirection = moveOnce(matrix, fastRow, fastCol, fastDirection)

		if slowRow == fastRow && slowCol == fastCol && slowDirection == fastDirection {
			return true
		}
	}

	return false
}

func moveOnce(matrix [][]rune, row int, col int, direction rune) (int, int, rune) {
	maxRows := len(matrix)
	maxCols := len(matrix[0])

	if direction == '^' {
		if row > 0 && matrix[row-1][col] == '#' {
			direction = '>'
		} else {
			row--
		}
	} else if direction == '>' {
		if col < maxCols-1 && matrix[row][col+1] == '#' {
			direction = 'v'
		} else {
			col++
		}
	} else if direction == 'v' {
		if row < maxRows-1 && matrix[row+1][col] == '#' {
			direction = '<'
		} else {
			row++
		}
	} else if direction == '<' {
		if col > 0 && matrix[row][col-1] == '#' {
			direction = '^'
		} else {
			col--
		}
	}

	return row, col, direction
}

// yes this is AI-generated, I just spent an hour debugging because I trusted slices.Clone()
func deepCloneMatrix(matrix [][]rune) [][]rune {
	if matrix == nil {
		return nil
	}

	// Create a new slice with the same length as the original matrix
	clone := make([][]rune, len(matrix))

	// Iterate over each row in the original matrix
	for i, row := range matrix {
		if row == nil {
			clone[i] = nil
			continue
		}

		// Create a new slice for the row with the same length as the original row
		rowClone := make([]rune, len(row))

		// Copy the contents of the original row to the cloned row
		copy(rowClone, row)

		// Assign the cloned row to the corresponding position in the cloned matrix
		clone[i] = rowClone
	}

	return clone
}
