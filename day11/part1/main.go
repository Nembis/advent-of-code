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
	numbers := parseFile("../input.txt")
	numbers = walkThroughSlice(numbers, 25)
	total := len(numbers)
	fmt.Println("Total:", total)
}

func walkThroughSlice(numbers []int, steps int) []int {

	for i := 0; i < steps; i++ {
		newSlice := make([]int, 0, len(numbers))
		for _, num := range numbers {
			if num == 0 {
				newSlice = append(newSlice, num+1)
			} else if len(fmt.Sprintf("%d", num))%2 == 0 {
				numStr := fmt.Sprintf("%d", num)
				leftHalf, err := strconv.Atoi(numStr[:len(numStr)/2])
				if err != nil {
					log.Fatal(err)
				}
				rightHalf, err := strconv.Atoi(numStr[len(numStr)/2:])
				if err != nil {
					log.Fatal(err)
				}
				newSlice = append(newSlice, leftHalf)
				newSlice = append(newSlice, rightHalf)
			} else {
				newSlice = append(newSlice, num*2024)
			}
		}

		numbers = newSlice
	}

	return numbers
}

func parseFile(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	items := strings.Split(scanner.Text(), " ")
	slice := make([]int, len(items))

	for i, item := range items {
		num, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		slice[i] = num
	}

	return slice
}
