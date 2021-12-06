package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ScanFish(input_file string) []int {

	var fish []int

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
			fish = append(fish, num)
		}
	}

	return fish

}
