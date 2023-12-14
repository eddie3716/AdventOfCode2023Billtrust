package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sets := parseFile("testinput.txt")

	part1(sets)

	part2(sets)
}

type Pattern []int
type Template string

func (t Template) replaceAtIndex(r rune, index int, number int) (string, int) {
	out := []rune(t)
	new := 0
	for i := index; i < index+number; i++ {
		if out[i] != r {
			out[i] = r
			new++
		}
	}
	return string(out), new
}

type Set struct {
	Pattern        []int
	Template       Template
	StartingPounds int
	ExpectedPounds int
}

func NewSet(pattern Pattern, template Template) Set {
	set := Set{pattern, template, 0, 0}

	for _, r := range template {
		if r == '#' {
			set.StartingPounds++
		}
	}

	for _, r := range pattern {
		set.ExpectedPounds += r
	}

	return set
}

func (s Set) Fold() Set {
	newTemplate := s.Template + "?" + s.Template + "?" + s.Template + "?" + s.Template + "?" + s.Template
	newPattern := []int{}
	newPattern = append(newPattern, s.Pattern...)
	newPattern = append(newPattern, s.Pattern...)
	newPattern = append(newPattern, s.Pattern...)
	newPattern = append(newPattern, s.Pattern...)
	newPattern = append(newPattern, s.Pattern...)
	newSet := NewSet(newPattern, Template(newTemplate))
	return newSet
}

func (s Set) PatternFits(iPattern, iTemplate int) bool {

	if iPattern >= len(s.Pattern) {
		return false
	}

	num := s.Pattern[iPattern]
	endIndex := iTemplate + num - 1
	if endIndex >= len(s.Template) {
		return false
	}

	if (s.Template[iTemplate] == '#' || s.Template[iTemplate] == '?') && (iTemplate == 0 || s.Template[iTemplate-1] != '#') && (endIndex == len(s.Template)-1 || s.Template[endIndex+1] != '#') {

		for i := iTemplate; i <= endIndex; i++ {
			if s.Template[i] == '.' {
				return false
			}
		}

		return true
	}
	return false
}

func (s Set) FindCombos() int {
	cache := make(map[int]bool)
	return s.FindCombinations(0, 0, 0, cache)
}

func (s Set) FindCombinations(iPattern int, iTemplate int, depth int, cache map[int]bool) int {
	if depth == 0 {
		fmt.Println("Set ", s, " starting...")
	}

	if iTemplate > len(s.Template) {
		return 0
	}
	combos := 0
	newDepth := depth + 1
	itFits := s.PatternFits(iPattern, iTemplate)
	if itFits {
		newTemplate, newPounds := s.Template.replaceAtIndex('#', iTemplate, s.Pattern[iPattern])
		if s.StartingPounds+newPounds == s.ExpectedPounds {
			//s.Template = Template(strings.ReplaceAll(string(s.Template), "?", "."))
			fmt.Println("Set ", s, " match")
			combos++
		}
		newSet := Set{s.Pattern, Template(newTemplate), s.StartingPounds + newPounds, s.ExpectedPounds}

		combos += newSet.FindCombinations(iPattern+1, iTemplate+s.Pattern[iPattern], newDepth, cache)
	}

	combos += s.FindCombinations(iPattern, iTemplate+1, newDepth, cache)

	if depth == 0 {
		fmt.Println("Set ", s, " has ", combos, " combinations")
	}
	return combos
}

type Sets []Set

func (ss Sets) PrintGridPatterns() {
	for _, s := range ss {
		fmt.Print(s.Template)
		fmt.Println(" ", s.Pattern)
	}
}

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
		set := NewSet(pattern, Template(tokens[0]))
		sets = append(sets, set)
	}

	//sets.PrintGridPatterns()
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
	// for _, set := range sets {
	// 	totalCombinations += set.Fold().FindCombos()
	// }

	fmt.Println("Part 1: ", totalCombinations)
}
