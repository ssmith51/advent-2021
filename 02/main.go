package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	direction string
	movement  int
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("Error occured: %s", err.Error())
	}
}

func readInput(fi *os.File) []Command {
	var orders []Command

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {

		parts := strings.Fields(strings.Trim(scanner.Text(), ""))
		move, err := strconv.Atoi(parts[1])
		handleError(err)

		command := Command{
			direction: parts[0],
			movement:  move,
		}
		orders = append(orders, command)

	}

	return orders

}

func part1(orders []Command) {
	log.Println("Puzzle 1: Starting")
	var hPos, depth int

	for _, order := range orders {

		if order.direction == "forward" {
			hPos = hPos + order.movement

		} else if order.direction == "up" {
			depth = depth - order.movement

		} else if order.direction == "down" {
			depth = depth + order.movement
		}

	}
	log.Printf("Total Horizontal: %d, Total Depth: %d", hPos, depth)
	log.Printf("Solution: %d", hPos*depth)

}

func part2(orders []Command) {
	log.Println("Puzzle 2: Starting")
	var hPos, depth, aim int

	for _, order := range orders {

		if order.direction == "forward" {
			hPos = hPos + order.movement
			depth = depth + (aim * order.movement)

		} else if order.direction == "up" {
			aim = aim - order.movement

		} else if order.direction == "down" {
			aim = aim + order.movement
		}

	}
	log.Printf("Total Horizontal: %d, Total Depth: %d, Total Aim: %d", hPos, depth, aim)
	log.Printf("Solution: %d", hPos*depth)
}

func main() {
	log.Println("Advent of Code - Day 2")
	log.Println("----------------------")
	fi, err := os.Open("input.txt")
	handleError(err)
	defer fi.Close()

	orders := readInput(fi)
	part1(orders)
	log.Println("----------------------")

	part2(orders)
	log.Println("----------------------")

}
