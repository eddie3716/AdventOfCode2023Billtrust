package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	x int
	m int
	a int
	s int
}

type Workflow struct {
}

func part1(parts []Part, workflows map[string]Workflow) int {
	return 0
}

func part2(parts []Part, workflows map[string]Workflow) int {
	return 0
}

func main() {
	parts, workflows := parseFile("input.txt")

	fmt.Println("answer part 1:", part1(parts, workflows))

	fmt.Println("answer part 2:", part2(parts, workflows))
}

func parseFile(fileName string) ([]Part, map[string]Workflow) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	digPlan := Digplan{realInstructions: []Instruction{}, instructions: []Instruction{}}

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		direction := tokens[0]
		steps, _ := strconv.Atoi(tokens[1])
		instruction := Instruction{direction: direction, steps: int64(steps)}
		digPlan.instructions = append(digPlan.instructions, instruction)
		hexSteps := tokens[2][2 : len(tokens[2])-2]
		realSteps, err := strconv.ParseInt(hexSteps, 16, 32)
		dirCode := tokens[2][len(tokens[2])-2 : len(tokens[2])-1]
		realDir := ""
		switch dirCode {
		case "0":
			realDir = "R"
			break
		case "1":
			realDir = "D"
			break
		case "2":
			realDir = "L"
			break
		case "3":
			realDir = "U"
			break
		}
		realInstruction := Instruction{direction: realDir, steps: realSteps}

		if err != nil {
			panic(err)
		}
		digPlan.realInstructions = append(digPlan.realInstructions, realInstruction)
	}

	fmt.Println(digPlan)

	return digPlan
}
