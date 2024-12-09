package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	North = iota
	East
	South
	West
)

func main() {
	maze, x, y := parseFile("../input.txt")

	count := countPath(maze, x, y)

	for _, value := range maze {
		for _, val := range value {
			fmt.Printf("%c", val)
		}
		fmt.Println()
	}

	fmt.Println("Total: ", count)
}

func countPath(maze [][]rune, startX, startY int) int {
	direction := North
	x, y := startX, startY
	count := 0

	for {
		if maze[x][y] == '.' || maze[x][y] == '^' {
			maze[x][y] = 'X'
			count++
		}

		if isLeavingMaze(maze, x, y, direction) {
			break
		}

		switch direction {
		case North:
			if maze[x-1][y] == '#' {
				direction = East
				continue
			}
			x -= 1

		case East:
			if maze[x][y+1] == '#' {
				direction = South
				continue
			}
			y += 1

		case South:
			if maze[x+1][y] == '#' {
				direction = West
				continue
			}
			x += 1

		case West:
			if maze[x][y-1] == '#' {
				direction = North
				continue
			}
			y -= 1

		default:
			log.Fatal("How did you break out of the maze and direction?")
		}
	}

	return count
}

func isLeavingMaze(maze [][]rune, x, y, direction int) bool {
	if direction == North && x == 0 {
		return true
	}

	if direction == East && y+1 == len(maze[x]) {
		return true
	}

	if direction == South && x+1 == len(maze) {
		return true
	}

	if direction == West && y == 0 {
		return true
	}

	return false
}

func parseFile(fileName string) ([][]rune, int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	maze := make([][]rune, 0)
	var x, y int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maze = append(maze, []rune{})
		i := len(maze) - 1
		for idx, char := range scanner.Text() {
			if char == '^' {
				x, y = i, idx
			}

			maze[i] = append(maze[i], char)
		}
	}

	return maze, x, y
}
