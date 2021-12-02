package main

import (
	"fmt"
	"os"
)

func getDay01Part1Result(fileName string) int {
	items, err := GetFileContents(fileName)
	numIncreases := 0
	previousItem := 0

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, item := range items {
		if i > 0 && item > previousItem {
			numIncreases++
		}

		previousItem = item
	}

	return numIncreases
}

func getDay02Part2Result(fileName string) int {
	items, err := GetFileContents(fileName)
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
