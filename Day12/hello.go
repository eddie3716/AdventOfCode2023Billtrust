package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	gridPatterns := parseFile("testinput.txt")

	part1(gridPatterns)

	part2(gridPatterns)
}

type Pattern []int
type Patterns []Pattern

type Templates []string
type Grid []Templates

type GridPatterns struct {
	grid    Grid
	pattern Patterns
}

func (gp GridPatterns) PrintGridPatterns() {
	for index, row := range gp.grid {
		fmt.Print(row)
		fmt.Println(" ", gp.pattern[index])
	}
}

func parseFile(fileName string) GridPatterns {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := Grid{}
	patterns := Patterns{}
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		templates := Templates{}
		for _, token := range strings.Split(tokens[0], ".") {
			if token != "" {
				templates = append(templates, token)
			}
		}
		grid = append(grid, templates)
		pattern := Pattern{}
		for _, token := range strings.Split(tokens[1], ",") {
			number, _ := strconv.Atoi(token)
			pattern = append(pattern, number)
		}
		patterns = append(patterns, pattern)
	}

	gridPatterns := GridPatterns{grid, patterns}
	gridPatterns.PrintGridPatterns()
	return gridPatterns
}

func part1(gridPatterns GridPatterns) {
	matches := 0
	//combinations := map[int]string{}

	for index, templates := range gridPatterns.grid {
		spacesRequired := len(gridPatterns.pattern[index]) - 1
		spacesAvailable := len(templates) - 1
		for _, template := range templates {
			spacesAvailable += len(template)
		}
		for _, template := range templates {
			for _, pattern := range gridPatterns.pattern[index] {
				if template == string(pattern) {
					matches++
				}
			}
		}
	}

	// for _, pattern := range gridPatterns.pattern {
	// 	for _, template := range gridPatterns.grid {

	// 	}
	// }

	fmt.Println(matches)

}
func part2(gridPatterns GridPatterns) {

}
