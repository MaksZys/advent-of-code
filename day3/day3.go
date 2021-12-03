package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ColumnValuesStorage struct {
	zero int
	one  int
}

func getFileFromRootPath(path string) *os.File {
	rootPath, rootPathError := os.Getwd()
	if rootPathError != nil {
		panic("Cannot get executable path")
	}

	file, fileError := os.Open(rootPath + "/" + path)
	if fileError != nil {
		panic(fileError)
	}

	return file
}

func scanRecords(scanner *bufio.Scanner) []string {
	binariesArray := make([]string, 0)

	for scanner.Scan() {
		binariesArray = append(binariesArray, scanner.Text())
	}

	return binariesArray
}

func getMaxLength(numbers []string) int {
	maxLength := 0
	for _, number := range numbers {
		maxLength = len(number)
	}

	return maxLength
}

func createBinaryFromMostCommonBits(storage []ColumnValuesStorage) string {
	var binary strings.Builder
	for _, storageRecord := range storage {
		if storageRecord.zero == storageRecord.one {
			panic("EQUAL!!!")
		}

		if storageRecord.zero > storageRecord.one {
			binary.WriteString("0")
		} else {
			binary.WriteString("1")
		}
	}

	return binary.String()
}

func createBinaryFromLeastCommonBits(storage []ColumnValuesStorage) string {
	var binary strings.Builder
	for _, storageRecord := range storage {
		if storageRecord.zero == storageRecord.one {
			panic("EQUAL!!!")
		}

		if storageRecord.zero < storageRecord.one {
			binary.WriteString("0")
		} else {
			binary.WriteString("1")
		}
	}

	return binary.String()
}

func parseToDecimal(binary string) int {
	decimal, parseError := strconv.ParseInt(binary, 2, 64)
	if parseError != nil {
		log.Fatal("Parse binary to decimal error")
	}

	return int(decimal)
}

func main() {
	file := getFileFromRootPath("day3/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	binaryNumbers := scanRecords(scanner)
	theLongest := getMaxLength(binaryNumbers)

	var columnsStats = make([]ColumnValuesStorage, theLongest)
	for _, number := range binaryNumbers {
		for j := theLongest - 1; j >= 0; j-- {
			if number[j] == '0' {
				columnsStats[j].zero += 1
			} else if number[j] == '1' {
				columnsStats[j].one += 1
			} else {
				panic("Unknown value")
			}
		}
	}

	binaryGamma := createBinaryFromMostCommonBits(columnsStats)
	binaryEpsilon := createBinaryFromLeastCommonBits(columnsStats)

	decimalGamma := parseToDecimal(binaryGamma)
	decimalEpsilon := parseToDecimal(binaryEpsilon)

	fmt.Printf("Gamma: %d, epsilon: %d, multiplied: %d",
		decimalGamma,
		decimalEpsilon,
		decimalGamma*decimalEpsilon,
	)
}
