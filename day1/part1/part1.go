package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error while reading input file: ", err)
	}
	defer file.Close()

	leftSlice := make([]int, 0)
	rightSlice := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		left, right := getValueFromLine(scanner.Text())
		leftSlice = append(leftSlice, left)
		rightSlice = append(rightSlice, right)
	}

	sort.Slice(leftSlice, func(i, j int) bool {
		return leftSlice[i] < leftSlice[j]
	})

	sort.Slice(rightSlice, func(i, j int) bool {
		return rightSlice[i] < rightSlice[j]
	})

	totalDif := 0
	for i := 0; i < len(leftSlice); i++ {
		if leftSlice[i] < rightSlice[i] {
			totalDif += rightSlice[i] - leftSlice[i]
		} else {
			totalDif += leftSlice[i] - rightSlice[i]
		}
	}

	fmt.Println("Total Dif:", totalDif)
}

func getValueFromLine(line string) (int, int) {
	slice := strings.Split(line, "   ")

	left, err := strconv.Atoi(slice[0])
	if err != nil {
		log.Fatal("Converting value in list error")
	}

	right, err := strconv.Atoi(slice[1])
	if err != nil {
		log.Fatal("Converting value in list error")
	}

	return left, right
}
