package main

import (
	"fmt"
	"math"

	"github.com/betester/aoc2023/utils"
	"golang.org/x/exp/slices"
)

type Pipe struct {
	possibleMovement []string
	allowedDirection []string
}

func getMovement() map[string][]int {
	movements := make(map[string][]int)

	movements["N"] = []int{-1, 0}
	movements["S"] = []int{1, 0}
	movements["E"] = []int{0, 1}
	movements["W"] = []int{0, -1}

	return movements
}

func getPipeMapping() map[byte]Pipe {
	pipeMap := make(map[byte]Pipe)

	pipeMap['|'] = Pipe{
		possibleMovement: []string{"N", "S"},
		allowedDirection: []string{"N", "S"},
	}

	pipeMap['-'] = Pipe{
		possibleMovement: []string{"W", "E"},
		allowedDirection: []string{"W", "E"},
	}

	pipeMap['L'] = Pipe{
		possibleMovement: []string{"N", "E"},
		allowedDirection: []string{"S", "W"},
	}

	pipeMap['J'] = Pipe{
		possibleMovement: []string{"N", "W"},
		allowedDirection: []string{"S", "E"},
	}

	pipeMap['7'] = Pipe{
		possibleMovement: []string{"S", "W"},
		allowedDirection: []string{"N", "E"},
	}

	pipeMap['F'] = Pipe{
		possibleMovement: []string{"S", "E"},
		allowedDirection: []string{"N", "W"},
	}

	pipeMap['.'] = Pipe{
		possibleMovement: make([]string, 0),
		allowedDirection: make([]string, 0),
	}

	pipeMap['S'] = Pipe{
		possibleMovement: []string{"S", "E", "N", "W"},
		allowedDirection: []string{},
	}

	return pipeMap
}

func pipeMaze(inputs []string) int {

	n, m := len(inputs), len(inputs[0])
	movements := getMovement()
	pipeMapper := getPipeMapping()
	pipes := make([][]Pipe, n)
	totalDistances := make([][]int, n)
	var currentPipePosition []int

	for i := 0; i < len(inputs); i++ {
		pipes[i] = make([]Pipe, m)
		totalDistances[i] = make([]int, m)
		for j := 0; j < len(inputs[i]); j++ {
			pipes[i][j] = pipeMapper[inputs[i][j]]
			if inputs[i][j] == 'S' {
				currentPipePosition = []int{i, j}
			}
		}
	}

	total := 0
	queue := make([][]int, 0)
	queue = append(queue, currentPipePosition)

	for len(queue) != 0 {
		i, j := queue[0][0], queue[0][1]

		queue = queue[1:]
		currentPipe := pipes[i][j]

		for _, direction := range currentPipe.possibleMovement {
			nextMovement := movements[direction]
			x, y := i+nextMovement[0], j+nextMovement[1]
			if x >= 0 && y >= 0 && x < n && y < m && totalDistances[x][y] == 0 {
				nextPipe := pipes[x][y]
				if slices.Contains(nextPipe.allowedDirection, direction) {
					total++
					totalDistances[x][y] = totalDistances[i][j] + 1
					queue = append(queue, []int{x, y})
				}
			} else if totalDistances[x][y] != 0 {
				totalDistances[x][y] = int(
					math.Max(
						float64(totalDistances[x][y]),
						float64(totalDistances[i][j]+1),
					),
				)
			}
		}

	}

	otherTotal := 0
	for _, total := range totalDistances {
		otherTotal = int(math.Max(float64(slices.Max(total)), float64(otherTotal)))
	}

	return (total + 1) / 2
}

func main() {
	inputs := utils.FileReader("./day10/day10.txt")
	fmt.Println(pipeMaze(inputs))
}
