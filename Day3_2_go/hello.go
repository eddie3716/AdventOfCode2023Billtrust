package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Gear struct {
	Number int
	Row    int
	Col    int
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	records := [][]rune{}

	potentialGears := []Gear{}

	for scanner.Scan() {
		line := scanner.Text()
		records = append(records, []rune(line))
	}

	for row := 0; row < len(records); row++ {
		for col := 0; col < len(records[row]); col++ {
			if unicode.IsDigit(records[row][col]) {
				tokenEnd := col
				for ; tokenEnd < len(records[row]) && unicode.IsDigit(records[row][tokenEnd]); tokenEnd++ {
				}
				digits := string(records[row][col:tokenEnd])
				number, _ := strconv.Atoi(digits)
				//fmt.Println(number)
				for scanRow := max(0, row-1); scanRow <= min(row+1, len(records)-1); scanRow++ {
					for scanCol := max(0, col-1); scanCol <= min(col+len(digits), len(records[row])-1); scanCol++ {
						if records[scanRow][scanCol] != '.' && !unicode.IsDigit(records[scanRow][scanCol]) {
							potentialGears = append(potentialGears, Gear{number, scanRow, scanCol})
							goto foundSymbol
						}
					}
				}
			foundSymbol:
				col = tokenEnd
			}
		}
	}

	totalPower := 0
	for index, gear := range potentialGears {
		for _, otherGear := range potentialGears[index+1:] {
			if gear.Col == otherGear.Col && gear.Row == otherGear.Row {
				totalPower += gear.Number * otherGear.Number
			}
		}
	}

	fmt.Println(totalPower)
}
