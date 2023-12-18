package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type CityBlock struct {
	Name                    string
	iRow                    int
	iCol                    int
	Heat                    int
	Direction               int
	stepsInCurrentDirection int
}

type PriorityQueue []*CityBlock

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Heat < pq[j].Heat
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*CityBlock))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func calculateDistances(graph *orderedmap.OrderedMap[string, orderedmap.OrderedMap[string, int]], startingCityBlock string) *orderedmap.OrderedMap[string, int] {
	cityBlockHeat := orderedmap.New[string, int]()

	for pair := graph.Oldest(); pair != nil; pair = pair.Next() {
		cityBlockHeat.Set(pair.Key, math.MaxInt)
	}

	cityBlockHeat.Set(startingCityBlock, 0)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &CityBlock{Name: startingCityBlock, Heat: 0})
	previousCityBlock := &CityBlock{Name: startingCityBlock, Heat: 0}
	// currentDirection := 0
	// countOfTimesInDirection := 0
	for len(pq) > 0 {
		currentCityBlock := heap.Pop(&pq).(*CityBlock)

		if distance, ok := cityBlockHeat.Get(currentCityBlock.Name); !ok || distance < currentCityBlock.Heat {
			continue
		}

		if neighbors, ok := graph.Get(currentCityBlock.Name); ok {
			for neighbor := neighbors.Oldest(); neighbor != nil; neighbor = neighbor.Next() {
				if neighbor.Key == previousCityBlock.Name {
					continue
				}
				heat := currentCityBlock.Heat + neighbor.Value

				if currentHeat, currentHeatOk := cityBlockHeat.Get(neighbor.Key); currentHeatOk && heat < currentHeat {
					cityBlockHeat.Set(neighbor.Key, heat)
					heap.Push(&pq, &CityBlock{Name: neighbor.Key, Heat: heat})
				} else if !currentHeatOk {
					panic("couldn't find distances from " + currentCityBlock.Name + " for neighbor " + neighbor.Key)
				}
			}
		} else {
			panic("couldn't find neighbors for: " + currentCityBlock.Name)
		}

		previousCityBlock = currentCityBlock
	}

	return cityBlockHeat
}

func heatMapToGraph(heatMap [][]rune) *orderedmap.OrderedMap[string, orderedmap.OrderedMap[string, int]] {
	graph := orderedmap.New[string, orderedmap.OrderedMap[string, int]]()

	for iRow := 0; iRow < len(heatMap); iRow++ {
		for iCol := 0; iCol < len(heatMap[0]); iCol++ {
			key := GetKey(iRow, iCol, 0)
			if _, ok := graph.Get(key); !ok {
				graph.Set(key, *orderedmap.New[string, int]())
			}
			if subGraph, ok := graph.Get(key); ok {
				if iCol < len(heatMap[0])-1 {
					otherKey := GetKey(iRow, iCol+1, 0)
					subGraph.Set(otherKey, int(heatMap[iRow][iCol+1]-'0')) // east node
				}
				if iRow < len(heatMap)-1 {
					otherKey := GetKey(iRow+1, iCol, 0)
					subGraph.Set(otherKey, int(heatMap[iRow+1][iCol]-'0')) // south node
				}
				if iRow > 0 {
					otherKey := GetKey(iRow-1, iCol, 0)
					subGraph.Set(otherKey, int(heatMap[iRow-1][iCol]-'0')) // west node
				}
				if iCol > 0 {
					otherKey := GetKey(iRow, iCol-1, 0)
					subGraph.Set(otherKey, int(heatMap[iRow][iCol-1]-'0')) // north node
				}
			} else {
				panic("couldn't find key: " + key)
			}

		}
	}

	return graph
}

func main() {
	sets := parseFile("testinput.txt")

	graph := heatMapToGraph(sets)
	//fmt.Println(graph)
	result := calculateDistances(graph, GetKey(0, 0, 0))

	for pair := result.Oldest(); pair != nil; pair = pair.Next() {
		fmt.Println("Neighbor: ", pair.Key, " at heat ", pair.Value)
	}

	// part1(sets)

	// part2(sets)
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
	return fmt.Sprintf("%d,%d", startX, startY)
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
