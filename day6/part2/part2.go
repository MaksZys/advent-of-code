package main

import (
	"adventOfCode2021"
	storage "adventOfCode2021/day6/part2/storage"
	"fmt"
	"strings"
)

func calcLeterfishForXDays(values []int, iterations int) int {
	store := storage.Create(values)

	for i := 0; i < iterations; i++ {
		store.Next()
	}

	return store.Summary()
}

func main() {
	file := adventOfCode2021.GetFile("/day6/input.txt")
	line := adventOfCode2021.ScanLines(file)[0]

	lineValues := strings.Split(line, ",")
	values := make([]int, 0)

	for _, val := range lineValues {
		values = append(values, adventOfCode2021.ParseToInt(val))
	}

	fmt.Printf("\nResult: %d \n", calcLeterfishForXDays(values, 256))
}
