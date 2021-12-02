package days

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type command struct {
	direction string
	unit int
}

func getFileContentsAsCommands(fileName string) ([]command, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var commands []command
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := regexp.MustCompile("\\s").Split(scanner.Text(), 2)
		amount, _ := strconv.Atoi(parts[1])
		if (err != nil) {
			return nil, err
		}

		commands = append(commands, command { direction: parts[0], unit: amount})
	}

	return commands, scanner.Err()
}

func GetDay02Part1Result(fileName string) int {
	commands, err := getFileContentsAsCommands(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	x := 0
	y := 0

	for _, c := range commands {
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

func GetDay02Part2Result(fileName string) int {
	commands, err := getFileContentsAsCommands(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	x := 0
	y := 0
	aim := 0

	for _, c := range commands {
		switch c.direction {
		case "forward":
			x = x + c.unit
			y = y + (aim * c.unit)
		case "up":
			aim =aim - c.unit
		case "down":
			aim = aim + c.unit
		}
	}

	return x * y
}