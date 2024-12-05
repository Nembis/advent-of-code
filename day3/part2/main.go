package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pattern := `mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`

	scanner := bufio.NewScanner(file)
	total := 0
	enabled := true
	for scanner.Scan() {
		matches := findMatches(scanner.Text(), pattern)
		total += calculateMul(matches, &enabled)
	}

	fmt.Println("Total:", total)
}

func calculateMul(muls []string, enabled *bool) int {
	pattern := `\d{1,3}`

	total := 0
	for _, item := range muls {
		switch item {
		case "do()":
			*enabled = true
		case "don't()":
			*enabled = false
		default:
			if *enabled {
				matches := findMatches(item, pattern)
				left, err := strconv.Atoi(matches[0])
				if err != nil {
					log.Fatal(err)
				}
				right, err := strconv.Atoi(matches[1])
				if err != nil {
					log.Fatal(err)
				}

				total += left * right

			}
		}

	}

	return total
}

func findMatches(line, pattern string) []string {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	return reg.FindAllString(line, -1)
}
