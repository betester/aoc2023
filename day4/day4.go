package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/betester/aoc2023/utils"
)

func winningScore(set1, set2 map[int]bool) int {
	score := 0
	for k, _ := range set1 {
		if _, ok := set2[k]; ok {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	return score
}

func getNumset(numbers []string) map[int]bool {
	numset := make(map[int]bool)

	for _, n := range numbers {
		n, err := strconv.Atoi(n)
		if err == nil {
			numset[n] = true
		}
	}

	return numset
}

func main() {
	inputs := utils.FileReader("./day4/day4.txt")
	total := 0
	for _, input := range inputs {
		games := strings.Split(input, "|")
		winningCards, cards :=
			getNumset(strings.Split(games[0], " ")[2:]),
			getNumset(strings.Split(games[1], " "))

		score := winningScore(cards, winningCards)
		total += score
	}

	fmt.Println(total)
}
