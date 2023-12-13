package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := parseFile("testinput.txt")

	part1(data)

	//part2(gridPatterns)
}

func WillFit(t string, number rune, index int) (bool, int) {
	skew := int(number) - 1

	if skew < 0 {
		fmt.Println(skew, " ", number, " ", index, " ", t)
		panic("will not fit")
	}

	endIndex := index + skew

	if endIndex >= len(t) {
		fmt.Println(skew, " ", number, " ", index, " ", t)
		panic("How does this even happen")
	}

	nextIndex := index + skew + 1
	testFirstSymbol := t[index]
	testEndSymbol := t[endIndex]
	switch testFirstSymbol {
	case '?':
		if testEndSymbol == '?' || (testEndSymbol == '#' && nextIndex < len(t) && t[nextIndex] != '#') {
			return true, nextIndex
		} else {
			return false, index + 1
		}
	case '#':
		if (testEndSymbol == '?') || (testEndSymbol == '#' && nextIndex < len(t) && t[nextIndex] != '#') {
			return true, nextIndex
		} else {
			return false, index + 1
		}
	default:
		fmt.Println(skew, " ", number, " ", index, " ", testFirstSymbol, " ", testEndSymbol, t)
		panic("Unknown symbol")
	}
	fmt.Println(skew, " ", number, " ", index, " ", testFirstSymbol, " ", testEndSymbol, t)
	panic("Should not be here")
}

type Pattern []int

type Set struct {
	Pattern  []int
	Template string
}

func (s Set) FindCombinations(start int) (int, bool) {
	found := true

	if found {

		newSet := Set{s.Pattern[start+1:], s.Template[start+1:]}
		return newSet.FindCombinations(start + 1)
	}
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
		set := Set{pattern, Template(tokens[0])}
		sets = append(sets, set)
	}

	sets.PrintGridPatterns()
	return sets
}

func part1(sets Sets) {

	for index, pattern := range sets.Patterns {
		templates := sets.Grid[index]
		iTemplates := 0
		iTemplate := 0
		iNumber := 0
		willFit := true
		currentTemplate := templates[iTemplates]
		for iNumber = 0; iNumber < len(pattern); iNumber++ {
			number := pattern[iNumber]
			if iTemplate >= len(currentTemplate) {
				iTemplates++
				currentTemplate = templates[iTemplates]
				iTemplate = 0
			}
			fmt.Println(number, " ", currentTemplate, " ", iTemplate)
			tempITemplate := iTemplate
			willFit, iTemplate = currentTemplate.WillFit(rune(number), iTemplate)
			if willFit {
				[]rune(currentTemplate)[tempITemplate] = '#'
			}
			if !willFit {
				fmt.Println("No match")
				break
			}
		}

		fmt.Println(gridPatterns)
	}
}

func part2(gridPatterns Sets) {
	fmt.Println("Part 2")
}
