package main

import (
	"adventOfCode2021"
	"fmt"
	"math"
	"strings"
)

func getMean(values []int) (int, int) {
	sum := float64(0)

	for _, val := range values {
		sum += float64(val)
	}

	length := float64(len(values))
	floorRes := int(math.Floor(float64(sum / length)))
	ceilRes := int(math.Ceil(float64(sum / length)))

	return floorRes, ceilRes
}

func calcDistance(value int, arithmeticPoint int) int {
	q := int(math.Abs(float64(value - arithmeticPoint)))

	distance := 0
	for i := 1; i <= q; i++ {
		distance += i
	}

	return distance
}

func calcFuelUsage(values []int, arithmeticPoint int) int {
	sum := 0

	for _, val := range values {
		sum += calcDistance(val, arithmeticPoint)
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

	floorMean, ceilMean := getMean(values)
	floorFuelUsage := calcFuelUsage(values, floorMean)
	ceilFuelUsage := calcFuelUsage(values, ceilMean)

	if floorFuelUsage < ceilFuelUsage {
		fmt.Println(floorFuelUsage)
	} else {
		fmt.Println(ceilFuelUsage)
	}
}

// incorrect
// too high: 98363819
