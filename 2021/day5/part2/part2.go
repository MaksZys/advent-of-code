package main

import (
	reader "adventOfCode2021/day5"
	"fmt"
	"math"
)

func createPoints(direction reader.Direction) []reader.Point {
	points := make([]reader.Point, 0)

	if direction.From.X == direction.To.X {
		diff := direction.To.Y - direction.From.Y
		absoluteDiff := int(math.Abs(float64(diff)))

		for i := 0; i <= absoluteDiff; i++ {
			points = append(points, reader.Point{
				X: direction.From.X,
				Y: direction.From.Y + calcNext(diff, i),
			})
		}
	} else if direction.From.Y == direction.To.Y {
		diff := direction.To.X - direction.From.X
		absoluteDiff := int(math.Abs(float64(diff)))

		for i := 0; i <= absoluteDiff; i++ {
			points = append(points, reader.Point{
				X: direction.From.X + calcNext(diff, i),
				Y: direction.From.Y,
			})
		}
	} else {
		diffX := direction.To.X - direction.From.X
		diffY := direction.To.Y - direction.From.Y

		absoluteDiff := int(math.Abs(float64(diffX)))

		for i := 0; i <= absoluteDiff; i++ {
			x := direction.From.X + calcNext(diffX, i)
			y := direction.From.Y + calcNext(diffY, i)

			points = append(points, reader.Point{
				X: x,
				Y: y,
			})
		}
	}

	return points
}

func calcNext(diff int, index int) int {
	if diff < 0 {
		return diff + index
	}

	return diff - index
}

func insertToDiagram(diagram *[][]int, direction reader.Direction) {
	points := createPoints(direction)

	for _, point := range points {
		(*diagram)[point.X][point.Y]++
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

		if maxY < direction.From.Y {
			maxY = direction.From.Y
		}
		if maxY < direction.To.Y {
			maxY = direction.To.Y
		}
	}

	maxX++
	maxY++
	diagram := make([][]int, 0)
	for i := 0; i < maxX; i++ {
		diagram = append(diagram, make([]int, maxY))
	}

	for _, direction := range directions {
		insertToDiagram(&diagram, direction)
	}

	// adventOfCode2021.Print2dArray(diagram)

	fmt.Printf("Result: %d \n", occurrencesBiggerThan2(&diagram))
}
