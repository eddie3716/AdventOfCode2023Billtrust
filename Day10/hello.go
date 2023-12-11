package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"sync"
)

type Position struct {
	Row    int
	Col    int
	Symbol rune
}

var North = Position{Row: -1, Col: 0}
var South = Position{Row: 1, Col: 0}
var West = Position{Row: 0, Col: -1}
var East = Position{Row: 0, Col: 1}
var Compass = []Position{North, South, West, East}
var VerticalPipes = []rune{'|', 'F', 'L', '7', 'J', 'S'}
var HorizontalPipes = []rune{'-', 'F', '7', 'L', 'J', 'S'}

func (p Position) Traverse(o Position, grid [][]rune) (Position, error) {
	newRow := p.Row + o.Row
	newCol := p.Col + o.Col
	if newRow < 0 || newCol < 0 || newRow >= len(grid) || newCol >= len(grid[newRow]) {
		return Position{}, errors.New("Out of bounds")
	}
	return Position{Row: newRow, Col: newCol, Symbol: grid[newRow][newCol]}, nil
}

func (p Position) GetNextPositions(currentSymbol rune, grid [][]rune, excludePostion Position, directions map[Position][]rune) ([]Position, error) {
	newPositions := []Position{}
	for _, traverse := range Compass {
		newPosition, err := p.Traverse(traverse, grid)
		if newPosition == excludePostion {
			continue
		}
		if err == nil && slices.Contains(directions[traverse], grid[newPosition.Row][newPosition.Col]) {
			newPositions = append(newPositions, newPosition)
		}
	}
	if len(newPositions) == 0 {
		//fmt.Println("No valid directions ", p, currentSymbol, directions)
		return newPositions, errors.New("No valid directions")
	}
	return newPositions, nil
}

func parseFile(fileName string) [][]rune {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	records := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		records = append(records, []rune(line))
	}

	return records
}

func printMain(main [][]Position) {
	fmt.Println()
	for _, row := range main {
		for _, position := range row {
			fmt.Print(string(position.Symbol))
		}
		fmt.Println()
	}
}

func part1(grid [][]rune, connectionPoints map[rune]map[Position][]rune) []Position {
	start := Position{Row: 0, Col: 0, Symbol: 'S'}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'S' {
				start.Col = col
				start.Row = row
				goto foundStart
			}
		}
	}
foundStart:

	startingPositions, err := start.GetNextPositions('S', grid, Position{Row: 0, Col: 0, Symbol: 'S'}, connectionPoints['S'])

	if err != nil {
		panic(err)
	}

	ch := make(chan []Position, len(startingPositions))

	wg := sync.WaitGroup{}
	for _, position := range startingPositions {
		wg.Add(1)
		go func(p Position) {
			defer wg.Done()
			count := 1
			var err error = nil
			oldPosition := p
			path := []Position{}
			path = append(path, p)
			for err == nil {
				currentSymbol := grid[p.Row][p.Col]
				ps, er := p.GetNextPositions(currentSymbol, grid, oldPosition, connectionPoints[currentSymbol])
				err = er
				if err != nil || len(ps) == 0 {
					//fmt.Println("No more directions ", p, currentSymbol, connectionPoints[currentSymbol])
					ch <- path
					//fmt.Println(path)
					return
				}
				if (len(ps)) > 1 {
					fmt.Println("More than one direction ", p, currentSymbol, connectionPoints[currentSymbol])
					panic("More than one direction")
				}
				oldPosition = p
				p = ps[0]
				path = append(path, p)
				count++
			}
		}(position)
	}

	wg.Wait()
	close(ch)

	allSteps := [][]Position{}
	for steps := range ch {
		allSteps = append(allSteps, steps)
	}
	var maxSteps int = 0
	for _, steps := range allSteps {
		for _, otherSteps := range allSteps {
			for i := 0; i < len(steps); i++ {
				for j := 0; j < len(otherSteps); j++ {
					if (min(i+1, len(steps)-1)-max(j-1, 0)) == 2 && steps[i] == otherSteps[j] && steps[min(i+1, len(steps)-1)] == otherSteps[max(j-1, 0)] {
						maxSteps = max(maxSteps, max(i, j)+1)
					}
				}
			}
		}
	}

	fmt.Println("Max steps ", maxSteps)

	allSteps = append(allSteps, []Position{start})
	loop := append(allSteps[0], start)
	loop = append(loop, start)
	return loop
}

func part2(loop []Position, grid [][]rune, connectionPoints map[rune]map[Position][]rune) {
	main := [][]Position{}

	rows := len(grid)
	cols := len(grid[0])

	//start with a clean slate
	for i := 0; i < rows; i++ {
		main = append(main, []Position{})
		for j := 0; j < cols; j++ {
			main[i] = append(main[i], Position{Row: i, Col: j, Symbol: '.'})
		}
	}

	//everything else besides the loop is trash, including broken pipes
	for _, position := range loop {
		main[position.Row][position.Col] = position
	}

	printMain(main)

	for _, row := range main {
		outside := true
		for index, position := range row {
			if position.Symbol == '.' && outside {
				position.Symbol = 'O'
				row[index] = position
			} else if position.Symbol == '.' && !outside {
				position.Symbol = 'I'
				row[index] = position
			} else if slices.Contains(VerticalPipes, position.Symbol) && position.Symbol != 'L' && position.Symbol != 'J' {
				outside = !outside
			}
		}
	}

	insideCount := 0
	for _, row := range main {
		for _, position := range row {
			if position.Symbol == 'I' {
				insideCount++
			}
		}
	}

	printMain(main)
	fmt.Println("Inside count ", insideCount)
}

func main() {
	grid := parseFile("input.txt")

	connectingPoints := make(map[rune]map[Position][]rune)
	connectionsS := make(map[Position][]rune)
	connectionsS[North] = []rune{'|', '7', 'F'}
	connectionsS[South] = []rune{'|', 'L', 'J'}
	connectionsS[West] = []rune{'-', 'F', 'L'}
	connectionsS[East] = []rune{'-', '7', 'J'}
	connectingPoints['S'] = connectionsS
	connectionsPipe := make(map[Position][]rune)
	connectionsPipe[North] = []rune{'|', '7', 'F'}
	connectionsPipe[South] = []rune{'|', 'L', 'J'}
	connectionsPipe[West] = []rune{}
	connectionsPipe[East] = []rune{}
	connectingPoints['|'] = connectionsPipe
	connectionsDash := make(map[Position][]rune)
	connectionsDash[North] = []rune{}
	connectionsDash[South] = []rune{}
	connectionsDash[West] = []rune{'-', 'F', 'L'}
	connectionsDash[East] = []rune{'-', '7', 'J'}
	connectingPoints['-'] = connectionsDash
	connectionsF := make(map[Position][]rune)
	connectionsF[North] = []rune{}
	connectionsF[South] = []rune{'|', 'L', 'J'}
	connectionsF[West] = []rune{}
	connectionsF[East] = []rune{'-', '7', 'J'}
	connectingPoints['F'] = connectionsF
	connectionsL := make(map[Position][]rune)
	connectionsL[North] = []rune{'|', '7', 'F'}
	connectionsL[South] = []rune{}
	connectionsL[West] = []rune{}
	connectionsL[East] = []rune{'-', '7', 'J'}
	connectingPoints['L'] = connectionsL
	connectionsJ := make(map[Position][]rune)
	connectionsJ[North] = []rune{'|', '7', 'F'}
	connectionsJ[South] = []rune{}
	connectionsJ[West] = []rune{'-', 'L', 'F'}
	connectionsJ[East] = []rune{}
	connectingPoints['J'] = connectionsJ
	connections7 := make(map[Position][]rune)
	connections7[North] = []rune{}
	connections7[South] = []rune{'|', 'L', 'J'}
	connections7[West] = []rune{'-', 'L', 'F'}
	connections7[East] = []rune{}
	connectingPoints['7'] = connections7
	connectionsDot := make(map[Position][]rune)
	connectionsDot[North] = []rune{'.'}
	connectionsDot[South] = []rune{'.'}
	connectionsDot[East] = []rune{'.'}
	connectionsDot[West] = []rune{'.'}
	connectingPoints['.'] = connectionsDot

	loop := part1(grid, connectingPoints)

	part2(loop, grid, connectingPoints)
}
