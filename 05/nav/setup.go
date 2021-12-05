package nav

import (
	"advent-05/utils"
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var PUZZEL_1_INPUT_REGEX = `^(\d*),(\d*)\s->\s(\d*),(\d*)$`

func LoadVents(input_file string) VentGrid {

	var readings VentGrid

	fi, err := os.Open(input_file)
	utils.HandleError(err)
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "")

		r := regexp.MustCompile(PUZZEL_1_INPUT_REGEX)
		parts := r.FindStringSubmatch(line)

		if len(parts) < 5 {
			utils.HandleError(errors.New("error parsing input"))
		}

		//Parse out X & Y values
		x1, err := strconv.Atoi(parts[1])
		utils.HandleError(err)

		y1, err := strconv.Atoi(parts[2])
		utils.HandleError(err)

		x2, err := strconv.Atoi(parts[3])
		utils.HandleError(err)

		y2, err := strconv.Atoi(parts[4])
		utils.HandleError(err)

		coord := NewReading(x1, y1, x2, y2)

		//Find the max of the grid
		readings.MaxX = evalMaxX(readings.MaxX, x1, x2)
		readings.MaxY = evalMaxY(readings.MaxY, y1, y2)
		readings.Coordinates = append(readings.Coordinates, coord)

	}

	//Increase max by 1 for proper calculations
	readings.MaxX = readings.MaxX + 1
	readings.MaxY = readings.MaxY + 1

	return readings
}

// Find the Max X value to establish Grid
func evalMaxX(maxX int, x1 int, x2 int) int {

	if x1 > x2 {
		if x1 > maxX {
			maxX = x1
		}

	} else {
		if x2 > maxX {
			maxX = x2
		}

	}

	return maxX

}

// Find the Min Y Value to Eastablish Grid
func evalMaxY(maxY int, y1 int, y2 int) int {

	if y1 > y2 {
		if y1 > maxY {
			maxY = y1
		}

	} else {
		if y2 > maxY {
			maxY = y2
		}
	}

	return maxY

}
