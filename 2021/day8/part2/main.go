package main

import (
	"adventOfCode2021"
	"fmt"
	"math"
	"sort"
	"strings"
)

type decodedPattern struct {
	digit     int
	codedVal  string
	sortedVal string
}

func sortChars(input string) string {
	sortedVal := strings.Split(input, "")
	sort.Strings(sortedVal)

	return strings.Join(sortedVal, "")
}

func decodeUnique(input string) decodedPattern {
	switch len(input) {
	case 2:
		return decodedPattern{
			digit:     1,
			codedVal:  input,
			sortedVal: sortChars(input),
		}
	case 3:
		return decodedPattern{
			digit:     7,
			codedVal:  input,
			sortedVal: sortChars(input),
		}
	case 4:
		return decodedPattern{
			digit:     4,
			codedVal:  input,
			sortedVal: sortChars(input),
		}
	case 7:
		return decodedPattern{
			digit:     8,
			codedVal:  input,
			sortedVal: sortChars(input),
		}
	}

	return decodedPattern{
		digit:     -1,
		codedVal:  input,
		sortedVal: sortChars(input),
	}
}

func splitRecord(record string) ([]string, []string) {
	split := strings.Split(record, " | ")

	return strings.Split(split[0], " "), strings.Split(split[1], " ")
}

func countOccurrences(uniquePattern decodedPattern, input string) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(uniquePattern.sortedVal); j++ {
			if input[i] == uniquePattern.sortedVal[j] {
				sum++
			}
		}
	}

	return sum
}

func findNeededValues(uniquePatterns []decodedPattern) (decodedPattern, decodedPattern) {
	var four decodedPattern
	var seven decodedPattern

	for _, pattern := range uniquePatterns {
		if pattern.digit == 4 {
			four = pattern
		} else if pattern.digit == 7 {
			seven = pattern
		}
	}

	return four, seven
}

func decodeValue(input string, four decodedPattern, seven decodedPattern) decodedPattern {
	sortedVal := sortChars(input)
	fourOccurrences := countOccurrences(four, sortedVal)
	sevenOccurrences := countOccurrences(seven, sortedVal)

	switch len(input) {
	case 5:
		if fourOccurrences == 3 {
			if sevenOccurrences == 3 {
				return decodedPattern{
					digit:     3,
					codedVal:  input,
					sortedVal: sortedVal,
				}
			}

			return decodedPattern{
				digit:     5,
				codedVal:  input,
				sortedVal: sortedVal,
			}
		}

		return decodedPattern{
			digit:     2,
			codedVal:  input,
			sortedVal: sortedVal,
		}
	case 6:
		if fourOccurrences == 3 {
			if sevenOccurrences == 3 {
				return decodedPattern{
					digit:     0,
					codedVal:  input,
					sortedVal: sortedVal,
				}
			}

			return decodedPattern{
				digit:     6,
				codedVal:  input,
				sortedVal: sortedVal,
			}
		}

		return decodedPattern{
			digit:     9,
			codedVal:  input,
			sortedVal: sortedVal,
		}
	}

	return decodedPattern{
		digit:     -1,
		codedVal:  input,
		sortedVal: sortedVal,
	}
}

func decodePatterns(signalPatterns []string) []decodedPattern {
	uniquePatterns := make([]decodedPattern, 0)
	allDecoded := make([]decodedPattern, 0)

	for _, pattern := range signalPatterns {
		unique := decodeUnique(pattern)

		if unique.digit != -1 {
			uniquePatterns = append(uniquePatterns, unique)
			allDecoded = append(allDecoded, unique)
		}
	}

	four, seven := findNeededValues(uniquePatterns)
	for _, pattern := range signalPatterns {
		decoded := decodeValue(pattern, four, seven)

		if decoded.digit != -1 {
			allDecoded = append(allDecoded, decoded)
		}
	}

	return allDecoded
}

func searchForOccurrences(patterns []decodedPattern, outputValues []string) []int {
	output := make([]int, 0)

	for _, value := range outputValues {
		for _, pattern := range patterns {
			sortedVal := strings.Split(value, "")
			sort.Strings(sortedVal)

			if strings.Join(sortedVal, "") == pattern.sortedVal {
				output = append(output, pattern.digit)
			}
		}
	}

	return output
}

func filterNegativePatterns(values []int) []int {
	positivePatterns := make([]int, 0)

	for _, val := range values {
		if val > -1 {
			positivePatterns = append(positivePatterns, val)
		}
	}

	return positivePatterns
}

func createValue(values []int) int {
	sum := 0
	length := len(values) - 1

	for i := 0; i <= length; i++ {
		sum += (values[length-i] * int(math.Pow(10, float64(i))))
	}

	return sum
}

func main() {
	file := adventOfCode2021.GetFile("/day8/input.txt")
	records := adventOfCode2021.ScanLines(file)

	sum := 0

	for _, record := range records {
		signalPatterns, outputValue := splitRecord(record)
		patterns := decodePatterns(signalPatterns)

		outputPatterns := searchForOccurrences(patterns, outputValue)

		sum += createValue(outputPatterns)
	}

	fmt.Println(sum)
}
