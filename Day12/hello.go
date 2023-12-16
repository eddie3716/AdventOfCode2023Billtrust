package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sets := parseFile("input.txt")

	part1(sets)

	part2(sets)
}

type Set struct {
	BrokenSpringBlocks []int
	SpringArray        string
}

func (s Set) Fold() Set {
	newSpringArray := s.SpringArray + "?" + s.SpringArray + "?" + s.SpringArray + "?" + s.SpringArray + "?" + s.SpringArray
	newBrokenSpringBlocks := []int{}
	newBrokenSpringBlocks = append(newBrokenSpringBlocks, s.BrokenSpringBlocks...)
	newBrokenSpringBlocks = append(newBrokenSpringBlocks, s.BrokenSpringBlocks...)
	newBrokenSpringBlocks = append(newBrokenSpringBlocks, s.BrokenSpringBlocks...)
	newBrokenSpringBlocks = append(newBrokenSpringBlocks, s.BrokenSpringBlocks...)
	newBrokenSpringBlocks = append(newBrokenSpringBlocks, s.BrokenSpringBlocks...)
	newSet := Set{newBrokenSpringBlocks, newSpringArray}
	return newSet
}

func (set Set) FindCombos() int {
	cache := make(map[string]int)
	result := set.FindCombinations(0, 0, 0, cache)
	return result
}

func (set Set) FindCombinations(iCurrentSpring, iCurrentBrokenBlock int, currentLengthOfBrokenBlock int, cache map[string]int) int {
	//this will be a unique key across all combinations
	key := fmt.Sprintf("%d_%d_%d", iCurrentBrokenBlock, iCurrentSpring, currentLengthOfBrokenBlock)

	//fmt.Println("Key", key)

	if val, ok := cache[key]; ok {
		//fmt.Println("Cache hit", key, val)
		return val
	}

	if iCurrentSpring == len(set.SpringArray) {
		// we've reached the end of our spring array

		if iCurrentBrokenBlock == len(set.BrokenSpringBlocks) && currentLengthOfBrokenBlock == 0 {
			// we've check all broken spring blocks and don't have any in progress that we're checking, so all other blocks were matched
			return 1
		} else if iCurrentBrokenBlock == len(set.BrokenSpringBlocks)-1 && set.BrokenSpringBlocks[iCurrentBrokenBlock] == currentLengthOfBrokenBlock {
			// we've been checking the last broken spring block and determined that it's a match
			return 1
		} else {
			// we've reached the end of our spring array but we still have broken spring blocks to check, so we weren't able to match all blocks
			return 0
		}
	}

	combos := 0
	currentSpring := set.SpringArray[iCurrentSpring]

	//possible match with our current broken block...increment and check next spring array index
	//keeps down this path until we either get to the end of the spring array or terminate at a '.'
	if currentSpring == '?' || currentSpring == '#' {
		combos += set.FindCombinations(iCurrentSpring+1, iCurrentBrokenBlock, currentLengthOfBrokenBlock+1, cache)
	}

	//need to check other branches...we don't branch on a '#' because we know that has to match, and to do so would skew our results
	if currentSpring != '#' {
		if currentLengthOfBrokenBlock == 0 {
			// We never had a match, so check the next spring for matches
			combos += set.FindCombinations(iCurrentSpring+1, iCurrentBrokenBlock, 0, cache)
		} else if currentLengthOfBrokenBlock > 0 && iCurrentBrokenBlock < len(set.BrokenSpringBlocks) && set.BrokenSpringBlocks[iCurrentBrokenBlock] == currentLengthOfBrokenBlock {

			// we thought we had a match with the current spring, but that turned out to be false.
			combos += set.FindCombinations(iCurrentSpring+1, iCurrentBrokenBlock+1, 0, cache)
		}
	}

	cache[key] = combos

	return combos
}

type Sets []Set

func parseFile(fileName string) Sets {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sets := Sets{}
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		brokenSpringArray := []int{}
		for _, token := range strings.Split(tokens[1], ",") {
			number, _ := strconv.Atoi(token)
			brokenSpringArray = append(brokenSpringArray, number)
		}
		set := Set{brokenSpringArray, tokens[0]}
		sets = append(sets, set)
	}

	return sets
}

func part1(sets Sets) {

	totalCombinations := 0
	for _, set := range sets {
		totalCombinations += set.FindCombos()
	}

	fmt.Println("Part 1: ", totalCombinations)
}

func part2(sets Sets) {
	totalCombinations := 0
	for _, set := range sets {
		totalCombinations += set.Fold().FindCombos()
	}

	fmt.Println("Part 1: ", totalCombinations)
}
