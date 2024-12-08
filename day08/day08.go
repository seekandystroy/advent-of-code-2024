package day08

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() {
	matrix, antennas := readMatrix("day08/input.txt")

	fmt.Println(calculateAntinodes(matrix, antennas))
}

func readMatrix(fileName string) ([][]rune, map[rune][][]int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix := [][]rune{}
	antennas := make(map[rune][][]int)

	rowNum := 0
	for scanner.Scan() {
		line := scanner.Text()

		row := []rune(line)
		matrix = append(matrix, row)

		for colNum, char := range row {
			if char != '.' {
				antennas[char] = append(antennas[char], []int{rowNum, colNum})
			}
		}

		rowNum++
	}

	return matrix, antennas
}

func calculateAntinodes(matrix [][]rune, antennas map[rune][][]int) int {
	maxRows := len(matrix)
	maxCols := len(matrix[0])

	antinodesTotal := 0
	for _, v := range antennas {
		for idx, pos := range v {
			for i := idx + 1; i < len(v); i++ {
				preRow, preCol := antinode(pos, v[i])
				posRow, posCol := antinode(v[i], pos)

				if preRow >= 0 && preRow < maxRows && preCol >= 0 && preCol < maxCols && matrix[preRow][preCol] != '#' {
					matrix[preRow][preCol] = '#'
					antinodesTotal++
				}

				if posRow >= 0 && posRow < maxRows && posCol >= 0 && posCol < maxCols && matrix[posRow][posCol] != '#' {
					matrix[posRow][posCol] = '#'
					antinodesTotal++
				}

			}
		}
	}

	return antinodesTotal
}

func antinode(pos1 []int, pos2 []int) (int, int) {
	r1 := pos1[0]
	c1 := pos1[1]
	r2 := pos2[0]
	c2 := pos2[1]

	row := r1 - r2 + r1
	col := c1 - c2 + c1

	return row, col
}
