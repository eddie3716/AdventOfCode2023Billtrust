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

type Pattern []int
type Template string

type Set struct {
	Pattern  []int
	Template Template
}

func (s Set) Fold() Set {
	newTemplate := s.Template + "?" + s.Template + "?" + s.Template + "?" + s.Template + "?" + s.Template
	newPattern := []int{}
	newPattern = append(newPattern, s.Pattern...)
	newPattern = append(newPattern, s.Pattern...)
	newPattern = append(newPattern, s.Pattern...)
	newPattern = append(newPattern, s.Pattern...)
	newPattern = append(newPattern, s.Pattern...)
	newSet := Set{newPattern, Template(newTemplate)}
	return newSet
}

func (set Set) FindCombos() int {
	cache := make(map[string]int)
	result := set.FindCombinations(0, 0, 0, cache)
	return result
}

func (set Set) FindCombinations(iTemplate, iPattern int, lengthOfPattern int, cache map[string]int) int {
	key := fmt.Sprintf("%d_%d_%d", iPattern, iTemplate, lengthOfPattern)

	//fmt.Println("Key", key)

	if val, ok := cache[key]; ok {
		//fmt.Println("Cache hit", key, val)
		return val
	}

	if iTemplate == len(set.Template) {
		if iPattern == len(set.Pattern) && lengthOfPattern == 0 {
			return 1
		} else if iPattern == len(set.Pattern)-1 && set.Pattern[iPattern] == lengthOfPattern {
			return 1
		} else {
			return 0
		}
	}

	combos := 0
	currentChar := set.Template[iTemplate]

	if currentChar == '?' || currentChar == '.' {
		if lengthOfPattern == 0 {
			combos += set.FindCombinations(iTemplate+1, iPattern, 0, cache)
		} else if lengthOfPattern > 0 && iPattern < len(set.Pattern) && set.Pattern[iPattern] == lengthOfPattern {
			combos += set.FindCombinations(iTemplate+1, iPattern+1, 0, cache)
		}
	}

	if currentChar == '?' || currentChar == '#' {
		combos += set.FindCombinations(iTemplate+1, iPattern, lengthOfPattern+1, cache)
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

		pattern := Pattern{}
		for _, token := range strings.Split(tokens[1], ",") {
			number, _ := strconv.Atoi(token)
			pattern = append(pattern, number)
		}
		set := Set{pattern, Template(tokens[0])}
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
