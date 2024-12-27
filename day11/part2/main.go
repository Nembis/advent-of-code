package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var cache = make(map[[2]uint64]uint64, 0)

func main() {
	numbers := parseFile("../input.txt")
	var total uint64 = 0
	for _, number := range numbers {
		total += count(uint64(number), uint64(75))
	}
	fmt.Println("Total:", total)
}

func count(stone, steps uint64) uint64 {
	if steps == 0 {
		return uint64(1)
	}

	if value, found := cache[[2]uint64{stone, steps}]; found {
		return value
	}

	if stone == 0 {
		return count(1, steps-1)
	}

	stoneString := fmt.Sprintf("%d", stone)
	if len(stoneString)%2 == 0 {
		leftHalf, err := strconv.Atoi(stoneString[:len(stoneString)/2])
		if err != nil {
			log.Fatal(err)
		}
		rightHalf, err := strconv.Atoi(stoneString[len(stoneString)/2:])
		if err != nil {
			log.Fatal(err)
		}
		leftResult := count(uint64(leftHalf), steps-1)
		rightResult := count(uint64(rightHalf), steps-1)
		cache[[2]uint64{uint64(leftHalf), steps - 1}] = leftResult
		cache[[2]uint64{uint64(rightHalf), steps - 1}] = rightResult
		return leftResult + rightResult
	}
	result := count(stone*2024, steps-1)
	cache[[2]uint64{stone * 2024, steps - 1}] = result
	return result
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
