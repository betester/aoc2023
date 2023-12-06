package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/betester/aoc2023/utils"
)

func toNum(input string) []int {
	arr := make([]int, 0)

	seperatedBySpace := strings.Split(input, " ")

	for _, s := range seperatedBySpace {
		num, err := strconv.Atoi(s)

		if err == nil {
			arr = append(arr, num)
		}
	}

	return arr
}

func concat(arr []int) int {
	num := ""

	for _, i := range arr {
		s := strconv.Itoa(i)
		num += s
	}

	numrepr, _ := strconv.Atoi(num)
	return numrepr
}

func parseInput(inputs []string) ([]int, []int) {

	time := toNum(inputs[0])
	distance := toNum(inputs[1])

	return time, distance
}

func partA(times, distances []int) int {

	multipliedTotal := 1
	for i := 0; i < len(times); i++ {
		total := 0
		for speed := 0; speed <= times[i]; speed++ {
			totalDistance := speed * (times[i] - speed)
			if totalDistance > distances[i] {
				total++
			}

		}

		multipliedTotal *= total
	}

	return multipliedTotal
}

func partB(times, distances []int) int {
	time := concat(times)
	distance := concat(distances)

	multipliedTotal := 1
	total := 0
	for speed := 14; speed <= time; speed++ {
		totalDistance := speed * (time - speed)
		if totalDistance > distance {
			total++
		}

	}

	multipliedTotal *= total

	return multipliedTotal

}

func main() {
	inputs := utils.FileReader("./day6/day6.txt")

	time, distance := parseInput(inputs)
	fmt.Println(partB(time, distance))
}
