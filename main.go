package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var data []string

	// DAY 01
	data = readLines("input/day_01_1.input")
	fmt.Printf("Day 01 - Part 1: %d\n", getDay01Part1Result(data))

	data = readLines("input/day_01_2.input")
	fmt.Printf("Day 01 - Part 2: %d\n", getDay01Part2Result(data))

	// DAY 02
	data = readLines("input/day_02_1.input")
	fmt.Printf("Day 02 - Part 1: %d\n", getDay02Part1Result(data))

	data = readLines("input/day_02_1.input")
	fmt.Printf("Day 02 - Part 2: %d\n", getDay02Part2Result(data))

	// DAY 03
	data = readLines("input/day_03_1.input")
	fmt.Printf("Day 03 - Part 1: %d\n", getDay03Part1Result(data))

}

func readLines(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Unable to read file:", path)
	}
	defer f.Close()

	var l []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		l = append(l, s.Text())
	}
	return l
}
