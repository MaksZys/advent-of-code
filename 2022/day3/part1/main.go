package main

import (
	"bufio"
	"fmt"
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

func prepareArray(firstCompartment []string, secondCompartment []string) ([]string, []string) {
	sort.Strings(firstCompartment)
	sort.Strings(secondCompartment)

	return removeDuplicatesFromSortedArr(firstCompartment),
		removeDuplicatesFromSortedArr(secondCompartment)
}

func findDuplicate(firstCompartment []string, secondCompartment []string) rune {
	for _, first := range firstCompartment {
		for _, second := range secondCompartment {
			if first == second {
				return []rune(first)[0]
			}
		}
	}

	fmt.Printf("Coudn't find any common character")
	os.Exit(1)

	return ' '
}

func calculateScore(character rune) int {
	if unicode.IsUpper(character) {
		return int(character) - 64 + 26
	}

	return int(character) - 96
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("File open error")
		os.Exit(1)
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		length := len(line)
		halfPoint := length / 2

		firstCompartment := strings.Split(line[0:halfPoint], "")
		secondCompartment := strings.Split(line[halfPoint:length], "")

		res1, res2 := prepareArray(firstCompartment, secondCompartment)
		duplicatedChar := findDuplicate(res1, res2)
		sum += calculateScore(duplicatedChar)
	}

	fmt.Println(sum)
}
