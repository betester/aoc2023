package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/betester/aoc2023/utils"
)

const (
	LEFT  = 'L'
	RIGHT = 'R'
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return (a * b) / gcd(a, b)
}

func lcmArray(numbers []int) int {
	result := 1

	for _, num := range numbers {
		result = lcm(num, result)
	}

	return result
}

func move(currentState string, graph map[string][]string, order byte) string {
	if order == LEFT {
		return graph[currentState][0]
	} else {
		return graph[currentState][1]
	}
}

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

func totalDistanceTraverse(
	order []byte,
	graph map[string][]string,
	initialState string,
	reachedFinalState func(string) bool,
) (int, string) {

	total := 0

	currentState := initialState

	for !reachedFinalState(currentState) {
		currentOrder := order[total%len(order)]
		move(currentState, graph, currentOrder)
		total++
	}

	return total, currentState
}

func reachedZZZState(state string) bool {
	return state == "ZZZ"
}

func reachedSuffixZState(state string) bool {
	return strings.HasSuffix(state, "Z")
}

type SuffixZState struct {
	Znode        string
	Distance     int
	ReachedCycle bool
}

func precomputeAllStateToGetSuffixZ(
	currentState string,
	order []byte,
	currentOrder int,
	graph map[string][]string,
	dp map[string][]*SuffixZState) *SuffixZState {

	nextState := move(currentState, graph, order[currentOrder%len(order)])
	if reachedSuffixZState(nextState) {
		return &SuffixZState{Znode: nextState, Distance: 1}
	} else {
		if _, ok := dp[currentState]; !ok {
			dp[currentState] = make([]*SuffixZState, len(order))
		}
		if dp[currentState][currentOrder%len(order)] == nil {
			nextStateZState := precomputeAllStateToGetSuffixZ(nextState, order, currentOrder+1, graph, dp)
			dp[currentState][currentOrder%len(order)] = &SuffixZState{Znode: nextStateZState.Znode, Distance: nextStateZState.Distance + 1}
		}
		return dp[currentState][currentOrder%len(order)]
	}
}

func multiSourceTotalDistanceTraverse(orders []byte, graph map[string][]string) int {
	dp := make(map[string][]*SuffixZState)
	initialStates := make([]SuffixZState, 0)

	for node := range graph {
		if strings.HasSuffix(node, "A") {
			initialStates = append(initialStates, SuffixZState{Znode: node, Distance: 0})
		}
	}
	for {
		for i, state := range initialStates {
			nextZState := precomputeAllStateToGetSuffixZ(state.Znode, orders, state.Distance, graph, dp)
			hasReachedZCycle := nextZState.Znode == initialStates[i].Znode
			initialStates[i] = SuffixZState{Znode: nextZState.Znode, Distance: nextZState.Distance + initialStates[i].Distance, ReachedCycle: hasReachedZCycle}
			if hasReachedZCycle {
				initialStates[i].Distance -= nextZState.Distance
			}
		}
		allStateReachCycle := true

		for i := 0; i < len(initialStates); i++ {
			allStateReachCycle = allStateReachCycle && initialStates[i].ReachedCycle
		}

		// honestly, this approach is lame
		if allStateReachCycle {
			fmt.Println(initialStates)
			distances := make([]int, 0)
			for _, dist := range initialStates {
				distances = append(distances, dist.Distance)
			}
			return lcmArray(distances)
		}
	}

}

func main() {
	inputs := utils.FileReader("./day8/day8.txt")
	order, graph := parseInput(inputs)

	fmt.Println(multiSourceTotalDistanceTraverse([]byte(order), graph))
}
