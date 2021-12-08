package main

import (
	"advent-07/utils"
	"log"
	"math"

	"github.com/montanaflynn/stats"
)

var INPUT_FILE = "input.txt"

func main() {

	log.Println("Advent of Code - Day 6")
	log.Println("----------------------")

	readings := utils.LoadReadings(INPUT_FILE)

	log.Println("Puzzle 1")
	fuel := puzzle_1(readings)
	log.Printf("Total Fuel: %f", fuel)

	log.Println("----------------------")
	log.Println("Puzzle 2")
	fuel = puzzle_2(readings)
	log.Printf("Total Fuel: %f", fuel)

}

func puzzle_1(readings []float64) float64 {

	avg, _ := stats.Median(readings)

	log.Println(avg)

	var fuel float64
	for _, r := range readings {

		if r > avg {
			fuel += r - avg
		} else {
			fuel += avg - r
		}
	}

	return fuel

}

func puzzle_2(readings []float64) float64 {

	stdev, _ := stats.StandardDeviation(readings)
	median, _ := stats.Median(readings)

	min := math.Round(median - stdev)
	max := math.Round(median + stdev)

	var fuel_cost []float64
	for i := min; i < max; i++ {
		var fuel float64
		for _, r := range readings {
			if r > i {
				fuel += factorial(r - i)
			} else {
				fuel += factorial(i - r)
			}
		}

		fuel_cost = append(fuel_cost, fuel)

	}

	min_fuel, _ := stats.Min(fuel_cost)

	return min_fuel
}

func factorial(num float64) float64 {

	var result float64

	for i := float64(0); i <= num; i++ {
		result += i
	}

	return result
}
