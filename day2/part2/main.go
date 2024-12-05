package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Smaller = iota
	Larger
	Undetermined
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("Error while reading input file: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		if validateReport(scanner.Text()) {
			total++
		}
	}

	fmt.Println("total:", total)

}

func validateReport(line string) bool {
	ints, succ := convertStringToIntSlice(line)
	if !succ {
		log.Fatal("Failed to convert string to ints.")
	}
	if isReportSafe(ints) {
		return true
	}

	for i := 0; i < len(ints); i++ {
		slice := make([]int, 0, len(ints))
		slice = append(slice, ints[:i]...)
		slice = append(slice, ints[i+1:]...)
		if isReportSafe(slice) {
			return true
		}
	}
	return false
}

func isReportSafe(report []int) bool {
	direction := Undetermined
	for i := 1; i < len(report); i++ {
		prev := report[i-1]
		curr := report[i]
		diff := absolute(prev - curr)
		if diff < 1 || diff > 3 {
			return false
		}
		if direction == Undetermined {
			switch {
			case prev > curr:
				direction = Smaller
			case prev < curr:
				direction = Larger
			default:
				return false
			}
		}

		switch direction {
		case Smaller:
			if prev < curr {
				return false
			}
		case Larger:
			if prev > curr {
				return false
			}
		default:
			return false
		}
	}

	return true
}

func convertStringToIntSlice(line string) ([]int, bool) {
	lineSlice := strings.Split(line, " ")
	ints := make([]int, 0, len(lineSlice))

	for i := 0; i < len(lineSlice); i++ {
		value, err := strconv.Atoi(lineSlice[i])
		if err != nil {
			log.Println("Failed to convert string to int slice", err)
			return ints, false
		}
		ints = append(ints, value)
	}

	return ints, true
}

func absolute(value int) int {
	if value > 0 {
		return value
	}
	return value * -1
}
