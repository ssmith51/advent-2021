package main

import (
	"advent-04/game"
	"advent-04/utils"
	"log"
	"os"
)

var INPUT_FILE = "input.txt"

func main() {
	log.Println("Advent of Code - Day 4")
	log.Println("----------------------")

	fi, err := os.Open(INPUT_FILE)
	utils.HandleError(err)
	defer fi.Close()

	//Execute Part 1
	part1(fi)

}

func part1(fi *os.File) {

	bingoNumbers, boards := game.LoadGame(fi)

	log.Printf("Bingo Numbers: %v", bingoNumbers)
	log.Printf("Number of Boards loaded: %d", len(boards))

	game.RunGame(bingoNumbers, boards)

}
