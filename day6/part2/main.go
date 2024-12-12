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

// answer 2266 incorrect
func main() {
	maze, x, y := parseFile("../input.txt")
	count := findInfiniteLoops(maze, x, y)
	fmt.Println("Count: ", count)
}

func findInfiniteLoops(maze [][]rune, x, y int) int {
	direction := North
	startX, startY, startDirection := x, y, direction
	prevX, prevY := x, y
	count := 0
	seenCords := make(map[string]struct{})

	for {
		cordString := fmt.Sprintf("%v,%v", x, y)

		if maze[x][y] != '^' && (prevX != x || prevY != y) {
			if _, found := seenCords[cordString]; !found {
				if isInfiniteLoop(maze, startX, startY, x, y, startDirection) {
					count++
				}

			}
		}

		seenCords[cordString] = struct{}{}

		prevX, prevY = x, y
		if !move(maze, &x, &y, &direction) {
			break
		}
	}

	return count
}

func isInfiniteLoop(maze [][]rune, x, y, obsX, obsY, direction int) bool {
	newMaze := copyMazeDeep(maze)
	prevSeen := make(map[string]struct{})
	newMaze[obsX][obsY] = '#'

	for {
		currCordDir := fmt.Sprintf("%v,%v,%v", x, y, direction)
		if _, found := prevSeen[currCordDir]; found {
			return true
		}

		prevSeen[currCordDir] = struct{}{}

		if !move(newMaze, &x, &y, &direction) {
			return false
		}
	}
}

func move(maze [][]rune, x, y, direction *int) bool {
	switch *direction {
	case North:
		if *x <= 0 {
			return false
		}
		if maze[*x-1][*y] == '#' {
			*direction = East
		} else {
			*x -= 1
		}

	case East:
		if *y+1 >= len(maze[*x]) {
			return false
		}
		if maze[*x][*y+1] == '#' {
			*direction = South
		} else {
			*y += 1
		}

	case South:
		if *x+1 >= len(maze) {
			return false
		}
		if maze[*x+1][*y] == '#' {
			*direction = West
		} else {
			*x += 1
		}

	case West:
		if *y <= 0 {
			return false
		}
		if maze[*x][*y-1] == '#' {
			*direction = North
		} else {
			*y -= 1
		}
	}

	return true
}

func copyMazeDeep(maze [][]rune) [][]rune {
	dist := make([][]rune, len(maze))
	for i, line := range maze {
		newLine := make([]rune, len(line))
		copy(newLine, line)
		dist[i] = newLine
	}
	return dist
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
