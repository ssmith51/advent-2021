package nav

import "log"

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
			log.Println("Eval H Lines too")
		}
	}

	numReadings := findMatches(ocean, readings.MaxX, readings.MaxY)

	log.Println("Completed Ocean:")
	for _, line := range ocean {
		log.Println(line)
	}

	log.Printf("Number of readings %d", numReadings)

}

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
