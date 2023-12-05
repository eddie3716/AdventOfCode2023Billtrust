package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Seed struct {
	Number int
	Count  int
}

type SourceToDestination struct {
	Source      int
	Destination int
	Range       int
}

func (std SourceToDestination) GetDestination(source int) (int, bool) {
	if source >= std.Source && source <= std.Source+std.Range-1 {
		offset := source - std.Source
		return std.Destination + offset, true
	}

	return source, false
}

func GetDestination(source int, records []SourceToDestination) int {
	for record := range records {
		destination, found := records[record].GetDestination(source)
		if found {
			return destination
		}
	}

	return source
}

func main() {

	seed_to_soil := "seed-to-soil map:"
	soil_to_fertilizer := "soil-to-fertilizer map:"
	fertilizer_to_water := "fertilizer-to-water map:"
	water_to_light := "water-to-light map:"
	light_to_temperature := "light-to-temperature map:"
	temperature_to_humidity := "temperature-to-humidity map:"
	humidity_to_location := "humidity-to-location map:"

	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	seeds := []int{}
	scanner.Scan()
	line := scanner.Text()
	for _, token := range strings.Split(line, " ") {
		seed, err := strconv.Atoi(token)
		if err == nil {
			seeds = append(seeds, seed)
		}
	}
	fmt.Println(seeds)
	scanner.Scan()

	analysis := make(map[string][]SourceToDestination)
	key := ""
	records := []SourceToDestination{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			//fmt.Println("Key Len was Zero")
			analysis[key] = records
			records = []SourceToDestination{}
			key = ""
		} else if strings.Contains(line, ":") {
			//fmt.Println("Key Len was not Zero")
			key = line
		} else {
			//fmt.Println("Tokenizing")
			tokens := strings.Split(line, " ")
			destination, _ := strconv.Atoi(tokens[0])
			source, _ := strconv.Atoi(tokens[1])
			rangeValue, _ := strconv.Atoi(tokens[2])
			std := SourceToDestination{Source: source, Destination: destination, Range: rangeValue}
			records = append(records, std)
		}
	}
	//fmt.Println(len(key))
	if len(key) > 0 {
		analysis[key] = records
	}

	seedLowestLocations := []int{}

	for _, seed := range seeds {
		fmt.Print("seed ", seed)
		destination := GetDestination(seed, analysis[seed_to_soil])
		fmt.Print("->", destination)
		destination = GetDestination(destination, analysis[soil_to_fertilizer])
		fmt.Print("->", destination)
		destination = GetDestination(destination, analysis[fertilizer_to_water])
		fmt.Print("->", destination)
		destination = GetDestination(destination, analysis[water_to_light])
		fmt.Print("->", destination)
		destination = GetDestination(destination, analysis[light_to_temperature])
		fmt.Print("->", destination)
		destination = GetDestination(destination, analysis[temperature_to_humidity])
		fmt.Print("->", destination)
		destination = GetDestination(destination, analysis[humidity_to_location])
		fmt.Println("->", destination)
		seedLowestLocations = append(seedLowestLocations, destination)
	}
	sort.Ints(seedLowestLocations)
	fmt.Println(seedLowestLocations[0])
	//346433842
	// for row := 0; row < len(records); row++ {
	// 	for col := 0; col < len(records[row]); col++ {
	// 		if unicode.IsDigit(records[row][col]) {
	// 			tokenEnd := col
	// 			for ; tokenEnd < len(records[row]) && unicode.IsDigit(records[row][tokenEnd]); tokenEnd++ {
	// 			}
	// 			digits := string(records[row][col:tokenEnd])
	// 			number, _ := strconv.Atoi(digits)
	// 			//fmt.Println(number)
	// 			for scanRow := max(0, row-1); scanRow <= min(row+1, len(records)-1); scanRow++ {
	// 				for scanCol := max(0, col-1); scanCol <= min(col+len(digits), len(records[row])-1); scanCol++ {
	// 					if records[scanRow][scanCol] != '.' && !unicode.IsDigit(records[scanRow][scanCol]) {
	// 						potentialGears = append(potentialGears, Gear{number, scanRow, scanCol})
	// 						goto foundSymbol
	// 					}
	// 				}
	// 			}
	// 		foundSymbol:
	// 			col = tokenEnd
	// 		}
	// 	}
	// }

	// totalPower := 0
	// for index, gear := range potentialGears {
	// 	for _, otherGear := range potentialGears[index+1:] {
	// 		if gear.Col == otherGear.Col && gear.Row == otherGear.Row {
	// 			totalPower += gear.Number * otherGear.Number
	// 		}
	// 	}
	// }

	// fmt.Println(totalPower)
}
