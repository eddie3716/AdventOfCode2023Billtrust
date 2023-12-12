package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

const (
	InflationFactor = 1000000
	//InflationFactor = 100
	//InflationFactor = 999999
)

type Inflation []int

type Row []rune

func (r Row) InsertAtIndex(index int, value rune) Row {
	if len(r) == index { // nil or empty slice or after last element
		return append(r, value)
	}
	r = append(r[:index+1], r[index:]...) // index < len(a)
	r[index] = value
	return r
}

type Grid []Row

func (g Grid) PrintGrid() {
	for _, row := range g {
		for _, symbol := range row {
			print(string(symbol))
		}
		println()
	}
}

type Galaxy struct {
	Symbol rune
	Row    int
	Col    int
}

type GalaxyPair struct {
	Galaxy1  Galaxy
	Galaxy2  Galaxy
	Distance int
}

func (g Galaxy) DistanceTo(otherGalaxy Galaxy) int {
	r := math.Abs(float64(g.Row - otherGalaxy.Row))
	c := math.Abs(float64(g.Col - otherGalaxy.Col))
	d := int(r + c)
	return d
}

func (g Galaxy) DistanceToWithInflation(otherGalaxy Galaxy, rowInflation Inflation, colInflation Inflation) int {

	rowInflationFactor := 0
	for _, inflation := range rowInflation {
		if Min(otherGalaxy.Row, g.Row) < inflation && inflation < Max(otherGalaxy.Row, g.Row) {
			rowInflationFactor++
		}
	}

	colInflationFactor := 0
	for _, inflation := range colInflation {
		if Min(otherGalaxy.Col, g.Col) < inflation && inflation < Max(otherGalaxy.Col, g.Col) {
			colInflationFactor++
		}
	}

	r := Abs(g.Row-otherGalaxy.Row) + Max(1, InflationFactor-1)*rowInflationFactor
	c := Abs(g.Col-otherGalaxy.Col) + Max(1, InflationFactor-1)*colInflationFactor
	d := r + c

	return d
}

func parseFile(fileName string) Grid {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	records := Grid{}

	for scanner.Scan() {
		line := scanner.Text()
		records = append(records, Row(line))
	}

	return records
}

func inflate(grid Grid) Grid {
	inflatedRows := Grid{}

	for _, row := range grid {
		inflatedRow := Row{}

		hasAllDots := true
		for _, symbol := range row {
			inflatedRow = append(inflatedRow, symbol)
			if hasAllDots && symbol != '.' {
				hasAllDots = false
			}
		}
		inflatedRows = append(inflatedRows, inflatedRow)
		if hasAllDots {
			newRow := Row{}
			for i := 0; i < len(grid[0]); i++ {
				newRow = append(newRow, '.')
			}
			inflatedRows = append(inflatedRows, newRow)
		}
	}

	colsWithAllDots := []int{}
	for iCol := 0; iCol < len(grid[0]); iCol++ {
		for _, row := range grid {
			if row[iCol] != '.' {
				goto nextCol
			}
		}
		colsWithAllDots = append(colsWithAllDots, iCol)
	nextCol:
	}

	for iRow := 0; iRow < len(inflatedRows); iRow++ {
		row := inflatedRows[iRow]
		for iCol := len(colsWithAllDots) - 1; iCol >= 0; iCol-- {
			col := colsWithAllDots[iCol]
			row = row.InsertAtIndex(col, '.')
		}
		inflatedRows[iRow] = row
	}

	return inflatedRows
}

func getInflation(grid Grid) (Inflation, Inflation) {
	rowInflation := Inflation{}
	colInflation := Inflation{}

	for iRow, row := range grid {
		inflatedRow := Row{}

		hasAllDots := true
		for _, symbol := range row {
			inflatedRow = append(inflatedRow, symbol)
			if hasAllDots && symbol != '.' {
				hasAllDots = false
			}
		}
		if hasAllDots {
			rowInflation = append(rowInflation, iRow)
		}
	}

	for iCol := 0; iCol < len(grid[0]); iCol++ {
		for _, row := range grid {
			if row[iCol] != '.' {
				goto nextCol
			}
		}
		colInflation = append(colInflation, iCol)
	nextCol:
	}
	return rowInflation, colInflation
}

func part1(grid Grid) {
	grid = inflate(grid)

	galaxyNumber := '1'
	galaxies := []Galaxy{}
	for rIndex, row := range grid {
		for cIndex, symbol := range row {
			if symbol == '#' {
				grid[rIndex][cIndex] = galaxyNumber
				galaxies = append(galaxies, Galaxy{Symbol: galaxyNumber, Row: rIndex, Col: cIndex})
				galaxyNumber++
			}
		}
	}

	galaxyPairs := []GalaxyPair{}
	galaxyPartnerIndex := 1
	for _, galaxy := range galaxies {
		for _, otherGalaxy := range galaxies[galaxyPartnerIndex:] {
			distance := galaxy.DistanceTo(otherGalaxy)
			galaxyPairs = append(galaxyPairs, GalaxyPair{Galaxy1: galaxy, Galaxy2: otherGalaxy, Distance: distance})
		}
		galaxyPartnerIndex++
	}

	totalDistance := 0
	for _, galaxyPair := range galaxyPairs {
		totalDistance += galaxyPair.Distance
	}

	fmt.Println(totalDistance)

}
func part2(grid Grid) {
	rowInflation, colInflation := getInflation(grid)

	galaxyNumber := '1'
	galaxies := []Galaxy{}
	for rIndex, row := range grid {
		for cIndex, symbol := range row {
			if symbol == '#' {
				grid[rIndex][cIndex] = galaxyNumber
				galaxies = append(galaxies, Galaxy{Symbol: galaxyNumber, Row: rIndex, Col: cIndex})
				galaxyNumber++
			}
		}
	}

	galaxyPairs := []GalaxyPair{}
	galaxyPartnerIndex := 1
	for _, galaxy := range galaxies {
		for _, otherGalaxy := range galaxies[galaxyPartnerIndex:] {
			distance := galaxy.DistanceToWithInflation(otherGalaxy, rowInflation, colInflation)
			galaxyPairs = append(galaxyPairs, GalaxyPair{Galaxy1: galaxy, Galaxy2: otherGalaxy, Distance: distance})
		}
		galaxyPartnerIndex++
	}

	totalDistance := 0
	for _, galaxyPair := range galaxyPairs {
		totalDistance += galaxyPair.Distance
	}

	fmt.Println(totalDistance)

}

func main() {
	grid := parseFile("input.txt")

	part1(grid)

	part2(grid)
}
