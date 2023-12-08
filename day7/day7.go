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

func getCardScore() map[byte]int {
	cardScore := make(map[byte]int)

	cardScore['A'] = 13
	cardScore['K'] = 12
	cardScore['Q'] = 11
	cardScore['J'] = 10
	cardScore['T'] = 9
	cardScore['9'] = 8
	cardScore['8'] = 7
	cardScore['7'] = 6
	cardScore['6'] = 5
	cardScore['5'] = 4
	cardScore['4'] = 3
	cardScore['3'] = 2
	cardScore['2'] = 1

	return cardScore
}

type Hand struct {
	Cards string
	Bid   int
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

func getHandPriority(s1 []byte, scoreMapping map[int]int) int {
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})

	duplicateCounter := make([]int, 0)
	currDuplicate, total := s1[0], 0

	for _, b := range s1 {
		if currDuplicate == b {
			total++
		} else {
			duplicateCounter = append(duplicateCounter, total)
			total = 1
			currDuplicate = b
		}
	}

	if total != 0 {
		duplicateCounter = append(duplicateCounter, total)
	}

	sort.Slice(duplicateCounter, func(i, j int) bool {
		return duplicateCounter[i] < duplicateCounter[j]
	})

	handNumberRepr := toNum(duplicateCounter)
	return scoreMapping[handNumberRepr]
}

func compareHands(s1, s2 string, handScoreMapping map[int]int, cardScoreMapping map[byte]int) bool {
	s1Score := getHandPriority([]byte(s1), handScoreMapping)
	s2Score := getHandPriority([]byte(s2), handScoreMapping)

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
	cardScoreMapping := getCardScore()

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i].Cards, hands[j].Cards, handScoreMapping, cardScoreMapping)
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
	fmt.Println(partA(hands))
}
