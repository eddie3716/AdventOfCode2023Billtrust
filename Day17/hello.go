package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	sets := parseFile("testinput.txt")

	part1(sets)

	part2(sets)
}

type HeatSteps struct {
	Steps int
	Heat  int
}

const (
	NORTH = iota*10 + 10
	SOUTH
	WEST
	EAST
)

func Min(nums ...HeatSteps) HeatSteps {
	min := HeatSteps{math.MaxInt32, 0}
	for _, num := range nums {
		if num.Heat < min.Heat {
			min = num
		}
	}
	return min
}

func GetKey(startX int, startY int, fromDirection int) string {
	return fmt.Sprintf("%d,%d,%d", startX, startY, fromDirection)
}

func setHeadSteps(visited map[string]HeatSteps, startX int, startY int, heatSteps HeatSteps, fromDirection int) {
	key := GetKey(startX, startY, fromDirection)
	visited[key] = heatSteps
}

func hasVisited(visited map[string]HeatSteps, startX int, startY int, fromDirection int) (bool, HeatSteps) {
	key := GetKey(startX, startY, fromDirection)
	heat, didVisit := visited[key]
	return didVisit, heat
}

func goNorth(goalX int, goalY int, visited map[string]HeatSteps, heatMap [][]rune, startX int, startY int, countInDirection int) HeatSteps {
	if startX < 0 || startX >= len(heatMap[0]) || startY < 0 || startY >= len(heatMap) {
		return HeatSteps{math.MaxInt32, 0}
	}
	if startX == goalX && startY == goalY {
		return HeatSteps{int(heatMap[goalX][goalY] - '0'), 1}
	}

	if visited, heatSteps := hasVisited(visited, startX, startY, SOUTH); visited {
		return heatSteps
	}
	northHeat := HeatSteps{math.MaxInt32, 0}
	if countInDirection < 3 {
		northHeat = goNorth(goalX, goalY, visited, heatMap, startX, startY-1, countInDirection+1)
	}

	westHeat := goWest(goalX, goalY, visited, heatMap, startX-1, startY, 0)
	eastHeat := goEast(goalX, goalY, visited, heatMap, startX+1, startY, 0)

	currentHeat := HeatSteps{int(heatMap[startY][startX] - '0'), 1}
	minHeat := Min(northHeat, westHeat, eastHeat)
	currentHeat.Heat += minHeat.Heat
	currentHeat.Steps += minHeat.Steps

	setHeadSteps(visited, startX, startY, currentHeat, SOUTH)

	return currentHeat
}

func goSouth(goalX int, goalY int, visited map[string]HeatSteps, heatMap [][]rune, startX int, startY int, countInDirection int) HeatSteps {
	if startX < 0 || startX >= len(heatMap[0]) || startY < 0 || startY >= len(heatMap) {
		return HeatSteps{math.MaxInt32, 0}
	}
	if startX == goalX && startY == goalY {
		return HeatSteps{int(heatMap[goalX][goalY] - '0'), 1}
	}

	if visited, heat := hasVisited(visited, startX, startY, NORTH); visited {
		return heat
	}
	southHeat := HeatSteps{math.MaxInt32, 0}
	if countInDirection < 3 {
		southHeat = goSouth(goalX, goalY, visited, heatMap, startX, startY+1, countInDirection+1)
	}

	westHeat := goWest(goalX, goalY, visited, heatMap, startX-1, startY, 0)
	eastHeat := goEast(goalX, goalY, visited, heatMap, startX+1, startY, 0)

	currentHeat := HeatSteps{int(heatMap[startY][startX] - '0'), 1}
	minHeat := Min(southHeat, westHeat, eastHeat)
	currentHeat.Heat += minHeat.Heat
	currentHeat.Steps += minHeat.Steps

	setHeadSteps(visited, startX, startY, currentHeat, NORTH)

	return currentHeat
}

func goWest(goalX int, goalY int, visited map[string]HeatSteps, heatMap [][]rune, startX int, startY int, countInDirection int) HeatSteps {
	if startX < 0 || startX >= len(heatMap[0]) || startY < 0 || startY >= len(heatMap) {
		return HeatSteps{math.MaxInt32, 0}
	}
	if startX == goalX && startY == goalY {
		return HeatSteps{int(heatMap[goalX][goalY] - '0'), 1}
	}

	if visited, heat := hasVisited(visited, startX, startY, EAST); visited {
		return heat
	}
	westHeat := HeatSteps{math.MaxInt32, 0}
	if countInDirection < 3 {
		westHeat = goWest(goalX, goalY, visited, heatMap, startX-1, startY, countInDirection+1)
	}

	northHeat := goNorth(goalX, goalY, visited, heatMap, startX, startY-1, 0)
	southHeat := goSouth(goalX, goalY, visited, heatMap, startX, startY+1, 0)

	currentHeat := HeatSteps{int(heatMap[startY][startX] - '0'), 1}
	minHeat := Min(westHeat, Min(northHeat, southHeat))
	currentHeat.Heat += minHeat.Heat
	currentHeat.Steps += minHeat.Steps

	setHeadSteps(visited, startX, startY, currentHeat, EAST)

	return currentHeat
}

func goEast(goalX int, goalY int, visited map[string]HeatSteps, heatMap [][]rune, startX int, startY int, countInDirection int) HeatSteps {
	if startX < 0 || startX >= len(heatMap[0]) || startY < 0 || startY >= len(heatMap) {
		return HeatSteps{math.MaxInt32, 0}
	}
	if startX == goalX && startY == goalY {
		return HeatSteps{int(heatMap[goalX][goalY] - '0'), 1}
	}

	if visited, heat := hasVisited(visited, startX, startY, WEST); visited {
		return heat
	}

	eastHeat := HeatSteps{math.MaxInt32, 0}
	if countInDirection < 3 {
		eastHeat = goEast(goalX, goalY, visited, heatMap, startX+1, startY, countInDirection+1)
	}

	northHeat := goNorth(goalX, goalY, visited, heatMap, startX, startY-1, 0)
	southHeat := goSouth(goalX, goalY, visited, heatMap, startX, startY+1, 0)

	currentHeat := HeatSteps{int(heatMap[startY][startX] - '0'), 1}

	minHeat := Min(eastHeat, northHeat, southHeat)
	currentHeat.Heat += minHeat.Heat
	currentHeat.Steps += minHeat.Steps

	setHeadSteps(visited, startX, startY, currentHeat, WEST)

	return currentHeat
}

func parseFile(fileName string) [][]rune {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	heatMap := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		heatMap = append(heatMap, []rune(line))
	}

	return heatMap
}

func printHeatMap(heatMap [][]rune) {
	for _, line := range heatMap {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func part1(heatMap [][]rune) {

	answer := goEast(len(heatMap[0])-1, len(heatMap)-1, map[string]HeatSteps{}, heatMap, 0, 0, 0)

	fmt.Println("Part 1: ", answer)
}

func part2(heatMap [][]rune) {

	answer := 0

	fmt.Println("Part 2: ", answer)
}
