package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	grid := parseFile("../input.txt")
	trackingGrid := createTrackingGrid(grid)

	total := 0

	for i, row := range trackingGrid {
		for j, seen := range row {
			if !seen {
				char := grid[i][j]
				area, perimeter := findArea(grid, trackingGrid, i, j, char)
				fmt.Printf("Character: %c, Area: %d, Perimeter: %d\n", char, area, perimeter)
				total += (area * perimeter)
			}
		}
	}

	fmt.Println("Total:", total)

}

func findArea(grid [][]rune, trackingGrid [][]bool, row, col int, char rune) (int, int) {
	trackingGrid[row][col] = true
	area := 1
	perimeter := 0

	if row == 0 || (row-1 >= 0 && grid[row-1][col] != char) {
		perimeter++
	}
	if col+1 == len(grid[row]) || (col+1 != len(grid[row]) && grid[row][col+1] != char) {
		perimeter++
	}
	if row+1 == len(grid) || (row+1 != len(grid) && grid[row+1][col] != char) {
		perimeter++
	}
	if col == 0 || (col-1 >= 0 && grid[row][col-1] != char) {
		perimeter++
	}

	if row-1 >= 0 && !trackingGrid[row-1][col] && grid[row-1][col] == char {
		addArea, addPerimeter := findArea(grid, trackingGrid, row-1, col, char)
		area += addArea
		perimeter += addPerimeter
	}
	if col+1 != len(grid[row]) && !trackingGrid[row][col+1] && grid[row][col+1] == char {
		addArea, addPerimeter := findArea(grid, trackingGrid, row, col+1, char)
		area += addArea
		perimeter += addPerimeter
	}
	if row+1 != len(grid) && !trackingGrid[row+1][col] && grid[row+1][col] == char {
		addArea, addPerimeter := findArea(grid, trackingGrid, row+1, col, char)
		area += addArea
		perimeter += addPerimeter
	}
	if col-1 >= 0 && !trackingGrid[row][col-1] && grid[row][col-1] == char {
		addArea, addPerimeter := findArea(grid, trackingGrid, row, col-1, char)
		area += addArea
		perimeter += addPerimeter
	}

	return area, perimeter
}

func createTrackingGrid(grid [][]rune) [][]bool {
	trackingGrid := make([][]bool, len(grid))

	for i, row := range grid {
		trackingGrid[i] = make([]bool, len(row))
	}

	return trackingGrid
}

func parseFile(filePath string) [][]rune {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, make([]rune, len(line)))
		for j, char := range line {
			grid[i][j] = char
		}
		i++
	}

	return grid
}
