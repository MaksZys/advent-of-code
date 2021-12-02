package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	direction string
	value     int
}

type Position struct {
	forward int
	down    int
}

const FORWARD = "forward"
const DOWN = "down"
const UP = "up"

func main() {
	file, error := os.Open("/Users/maks/Dev/maks/Golang/adventOfCode/day2/input.txt")
	if error != nil {
		log.Fatalf("Unable to read file %s", error)
	}

	defer file.Close()

	var directions = make([]Record, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		var newRecord Record
		value, error := strconv.ParseUint(splitLine[1], 10, 64)
		if error != nil {
			panic(error)
		}

		newRecord.direction = splitLine[0]
		newRecord.value = int(value)

		directions = append(directions, newRecord)
	}

	var position = Position{
		forward: 0,
		down:    0,
	}

	for i := 0; i < len(directions); i++ {
		var step Record = directions[i]

		switch step.direction {
		case FORWARD:
			position.forward += step.value
		case UP:
			position.down -= step.value
		case DOWN:
			position.down += step.value
		}
	}

	fmt.Printf("Values are: forward -> %d down -> %d and multiplied is: %d",
		position.forward,
		position.down,
		position.forward*position.down,
	)

	//fmt.Println(directions)
}
