package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("Error while reading input file: ", err)
	}
	defer file.Close()

	mutationDict := make(map[string][]string)
	moleculesSet := make(map[string]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		left, right := getKeyValueFromLine(line)
		mutationDict[left] = append(mutationDict[left], right)
	}
	scanner.Scan()
	mainStr := scanner.Text()
	mainStrLen := len(mainStr)
	fmt.Println(mainStrLen)

	for key, value := range mutationDict {
		keyLen := len(key)
		fmt.Println("keyLen:", keyLen)
		fmt.Println("ValueLen:", value)
		for i := 0; i < mainStrLen; i++ {
			if keyLen+i > mainStrLen {
				fmt.Println("Stopping at ", i)
				fmt.Println()
				break
			}

			if mainStr[i:i+keyLen] == key {
				for _, replace := range value {
					newMolecule := replaceStringSection(mainStr, i, i+keyLen, replace)
					moleculesSet[newMolecule] = struct{}{}
				}
			}
		}
	}

	fmt.Println("Total:", len(moleculesSet))
}

func getKeyValueFromLine(line string) (string, string) {
	stringSlice := strings.Split(line, " => ")
	return stringSlice[0], stringSlice[1]
}

func replaceStringSection(original string, start int, end int, replaceString string) string {
	return original[:start] + replaceString + original[end:]
}
