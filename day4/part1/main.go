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
			if rune(slice[i][j]) == 'X' {
				total += containXmasCount(slice, i, j)

			}
		}
	}

	fmt.Println("total", total)
	return total
}

func containXmasCount(slice [][]rune, i, j int) int {
	count := 0

	if i >= 3 {
		if rune(slice[i-1][j]) == 'M' && rune(slice[i-2][j]) == 'A' && rune(slice[i-3][j]) == 'S' {
			count++
		}

		if j >= 3 {
			if rune(slice[i-1][j-1]) == 'M' && rune(slice[i-2][j-2]) == 'A' && rune(slice[i-3][j-3]) == 'S' {
				count++
			}
		}

		if j+3 < len(slice[i]) {
			if rune(slice[i-1][j+1]) == 'M' && rune(slice[i-2][j+2]) == 'A' && rune(slice[i-3][j+3]) == 'S' {
				count++
			}
		}
	}

	if i+3 < len(slice) {
		if rune(slice[i+1][j]) == 'M' && rune(slice[i+2][j]) == 'A' && rune(slice[i+3][j]) == 'S' {
			count++
		}

		if j >= 3 {
			if rune(slice[i+1][j-1]) == 'M' && rune(slice[i+2][j-2]) == 'A' && rune(slice[i+3][j-3]) == 'S' {
				count++
			}
		}

		if j+3 < len(slice[i]) {
			if rune(slice[i+1][j+1]) == 'M' && rune(slice[i+2][j+2]) == 'A' && rune(slice[i+3][j+3]) == 'S' {
				count++
			}
		}
	}

	if j >= 3 {
		if rune(slice[i][j-1]) == 'M' && rune(slice[i][j-2]) == 'A' && rune(slice[i][j-3]) == 'S' {
			count++
		}
	}

	if j+3 < len(slice[i]) {
		if rune(slice[i][j+1]) == 'M' && rune(slice[i][j+2]) == 'A' && rune(slice[i][j+3]) == 'S' {
			count++
		}
	}

	return count
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
