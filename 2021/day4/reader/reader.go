package reader

import (
	"adventOfCode2021"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingoRecord struct {
	value    uint16
	isMarked bool
}

func (record *bingoRecord) SetMarked() {
	(*record).isMarked = true
}

func (record *bingoRecord) SetValue(value uint16) {
	(*record).value = value
}

func (record *bingoRecord) GetValue() uint16 {
	return (*record).value
}

type BingoBoard struct {
	board   [5][5]bingoRecord
	isBingo bool
}

func (board *BingoBoard) Bingo() {
	board.isBingo = true
}

func (board *BingoBoard) AlreadyWon() bool {
	return board.isBingo
}

func (board *BingoBoard) checkBoard(x int, y int) bool {
	var xResult, yResult bool = true, true

	for i := 0; i < 5; i++ {
		if !board.board[x][i].isMarked {
			xResult = false
			break
		}
	}

	for j := 0; j < 5; j++ {
		if !board.board[j][y].isMarked {
			yResult = false
			break
		}
	}

	return xResult || yResult
}

func (board *BingoBoard) PrintArray() {
	for _, x := range (*board).board {
		for i := 0; i < 5; i++ {
			fmt.Printf("%3d", x[i].value)
		}

		fmt.Println()
	}
}

func (board *BingoBoard) MarkValue(value uint16) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if (*board).board[i][j].value == value {
				(*board).board[i][j].SetMarked()

				if board.checkBoard(i, j) {
					return true
				}
			}
		}
	}

	return false
}

func (board *BingoBoard) SumUnmarked() int {
	sum := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if (*board).board[i][j].isMarked == false {
				sum += int((*board).board[i][j].value)
			}
		}
	}

	return sum
}

func filterEmpty(values []string) []string {
	result := make([]string, 0)

	for _, element := range values {
		if element != "" {
			result = append(result, element)
		}
	}

	return result
}

// Return array of BingoBoards array (1) and bingo values (2)
func ReadFromFile(path string) ([]BingoBoard, []uint16) {
	filePath := adventOfCode2021.GetPath(path)
	fileIO, error := os.OpenFile(filePath, os.O_RDONLY, 0600)
	defer fileIO.Close()
	if error != nil {
		log.Fatalf("Error appeared in ReadFromFile. \nPath: %s.\n Error: %s", path, error)
	}

	rawBytes, error := ioutil.ReadAll(fileIO)
	if error != nil {
		panic(error)
	}

	bingoBoards := make([]BingoBoard, 0)
	bingoValues := make([]uint16, 0)
	lines := strings.Split(string(rawBytes), "\n")

	bingoBoardIndex := -1
	bingoBoardLine := 0
	for index, line := range lines {
		if index == 0 {
			bingoStringValues := strings.Split(line, ",")

			for _, val := range bingoStringValues {
				parsed, error := strconv.Atoi(val)
				if error != nil {
					panic(error)
				}

				bingoValues = append(bingoValues, uint16(parsed))
			}

			continue
		}

		if line == "" {
			bingoBoardIndex++
			bingoBoardLine = index + 1
			bingoBoards = append(bingoBoards, BingoBoard{
				board:   [5][5]bingoRecord{},
				isBingo: false,
			})

			continue
		}

		splitLineRows := strings.Split(line, " ")
		lineRows := filterEmpty(splitLineRows)
		for rowIndex, rowValue := range lineRows {
			rowInt, error := strconv.Atoi(rowValue)
			if error != nil {
				continue
			}

			bingoBoards[bingoBoardIndex].board[index-bingoBoardLine][rowIndex].value = uint16(rowInt)
			bingoBoards[bingoBoardIndex].board[index-bingoBoardLine][rowIndex].isMarked = false
		}
	}

	return bingoBoards, bingoValues
}
