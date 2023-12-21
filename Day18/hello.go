package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	direction string
	steps     int64
}

type Digplan struct {
	instructions     []Instruction
	realInstructions []Instruction
}

type point struct{ x, y float64 }

func shoelace(points []point) float64 {
	n := len(points)
	if n < 3 {
		return 0 // not a polygon
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		sum += (points[i].x * points[j].y) - (points[j].x * points[i].y)
	}
	return 0.5 * sum
}

func computeArea(instructions []Instruction) float64 {
	var x int64 = 0
	var y int64 = 0
	perimeter := 0.

	points := []point{}
	points = append(points, point{float64(x), float64(y)})
	for _, instruction := range instructions {
		perimeter += float64(instruction.steps)
		points = append(points, point{float64(x), float64(y)})
		switch instruction.direction {
		case "R":
			x += instruction.steps
		case "L":
			x -= instruction.steps
		case "U":
			y -= instruction.steps
		case "D":
			y += instruction.steps
		}
	}

	//Pick's theorem having a baby with shoelace
	area := shoelace(points) + perimeter/2 + 1

	return area
}

func main() {
	digPlan := parseFile("input.txt")

	fmt.Println("answer part 1:", int64(computeArea(digPlan.instructions)))

	fmt.Println("answer part 2:", int64(computeArea(digPlan.realInstructions)))
}

func parseFile(fileName string) Digplan {
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
