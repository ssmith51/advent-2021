package game

import (
	"advent-04/utils"
	"fmt"
	"log"
	"strconv"
)

type Space struct {
	value    int
	isCalled bool
}

type Row struct {
	spaces [5]Space
}

type Board struct {
	rows   [5]Row
	hasWon bool
}

//Parse input and generate a new space on the board
func GenerateSpace(input string) Space {
	val, err := strconv.Atoi(input)
	utils.HandleError(err)

	newSpace := Space{
		value:    val,
		isCalled: false,
	}

	return newSpace

}

func PrintBoard(board *Board) {
	for _, row := range board.rows {
		var vals string
		for _, space := range row.spaces {
			if space.isCalled {
				vals = vals + fmt.Sprintf("%d* ", space.value)
			} else {
				vals = vals + fmt.Sprintf("%d  ", space.value)
			}
		}
		log.Println(vals)
	}
}
