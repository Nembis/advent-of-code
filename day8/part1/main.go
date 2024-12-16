package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	grid, towers := parseFile("../input.txt")
	printGrid := func() {
		for i, row := range grid {
			fmt.Print(i)
			for _, col := range row {
				fmt.Printf("%c", col)
			}
			fmt.Println()
		}
		fmt.Println("=================================")
	}

	printGrid()
	total := findBadTowers(grid, towers)
	printGrid()

	fmt.Println("Total:", total)
}

func findBadTowers(grid [][]rune, towers map[rune][][]int) int {
	countDict := make(map[string]struct{})

	for _, value := range towers {
		for i, curr := range value {
			for j, next := range value {
				if i == j {
					continue
				}

				row := curr[0] + (curr[0] - next[0])
				col := curr[1] + (curr[1] - next[1])

				if isInGrid(grid, row, col) {
					countDict[fmt.Sprintf("%v,%v", row, col)] = struct{}{}
					if grid[row][col] == '.' {
						grid[row][col] = '#'
					} else {
						grid[row][col] = '@'
					}
				}
			}
		}
	}

	return len(countDict)
}

func isInGrid(grid [][]rune, row, col int) bool {
	return (row >= 0 && col >= 0 &&
		row < len(grid) && col < len(grid[row]))
}

func parseFile(filePath string) ([][]rune, map[rune][][]int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make([][]rune, 0)
	tower := make(map[rune][][]int)
	scanner := bufio.NewScanner(file)

	rowIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for i, char := range scanner.Text() {
			if char != '.' {
				tower[char] = append(tower[char], []int{rowIdx, i})
			}
			row[i] = char
		}
		grid = append(grid, row)
		rowIdx++
	}

	return grid, tower
}
