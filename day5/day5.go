package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/betester/aoc2023/utils"
	"golang.org/x/exp/slices"
)

func toInt(input []string) []int {
	convertedInput := make([]int, 0)
	for _, strNum := range input {
		num, err := strconv.Atoi(strNum)
		if err == nil {
			convertedInput = append(convertedInput, num)
		}
	}

	return convertedInput
}

func getSeeds(input string) []int {
	seedsNumber := strings.Split(strings.Split(input, "seeds:")[1], " ")

	return toInt(seedsNumber)
}

func parseMapping(requirements map[string]map[string][][]int, inputs []string, i int) int {
	r := strings.Split(strings.Split(inputs[i], " ")[0], "-")
	requirements[r[0]] = make(map[string][][]int)
	requirements[r[0]][r[2]] = make([][]int, 0)
	i++
	for i < len(inputs) && inputs[i] != "" {
		requirements[r[0]][r[2]] = append(requirements[r[0]][r[2]], toInt(strings.Split(inputs[i], " ")))
		i++
	}

	return i
}

func parseInput(inputs []string) ([]int, map[string]map[string][][]int) {
	seeds := make([]int, 0)
	requirements := make(map[string]map[string][][]int, 0)
	i := 0
	for i < len(inputs) {
		if strings.HasPrefix(inputs[i], "seeds:") {
			seeds = getSeeds(inputs[i])
		} else if inputs[i] != "" {
			newI := parseMapping(requirements, inputs, i)
			i = newI
		}
		i++
	}
	return seeds, requirements
}

func binarySearch(val int, ranges [][]int) int {
	l, r := -1, len(ranges)

	for r-l > 1 {
		mid := (l + r) / 2
		sp, rng := ranges[mid][1], ranges[mid][2]
		if val >= sp+rng {
			l = mid
		} else if val < sp {
			r = mid
		} else {
			return mid
		}
	}

	return -1
}

func partA(almanac string, currentRequirements []int, requirements map[string]map[string][][]int) {

	for requirements[almanac] != nil {
		for nextAlmanac, ranges := range requirements[almanac] {

			sort.Slice(ranges, func(i, j int) bool {
				return ranges[i][1] < ranges[j][1]
			})

			nextRequirements := make([]int, 0)
			for _, req := range currentRequirements {
				foundPosition := binarySearch(req, ranges)
				if foundPosition > -1 {
					d, s := ranges[foundPosition][0], ranges[foundPosition][1]
					nextRequirements = append(nextRequirements, req-s+d)
				} else {
					nextRequirements = append(nextRequirements, req)
				}
			}
			currentRequirements = nextRequirements
			almanac = nextAlmanac
		}
	}

	fmt.Println(slices.Min(currentRequirements))

}

func min(arr []int, s int) int {
	min := arr[0]

	for i := 0; i < len(arr); i += s {
		if min > arr[i] {
			min = arr[i]
		}
	}

	return min
}

func partB(almanac string, currentRequirements []int, requirements map[string]map[string][][]int) {

	for requirements[almanac] != nil {
		for nextAlmanac, ranges := range requirements[almanac] {

			sort.Slice(ranges, func(i, j int) bool {
				return ranges[i][1] < ranges[j][1]
			})

			nextRequirements := make([]int, 0)
			for i := 0; i < len(currentRequirements); i += 2 {
				lb := currentRequirements[i]
				ub := lb + currentRequirements[i+1]

				lbPosition := binarySearch(lb, ranges)
				ubPosition := binarySearch(ub, ranges)

				if lbPosition > -1 {
					d, s := ranges[lbPosition][0], ranges[lbPosition][1]
					lb = lb - s + d
				}

				if ubPosition > -1 {
					d, s := ranges[ubPosition][0], ranges[ubPosition][1]
					ub = ub - s + d
					if lb > d && lb > ub {
						lb = d
					}
				}
				nextRequirements = append(nextRequirements, lb, ub-lb)
			}
			currentRequirements = nextRequirements
			almanac = nextAlmanac
		}
	}

	fmt.Println(min(currentRequirements, 2))

}

func main() {
	inputs := utils.FileReader("./day5/day5.txt")
	seeds, requirements := parseInput(inputs)
	start := time.Now()
	partB("seed", seeds, requirements)
	end := time.Since(start)
	fmt.Println(end)
}
