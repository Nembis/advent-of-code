package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "../input.txt"
	slice := sliceFromFile(filePath)
	countXmas(slice)
}

func countXmas(slice [][]rune) int {
	total := 0

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			if rune(slice[i][j]) == 'A' && containXmas(slice, i, j) {
				total++

			}
		}
	}

	fmt.Println("total", total)
	return total
}

func containXmas(slice [][]rune, i, j int) bool {
	if i == 0 || j == 0 || i+1 == len(slice) || j+1 == len(slice[i]) {
		return false
	}

	if !((slice[i-1][j-1] == 'M' && slice[i+1][j+1] == 'S') ||
		(slice[i-1][j-1] == 'S' && slice[i+1][j+1] == 'M')) {
		return false
	}

	if !((slice[i-1][j+1] == 'M' && slice[i+1][j-1] == 'S') ||
		(slice[i-1][j+1] == 'S' && slice[i+1][j-1] == 'M')) {
		return false
	}

	return true
}

func sliceFromFile(filePath string) [][]rune {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	slice2D := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		slice2D = append(slice2D, []rune{})
		for _, char := range line {
			slice2D[i] = append(slice2D[i], char)
		}
		i++
	}

	return slice2D
}
