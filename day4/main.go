package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bingoBoard struct {
	rows []row
	ID   int64
}

type row struct {
	row []bingoNumber
}

type bingoNumber struct {
	value int64
	drawn bool
}

var boardWinningOrder []int64

func main() {
	bingoNumbers := []int64{91, 17, 64, 45, 8, 13, 47, 19, 52, 68, 63, 76, 82, 44, 28, 56, 37, 2, 78, 48, 32, 58, 72, 53, 9, 85, 77, 89, 36, 22, 49, 86, 51, 99, 6, 92, 80, 87, 7, 25, 31, 66, 84, 4, 98, 67, 46, 61, 59, 79, 0, 3, 38, 27, 23, 95, 20, 35, 14, 30, 26, 33, 42, 93, 12, 57, 11, 54, 50, 75, 90, 41, 88, 96, 40, 81, 24, 94, 18, 39, 70, 34, 21, 55, 5, 29, 71, 83, 1, 60, 74, 69, 10, 62, 43, 73, 97, 65, 15, 16}

	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("cannot read input.txt. Error: %v", err)
		return
	}
	lines := strings.Split(string(contents), "\n")
	boards := buildBoards(lines)

	for _, bingoNumber := range bingoNumbers {
		boards = markBoards(boards, bingoNumber)
		checkBoards(boards)
		// fmt.Println(boardWinningOrder)
		// fmt.Println(len(boardWinningOrder))
		// fmt.Println(bingoNumber)
		if len(boards) == len(boardWinningOrder) {
			for _, board := range boards {
				if board.ID == boardWinningOrder[len(boardWinningOrder)-1] {
					calculateScore(board, bingoNumber)
				}
			}
			break
		}
	}
}

func makeHumanReadableBoard(board bingoBoard) {
	for _, rows := range board.rows {
		var buffer bytes.Buffer
		for _, bingoNumber := range rows.row {
			buffer.WriteString(strconv.FormatBool(bingoNumber.drawn) + " ")
		}
		fmt.Println(buffer.String())
	}
}

func buildBoards(lines []string) []bingoBoard {
	bingoBoards := []bingoBoard{}
	board := bingoBoard{}
	ID := int64(0)
	for _, line := range lines {
		rowNumbers := strings.Split(line, " ")
		if len(rowNumbers) == 1 {
			board.ID = ID
			bingoBoards = append(bingoBoards, board)
			board = bingoBoard{}
			ID += 1
			continue
		}
		row := row{}
		for _, number := range rowNumbers {
			if number == "" {
				continue
			}
			n, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				fmt.Printf("error: %v", err)
			}
			row.row = append(row.row, bingoNumber{
				value: n,
				drawn: false,
			})
		}
		board.rows = append(board.rows, row)
	}
	return bingoBoards
}

func markBoards(boards []bingoBoard, number int64) []bingoBoard {
	for _, board := range boards {
		for _, rows := range board.rows {
			for i := 0; i < len(rows.row); i++ {
				if rows.row[i].value == number {
					rows.row[i].drawn = true
				}
			}
		}
	}
	return boards
}

func checkBoards(boards []bingoBoard) {
	// check rows
	for _, board := range boards {
		for _, rows := range board.rows {
			completedRow := true
			for _, bingoNumber := range rows.row {
				if !bingoNumber.drawn {
					completedRow = false
					break
				}
			}
			if completedRow && !contains(boardWinningOrder, board.ID) {
				fmt.Println("completed row: ", board.ID)
				boardWinningOrder = append(boardWinningOrder, board.ID)
				makeHumanReadableBoard(board)
			}
		}
	}
	// check columns
	for _, board := range boards {
		for i := 0; i < 5; i++ {
			completedColumn := true
			for _, row := range board.rows {
				if !row.row[i].drawn {
					completedColumn = false
					break
				}
			}
			if completedColumn && !contains(boardWinningOrder, board.ID) {
				fmt.Println("completed column: ", board.ID)
				boardWinningOrder = append(boardWinningOrder, board.ID)
				makeHumanReadableBoard(board)
			}
		}
	}
}

func calculateScore(bingoBoard bingoBoard, winningNumber int64) {
	var sumOfUnmarkedNumbers int64
	for _, row := range bingoBoard.rows {
		for _, bingoNumber := range row.row {
			if !bingoNumber.drawn {
				sumOfUnmarkedNumbers += bingoNumber.value
			}
		}
	}
	fmt.Println("sum: ", sumOfUnmarkedNumbers*winningNumber)
}

func contains(slice []int64, value int64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
