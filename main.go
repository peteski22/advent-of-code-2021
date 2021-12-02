package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	day01Part1Result := getDay01Part1Result("day_01_1.input")
	fmt.Printf("Day 01 - Part 1: %d\n", day01Part1Result)
	day01Part2Result := getDay02Part2Result("day_01_2.input")
	fmt.Printf("Day 01 - Part 2: %d\n", day01Part2Result)
}

func GetFileContents(fileName string) ([]int, error) {
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
