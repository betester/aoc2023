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

func contain(array []int, a int) bool {
	for _, num := range array {
		if num ==  a {
			return true;
		}
	}
	return false
}

func getAdjacentSymbols(matrix [][]byte, consecutiveNumbers [][]int, i, j int) [][]int {
	adjacentSymbols := make([][]int, 0)
	directions := [][]int{
		{-1, 0}, {-1, -1}, {0, -1}, {1, -1},
		{1, 0}, {1, 1}, {0, 1}, {-1, 1},
	}

	m, n := len(matrix), len(matrix[0]) // assuming matrix is non-empty

	for _, direction := range directions {
		a, b := direction[0], direction[1]
		newI, newJ := i+a, j+b

		// Check if the new coordinates are within the matrix bounds
		if newI >= 0 && newI < m && newJ >= 0 && newJ < n {
			if !inside(consecutiveNumbers, newI, newJ) && isSymbol(matrix[newI][newJ]) {
				adjacentSymbols = append(adjacentSymbols, []int{newI, newJ})
			}
		}
	}

	return adjacentSymbols 
}

func getNumber(matrix [][]byte, consecutiveNumbers [][]int) (int, [][]int) {
	numberRepr := 0
	adjacentSymbols := make([][]int, 0)
	for _, numberPos := range consecutiveNumbers {
		adjacentSymbols = append(adjacentSymbols, getAdjacentSymbols(matrix, consecutiveNumbers, numberPos[0], numberPos[1])...)
	}

	if len(adjacentSymbols) > 0 {
		for i := 0; i < len(consecutiveNumbers); i++ {
			a, b := consecutiveNumbers[i][0], consecutiveNumbers[i][1]
			muliplicationFactor := math.Pow(10, float64(len(consecutiveNumbers)-i-1))
			numberRepr += convert(matrix[a][b]) * int(muliplicationFactor)
		}
	}
	
	return numberRepr, adjacentSymbols
}

func addSymbolNumber(symbolsNumber map[int]map[int][]int, symbolPositions [][]int, number int) {

	for _, sp := range symbolPositions {
		i, j := sp[0], sp[1]
		if _, ok := symbolsNumber[i]; !ok {
			symbolsNumber[i] = make(map[int][]int)
		} 

		if _, ok := symbolsNumber[i][j] ;!ok {
			symbolsNumber[i][j] = make([]int, 0)
		}

		if !contain(symbolsNumber[i][j], number) {
			symbolsNumber[i][j] = append(symbolsNumber[i][j], number)
		}

	}
}

func getTotalByAdjacentSymbols(matrix [][]byte, n int) int {
	
	symbolsNumber := make(map[int]map[int][]int)

	for i := 0; i < len(matrix); i++ {
		consecutiveDigits := make([][]int, 0)
		for j := 0; j < len(matrix[i]); j++ {
			if isDigit(matrix[i][j]) {
				consecutiveDigits = append(consecutiveDigits, []int{i, j})
			} else {
				number, adjacentSymbols := getNumber(matrix, consecutiveDigits)
				addSymbolNumber(symbolsNumber, adjacentSymbols, number)
				consecutiveDigits = make([][]int, 0)
			}
		}
		number, adjacentSymbols := getNumber(matrix, consecutiveDigits)
		addSymbolNumber(symbolsNumber, adjacentSymbols, number)
	}
	total := 0


	for _, v := range symbolsNumber {
		for _, arr := range v {
			if len(arr) == n {
				ratio := 1
				for _, num := range arr {
					ratio *= num
				}
				fmt.Println()
				total += ratio
			}
		}
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

	fmt.Println(getTotalByAdjacentSymbols(matrix, 2))
}
