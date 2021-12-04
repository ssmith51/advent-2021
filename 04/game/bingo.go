package game

import (
	"log"
)

func RunGame(bingoNumbers []int, boards []Board) {

evalute:
	for _, num := range bingoNumbers {
		log.Println(num)
		for boardIndex, board := range boards {
			evaluteBoard(num, &boards[boardIndex])
			checkWinConditions(&boards[boardIndex])
			if boards[boardIndex].hasWon {
				log.Printf("Board Won %d with number %d", boardIndex, num)
				PrintBoard(board)
				calculateScore(num, board)
				break evalute
			}
		}
	}

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
	// checkDownDiagonalWin(board)
	// checkUpDiagonalWin(board)
}

func calculateScore(num int, board Board) {

	var sum int
	for _, row := range board.rows {
		for _, space := range row.spaces {
			if !space.isCalled {
				sum = sum + space.value

			}
		}
	}
	sum = sum - num
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
				log.Println("Horizontal Win!")
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
				log.Println("Vertical Win!")
			}
			numMatched = 0
		}

	}
}

// func checkDownDiagonalWin(board *Board) {
// 	if !board.hasWon {
// 		var numMatched int
// 		for i := 0; i <= 4; i++ {
// 			if board.rows[i].spaces[i].isCalled {
// 				numMatched++
// 			}
// 		}

// 		if numMatched == 5 {
// 			board.hasWon = true
// 		}
// 		numMatched = 0
// 	}
// }

// func checkUpDiagonalWin(board *Board) {
// 	if !board.hasWon {
// 		var numMatched int
// 		var x int
// 		for y := 4; y >= 0; y-- {
// 			if board.rows[x].spaces[y].isCalled {
// 				numMatched++
// 			}
// 			x++
// 		}

// 		if numMatched == 5 {
// 			board.hasWon = true
// 		}
// 		numMatched = 0
// 	}
// }
