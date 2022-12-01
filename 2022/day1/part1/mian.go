package main

import (
	common "2022"
	"bufio"
	"fmt"
)

func main() {
	file := common.GetFile("/day1/input.txt")

	maxValue := 0
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			sum += common.ParseToInt(text)
		} else {
			sum = 0
		}

		if maxValue < sum {
			maxValue = sum
		}
	}
	defer file.Close()

	fmt.Println(maxValue)
}
