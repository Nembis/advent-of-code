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
	Undetermined = iota
	Smaller
	Larger
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

	fmt.Println("Total:", total)
}

func validateReport(line string) bool {
	reportSlice := strings.Split(line, " ")
	item, err := strconv.Atoi(reportSlice[0])
	if err != nil {
		log.Println("Failed to convert string to int:", err)
		return false
	}
	prevItem := item
	direction := Undetermined

	for i := 1; i < len(reportSlice); i++ {
		currItem, err := strconv.Atoi(reportSlice[i])
		if err != nil {
			log.Println("Failed to convert string to int:", err)
			return false
		}

		diff := prevItem - currItem
		if diff < -3 || diff == 0 || diff > 3 {
			return false
		}

		if direction == Undetermined {
			if prevItem > currItem {
				direction = Smaller
			} else {
				direction = Larger
			}
		}

		if (direction == Smaller && prevItem < currItem) ||
			(direction == Larger && prevItem > currItem) {
			return false
		}

		prevItem = currItem
	}

	return true
}
