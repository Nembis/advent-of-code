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
			if len(value) > 1 {
				countDict[fmt.Sprintf("%v,%v", curr[0], curr[1])] = struct{}{}
			}
			for j, next := range value {
				if i == j {
					continue
				}

				rowDif := curr[0] - next[0]
				colDif := curr[1] - next[1]

				row := curr[0] + rowDif
				col := curr[1] + colDif

				for {
					if isInGrid(grid, row, col) {
						countDict[fmt.Sprintf("%v,%v", row, col)] = struct{}{}
						if grid[row][col] == '.' {
							grid[row][col] = '#'
						} else {
							grid[row][col] = '@'
						}
						row += rowDif
						col += colDif
					} else {
						break
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
