package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func LoadReadings(input_file string) []float64 {

	var readings []float64

	fi, err := os.Open(input_file)
	HandleError(err)
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "")
		s := strings.Split(line, ",")

		for _, n := range s {
			num, err := strconv.Atoi(n)
			HandleError(err)
			readings = append(readings, float64(num))
		}
	}

	return readings

}
