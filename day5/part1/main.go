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
	rules, lines := parseFile("../input.txt")

	total := 0

	for _, line := range lines {
		if isValidLine(rules, line) {
			middle := len(line) / 2
			middleValue, err := strconv.Atoi(line[middle])
			if err != nil {
				log.Fatal()
			}
			total += middleValue
		}
	}

	fmt.Println("Total: ", total)
}

func isValidLine(rules [][]string, line []string) bool {
	for i := 0; i < len(line); i++ {
		char := line[i]
		for j := i; j < len(line); j++ {
			for _, rule := range rules {
				if rule[1] == char && rule[0] == line[j] {
					return false
				}
			}
		}
	}

	return true
}

func parseFile(fileName string) ([][]string, [][]string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rules := make([][]string, 0)
	lines := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	isRule := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isRule = false
			continue
		}

		if isRule {
			ruleParts := strings.Split(line, "|")
			rules = append(rules, []string{ruleParts[0], ruleParts[1]})
		} else {
			lines = append(lines, strings.Split(line, ","))
		}
	}

	return rules, lines
}
