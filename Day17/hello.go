package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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

type CityBlock struct {
	iRow                    int
	iCol                    int
	totalHeat               int
	stepsInCurrentDirection int
	direction               byte
}

func (cb *CityBlock) NextCityBlock(mnemonicDirection byte) *CityBlock {
	var direction []int = directions[mnemonicDirection]

	stepsInCurrentDirection := 0
	if mnemonicDirection == cb.direction {
		stepsInCurrentDirection = cb.stepsInCurrentDirection + 1
	} else {
		stepsInCurrentDirection = 1
	}
	return &CityBlock{iRow: cb.iRow + direction[0], iCol: cb.iCol + direction[1], totalHeat: cb.totalHeat + heatMap[cb.iRow+direction[0]][cb.iCol+direction[1]], direction: mnemonicDirection, stepsInCurrentDirection: stepsInCurrentDirection}
}

func (cb *CityBlock) GetNextDirections(heatMap [][]int) []*CityBlock {
	directions := []*CityBlock{}
	if cb.direction == NORTH {
		if cb.stepsInCurrentDirection < 3 && cb.iRow > 0 {
			directions = append(directions, cb.NextCityBlock(NORTH))
		}
		if cb.iCol > 0 {
			directions = append(directions, cb.NextCityBlock(WEST))
		}
		if cb.iCol < len(heatMap[0])-1 {
			directions = append(directions, cb.NextCityBlock(EAST))
		}
	} else if cb.direction == SOUTH {
		if cb.stepsInCurrentDirection < 3 && cb.iRow < len(heatMap)-1 {
			directions = append(directions, cb.NextCityBlock(SOUTH))
		}
		if cb.iCol > 0 {
			directions = append(directions, cb.NextCityBlock(WEST))
		}
		if cb.iCol < len(heatMap[0])-1 {
			directions = append(directions, cb.NextCityBlock(EAST))
		}
	} else if cb.direction == WEST {
		if cb.stepsInCurrentDirection < 3 && cb.iCol > 0 {
			directions = append(directions, cb.NextCityBlock(WEST))
		}
		if cb.iRow > 0 {
			directions = append(directions, cb.NextCityBlock(NORTH))
		}
		if cb.iRow < len(heatMap)-1 {
			directions = append(directions, cb.NextCityBlock(SOUTH))
		}
	} else if cb.direction == EAST {
		if cb.stepsInCurrentDirection < 3 && cb.iCol < len(heatMap[0])-1 {
			directions = append(directions, cb.NextCityBlock(EAST))
		}
		if cb.iRow > 0 {
			directions = append(directions, cb.NextCityBlock(NORTH))
		}
		if cb.iRow < len(heatMap)-1 {
			directions = append(directions, cb.NextCityBlock(SOUTH))
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
	parseFile("input.txt")
	printHeatMap(heatMap)

	part1()

	// part2()
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

func (cb *CityBlock) String() string {
	return fmt.Sprintf("(col:%d,row:%d,dir:%q,theat:%d)", cb.iRow, cb.iCol, cb.direction, cb.totalHeat)
}

func (cb *CityBlock) GetKey() string {
	return fmt.Sprintf("%d,%d,%d,%d", cb.iRow, cb.iCol, cb.direction, cb.stepsInCurrentDirection)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func part1() {

	answer := math.MaxInt
	visited := map[string]int{}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &CityBlock{iRow: 0, iCol: 0, direction: EAST, stepsInCurrentDirection: 0, totalHeat: 0})

	for len(pq) > 0 {
		currentCityBlock := heap.Pop(&pq).(*CityBlock)
		newDirections := currentCityBlock.GetNextDirections(heatMap)

		key := currentCityBlock.GetKey()
		if visitedHeat, didVisit := visited[key]; didVisit && visitedHeat < currentCityBlock.totalHeat {
			continue
		}

		visited[key] = currentCityBlock.totalHeat

		for _, newDirection := range newDirections {
			heap.Push(&pq, newDirection)
		}
	}
	//fmt.Println("Visited: ", visited)
	for key, value := range visited {
		tokens := strings.Split(key, ",")
		row, _ := strconv.Atoi(tokens[0])
		col, _ := strconv.Atoi(tokens[1])
		if row == len(heatMap[0])-1 && col == len(heatMap)-1 {
			answer = Min(answer, value)
		}
	}

	fmt.Println()
	printHeatMap(heatMap)
	fmt.Println("Part 1: ", answer)
}

func part2(heatMap [][]int) {

	answer := 0

	fmt.Println("Part 2: ", answer)
}
