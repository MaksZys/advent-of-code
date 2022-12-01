package main

import (
	"adventOfCode2021/day4/reader"
	"fmt"
)

func main() {
	bingoBoards, bingoValues := reader.ReadFromFile("/day4/input.txt")
	// bingoBoards, bingoValues := reader.ReadFromFile("/../../day4/test.txt")

	var bingoBoard *reader.BingoBoard
	var lastValue uint16

	for _, val := range bingoValues {
		for i := 0; i < len(bingoBoards); i++ {
			if !(bingoBoards[i]).AlreadyWon() && (bingoBoards[i]).MarkValue(val) {
				bingoBoards[i].Bingo()
				lastValue = val
				bingoBoard = &bingoBoards[i]
			}
		}
	}

	unmarkedSum := (*bingoBoard).SumUnmarked()
	fmt.Printf("Result: %d * %d = %d \n", lastValue, unmarkedSum, int(lastValue)*unmarkedSum)
}
