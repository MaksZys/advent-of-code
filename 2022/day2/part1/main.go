package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculatePoints(round []string) int {
	moves := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	oponent := moves[round[0]]
	you := moves[round[1]]

	if you == oponent {
		return you + 3
	}

	if (oponent+1) == you || (oponent-2) == you {
		return you + 6
	}

	return you
}

func main() {
	//file, err := os.Open("./ownTests.txt")
	//file, err := os.Open("./test.txt")
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
