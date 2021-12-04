package game

import (
	"log"
)

func RunGame(bingoNumbers []int, boards []Board) {

evalute:
	for _, num := range bingoNumbers {
		for boardIndex := range boards {
			evaluteBoard(num, &boards[boardIndex])
			checkWinConditions(&boards[boardIndex])
			if boards[boardIndex].hasWon {
				log.Printf("Board %d won with number %d:", boardIndex, num)
				PrintBoard(&boards[boardIndex])
				calculateScore(num, &boards[boardIndex])
				break evalute
			}
		}
	}

}

//Bad Logic - should re-write (sleepy)
func RunGameLastWin(bingoNumbers []int, boards []Board) {

	var lastBoard Board
	var lastNum int

	for _, num := range bingoNumbers {
		for boardIndex := range boards {
			var won bool
			if !boards[boardIndex].hasWon {
				evaluteBoard(num, &boards[boardIndex])
				checkWinConditions(&boards[boardIndex])
				if boards[boardIndex].hasWon {
					won = true
				}
			}

			if won {
				lastBoard = boards[boardIndex]
				lastNum = num
			}
		}
	}
	log.Println("Winning Board:")
	PrintBoard(&lastBoard)
	calculateScore(lastNum, &lastBoard)

}

func evaluteBoard(num int, board *Board) {

	for rowIndex, row := range board.rows {
		for spaceIndex, space := range row.spaces {
			if space.value == num {
				board.rows[rowIndex].spaces[spaceIndex].isCalled = true
			}
		}
	}

}

func checkWinConditions(board *Board) {
	checkHorizontalWin(board)
	checkVerticalWin(board)
}

func calculateScore(num int, board *Board) {

	var sum int
	for _, row := range board.rows {
		for _, space := range row.spaces {
			if !space.isCalled {
				sum = sum + space.value

			}
		}
	}
	log.Printf("Sum of unmarked numbers %d", sum)
	log.Printf("Final Score %d", sum*num)
}

func checkHorizontalWin(board *Board) {

	if !board.hasWon {
		var numMatched int
		for _, row := range board.rows {
			for _, space := range row.spaces {
				if space.isCalled {
					numMatched++
				}
			}

			if numMatched == 5 {
				board.hasWon = true
			}
			numMatched = 0
		}
	}

}

func checkVerticalWin(board *Board) {
	if !board.hasWon {
		var numMatched int
		for i := 0; i <= 4; i++ {
			for _, row := range board.rows {
				if row.spaces[i].isCalled {
					numMatched++
				}
			}

			if numMatched == 5 {
				board.hasWon = true
			}
			numMatched = 0
		}

	}
}
