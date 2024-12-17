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
			continue
		}
		checksum += i * value
	}

	return checksum
}

func reorderData(data []int) []int {
	j, currId := findLastId(data)

	for currId > -1 {
		j, fileSize := findFileSize(data, currId, j)
		i, foundSpace := findFreeSpace(data, fileSize, j)

		if foundSpace {
			for curr := 0; curr < fileSize; curr++ {
				data[i+curr], data[j+curr] = data[j+curr], data[i+curr]
			}
		}

		if currId >= 0 {
			currId--
		}
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

func findFileSize(data []int, currId, rightIdx int) (int, int) {
	fileSize := 0
	foundId := false
	for i := rightIdx; i >= 0; i-- {
		if data[i] == currId {
			fileSize++
			foundId = true
		} else if foundId {
			return i + 1, fileSize
		}
	}

	return 0, 0
}

func findFreeSpace(data []int, fileSize, rightIdx int) (int, bool) {
	freeSpaceIdx := 0
	freeSpace := 0
	for i, value := range data {
		if freeSpace == fileSize {
			return freeSpaceIdx, true
		}
		if i >= rightIdx {
			return -1, false
		}

		if value == -1 {
			if freeSpace == 0 {
				freeSpaceIdx = i
			}
			freeSpace++
		} else {
			freeSpace = 0
		}
	}
	return -1, false
}

func findLastId(data []int) (int, int) {
	idx := len(data) - 1
	for idx > 0 && data[idx] == -1 {
		idx--
	}

	return idx, data[idx]
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
