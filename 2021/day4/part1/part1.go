package main

import (
	"adventOfCode2021/day4/reader"
	"fmt"
	"os"
)

func main() {
	bingoBoards, bingoValues := reader.ReadFromFile("/day4/input.txt")
	// bingoBoards, bingoValues := reader.ReadFromFile("/../../day4/test.txt")

	for _, val := range bingoValues {
		for i := 0; i < len(bingoBoards); i++ {
			if (bingoBoards[i]).MarkValue(val) {
				unmarkedSum := bingoBoards[i].SumUnmarked()
				fmt.Printf("Result: %d * %d = %d \n", val, unmarkedSum, val*uint16(unmarkedSum))

				os.Exit(0)
			}
		}
	}
}
