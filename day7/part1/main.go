package main

import (
	"adventOfCode2021"
	"fmt"
	"math"
	"sort"
	"strings"
)

func getMedianValue(values []int) int {
	length := len(values)
	sort.Ints(values)

	if length%2 == 0 {
		return (values[length/2-1] + values[length/2]) / 2
	}

	return values[length/2] / 2
}

func calcUsage(values []int, point int) int {
	sum := 0

	for _, value := range values {
		sum += int(math.Abs(float64(value - point)))
	}

	return sum
}

func main() {
	file := adventOfCode2021.GetFile("/day7/input.txt")
	row := adventOfCode2021.ScanLines(file)[0]

	stringValues := strings.Split(row, ",")
	values := make([]int, 0)

	for _, val := range stringValues {
		values = append(values, adventOfCode2021.ParseToInt(val))
	}

	medianPoint := getMedianValue(values)
	fmt.Println(calcUsage(values, medianPoint))
}
