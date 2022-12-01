package day5_reader

import (
	"adventOfCode2021"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Direction struct {
	From Point
	To   Point
}

func GetDirectionsFromFile(path string) []Direction {
	file := adventOfCode2021.GetFile(path)
	rows := adventOfCode2021.ScanLines(file)

	direction := make([]Direction, 0)

	for _, row := range rows {
		pointsString := strings.Split(row, " -> ")

		fromStringPoint := strings.Split(pointsString[0], ",")
		from := Point{
			X: adventOfCode2021.ParseToInt(fromStringPoint[0]),
			Y: adventOfCode2021.ParseToInt(fromStringPoint[1]),
		}

		toStringPoint := strings.Split(pointsString[1], ",")
		to := Point{
			X: adventOfCode2021.ParseToInt(toStringPoint[0]),
			Y: adventOfCode2021.ParseToInt(toStringPoint[1]),
		}

		direction = append(direction, Direction{
			From: from,
			To:   to,
		})
	}

	return direction
}
