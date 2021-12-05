package main

import (
	"adventOfCode2021"
	"fmt"
	"strconv"
	"strings"
)

type submarinePosition struct {
	aim        int
	depth      int
	horizontal int
}

type record struct {
	direction string
	value     int
}

const FORWARD = "forward"
const DOWN = "down"
const UP = "up"

func parseToRecord(textRecords []string) []record {
	recordsArray := make([]record, len(textRecords))

	for index, value := range textRecords {
		splitString := strings.Split(value, " ")

		recordValue, err := strconv.Atoi(splitString[1])
		if err != nil {
			panic(err)
		}

		recordsArray[index] = record{
			value:     recordValue,
			direction: splitString[0],
		}
	}

	return recordsArray
}

func main() {
	file := adventOfCode2021.GetFile("/day2/input.txt")
	fileLines := adventOfCode2021.ScanLines(file)

	recordsArray := parseToRecord(fileLines)

	position := submarinePosition{
		aim:        0,
		depth:      0,
		horizontal: 0,
	}

	for _, step := range recordsArray {
		switch step.direction {
		case FORWARD:
			position.depth += step.value * position.aim
			position.horizontal += step.value
		case UP:
			position.aim -= step.value
		case DOWN:
			position.aim += step.value
		}

		fmt.Printf("Step: %s, value: %d. Position -> horizontal: %d, depth: %d, aim: %d \n",
			step.direction,
			step.value,
			position.horizontal,
			position.depth,
			position.aim,
		)
	}
	fmt.Printf("%d \n", position.horizontal*position.depth)

	defer file.Close()
}
