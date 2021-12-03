package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var INPUT_FILE = "input.txt"

func handleError(err error) {
	if err != nil {
		log.Fatalf("Error occured: %s", err.Error())
	}
}

//Read in the file as an array to make up the 'Report'
func readInput(fi *os.File) []string {
	var report []string

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "")
		report = append(report, line)
	}

	return report
}

//Puzzle #1 Entry
func part1(report []string) {
	log.Println("Puzzle 1: Starting")

	var pos, oneCount, zeroCount int
	var gamma, eplison string

	//Iterate over each column in the report
	for pos < len(report[0]) {

		//Iterate over every row in the report and read the character at the current column
		for _, line := range report {

			if line[pos] == '0' {
				zeroCount++

			} else if line[pos] == '1' {
				oneCount++
			}

		}

		//Determine Gamma & Epilson Values for the current colum & position
		if zeroCount > oneCount {
			gamma = gamma + "0"
			eplison = eplison + "1"
		} else if oneCount > zeroCount {
			gamma = gamma + "1"
			eplison = eplison + "0"
		}

		//Reset counts and increment Columnn position
		zeroCount = 0
		oneCount = 0
		pos++

	}

	log.Printf("Gamma (Binary): %s, Epilson (Binary): %s", gamma, eplison)

	//Convert the binary Gamma to Int
	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	handleError(err)

	//Convert the binary Eplison to Int
	eplisonInt, err := strconv.ParseInt(eplison, 2, 64)
	handleError(err)

	log.Printf("Gamma (Int): %d, Epilson (Int) %d", gammaInt, eplisonInt)
	log.Printf("Power Consumpiton: %d", gammaInt*eplisonInt)

}

func part2(report []string) {
	log.Println("Puzzle 2: Starting")

	var pos int
	var oxygenReport = report
	var co2Report = report

	for pos < len(report[0]) {

		//Determine the value to search for from each report
		oxygenVal := findValue(oxygenReport, pos, false)
		co2Val := findValue(co2Report, pos, true)

		//Recursively reduces the values of the report
		oxygenReport = readReport(oxygenReport, pos, oxygenVal)
		co2Report = readReport(co2Report, pos, co2Val)

		pos++
	}

	log.Printf("Oxygen Value: %s, C02 Value: %s", oxygenReport[0], co2Report[0])

	//Convert the binary Gamma to Int
	oxygenInt, err := strconv.ParseInt(oxygenReport[0], 2, 64)
	handleError(err)

	//Convert the binary Eplison to Int
	co2Int, err := strconv.ParseInt(co2Report[0], 2, 64)
	handleError(err)

	log.Printf("Oxygen (Int): %d, Co2 (Int) %d", oxygenInt, co2Int)
	log.Printf("Life Support Rating: %d", oxygenInt*co2Int)
}

//Reduce the values in the report
func readReport(report []string, pos int, val byte) []string {
	var reducedReport []string

	if len(report) == 1 {
		reducedReport = report

	} else {
		for _, line := range report {
			if line[pos] == val {
				reducedReport = append(reducedReport, line)
			}
		}

		if len(reducedReport) == 0 {
			reducedReport = report
		}

	}

	return reducedReport
}

//Find the most commmon number in the current position to filter the report by
func findValue(report []string, pos int, invert bool) byte {

	var zeroCount, oneCount int
	var val byte

	//Iterate over every row in the report and read the character at the current column
	for _, line := range report {
		if line[pos] == '0' {
			zeroCount++
		} else if line[pos] == '1' {
			oneCount++
		}
	}

	//If values are equal default to a '1'
	if zeroCount > oneCount {
		val = '0'
	} else if oneCount > zeroCount {
		val = '1'
	} else {
		val = '1'
	}

	if invert && val == '0' {
		val = '1'
	} else if invert && val == '1' {
		val = '0'
	}

	return val
}

func main() {
	log.Println("Advent of Code - Day 3")
	log.Println("----------------------")

	fi, err := os.Open(INPUT_FILE)
	handleError(err)
	defer fi.Close()

	report := readInput(fi)

	part1(report)
	log.Println("----------------------")

	part2(report)
	log.Println("----------------------")

}
