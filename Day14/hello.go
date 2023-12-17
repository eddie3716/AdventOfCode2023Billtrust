package main

import (
	"bufio"
	"fmt"
	"os"
)

const Cycles = 1000000000

//const Cycles = 3

func main() {
	sets := parseFile("input.txt")

	//part1(sets)

	part2(sets)
}

func PrintGrid(g [][]rune) {
	for row := 0; row < len(g); row++ {
		fmt.Println(string(g[row]))
	}
	fmt.Println()
}

func castNorth(g [][]rune, row int, col int) {
	if row == 0 {
		return
	}

	if g[row-1][col] == '.' {
		g[row-1][col] = 'O'
		g[row][col] = '.'
		castNorth(g, row-1, col)
	}
}

func castSouth(g [][]rune, row int, col int) {
	if row == len(g)-1 {
		return
	}

	if g[row+1][col] == '.' {
		g[row+1][col] = 'O'
		g[row][col] = '.'
		castSouth(g, row+1, col)
	}
}

func castEast(g [][]rune, row int, col int) {
	if col == len(g[row])-1 {
		return
	}

	if g[row][col+1] == '.' {
		g[row][col+1] = 'O'
		g[row][col] = '.'
		castEast(g, row, col+1)
	}
}

func castWest(g [][]rune, row int, col int) {
	if col == 0 {
		return
	}

	if g[row][col-1] == '.' {
		g[row][col-1] = 'O'
		g[row][col] = '.'
		castWest(g, row, col-1)
	}
}

func spin(g [][]rune) {
	for row := 0; row < len(g); row++ {
		for col := 0; col < len(g[row]); col++ {
			if g[row][col] == 'O' {
				castNorth(g, row, col)
			}
		}
	}

	//PrintGrid(g)

	for row := 0; row < len(g); row++ {
		for col := 0; col < len(g[row]); col++ {
			if g[row][col] == 'O' {
				castWest(g, row, col)
			}
		}
	}

	//PrintGrid(g)

	for row := len(g) - 1; row >= 0; row-- {
		for col := 0; col < len(g[row]); col++ {
			if g[row][col] == 'O' {
				castSouth(g, row, col)
			}
		}
	}

	//PrintGrid(g)

	for row := 0; row < len(g); row++ {
		for col := len(g[row]) - 1; col >= 0; col-- {
			if g[row][col] == 'O' {
				castEast(g, row, col)
			}
		}
	}

	//PrintGrid(g)
}
func GetPart1Answer(g [][]rune) int {

	answer := 0

	//PrintGrid(g)

	for row := 0; row < len(g); row++ {
		for col := 0; col < len(g[row]); col++ {
			if g[row][col] == 'O' {
				castNorth(g, row, col)
			}
		}
	}

	//PrintGrid(g)

	factor := len(g)
	for row := 0; row < len(g); row++ {
		for col := 0; col < len(g[row]); col++ {
			if g[row][col] == 'O' {
				answer += factor - row
			}
		}
	}

	return answer
}

func GetPart2Answer(g [][]rune) int {

	mapStates := make(map[string]int)
	answer := 0

	PrintGrid(g)

	for i := 1; i <= Cycles; i++ {
		spin(g)
		key := fmt.Sprintf("%v", g)
		if oldCycle, ok := mapStates[key]; ok {
			fmt.Println("Found cycle ", oldCycle, " at: ", i, "...we're starting to repeat")
			if (Cycles-oldCycle)%(i-oldCycle) == 0 {
				fmt.Println("We're at the same cycle as before, so we can just return the answer")
				goto done
			}
		} else {
			mapStates[key] = i
		}
	}
done:
	factor := len(g)
	for row := 0; row < len(g); row++ {
		for col := 0; col < len(g[row]); col++ {
			if g[row][col] == 'O' {
				answer += factor - row
			}
		}
	}

	PrintGrid(g)

	return answer
}

func parseFile(fileName string) [][][]rune {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	grids := [][][]rune{}
	currentGrid := [][]rune{}
	line := ""
	for scanner.Scan() {
		//fmt.Println(line)
		line = scanner.Text()
		if len(line) == 0 {
			grids = append(grids, currentGrid)
			currentGrid = [][]rune{}
			continue
		}
		currentGrid = append(currentGrid, []rune(line))
	}
	grids = append(grids, currentGrid)

	return grids
}

func part1(grids [][][]rune) {

	answer := 0
	for _, grid := range grids {
		answer += GetPart1Answer(grid)
	}

	fmt.Println("Part 1: ", answer)
}

func part2(grids [][][]rune) {

	answer := 0
	for _, grid := range grids {
		answer += GetPart2Answer(grid)
	}

	fmt.Println("Part 2: ", answer)
}
