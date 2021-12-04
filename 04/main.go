package main

import (
	"advent-04/game"
	"log"
)

func main() {
	log.Println("Advent of Code - Day 4")
	log.Println("----------------------")

	log.Println("Loading Game")

	log.Println("Puzzle 1: Starting")
	bingoNumbers1, boards1 := game.LoadGame()
	game.RunGame(bingoNumbers1, boards1)
	log.Println("----------------------")

	log.Println("Puzzle 2: Starting")
	bingoNumbers2, boards2 := game.LoadGame()
	game.RunGameLastWin(bingoNumbers2, boards2)
	log.Println("----------------------")

}
