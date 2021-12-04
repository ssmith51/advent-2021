package game

import (
	"advent-04/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var INPUT_FILE = "input.txt"

func LoadGame() ([]int, []Board) {

	fi, err := os.Open(INPUT_FILE)
	utils.HandleError(err)
	defer fi.Close()

	var boards []Board
	var bingoNumbers []int

	scanner := bufio.NewScanner(fi)

	var rowCount int
	var isFirst = true
	newBoard := Board{
		hasWon: false,
	}
	for scanner.Scan() {
		line := scanner.Text()

		//Read first line of Bingo Numbers
		if isFirst {
			bingoNumbers = parseBingoNumbers(line)
			isFirst = false

		} else if len(line) > 0 { //Read boards & skip empty lines

			row := parseBoardRow(line)
			newBoard.rows[rowCount] = row
			rowCount++

			//If we have loaded 5 boards, time to push & create a new board
			if rowCount == 5 {
				boards = append(boards, newBoard)

				//Reset board & row counter
				newBoard = Board{
					hasWon: false,
				}
				rowCount = 0
			}

		}

	}

	return bingoNumbers, boards
}

//Spilt the input on ',' and convert to int to genearte array of input numbers
func parseBingoNumbers(input string) []int {
	var bingoNumbers []int

	for _, n := range strings.Split(input, ",") {
		val, err := strconv.Atoi(n)
		utils.HandleError(err)
		bingoNumbers = append(bingoNumbers, val)
	}

	return bingoNumbers
}

//Generate a row of values
func parseBoardRow(row string) Row {
	newRow := Row{}
	parts := strings.Fields(strings.Trim(row, ""))

	for i, part := range parts {
		newRow.spaces[i] = GenerateSpace(part)
	}

	return newRow

}
