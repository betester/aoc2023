package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/betester/aoc2023/utils"
)

func getIntersection(set1, set2 map[int]bool) int {

	totalIntersection := 0

	for k, _ := range set1 {
		if _, ok := set2[k]; ok {
			totalIntersection++
		}
	}

	return totalIntersection
}

func winningScore(set1, set2 map[int]bool) int {
	sameCards := getIntersection(set1, set2)
	if sameCards == 0 {
		return 0
	}
	return int(math.Pow(2, float64(sameCards-1)))
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

func parseInput(inputs []string) ([]map[int]bool, []map[int]bool) {

	winningCards := make([]map[int]bool, 0)
	cards := make([]map[int]bool, 0)

	for _, input := range inputs {
		games := strings.Split(input, "|")
		winningCard, card :=
			getNumset(strings.Split(games[0], " ")[2:]),
			getNumset(strings.Split(games[1], " "))

		winningCards = append(winningCards, winningCard)
		cards = append(cards, card)
	}

	return winningCards, cards
}

func partA(inputs []string) int {
	total := 0

	winningCards, cards := parseInput(inputs)

	for i := 0; i < len(cards); i++ {
		total += winningScore(cards[i], winningCards[i])
	}

	return total
}

func sum(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}

	return total
}

func partB(inputs []string) int {
	cardsCopy := make([]int, len(inputs)+1)
	winningCards, cards := parseInput(inputs)

	for i := 1; i <= len(inputs); i++ {
		matchingCards := getIntersection(cards[i-1], winningCards[i-1])
		cardsCopy[i] += cardsCopy[i-1]

		if i+1 <= len(inputs) {
			cardsCopy[i+1] += cardsCopy[i] + 1
		}
		if matchingCards+1+i <= len(inputs) {
			cardsCopy[i+matchingCards+1] -= cardsCopy[i] + 1
		}
	}

	fmt.Println(cardsCopy)
	return sum(cardsCopy) + len(inputs)
}

func main() {
	inputs := utils.FileReader("./day4/day4.txt")
	fmt.Println(partB(inputs))
}
