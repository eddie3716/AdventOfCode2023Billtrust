package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const (
	NORTH = iota*10 + 10
	SOUTH
	WEST
	EAST
)

var toNorth = []int{-1, 0}
var toSouth = []int{1, 0}
var toWest = []int{0, -1}
var toEast = []int{0, 1}

type CityBlock struct {
	iRow                    int
	iCol                    int
	totalHeat               int
	totalSteps              int
	direction               int
	stepsInCurrentDirection int
	previousCityBlock       *CityBlock
}

func (cb *CityBlock) NextCityBlock(mnemonicDirection int, direction []int, heatMap [][]rune) *CityBlock {
	stepsInCurrentDirection := 0
	if mnemonicDirection == cb.direction {
		stepsInCurrentDirection = cb.stepsInCurrentDirection + 1
	} else {
		stepsInCurrentDirection = 1
	}
	return &CityBlock{iRow: cb.iRow + direction[0], iCol: cb.iCol + direction[1], totalHeat: cb.totalHeat + int(heatMap[cb.iRow+direction[0]][cb.iCol+direction[1]]-'0'), direction: mnemonicDirection, totalSteps: 1 + cb.totalSteps, stepsInCurrentDirection: stepsInCurrentDirection}
}

func (cb *CityBlock) GetNextDirections(heatMap [][]rune) []*CityBlock {
	directions := []*CityBlock{}
	if cb.direction == NORTH {
		if cb.stepsInCurrentDirection < 3 && cb.iRow > 0 {
			directions = append(directions, cb.NextCityBlock(NORTH, toNorth, heatMap))
		}
		if cb.iCol > 0 {
			directions = append(directions, cb.NextCityBlock(WEST, toWest, heatMap))
		}
		if cb.iCol < len(heatMap[0])-1 {
			directions = append(directions, cb.NextCityBlock(EAST, toEast, heatMap))
		}
	} else if cb.direction == SOUTH {
		if cb.stepsInCurrentDirection < 3 && cb.iRow < len(heatMap)-1 {
			directions = append(directions, cb.NextCityBlock(SOUTH, toSouth, heatMap))
		}
		if cb.iCol > 0 {
			directions = append(directions, cb.NextCityBlock(WEST, toWest, heatMap))
		}
		if cb.iCol < len(heatMap[0])-1 {
			directions = append(directions, cb.NextCityBlock(EAST, toEast, heatMap))
		}
	} else if cb.direction == WEST {
		if cb.stepsInCurrentDirection < 3 && cb.iCol > 0 {
			directions = append(directions, cb.NextCityBlock(WEST, toWest, heatMap))
		}
		if cb.iRow > 0 {
			directions = append(directions, cb.NextCityBlock(NORTH, toNorth, heatMap))
		}
		if cb.iRow < len(heatMap)-1 {
			directions = append(directions, cb.NextCityBlock(SOUTH, toSouth, heatMap))
		}
	} else if cb.direction == EAST {
		if cb.stepsInCurrentDirection < 3 && cb.iCol < len(heatMap[0])-1 {
			directions = append(directions, cb.NextCityBlock(EAST, toEast, heatMap))
		}
		if cb.iRow > 0 {
			directions = append(directions, cb.NextCityBlock(NORTH, toNorth, heatMap))
		}
		if cb.iRow < len(heatMap)-1 {
			directions = append(directions, cb.NextCityBlock(SOUTH, toSouth, heatMap))
		}
	}
	return directions
}

type PriorityQueue []*CityBlock

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].totalHeat < pq[j].totalHeat
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

func main() {
	heatMap := parseFile("testinput.txt")
	//fmt.Println(heatMap)

	part1(heatMap)

	// part2(heatMap)
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

func (cb *CityBlock) GetKey() string {
	return fmt.Sprintf("%d,%d,%d", cb.iRow, cb.iCol, cb.direction)
}

func part1(heatMap [][]rune) {

	answer := 0
	visited := map[string]struct{}{}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &CityBlock{iRow: 0, iCol: 0, direction: EAST, stepsInCurrentDirection: 0, totalHeat: int(heatMap[0][0]) - '0'})

	for len(pq) > 0 {
		currentCityBlock := heap.Pop(&pq).(*CityBlock)
		newDirections := currentCityBlock.GetNextDirections(heatMap)

		if currentCityBlock.iRow == len(heatMap)-1 && currentCityBlock.iCol == len(heatMap[0])-1 {
			answer = currentCityBlock.totalHeat
			goto done
		}

		key := currentCityBlock.GetKey()
		if _, didVisit := visited[key]; didVisit {
			continue
		}
		visited[key] = struct{}{}

		for _, newDirection := range newDirections {
			heap.Push(&pq, newDirection)
		}
	}
done:
	fmt.Println("Part 1: ", answer)
}

func part2(heatMap [][]rune) {

	answer := 0

	fmt.Println("Part 2: ", answer)
}
