package main

import (
	"advent-05/nav"
	"log"
)

var INPUT_FILE = "input.txt"

func main() {
	log.Println("Advent of Code - Day 4")
	log.Println("----------------------")

	readings := nav.LoadVents(INPUT_FILE)
	log.Printf("Max X: %d, Max Y: %d, Number of Coordinates: %d", readings.MaxX, readings.MaxY, len(readings.Coordinates))

	log.Println("Puzzle 1: Starting")
	nav.Navigate(readings, true)
	log.Println("----------------------")

	log.Println("Puzzle 2: Starting")
	nav.Navigate(readings, false)
	log.Println("----------------------")

}
