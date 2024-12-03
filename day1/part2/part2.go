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
		log.Fatal("Error while reading input file: ", err)
	}
	defer file.Close()

	leftSlice := make([]int, 0)
	dict := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		left, right := getValueFromLine(scanner.Text())
		leftSlice = append(leftSlice, left)
		dict[right]++
	}

	total := 0
	for _, value := range leftSlice {
		freqency := dict[value]
		total += freqency * value
	}

	fmt.Println("Total: ", total)
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
