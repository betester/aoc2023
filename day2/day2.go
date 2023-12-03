package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/betester/aoc2023/utils"
	"golang.org/x/exp/slices"
)

type Game struct {
	Id   int
	Cube CubeInformation
}

type CubeInformation struct {
	ColorCountOccurences map[string][]int
}

func totalValidGame(games []Game) int {
	total := 0

	gameRule := make(map[string]int)
	gameRule["red"] = 12
	gameRule["green"] = 13
	gameRule["blue"] = 14

	isValidGame := func(occurences []int, max int) bool {
		allValid := true
		for _, num := range occurences {
			allValid = allValid && num <= max
		}
		return allValid
	}

	for _, game := range games {
		cube := game.Cube
		allColorIsValid := true
		for color, occurence := range gameRule {
			colorOcurreces := cube.ColorCountOccurences[color]
			allColorIsValid = allColorIsValid && isValidGame(colorOcurreces, occurence)
		}
		if allColorIsValid {
			total += game.Id
		}
	}

	return total
}

func totalPower(games []Game) int {
	colors := make([]string, 0)
	colors = append(colors, "red")
	colors = append(colors, "blue")
	colors = append(colors, "green")
	total := 0
	for _, game := range games {
		cube := game.Cube
		colorTotal := 1
		for _, color := range colors {
			occurences := cube.ColorCountOccurences[color]
			colorTotal *= slices.Max(occurences)
		}

		total += colorTotal

	}

	return total
}

func main() {
	games := make([]Game, 0)
	inputs := utils.FileReader("./day2/day2.txt")
	for _, input := range inputs {
		colorPattern := regexp.MustCompile(`(\d+) (\w+)`)
		gamePattern := regexp.MustCompile(`Game (\d+)`)
		gameMatch := gamePattern.FindStringSubmatch(input)
		colorMatch := colorPattern.FindAllString(input, -1)
		id, _ := strconv.Atoi(gameMatch[1])

		game := Game{
			Id: id,
		}

		cubeInformation := CubeInformation{
			ColorCountOccurences: make(map[string][]int),
		}

		for _, colorOccurence := range colorMatch {
			spaceSeparated := strings.Split(colorOccurence, " ")
			for i := 0; i < len(spaceSeparated)-1; i += 2 {
				occ, color := spaceSeparated[i], spaceSeparated[i+1]
				occurence, _ := strconv.Atoi(occ)

				if _, ok := cubeInformation.ColorCountOccurences[color]; !ok {
					cubeInformation.ColorCountOccurences[color] = make([]int, 0)
				}

				cubeInformation.ColorCountOccurences[color] = append(cubeInformation.ColorCountOccurences[color], occurence)
			}
		}

		game.Cube = cubeInformation
		games = append(games, game)
	}
	totalScore := totalPower(games)
	fmt.Println(totalScore)
}
