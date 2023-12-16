package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sets := parseFile("input.txt")

	part1(sets)

	part2(sets)
}

func bound(lowerIndex int, higherIndex int, recordLength int) int {
	//fmt.Println("lowerIndex: ", lowerIndex, " higherIndex: ", higherIndex, " recordLength: ", recordLength)
	if lowerIndex > (recordLength - higherIndex - 1) {
		//fmt.Println("bound: ", recordLength-higherIndex-1)
		return recordLength - higherIndex - 1
	}
	//fmt.Println("bound: ", lowerIndex)
	return lowerIndex
}

func GetPart1Answer(g [][]rune) int {

	answer := 0
	verticalCol := 0

	//check all the top row columns to see if they are the same..keep doing this until we find a match
	for cI := 1; cI < len(g[0]) && verticalCol == 0; cI++ {
		//found a top column set of tokens that are the same
		if g[0][cI-1] == g[0][cI] {
			//this could be the one we're looking for
			verticalCol = cI

			//figure out how far we can go to the left and right
			boundary := bound(cI-1, cI, len(g[0]))
			//now check all the rows to see if they reflect the same
			for rI := 0; rI < len(g); rI++ {
				for b := 0; b <= boundary; b++ {
					if g[rI][cI-1-b] != g[rI][cI+b] {
						//this row's columns do not reflect, so this entire column is not a reflecting
						verticalCol = 0
						goto tryNextColumn
					}
				}
			}
		}
	tryNextColumn:
	}

	answer += verticalCol

	horiontalRow := 0

	//check all the top row columns to see if they are the same..keep doing this until we find a match
	for rI := 1; rI < len(g) && horiontalRow == 0; rI++ {
		//found a top row set of tokens that are the same
		if g[rI-1][0] == g[rI][0] {
			//this could be the one we're looking for
			horiontalRow = rI
			//figure out how far we can go to the up and down
			boundary := bound(rI-1, rI, len(g))
			for cI := 0; cI < len(g[rI]); cI++ {
				for b := 0; b <= boundary; b++ {
					if g[rI-1-b][cI] != g[rI+b][cI] {
						//this column's rows do not reflect, so this entire colums is not a reflecting
						horiontalRow = 0
						goto tryNextRow
					}
				}
			}
		}
	tryNextRow:
	}

	answer += horiontalRow * 100

	return answer
}

func GetPart2Answer(g [][]rune) int {

	answer := 0

	//check all the top row columns to see if they are the same..keep doing this until we find a match
	for cI := 1; cI < len(g[0]); cI++ {
		//found a top column set of tokens that are the same
		//figure out how far we can go to the left and right
		boundary := bound(cI-1, cI, len(g[0]))
		//now check all the rows to see if they reflect the same
		offBy := 0
		for rI := 0; rI < len(g); rI++ {
			for b := 0; b <= boundary; b++ {
				if g[rI][cI-1-b] != g[rI][cI+b] {
					offBy++
				}
			}
		}
		if offBy == 1 {
			//this row do not reflect without this smudge adjustment
			answer += cI
			goto doHorizontal
		}
	}
doHorizontal:

	//check all the top row columns to see if they are the same..keep doing this until we find a match
	for rI := 1; rI < len(g); rI++ {
		offBy := 0
		boundary := bound(rI-1, rI, len(g))
		for cI := 0; cI < len(g[rI]); cI++ {
			for b := 0; b <= boundary; b++ {
				if g[rI-1-b][cI] != g[rI+b][cI] {
					offBy++
				}
			}
		}
		if offBy == 1 {
			//this column do not reflect without this smudge adjustment
			answer += rI * 100
			goto done
		}
	}
done:
	//fmt.Println("verticalCol: ", verticalCol)
	//fmt.Println("horiontalRow: ", horiontalRow)

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
