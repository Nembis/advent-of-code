package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	line := parseFile("../input.txt")

	data := uncompressLine(line)
	data = reorderData(data)

	checksum := calcChecksum(data)
	fmt.Println("total: ", checksum)
}

func calcChecksum(data []int) int {
	checksum := 0

	for i, value := range data {
		if value == -1 {
			break
		}
		checksum += i * value
	}

	return checksum
}

func reorderData(data []int) []int {
	i := 0
	j := len(data) - 1

	for {
		for data[i] != -1 && i+1 < len(data) {
			i++
		}
		for data[j] == -1 && j >= 0 {
			j--
		}
		if i > j {
			break
		}
		data[i], data[j] = data[j], data[i]
	}

	return data
}

func uncompressLine(line string) []int {
	data := make([]int, 0)

	currId := 0
	isFileData := true
	for _, char := range line {
		if isFileData {
			amount := int(char - '0')
			for i := 0; i < amount; i++ {
				data = append(data, currId)
			}
			currId++
			isFileData = false
		} else {
			amount := int(char - '0')
			for i := 0; i < amount; i++ {
				data = append(data, -1)
			}
			isFileData = true
		}
	}

	return data
}

func parseFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}
