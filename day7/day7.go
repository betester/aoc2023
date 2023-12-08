package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/betester/aoc2023/utils"
)

// denotes the mapping for each of camel card based on their
// duplicate count sorted
const (
	FIVE_OF_KIND  = 5
	FOUR_OF_KIND  = 14
	FULL_HOUSE    = 23
	THREE_OF_KIND = 113
	TWO_PAIR      = 122
	ONE_PAIR      = 1112
	HIGH_CARD     = 11111
)

const (
	FIVE_OF_KIND_SCORE  = 7
	FOUR_OF_KIND_SCORE  = 6
	FULL_HOUSE_SCORE    = 5
	THREE_OF_KIND_SCORE = 4
	TWO_PAIR_SCORE      = 3
	ONE_PAIR_SCORE      = 2
	HIGH_CARD_SCORE     = 1
)

type Hand struct {
	Cards string
	Bid   int
}

func getHandScore() map[int]int {
	scoreMapping := make(map[int]int)
	scoreMapping[FIVE_OF_KIND] = FIVE_OF_KIND_SCORE
	scoreMapping[FOUR_OF_KIND] = FOUR_OF_KIND_SCORE
	scoreMapping[FULL_HOUSE] = FULL_HOUSE_SCORE
	scoreMapping[THREE_OF_KIND] = THREE_OF_KIND_SCORE
	scoreMapping[TWO_PAIR] = TWO_PAIR_SCORE
	scoreMapping[ONE_PAIR] = ONE_PAIR_SCORE
	scoreMapping[HIGH_CARD] = HIGH_CARD_SCORE

	return scoreMapping
}

func getCardScore(priorityNumbers []int) map[byte]int {
	cardScore := make(map[byte]int)
	cardTypes := "23456789TJQKA"

	for i := 0; i < len(priorityNumbers); i++ {
		cardScore[cardTypes[i]] = priorityNumbers[i]
	}

	return cardScore
}
func getPriorityScores() []int {
	priorityScores := make([]int, 13)

	for i := 1; i <= len(priorityScores); i++ {
		priorityScores[i-1] = i
	}
	return priorityScores
}

func parseInput(inputs []string) []Hand {
	hands := make([]Hand, 0)
	for _, input := range inputs {
		spaceSeparated := strings.Split(input, " ")
		card, bidStr := spaceSeparated[0], spaceSeparated[1]
		bid, _ := strconv.Atoi(bidStr)

		hands = append(hands, Hand{Cards: card, Bid: bid})
	}
	return hands
}

func toNum(numbers []int) int {
	numRepr := 0

	for i := 0; i < len(numbers); i++ {
		numRepr *= int(math.Pow10(int(math.Log10(float64(numbers[i])) + 1)))
		numRepr += numbers[i]
	}

	return numRepr
}

func wildCardJoker(duplicateMapCounter map[byte]int) {
	const JOKER = 'J'

	if _, ok := duplicateMapCounter[JOKER]; ok {
		var maxCard byte = 0
		var maxCardDuplicate int = 0

		for card, duplicateCount := range duplicateMapCounter {
			if card != JOKER && duplicateCount > maxCardDuplicate {
				maxCard = card
				maxCardDuplicate = duplicateCount
			}
		}

		// we can be sure that at least the minimum duplicate would be 1
		if maxCardDuplicate != 0 {
			duplicateMapCounter[maxCard] += duplicateMapCounter[JOKER]
			delete(duplicateMapCounter, JOKER)
		}
	}
}

func getHandScoreByDuplicateCount(duplicateCount map[byte]int, scoreMapping map[int]int) int {
	duplicateList := make([]int, 0)
	for _, v := range duplicateCount {
		duplicateList = append(duplicateList, v)
	}

	sort.Slice(duplicateList, func(i, j int) bool {
		return duplicateList[i] < duplicateList[j]
	})

	handNumberRepr := toNum(duplicateList)
	return scoreMapping[handNumberRepr]
}

func getDuplicateCount(s1 []byte, scoreMapping map[int]int) map[byte]int {

	duplicateMapCounter := make(map[byte]int, 0)

	for _, b := range s1 {
		if _, ok := duplicateMapCounter[b]; !ok {
			duplicateMapCounter[b] = 0
		}

		duplicateMapCounter[b]++
	}

	return duplicateMapCounter
}

func compareHandsPartA(s1, s2 string, handScoreMapping map[int]int, cardScoreMapping map[byte]int) bool {
	s1DuplicateMap := getDuplicateCount([]byte(s1), handScoreMapping)
	s2DuplicateMap := getDuplicateCount([]byte(s2), handScoreMapping)

	s1Score := getHandScoreByDuplicateCount(s1DuplicateMap, handScoreMapping)
	s2Score := getHandScoreByDuplicateCount(s2DuplicateMap, handScoreMapping)
	if s1Score == s2Score {
		for i := 0; i < len(s1); i++ {
			if cardScoreMapping[s1[i]] == cardScoreMapping[s2[i]] {
				continue
			}
			return cardScoreMapping[s1[i]] < cardScoreMapping[s2[i]]
		}
	}

	return s1Score < s2Score
}

func compareHandsPartB(s1, s2 string, handScoreMapping map[int]int, cardScoreMapping map[byte]int) bool {
	s1DuplicateMap := getDuplicateCount([]byte(s1), handScoreMapping)
	s2DuplicateMap := getDuplicateCount([]byte(s2), handScoreMapping)

	wildCardJoker(s1DuplicateMap)
	wildCardJoker(s2DuplicateMap)

	s1Score := getHandScoreByDuplicateCount(s1DuplicateMap, handScoreMapping)
	s2Score := getHandScoreByDuplicateCount(s2DuplicateMap, handScoreMapping)
	if s1Score == s2Score {
		for i := 0; i < len(s1); i++ {
			if cardScoreMapping[s1[i]] == cardScoreMapping[s2[i]] {
				continue
			}
			return cardScoreMapping[s1[i]] < cardScoreMapping[s2[i]]
		}
	}

	return s1Score < s2Score
}

func partA(hands []Hand) int {
	total, rank := 0, 1
	handScoreMapping := getHandScore()
	cardScoreMapping := getCardScore(getPriorityScores())

	sort.Slice(hands, func(i, j int) bool {
		return compareHandsPartA(hands[i].Cards, hands[j].Cards, handScoreMapping, cardScoreMapping)
	})

	for _, hand := range hands {
		total += rank * hand.Bid
		rank++
	}

	return total
}

func partB(hands []Hand) int {
	total, rank := 0, 1
	handScoreMapping := getHandScore()
	priorityScores := getPriorityScores()
	// modify joker to make it the lowest priority
	priorityScores[9] = 0
	cardScoreMapping := getCardScore(priorityScores)

	sort.Slice(hands, func(i, j int) bool {
		return compareHandsPartB(hands[i].Cards, hands[j].Cards, handScoreMapping, cardScoreMapping)
	})

	for _, hand := range hands {
		total += rank * hand.Bid
		rank++
	}

	return total
}

func main() {
	inputs := utils.FileReader("./day7/day7.txt")
	hands := parseInput(inputs)
	partAScore := partA(hands)
	partBScore := partB(hands)

	fmt.Println("Part A: ", partAScore)
	fmt.Println("Part B: ", partBScore)
}
