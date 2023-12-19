package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type validDirection func(int) bool

var endRow int
var endCol int
var heatMap [][]int = [][]int{}
var directions = map[byte][]int{NORTH: toNorth, SOUTH: toSouth, WEST: toWest, EAST: toEast}

const (
	NORTH = byte('N')
	SOUTH = byte('S')
	WEST  = byte('W')
	EAST  = byte('E')
)

var toNorth = []int{-1, 0}
var toSouth = []int{1, 0}
var toWest = []int{0, -1}
var toEast = []int{0, 1}

type CityBlockState struct {
	iRow                    int
	iCol                    int
	stepsInCurrentDirection int
	direction               byte
}

type CityBlock struct {
	totalHeat int
	state     CityBlockState
}

func (cb *CityBlock) NextCityBlock(mnemonicDirection byte) *CityBlock {
	var direction []int = directions[mnemonicDirection]

	stepsInCurrentDirection := 1
	if mnemonicDirection == cb.state.direction {
		stepsInCurrentDirection += cb.state.stepsInCurrentDirection
	}

	return &CityBlock{state: CityBlockState{iRow: cb.state.iRow + direction[0], iCol: cb.state.iCol + direction[1], direction: mnemonicDirection, stepsInCurrentDirection: stepsInCurrentDirection}, totalHeat: cb.totalHeat + heatMap[cb.state.iRow+direction[0]][cb.state.iCol+direction[1]]}
}

func (cb *CityBlock) GetNextDirections(heatMap [][]int, min int, max int) *[]*CityBlock {
	directions := []*CityBlock{}
	if cb.state.direction == NORTH {
		if cb.state.stepsInCurrentDirection < max && cb.state.iRow > 0 {
			directions = append(directions, cb.NextCityBlock(NORTH))
		}
		if cb.state.stepsInCurrentDirection >= min && cb.state.iCol > 0 {
			directions = append(directions, cb.NextCityBlock(WEST))
		}
		if cb.state.stepsInCurrentDirection >= min && cb.state.iCol < len(heatMap[0])-1 {
			directions = append(directions, cb.NextCityBlock(EAST))
		}
	} else if cb.state.direction == SOUTH {
		if cb.state.stepsInCurrentDirection < max && cb.state.iRow < endRow {
			directions = append(directions, cb.NextCityBlock(SOUTH))
		}
		if cb.state.stepsInCurrentDirection >= min && cb.state.iCol > 0 {
			directions = append(directions, cb.NextCityBlock(WEST))
		}
		if cb.state.stepsInCurrentDirection >= min && cb.state.iCol < len(heatMap[0])-1 {
			directions = append(directions, cb.NextCityBlock(EAST))
		}
	} else if cb.state.direction == WEST {
		if cb.state.stepsInCurrentDirection < max && cb.state.iCol > 0 {
			directions = append(directions, cb.NextCityBlock(WEST))
		}
		if cb.state.stepsInCurrentDirection >= min && cb.state.iRow > 0 {
			directions = append(directions, cb.NextCityBlock(NORTH))
		}
		if cb.state.stepsInCurrentDirection >= min && cb.state.iRow < len(heatMap)-1 {
			directions = append(directions, cb.NextCityBlock(SOUTH))
		}
	} else if cb.state.direction == EAST {
		if cb.state.stepsInCurrentDirection < max && cb.state.iCol < endCol {
			directions = append(directions, cb.NextCityBlock(EAST))
		}
		if cb.state.stepsInCurrentDirection >= min && cb.state.iRow > 0 {
			directions = append(directions, cb.NextCityBlock(NORTH))
		}
		if cb.state.stepsInCurrentDirection >= min && cb.state.iRow < len(heatMap)-1 {
			directions = append(directions, cb.NextCityBlock(SOUTH))
		}
	}
	return &directions
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
	parseFile("input.txt")
	//printHeatMap(heatMap)

	fmt.Println("answer part 1:", findMinHeat(0, 3))

	fmt.Println("answer part 2:", findMinHeat(4, 10))
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
		record := []int{}
		for _, char := range line {
			record = append(record, int(char)-'0')
		}
		heatMap = append(heatMap, record)
	}

	endRow = len(heatMap) - 1
	endCol = len(heatMap[0]) - 1
}

func printHeatMap(heatMap [][]int) {
	for _, line := range heatMap {
		for _, char := range line {
			fmt.Print(char)
		}
		fmt.Println()
	}
	fmt.Println()
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findMinHeat(min int, max int) int {

	answer := math.MaxInt
	visited := map[CityBlockState]int{}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &CityBlock{state: CityBlockState{iRow: 0, iCol: 0, direction: EAST, stepsInCurrentDirection: 0}, totalHeat: 0})

	for len(pq) > 0 {
		currentCityBlock := heap.Pop(&pq).(*CityBlock)

		if _, didVisit := visited[currentCityBlock.state]; didVisit {
			continue
		}

		visited[currentCityBlock.state] = currentCityBlock.totalHeat

		newDirections := currentCityBlock.GetNextDirections(heatMap, min, max)

		for _, newDirection := range *newDirections {
			//fmt.Println(newDirection.state, newDirection.totalHeat)
			heap.Push(&pq, newDirection)
		}
	}
	//fmt.Println("Visited: ", visited)
	for state, value := range visited {
		//fmt.Println(state, value)
		if state.iRow == len(heatMap)-1 && state.iCol == len(heatMap[0])-1 {
			answer = Min(answer, value)
		}
	}

	//fmt.Println()
	//printHeatMap(heatMap)
	return answer
}
