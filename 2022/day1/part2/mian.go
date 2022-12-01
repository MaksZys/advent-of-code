package main

import (
	common "2022"
	"bufio"
	"fmt"
)

func findTheSmallestValueIndex(arr *[3]int, value int) (int, bool) {
	index := 0

	for i := 0; i < len(arr); i++ {
		if (*arr)[index] > (*arr)[i] {
			index = i
		}
	}

	smallestFound := (*arr)[index]
	if smallestFound < value {
		return index, true
	}

	return -1, false
}

func main() {
	file := common.GetFile("/day1/input.txt")
	defer file.Close()

	maxValuesSlice := [3]int{0, 0, 0}
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			sum += common.ParseToInt(text)
		} else {
			index, wasFound := findTheSmallestValueIndex(&maxValuesSlice, sum)
			fmt.Println()
			if wasFound {
				maxValuesSlice[index] = sum
			}

			sum = 0
		}
	}

	if sum > 0 {
		index, wasFound := findTheSmallestValueIndex(&maxValuesSlice, sum)
		fmt.Println()
		if wasFound {
			maxValuesSlice[index] = sum
		}
	}

	sum = 0
	for i := 0; i < len(maxValuesSlice); i++ {
		sum += maxValuesSlice[i]
	}

	fmt.Println(sum)
}
