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

func (std SourceToDestination) GetRanges(tests []Range, elegantSolution bool) ([]Range, []Range) {
	destinationRanges := []Range{}
	noFoundSourceRanges := []Range{}
	sourceStart := std.Source
	sourceEnd := std.Source + std.Range - 1
	if elegantSolution {
		for _, test := range tests {
			testStart := test.Start
			testEnd := test.End
			if testStart >= sourceStart && testEnd <= sourceEnd {
				//fits neatly within range
				offset := testStart - sourceStart
				destinationRange := Range{std.Destination + offset, std.Destination + offset + testEnd - testStart}
				destinationRanges = append(destinationRanges, destinationRange)
				if destinationRange.Start > destinationRange.End {
					fmt.Println("destination wrong: ", test, destinationRange, std)
					panic("testStart >= sourceStart && testEnd <= sourceEnd")
				}
			} else if testEnd < sourceStart || testStart > sourceEnd {
				//does not fit at all within range
				noFoundSourceRanges = append(noFoundSourceRanges, test)
				if test.Start > test.End {
					fmt.Println("test wrong: ", test, std)
					panic("testEnd < sourceStart || testStart > sourceEnd")
				}
			} else if testStart >= sourceStart && testEnd > sourceEnd {
				//start fits in range, but end does not
				noFoundSourceRange := Range{sourceEnd + 1, testEnd}
				noFoundSourceRanges = append(noFoundSourceRanges, noFoundSourceRange)
				offset := testStart - sourceStart
				destinationRange := Range{std.Destination + offset, std.Destination + std.Range - 1}
				destinationRanges = append(destinationRanges, destinationRange)
				if destinationRange.Start > destinationRange.End {
					fmt.Println("destination wrong: ", test, destinationRange, std)
					panic("testStart >= sourceStart && testEnd > sourceEnd")
				}
				if noFoundSourceRange.Start > noFoundSourceRange.End {
					fmt.Println("noFoundSourceRange wrong: ", test, noFoundSourceRange, std)
					panic("testStart >= sourceStart && testEnd > sourceEnd")
				}
			} else if testStart < sourceStart && testEnd <= sourceEnd {
				//end fits in range, but start does not
				noFoundSourceRange := Range{testStart, sourceStart - 1}
				noFoundSourceRanges = append(noFoundSourceRanges, noFoundSourceRange)
				offset := testEnd - sourceStart
				destinationRange := Range{std.Destination, std.Destination + offset}
				destinationRanges = append(destinationRanges, destinationRange)
				if destinationRange.Start > destinationRange.End {
					fmt.Println("destination wrong: ", test, destinationRange, std)
					panic("testStart < sourceStart && testEnd <= sourceEnd")
				}
				if noFoundSourceRange.Start > noFoundSourceRange.End {
					fmt.Println("noFoundSourceRange wrong: ", test, noFoundSourceRange, std)
					panic("testStart < sourceStart && testEnd <= sourceEnd")
				}
			} else if testStart < sourceStart && testEnd > sourceEnd {
				//super wide....source extends past the start and end, but does overlap
				noFoundSourceRangeAtStart := Range{testStart, sourceStart - 1}
				noFoundSourceRanges = append(noFoundSourceRanges, noFoundSourceRangeAtStart)
				noFoundSourceRangeAtEnd := Range{sourceEnd + 1, testEnd}
				noFoundSourceRanges = append(noFoundSourceRanges, noFoundSourceRangeAtEnd)
				destinationRange := Range{std.Destination, std.Destination + std.Range - 1}
				destinationRanges = append(destinationRanges, destinationRange)
				if destinationRange.Start > destinationRange.End {
					fmt.Println("destination wrong: ", test, destinationRange, std)
					panic("testStart < sourceStart && testEnd <= sourceEnd")
				}
				if noFoundSourceRangeAtStart.Start > noFoundSourceRangeAtStart.End {
					fmt.Println("noFoundSourceRangeAtStart wrong: ", test, noFoundSourceRangeAtStart, std)
					panic("testStart < sourceStart && testEnd <= sourceEnd")
				}
				if noFoundSourceRangeAtEnd.Start > noFoundSourceRangeAtEnd.End {
					fmt.Println("noFoundSourceRangeAtEnd wrong: ", test, noFoundSourceRangeAtEnd, std)
					panic("testStart < sourceStart && testEnd <= sourceEnd")
				}
			} else {
				fmt.Println("ALL WRONG: ", test, std)
				fmt.Println(sourceStart, sourceEnd)
				panic("why are we here?")
			}
		}
	}
	return destinationRanges, noFoundSourceRanges
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
	scanner.Scan()

	analysis := make(map[string][]SourceToDestination)
	key := ""
	records := []SourceToDestination{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			analysis[key] = records
			records = []SourceToDestination{}
			key = ""
		} else if strings.Contains(line, ":") {
			key = line
		} else {
			tokens := strings.Split(line, " ")
			destination, _ := strconv.Atoi(tokens[0])
			source, _ := strconv.Atoi(tokens[1])
			rangeValue, _ := strconv.Atoi(tokens[2])
			std := SourceToDestination{Source: source, Destination: destination, Range: rangeValue}
			records = append(records, std)
		}
	}
	if len(key) > 0 {
		analysis[key] = records
	}

	seedStart := 0
	seedtosoilmap := analysis[seed_to_soil]
	soiltofertilizermap := analysis[soil_to_fertilizer]
	ferttowatermap := analysis[fertilizer_to_water]
	watertolight := analysis[water_to_light]
	lighttotemp := analysis[light_to_temperature]
	temptohumid := analysis[temperature_to_humidity]
	humidtoloc := analysis[humidity_to_location]

	noDestinationDetermined := []Range{}
	for index, value := range seeds {
		if index%2 == 0 {
			seedStart = value
		} else {
			seedRange := Range{seedStart, seedStart + value - 1}
			noDestinationDetermined = append(noDestinationDetermined, seedRange)
		}
	}

	destinationMaps := [][]SourceToDestination{seedtosoilmap, soiltofertilizermap, ferttowatermap, watertolight, lighttotemp, temptohumid, humidtoloc}

	destinations := []Range{}
	for _, destinationMap := range destinationMaps {
		for _, std := range destinationMap {
			tempDest, tempNoDestinationDetermined := std.GetRanges(noDestinationDetermined, true)
			noDestinationDetermined = tempNoDestinationDetermined
			for _, tempdest := range tempDest {
				destinations = append(destinations, tempdest)
			}
		}
		//get ready for next round of destination checking
		for _, dest := range destinations {
			noDestinationDetermined = append(noDestinationDetermined, dest)
		}
		destinations = nil
	}

	lowestSeedvalue := 999999999
	for _, source := range noDestinationDetermined {
		if source.Start < lowestSeedvalue {
			lowestSeedvalue = source.Start
		}
	}
	fmt.Println(lowestSeedvalue)
}
