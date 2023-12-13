package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/betester/aoc2023/utils"
)

func allZero(sequence []int) bool {
	for _, num := range sequence {
		if num != 0 {
			return false
		}
	}

	return true
}

func getSequenceDifference(sequence []int) []int {
	diffSeq := make([]int, len(sequence)-1)

	for i := 0; i < len(sequence)-1; i++ {
		diffSeq[i] = sequence[i+1] - sequence[i]
	}

	return diffSeq
}

func getListOfDiffSequences(sequence []int) [][]int {
	diffSequences := make([][]int, 0)
	diffSequences = append(diffSequences, sequence)

	i := 0

	for !allZero(diffSequences[i]) {
		diffSequences = append(diffSequences, getSequenceDifference(diffSequences[i]))
		i++
	}

	return diffSequences
}

func predictNextSequence(diffSeqeunces [][]int) int {

	diffSeqeunces[len(diffSeqeunces)-1] = append(diffSeqeunces[len(diffSeqeunces)-1], 0)
	for i := len(diffSeqeunces) - 2; i >= 0; i-- {
		diffSeqeunces[i] = append(diffSeqeunces[i], 0)
		j := len(diffSeqeunces[i]) - 1
		diffSeqeunces[i][j] = diffSeqeunces[i][j-1] + diffSeqeunces[i+1][j-1]
	}

	return diffSeqeunces[0][len(diffSeqeunces[0])-1]
}

func predictPreviousSequence(diffSeqeunces [][]int) int {

	diffSeqeunces[len(diffSeqeunces)-1] = append([]int{0}, diffSeqeunces[len(diffSeqeunces)-1]...)
	for i := len(diffSeqeunces) - 2; i >= 0; i-- {
		diffSeqeunces[i] = append([]int{0}, diffSeqeunces[i]...)
		diffSeqeunces[i][0] = diffSeqeunces[i][1] - diffSeqeunces[i+1][0]
	}

	return diffSeqeunces[0][0]
}

func parseInput(inputs []string) [][]int {
	sequences := make([][]int, len(inputs))

	for i, input := range inputs {
		sepSpace := strings.Split(input, " ")
		sequences[i] = make([]int, len(sepSpace))
		for j, s := range sepSpace {
			num, _ := strconv.Atoi(s)
			sequences[i][j] = num
		}
	}

	return sequences
}

func mirageMaintenance(sequences [][]int) int {
	total := 0

	for _, sequence := range sequences {
		diffSequence := getListOfDiffSequences(sequence)
		total += predictPreviousSequence(diffSequence)
	}

	return total
}

func main() {
	inputs := utils.FileReader("./day9/day9.txt")
	sequences := parseInput(inputs)
	fmt.Println(mirageMaintenance(sequences))
}
