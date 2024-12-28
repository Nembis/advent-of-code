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
				area, perimeter := findAreaSide(grid, trackingGrid, i, j, char)
				fmt.Printf("Character: %c, Area: %d, Side: %d\n", char, area, perimeter)
				total += (area * perimeter)
			}
		}
	}

	fmt.Println("Total:", total)

}

func findAreaSide(grid [][]rune, trackingGrid [][]bool, row, col int, char rune) (int, int) {
	trackingGrid[row][col] = true
	area := 1
	corners := 0

	if row == 0 && col == 0 {
		corners++
	} else if row == 0 && col+1 == len(grid[row]) {
		corners++
	} else if row+1 == len(grid) && col == 0 {
		corners++
	} else if row+1 == len(grid) && col+1 == len(grid[row]) {
		corners++
	}

	if row == 0 {
		if col+1 != len(grid[row]) && grid[row][col+1] != char {
			corners++
		}
		if col-1 >= 0 && grid[row][col-1] != char {
			corners++
		}
	}
	if col+1 == len(grid[row]) {
		if row+1 != len(grid) && grid[row+1][col] != char {
			corners++
		}
		if row-1 >= 0 && grid[row-1][col] != char {
			corners++
		}
	}
	if row+1 == len(grid) {
		if col+1 != len(grid[col]) && grid[row][col+1] != char {
			corners++
		}
		if col-1 >= 0 && grid[row][col-1] != char {
			corners++
		}
	}
	if col == 0 {
		if row+1 != len(grid) && grid[row+1][col] != char {
			corners++
		}
		if row-1 >= 0 && grid[row-1][col] != char {
			corners++
		}
	}

	if row-1 >= 0 && col+1 != len(grid[row]) &&
		((grid[row-1][col] != char && grid[row][col+1] != char) ||
			(grid[row-1][col] == char && grid[row][col+1] == char && grid[row-1][col+1] != char)) {
		corners++
	}
	if row+1 != len(grid) && col+1 != len(grid[row]) &&
		((grid[row+1][col] != char && grid[row][col+1] != char) ||
			(grid[row+1][col] == char && grid[row][col+1] == char && grid[row+1][col+1] != char)) {
		corners++
	}
	if row+1 != len(grid) && col-1 >= 0 &&
		((grid[row+1][col] != char && grid[row][col-1] != char) ||
			(grid[row+1][col] == char && grid[row][col-1] == char && grid[row+1][col-1] != char)) {
		corners++
	}
	if row-1 >= 0 && col-1 >= 0 &&
		((grid[row-1][col] != char && grid[row][col-1] != char) ||
			(grid[row-1][col] == char && grid[row][col-1] == char && grid[row-1][col-1] != char)) {
		corners++
	}

	if row-1 >= 0 && !trackingGrid[row-1][col] && grid[row-1][col] == char {
		addArea, addPerimeter := findAreaSide(grid, trackingGrid, row-1, col, char)
		area += addArea
		corners += addPerimeter
	}
	if col+1 != len(grid[row]) && !trackingGrid[row][col+1] && grid[row][col+1] == char {
		addArea, addPerimeter := findAreaSide(grid, trackingGrid, row, col+1, char)
		area += addArea
		corners += addPerimeter
	}
	if row+1 != len(grid) && !trackingGrid[row+1][col] && grid[row+1][col] == char {
		addArea, addPerimeter := findAreaSide(grid, trackingGrid, row+1, col, char)
		area += addArea
		corners += addPerimeter
	}
	if col-1 >= 0 && !trackingGrid[row][col-1] && grid[row][col-1] == char {
		addArea, addPerimeter := findAreaSide(grid, trackingGrid, row, col-1, char)
		area += addArea
		corners += addPerimeter
	}

	return area, corners
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
