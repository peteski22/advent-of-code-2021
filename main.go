package main

import (
	"fmt"
	"advent-of-code-2021/days"
)

func main() {
	// DAY 01
	day01Part1Result := days.GetDay01Part1Result("days/day_01_1.input")
	fmt.Printf("Day 01 - Part 1: %d\n", day01Part1Result)
	day01Part2Result := days.GetDay01Part2Result("days/day_01_2.input")
	fmt.Printf("Day 01 - Part 2: %d\n", day01Part2Result)

	// DAY 02
	day02Part1Result := days.GetDay02Part1Result("days/day_02_1.input")
	fmt.Printf("Day 02 - Part 1: %d\n", day02Part1Result)
	day02Part2Result := days.GetDay02Part2Result("days/day_02_1.input")
	fmt.Printf("Day 02 - Part 2: %d\n", day02Part2Result)
}
