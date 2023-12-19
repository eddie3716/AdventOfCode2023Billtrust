package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instructions struct {
	direction string
	steps     int
}

type Digplan struct {
	instructions []Instructions
	colors       []string
	maxRows      int
	maxCols      int
	grid         [][]rune
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (digPlan *Digplan) generateGrid() {
	maxRows := 0
	maxCols := 0
	deltaRows := 0
	deltaCols := 0

	for _, instruction := range digPlan.instructions {
		switch instruction.direction {
		case "R":
			deltaCols += instruction.steps
			maxCols = max(maxCols, deltaCols)
		case "L":
			deltaCols -= instruction.steps
		case "U":
			deltaRows -= instruction.steps
		case "D":
			deltaRows += instruction.steps
			maxRows = max(maxRows, deltaRows)
		}
	}

	if maxRows <= 0 {
		panic("negative rows")
	}

	if maxCols <= 0 {
		panic("negative cols")
	}

	digPlan.maxRows = maxRows
	digPlan.maxCols = maxCols

	digPlan.grid = make([][]rune, digPlan.maxRows)
	for i := range digPlan.grid {
		digPlan.grid[i] = make([]rune, digPlan.maxCols)
		for j := range digPlan.grid[i] {
			digPlan.grid[i][j] = '.'
		}
	}
}

var digPlan Digplan = Digplan{colors: []string{}, instructions: []Instructions{}}

func main() {
	parseFile("testinput.txt")

	fmt.Println("answer part 1:", findMinHeat())

	fmt.Println("answer part 2:", findMinHeat())
}

func parseFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		direction := tokens[0]
		steps, _ := strconv.Atoi(tokens[1])
		instruction := Instructions{direction: direction, steps: steps}
		digPlan.instructions = append(digPlan.instructions, instruction)
		color := tokens[2][1:]
		digPlan.colors = append(digPlan.colors, color)
	}
	digPlan.generateGrid()
}

func findMinHeat() int {

	answer := 0

	fmt.Println(digPlan)
	return answer
}
