package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	grid, heads := parseFile("../input.txt")
	total := findValidTrails(grid, heads)
	fmt.Println("Total:", total)
}

func findValidTrails(grid, heads [][]int) int {
	count := 0
	for _, cords := range heads {
		tracker := make(map[string]struct{})
		followTrailRecur(tracker, grid, cords[0], cords[1])
		count += len(tracker)
	}
	return count
}

func followTrailRecur(tracker map[string]struct{}, grid [][]int, row, col int) {
	currValue := grid[row][col]
	if currValue == 9 {
		tracker[fmt.Sprintf("%v,%v", row, col)] = struct{}{}
		return
	}

	if row-1 >= 0 && currValue+1 == grid[row-1][col] {
		followTrailRecur(tracker, grid, row-1, col)
	}
	if col+1 < len(grid[row]) && currValue+1 == grid[row][col+1] {
		followTrailRecur(tracker, grid, row, col+1)
	}
	if row+1 < len(grid) && currValue+1 == grid[row+1][col] {
		followTrailRecur(tracker, grid, row+1, col)
	}
	if col-1 >= 0 && currValue+1 == grid[row][col-1] {
		followTrailRecur(tracker, grid, row, col-1)
	}
}

func parseFile(filePath string) ([][]int, [][]int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make([][]int, 0)
	heads := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		grid = append(grid, []int{})
		line := scanner.Text()
		for j, char := range line {
			num := int(char - '0')
			grid[i] = append(grid[i], num)
			if num == 0 {
				heads = append(heads, []int{i, j})
			}
		}
		i++
	}

	return grid, heads
}
