package nav

import (
	"log"
	"math"
)

//Navigate evaluates the reads from the system to deremine how many thermal vents intersect each other
func Navigate(readings VentGrid, xyOnly bool) {

	//Create basic Grid
	ocean := make([][]int, readings.MaxX)

	for i := 0; i < readings.MaxY; i++ {
		ocean[i] = make([]int, readings.MaxY)
	}

	for _, coord := range readings.Coordinates {
		evalHorizontalLines(coord, &ocean)
		evalVerticalLines(coord, &ocean)

		if !xyOnly {
			evalDiagonalLines(coord, &ocean)
		}
	}

	numReadings := findMatches(ocean, readings.MaxX, readings.MaxY)

	//Debug Ocean Grid
	// log.Println("Completed Ocean:")
	// for _, line := range ocean {
	// 	log.Println(line)
	// }

	log.Printf("Number of readings %d", numReadings)

}

//evalHorizontalLines Evalutes all the points in an horizatonal line between two coordinates
func evalHorizontalLines(coord Coordinate, ocean *[][]int) {
	if coord.X1 == coord.X2 {
		x := coord.X1

		if coord.Y2 < coord.Y1 {

			for y := coord.Y1; y >= coord.Y2; y-- {
				(*ocean)[y][x] = (*ocean)[y][x] + 1
			}
		}

		if coord.Y1 < coord.Y2 {
			for y := coord.Y2; y >= coord.Y1; y-- {
				(*ocean)[y][x] = (*ocean)[y][x] + 1
			}
		}
	}

}

//evalVerticalLines Evalutes all the points in an vertical line between two coordinates
func evalVerticalLines(coord Coordinate, ocean *[][]int) {

	if coord.Y1 == coord.Y2 {
		y := coord.Y1

		if coord.X2 < coord.X1 {

			for x := coord.X1; x >= coord.X2; x-- {
				(*ocean)[y][x] = (*ocean)[y][x] + 1
			}
		}

		if coord.X1 < coord.X2 {
			for x := coord.X2; x >= coord.X1; x-- {
				(*ocean)[y][x] = (*ocean)[y][x] + 1
			}
		}
	}

}

func evalDiagonalLines(coord Coordinate, ocean *[][]int) {

	angle := calculateAngle(coord)

	if angle == 45 {

		//Mark the starting coordinate
		(*ocean)[coord.Y1][coord.X1] = (*ocean)[coord.Y1][coord.X1] + 1
		(*ocean)[coord.Y2][coord.X2] = (*ocean)[coord.Y2][coord.X2] + 1

		//Debug Messages
		// log.Printf("Positive 45 Degree Angle Found: %f", angle)
		// log.Printf("Corrdinate: %d,%d -> %d,%d", coord.X1, coord.Y1, coord.X2, coord.Y2)

		var count = 1
		if coord.X1 < coord.X2 {
			for i := coord.X1; i < coord.X2-1; i++ {
				x := coord.X1 + count
				y := coord.Y1 + count
				count++
				(*ocean)[y][x] = (*ocean)[y][x] + 1

				//Debug Message
				// log.Printf("New Coord: %d,%d", x, y)
			}
		} else {
			for i := coord.X2; i < coord.X1-1; i++ {
				x := coord.X1 - count
				y := coord.Y1 - count
				count++
				(*ocean)[y][x] = (*ocean)[y][x] + 1

				//Debug Message
				// log.Printf("New Coord: %d,%d", x, y)
			}

		}

	} else if angle == -45 {
		//Debug Message
		// log.Printf("Negative 45 Degree Angle Found: %f", angle)
		// log.Printf("Corrdinate: %d,%d -> %d,%d", coord.X1, coord.Y1, coord.X2, coord.Y2)
		(*ocean)[coord.Y1][coord.X1] = (*ocean)[coord.Y1][coord.X1] + 1
		(*ocean)[coord.Y2][coord.X2] = (*ocean)[coord.Y2][coord.X2] + 1
		var count = 1

		if coord.X1 < coord.X2 {
			for i := coord.X1; i < coord.X2-1; i++ {
				x := coord.X1 + count
				y := coord.Y1 - count
				count++
				(*ocean)[y][x] = (*ocean)[y][x] + 1

				//Debug Messages
				// log.Printf("New Coord: %d,%d", x, y)
			}
		} else {
			for i := coord.X2; i < coord.X1-1; i++ {
				x := coord.X2 + count
				y := coord.Y2 - count
				count++
				(*ocean)[y][x] = (*ocean)[y][x] + 1

				//Debug Messages
				// log.Printf("New Coord: %d,%d", x, y)
			}
		}

	}
}

func calculateAngle(coord Coordinate) float64 {

	//Calculate Slope of the line
	var deltay = float64(coord.Y1) - float64(coord.Y2)
	var deltaX = float64(coord.X1) - float64(coord.X2)
	var slope float64

	if deltaX != 0 {
		slope = deltay / deltaX
	}

	//Calculate Angle from Slope
	angle := math.Atan(slope)

	angle = angle * (180 / math.Pi)

	return angle
}

//findMatches Looks for all the lines that have more than two intersections with each other
func findMatches(ocean [][]int, maxX int, maxY int) int {
	numReadings := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if ocean[y][x] > 1 {
				numReadings++
			}
		}
	}

	return numReadings
}
