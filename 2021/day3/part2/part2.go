package main

import (
	"adventOfCode2021"
	"adventOfCode2021/day3/part2/helper"
	"adventOfCode2021/day3/part2/unit"
	"fmt"
)

const oxygen_generator = "oxygenGenerator"
const co2_scrubber = "co2Scrubber"

func main() {
	file := adventOfCode2021.GetFile("/day3/input.txt")
	fileRecords := adventOfCode2021.ScanLines(file)

	defer file.Close()

	oxygenGeneratorUnit := unit.Unit{
		Name:     oxygen_generator,
		BitValue: "",
	}
	oxygenGeneratorUnit.CalcValues(
		fileRecords,
		helper.OxygenComparator,
	)

	co2ScrubberUnit := unit.Unit{
		Name:     co2_scrubber,
		BitValue: "",
	}
	co2ScrubberUnit.CalcValues(
		fileRecords,
		helper.Co2ScrubberComparator,
	)

	fmt.Printf("oxygenGenerator: %d (%s)\n", oxygenGeneratorUnit.CalcDecimal(), oxygenGeneratorUnit.BitValue)
	fmt.Printf("co2Scrubber: %d (%s)\n", co2ScrubberUnit.CalcDecimal(), co2ScrubberUnit.BitValue)

	fmt.Println()
	fmt.Printf("Result %d \n", co2ScrubberUnit.CalcDecimal()*oxygenGeneratorUnit.CalcDecimal())
}
