package main

import (
	"fmt"
	"math"

	"github.com/betester/aoc2023/utils"
)

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func isSymbol(b byte) bool {
	return b != '.'
}

func convert(b byte) int {
	if isDigit(b) {
		return int(b - '0')
	}

	return -1
}

func inside(consecutiveNumbers [][]int, i, j int) bool {
	for _, cn := range consecutiveNumbers {
		if cn[0] == i && cn[1] == j {
			return true
		}
	}
	return false
}

func adjacentContainsSymbol(matrix [][]byte, consecutiveNumbers [][]int, i, j int) bool {
	directions := [][]int{
		{-1, 0}, {-1, -1}, {0, -1}, {1, -1},
		{1, 0}, {1, 1}, {0, 1}, {-1, 1},
	}

	m, n := len(matrix), len(matrix[0]) // assuming matrix is non-empty

	containsSymbol := false

	for _, direction := range directions {
		a, b := direction[0], direction[1]
		newI, newJ := i+a, j+b

		// Check if the new coordinates are within the matrix bounds
		if newI >= 0 && newI < m && newJ >= 0 && newJ < n {
			if !inside(consecutiveNumbers, newI, newJ) {
				containsSymbol = isSymbol(matrix[newI][newJ]) || containsSymbol
			}
		}
	}

	return containsSymbol
}

func getNumber(matrix [][]byte, consecutiveNumbers [][]int) int {
	numberRepr := 0
	containsSymbol := false
	for _, numberPos := range consecutiveNumbers {
		containsSymbol = containsSymbol || adjacentContainsSymbol(matrix, consecutiveNumbers, numberPos[0], numberPos[1])
	}

	if containsSymbol {
		for i := 0; i < len(consecutiveNumbers); i++ {
			a, b := consecutiveNumbers[i][0], consecutiveNumbers[i][1]
			muliplicationFactor := math.Pow(10, float64(len(consecutiveNumbers)-i-1))
			numberRepr += convert(matrix[a][b]) * int(muliplicationFactor)
		}
	}

	return numberRepr
}

func getTotal(matrix [][]byte) int {
	total := 0
	for i := 0; i < len(matrix); i++ {
		consecutiveDigits := make([][]int, 0)
		for j := 0; j < len(matrix[i]); j++ {
			if isDigit(matrix[i][j]) {
				consecutiveDigits = append(consecutiveDigits, []int{i, j})
			} else {
				total += getNumber(matrix, consecutiveDigits)
				consecutiveDigits = make([][]int, 0)
			}
		}
		total += getNumber(matrix, consecutiveDigits)
	}

	return total
}

func main() {
	inputs := utils.FileReader("./day3/day3.txt")
	matrix := make([][]byte, 0)

	for _, input := range inputs {
		row := []byte(input)
		matrix = append(matrix, row)
	}

	fmt.Println(getTotal(matrix))
}
