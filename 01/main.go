package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		log.Fatalf("Error occured: %s", err.Error())
	}

}

func readInput(fi *os.File) []int {
	var depths []int

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		handleError(err)
		depths = append(depths, depth)
	}
	if len(depths) <= 0 {
		log.Fatal("Unable to get depths")
	}
	return depths

}

func part1(depths []int) {
	log.Println("Puzzle 1: Starting")
	var previousDepth, count int = -1, 0

	for _, currentDepth := range depths {
		if previousDepth != -1 {
			if currentDepth > previousDepth {
				count++
			}
		}
		previousDepth = currentDepth
	}
	log.Printf("Puzzle 1: Total Depth Increase Count: %d", count)

}

func part2(depths []int) {
	log.Println("Puzzle 2: Starting")

	lenDepts := len(depths) - 3

	var i, count int

	for i <= lenDepts {
		a := depths[i : i+3]
		b := depths[i+1 : i+4]

		sumA := a[0] + a[1] + a[2]
		sumB := b[0] + b[1] + b[2]

		if sumB > sumA {
			count++
		}

		i++
	}

	log.Printf("Puzzle 2: Total Depth Increase Count: %d", count)

}

func main() {
	log.Println("Advent of Code - Day 1")
	log.Println("----------------------")
	fi, err := os.Open("input.txt")
	handleError(err)
	defer fi.Close()

	//Read in the file to get a list of depths
	depths := readInput(fi)

	//Calculate Part 1
	part1(depths)
	log.Println("----------------------")

	//Calculate Part 2
	part2(depths)
	log.Println("----------------------")

}
