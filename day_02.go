package main

import (
	"regexp"
	"strconv"
)

type command struct {
	direction string
	unit      int
}

func mapToCmds(vs []string) []command {
	vsm := make([]command, len(vs))

	for i, v := range vs {
		p := regexp.MustCompile("\\s").Split(v, 2)
		a, _ := strconv.Atoi(p[1])
		vsm[i] = command{direction: p[0], unit: a}
	}

	return vsm
}

func getDay02Part1Result(vs []string) int {
	cmds := mapToCmds(vs)

	x := 0
	y := 0

	for _, c := range cmds {
		switch c.direction {
		case "forward":
			x = x + c.unit
		case "up":
			y = y - c.unit
		case "down":
			y = y + c.unit
		}
	}

	return x * y
}

func getDay02Part2Result(vs []string) int {
	cmds := mapToCmds(vs)

	x := 0
	y := 0
	aim := 0

	for _, c := range cmds {
		switch c.direction {
		case "forward":
			x = x + c.unit
			y = y + (aim * c.unit)
		case "up":
			aim = aim - c.unit
		case "down":
			aim = aim + c.unit
		}
	}

	return x * y
}
