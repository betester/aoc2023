package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/betester/aoc2023/utils"
)

const (
	INITIAL_STATE = "AAA"
	FINAL_STATE   = "ZZZ"
	LEFT          = 'L'
	RIGHT         = 'R'
)

func parseInput(inputs []string) (string, map[string][]string) {
	order := inputs[0]
	nodePattern := regexp.MustCompile(`(\w+)`)
	graph := make(map[string][]string)

	for i := 2; i < len(inputs); i++ {
		nodeMatches := nodePattern.FindAllString(inputs[i], -1)
		sourceNode, leftNode, rightNode := nodeMatches[0], nodeMatches[1], nodeMatches[2]
		graph[sourceNode] = make([]string, 0)
		graph[sourceNode] = append(graph[sourceNode], leftNode, rightNode)
	}

	return order, graph
}

func totalDistanceTraverse(order []byte, graph map[string][]string) int {

	total := 0

	currentState := INITIAL_STATE

	for currentState != FINAL_STATE {
		leftState, rightState := graph[currentState][0], graph[currentState][1]
		currentOrder := order[total%len(order)]

		if currentOrder == LEFT {
			currentState = leftState
		} else {
			currentState = rightState
		}

		total++
	}

	return total
}

func multiSourceTotalDistanceTravel(order []byte, graph map[string][]string) int {

	totalSteps := 0
	startingStates := make([]string, 0)

	for node := range graph {
		if strings.HasSuffix(node, "A") {
			startingStates = append(startingStates, node)
		}
	}

	// this assumes that there is no cycle found
	for {
		allEndsWithZ := true

		for i := 0; i < len(startingStates); i++ {
			leftState, rightState := graph[startingStates[i]][0], graph[startingStates[i]][1]

			if order[totalSteps%len(order)] == LEFT {
				startingStates[i] = leftState
			} else {
				startingStates[i] = rightState
			}
			allEndsWithZ = allEndsWithZ && strings.HasSuffix(startingStates[i], "Z")
		}

		totalSteps++
		if allEndsWithZ {
			break
		}
	}
	return totalSteps
}

func main() {
	inputs := utils.FileReader("./day8/day8.txt")
	order, graph := parseInput(inputs)
	fmt.Println(totalDistanceTraverse([]byte(order), graph))
	fmt.Println(multiSourceTotalDistanceTravel([]byte(order), graph))
}
