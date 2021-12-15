package main

import (
	reader "adventOfCode2021/day5"
	"fmt"
	"math"
)

func calcNext(diff int, index int) int {
	if diff < 0 {
		return diff + index
	}

	return diff - index
}

func insertToDiagram(diagram *[][]int, direction reader.Direction) {
	x := direction.From.X
	y := direction.From.Y

	if direction.From.X == direction.To.X {
		diff := direction.To.Y - direction.From.Y
		absoluteDiff := int(math.Abs(float64(diff)))

		for i := 0; i <= absoluteDiff; i++ {
			val := calcNext(diff, i)
			(*diagram)[x][y+val]++
		}
	}

	if direction.From.Y == direction.To.Y {
		diff := direction.To.X - direction.From.X
		absoluteDiff := int(math.Abs(float64(diff)))

		for i := 0; i <= absoluteDiff; i++ {
			val := calcNext(diff, i)
			(*diagram)[x+val][y]++
		}
	}
}

func occurrencesBiggerThan2(diagram *[][]int) int {
	sum := 0

	for i := 0; i < len(*diagram); i++ {
		for j := 0; j < len((*diagram)[0]); j++ {
			if (*diagram)[i][j] >= 2 {
				sum++
			}
		}
	}

	return sum
}

func main() {
	directions := reader.GetDirectionsFromFile("/day5/input.txt")

	maxX, maxY := 0, 0
	for _, direction := range directions {
		if direction.From.X > maxX {
			maxX = direction.From.X
		}
		if direction.To.X > maxX {
			maxX = direction.To.X
		}

		if direction.From.Y > maxY {
			maxY = direction.From.Y
		}
		if direction.To.X > maxY {
			maxY = direction.To.Y
		}
	}

	maxX++
	maxY++
	diagram := make([][]int, 0)
	for i := 0; i < maxY; i++ {
		diagram = append(diagram, make([]int, maxY))
	}

	for _, direction := range directions {
		insertToDiagram(&diagram, direction)
	}

	fmt.Printf("Result: %d \n", occurrencesBiggerThan2(&diagram))
}
