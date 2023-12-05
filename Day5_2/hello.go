package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SourceToDestination struct {
	Source      int
	Destination int
	Range       int
}

type Range struct {
	Start int
	End   int
}

func (std []SourceToDestination) GetDestinationRange(source Range) []Range []Range {
	ranges := []Range{}
	if source.Source >= std.Source && source.Source <= std.Source+std.Range-1 {
		offset := source.Source - std.Source
		return Range{Start: std.Destination + offset, End: std.Destination + offset + source.Range - 1}
	}

	return Range{Start: source.Source, End: source.Source}
}

func (std SourceToDestination) GetDestination(source *int) (int, bool) {
	if *source >= std.Source && *source <= std.Source+std.Range-1 {
		offset := *source - std.Source
		return std.Destination + offset, true
	}

	return *source, false
}

func GetDestination(source *int, records *[]SourceToDestination) int {
	for _, record := range *records {
		destination, found := record.GetDestination(source)
		if found {
			return destination
		}
	}

	return *source
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

	lowestSeedvalue := 999999999

	seedStart := 0
	seedRange := 0
	seedtosoilmap := analysis[seed_to_soil]
	soiltofertilizermap := analysis[soil_to_fertilizer]
	ferttowatermap := analysis[fertilizer_to_water]
	watertolight := analysis[water_to_light]
	lighttotemp := analysis[light_to_temperature]
	temptohumid := analysis[temperature_to_humidity]
	humidtoloc := analysis[humidity_to_location]

	// for _, record := range humidtoloc {

	// }
	for index, value := range seeds {
		if index%2 == 0 {
			seedStart = value
		} else {
			seedRange = value

			for seed := seedStart; seed < seedStart+seedRange; seed++ {
				fmt.Println("seed ", seed)
				destination := GetDestination(&seed, &seedtosoilmap)
				//fmt.Print("->", destination)
				destination = GetDestination(&destination, &soiltofertilizermap)
				//fmt.Print("->", destination)
				destination = GetDestination(&destination, &ferttowatermap)
				//fmt.Print("->", destination)
				destination = GetDestination(&destination, &watertolight)
				//fmt.Print("->", destination)
				destination = GetDestination(&destination, &lighttotemp)
				//fmt.Print("->", destination)
				destination = GetDestination(&destination, &temptohumid)
				//fmt.Print("->", destination)
				destination = GetDestination(&destination, &humidtoloc)
				//fmt.Println("->", destination)
				if destination < lowestSeedvalue {
					lowestSeedvalue = destination
				}
			}
		}
	}
	fmt.Println(lowestSeedvalue)
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
