package main

import (
	"fmt"
	"os"
	"strconv"
)

func mapToInt(vs []string) []int {
	vsm := make([]int, len(vs))
	var err error

	for i, v := range vs {
		vsm[i], err = strconv.Atoi(v)
		if err != nil {
			fmt.Println("Unable to map string to int:", v)
			os.Exit(1)
		}
	}
	return vsm
}

func getDay01Part1Result(vs []string) int {
	items := mapToInt(vs)

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

func getDay01Part2Result(vs []string) int {
	items := mapToInt(vs)
	totalItems := len(items)
	numIncreases := 0
	previousValue := 0

	for i := range items {
		// Rolling sum of 3 items
		if i > 2 && i <= totalItems {
			a := items[i]
			b := items[i-1]
			c := items[i-2]
			value := a + b + c

			if value > previousValue {
				numIncreases++
			}

			previousValue = value
		}
	}

	return numIncreases
}
