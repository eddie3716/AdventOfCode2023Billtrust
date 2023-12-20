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
	startingRow  int
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

	deltaFromStart := 0
	maxDeltaFromStart := 0
	for _, instruction := range digPlan.instructions {
		switch instruction.direction {
		case "R":
			deltaCols += instruction.steps
			maxCols = max(maxCols, deltaCols)
		case "L":
			deltaCols -= instruction.steps
		case "U":
			deltaFromStart += instruction.steps
			maxDeltaFromStart = max(maxDeltaFromStart, deltaFromStart)
			deltaRows -= instruction.steps
		case "D":
			deltaFromStart -= instruction.steps
			maxDeltaFromStart = max(maxDeltaFromStart, deltaFromStart)
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

	digPlan.maxRows = maxRows + 1
	digPlan.maxCols = maxCols + 1
	digPlan.startingRow = maxDeltaFromStart

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
	parseFile("input.txt")

	fmt.Println("answer part 1:", findMinHeat())

	//fmt.Println("answer part 2:", findMinHeat())
}

func (digPlan *Digplan) String() string {
	var sb strings.Builder
	for _, row := range digPlan.grid {
		for _, col := range row {
			sb.WriteRune(col)
		}
		sb.WriteString("\n")
	}
	return sb.String()
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
	fmt.Println(digPlan.String())
	answer := 0

	currentRow := digPlan.startingRow
	currentCol := 0
	digPlan.grid[currentRow][currentCol] = '#'
	for _, instruction := range digPlan.instructions {
		switch instruction.direction {
		case "R":
			for i := 0; i < instruction.steps; i++ {
				currentCol++
				digPlan.grid[currentRow][currentCol] = '#'
				answer++
			}
		case "L":
			for i := 0; i < instruction.steps; i++ {
				currentCol--
				digPlan.grid[currentRow][currentCol] = '#'
				answer++
			}
		case "U":
			for i := 0; i < instruction.steps; i++ {
				currentRow--
				digPlan.grid[currentRow][currentCol] = '#'
				answer++
			}
		case "D":
			for i := 0; i < instruction.steps; i++ {
				currentRow++
				digPlan.grid[currentRow][currentCol] = '#'
				answer++
			}
		}
	}
	fmt.Println(digPlan.String())
	for i := 0; i < digPlan.maxRows; i++ {
		previousChar := '?'
		inside := false
		for j := 0; j < digPlan.maxCols; j++ {

			currentChar := digPlan.grid[i][j]

			if previousChar == '#' && currentChar == '.' {
				inside = !inside
			}

			previousChar = currentChar
			if inside && currentChar == '.' {
				digPlan.grid[i][j] = '#'
				answer++
			}
		}
	}

	fmt.Println(digPlan.String())

	//fmt.Println(digPlan)
	return answer
}
