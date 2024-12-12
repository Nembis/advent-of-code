package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		if value, match := evaluateEquation(scanner.Text()); match {
			total += value
		}
	}

	fmt.Println("total: ", total)
}

func evaluateEquation(line string) (int, bool) {
	leftSide := strings.Split(line, ":")[0]
	rightSide := strings.TrimSpace(strings.Split(line, ":")[1])
	rightSlice := strings.Split(rightSide, " ")

	leftNumber, err := strconv.Atoi(leftSide)
	if err != nil {
		log.Fatal(err)
	}
	rightNumbers := make([]int, 0, len(rightSlice))
	for _, value := range rightSlice {
		num, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		rightNumbers = append(rightNumbers, num)
	}

	optionMap := make(map[int]struct{})
	calculateAllOptions(rightNumbers, optionMap, 1, rightNumbers[0])

	if _, found := optionMap[leftNumber]; found {
		return leftNumber, true
	}

	return 0, false
}

func calculateAllOptions(numbers []int, calcOptions map[int]struct{}, i, total int) {
	if i+1 == len(numbers) {
		calcOptions[total+numbers[i]] = struct{}{}
		calcOptions[total*numbers[i]] = struct{}{}
	} else {
		calculateAllOptions(numbers, calcOptions, i+1, total+numbers[i])
		calculateAllOptions(numbers, calcOptions, i+1, total*numbers[i])
	}
}
