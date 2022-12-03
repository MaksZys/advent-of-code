package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode"
)

func removeDuplicatesFromSortedArr(source []string) []string {
	target := make([]string, 0)

	actualValue := source[0]
	target = append(target, actualValue)

	for i := 1; i < len(source); i++ {
		if actualValue != source[i] {
			actualValue = source[i]
			target = append(target, actualValue)
		}
	}

	return target
}

func prepareArray(source string) []string {
	target := strings.Split(source, "")
	sort.Strings(target)

	return removeDuplicatesFromSortedArr(target)
}

func findDuplicate(firstCompartment []string, secondCompartment []string) []string {
	foundDuplicates := make([]string, 0)
	iStartPoint := 0

	for _, first := range firstCompartment {
		for i := iStartPoint; i < len(secondCompartment); i++ {
			second := secondCompartment[i]
			if first == second {
				foundDuplicates = append(foundDuplicates, first)
			}

			if first > second {
				iStartPoint = i
			}
		}
	}

	return foundDuplicates
}

func calculateScore(character rune) int {
	if unicode.IsUpper(character) {
		return int(character) - 64 + 26
	}

	return int(character) - 96
}

func get3Rows(scanner *bufio.Scanner) [3]string {
	var rows [3]string

	for i := 0; i < 3; i++ {
		if scanner.Scan() {
			rows[i] = scanner.Text()
		}
	}

	return rows
}

func fileRowsCounter(file *os.File) int {
	size := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		size++
	}

	// reset file pointer to the beginning
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return 0
	}

	return size
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("File open error")
		os.Exit(1)
	}

	fileRowsSize := fileRowsCounter(file)
	scanner := bufio.NewScanner(file)

	sum := 0
	for i := 0; i < fileRowsSize/3; i++ {
		lines := get3Rows(scanner)

		row1 := prepareArray(lines[0])
		row2 := prepareArray(lines[1])
		row3 := prepareArray(lines[2])

		duplicateChar := findDuplicate(row3, findDuplicate(row1, row2))

		sum += calculateScore([]rune(duplicateChar[0])[0])
	}

	fmt.Println(sum)
}
