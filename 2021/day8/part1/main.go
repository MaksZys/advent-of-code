package main

import (
	"adventOfCode2021"
	"fmt"
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

func decodeByLength(input string) decodedPattern {
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
	// case 5:
	// more than 1
	// case 6:
	// more than 1
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

func decodePatterns(signalPatterns []string) []decodedPattern {
	patterns := make([]decodedPattern, 0)

	for _, pattern := range signalPatterns {
		patterns = append(patterns, decodeByLength(pattern))
	}

	return patterns
}

func filterNegativePatterns(values []decodedPattern) []decodedPattern {
	positivePatterns := make([]decodedPattern, 0)

	for _, val := range values {
		if val.digit > -1 {
			positivePatterns = append(positivePatterns, val)
		}
	}

	return positivePatterns
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

func sumArrVal(arr [9]int) int {
	sum := arr[0]

	for i := 1; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}

func main() {
	file := adventOfCode2021.GetFile("/day8/input.txt")
	records := adventOfCode2021.ScanLines(file)

	var occurrences [9]int
	for _, record := range records {
		signalPatterns, outputValue := splitRecord(record)
		patterns := filterNegativePatterns(decodePatterns(signalPatterns))

		outputPatterns := searchForOccurrences(patterns, outputValue)

		for _, val := range outputPatterns {
			occurrences[val]++
		}
	}

	fmt.Println(sumArrVal(occurrences))
}
