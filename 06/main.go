package main

import (
	"advent-06/utils"
	"log"
)

var INPUT_FILE = "test.txt"

func main() {
	log.Println("Advent of Code - Day 6")
	log.Println("----------------------")

	initFish := utils.ScanFish(INPUT_FILE)

	log.Println("Puzzle 1: Starting")
	calcFish(80, initFish)
	log.Println("----------------------")

	log.Println("Puzzle 1: Starting")
	calcFish(256, initFish)
	log.Println("----------------------")

}

func calcFish(days int, initFish []int) {

	//Fish go on 7 or 9 day cycles, so store up to 9 days of fishes
	fishCountByDay := make([]int, 9)

	//Create an initial count of fish in each day of the cycle
	for _, f := range initFish {
		fishCountByDay[f]++
	}

	log.Printf("Initial Fish: %v", fishCountByDay)

	//Iterate over each day
	for day := 0; day < days; day++ {

		//Save how many fish are going to give birth and be born in the first position
		var newFish = fishCountByDay[0]

		//Iterate over each fish and shift them left to 'count down' to birth (e.g. fish move from init position -> 6 -> 0...)
		for i := 0; i < len(fishCountByDay)-1; i++ {
			fishCountByDay[i] = fishCountByDay[i+1]
		}

		//Birth new fish - parent fish go to position 6, baby fish go to position 8
		fishCountByDay[6] += newFish
		fishCountByDay[8] = newFish
	}

	//Sum the total number of fish
	count := 0
	for _, day := range fishCountByDay {
		count += day
	}

	log.Printf("Total number of fish: %d", count)

}
