package utils

import (
	"bufio"
	"os"
)

func FileReader(filepath string) []string {

	file, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	array := make([]string, 0)
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}

	return array
}

func FileWriter(filepath string, output []byte) {
	os.WriteFile(filepath, output, 0644)
}
