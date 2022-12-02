package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculatePoints(round []string) int {
	const DRAW int = 3
	const WIN int = 6

	moves := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	oponent := moves[round[0]]
	you := round[1]

	if you == "Y" {
		return oponent + DRAW
	}

	if you == "Z" {
		if oponent == 3 {
			return 1 + WIN
		}

		return oponent + 1 + WIN
	}

	if oponent == 1 {
		return moves["C"]
	}

	return oponent - 1
}

func main() {
	//file, err := os.Open("./test.txt")
	//file, err := os.Open("./ownTests.txt")
	file, err := os.Open("./input.txt")
	if err != nil {
		os.Exit(1)
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")

		if len(values) != 2 {
			fmt.Println(values)
			os.Exit(2)
		}

		sum += calculatePoints(values)
	}

	fmt.Printf("Game result: %d \n", sum)
}
