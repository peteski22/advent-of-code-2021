package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getFileContentsAsIntegers(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, line)
	}

	return lines, scanner.Err()
}

func GetDay01Part1Result(fileName string) int {
	items, err := getFileContentsAsIntegers(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	numIncreases := 0
	previousItem := 0

	for i, item := range items {
		if i > 0 && item > previousItem {
			numIncreases++
		}

		previousItem = item
	}

	return numIncreases
}

func GetDay01Part2Result(fileName string) int {
	items, err := getFileContentsAsIntegers(fileName)
	totalItems := len(items)
	numIncreases := 0
	previousValue := 0

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, _ := range items {
		// Rolling sum of 3 items
		if i > 2 && i <= totalItems {
			a := items[i]
			b := items[i - 1]
			c := items[i - 2]
			value := a + b + c

			if (value > previousValue) {
				numIncreases++
			}

			previousValue = value
		}
	}

	return numIncreases
}
