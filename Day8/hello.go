package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	Left  string
	Right string
}

func parseFile(fileName string) (string, map[string]Entry) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	movement := scanner.Text()

	entries := map[string]Entry{}

	for scanner.Scan() {
		metatokens := strings.Split(scanner.Text(), " = ")
		if len(metatokens) != 2 {
			continue
		}
		key := metatokens[0]

		entrytokens := strings.Split(metatokens[1], ", ")
		if len(entrytokens) != 2 {
			panic("bad entry")
		}
		lefttoken := entrytokens[0][1:]
		righttoken := entrytokens[1][0 : len(entrytokens[1])-1]
		entry := Entry{Left: lefttoken, Right: righttoken}
		entries[key] = entry
	}

	return movement, entries
}

func part1(movement string, currentKey string, terminatingValue string, entries map[string]Entry) uint {
	var steps uint = 0
	for !strings.HasSuffix(currentKey, terminatingValue) {
		for i := 0; i < len(movement) && currentKey != terminatingValue; i++ {
			entry := entries[currentKey]
			currentDirection := movement[i]
			if currentDirection == 'L' {
				currentKey = entry.Left
			} else if currentDirection == 'R' {
				currentKey = entry.Right
			} else {
				panic("bad direction")
			}
			steps++
		}
	}

	return steps
}

func part2(movement string, entries map[string]Entry) uint {
	keysThatEndWithA := []string{}
	for key := range entries {
		if strings.HasSuffix(key, "A") {
			keysThatEndWithA = append(keysThatEndWithA, key)
		}
	}

	answers := []uint{}
	for _, key := range keysThatEndWithA {
		steps := part1(movement, key, "Z", entries)
		answers = append(answers, steps)
	}

	finalAnswer := lcmAll(answers[0], answers[1:]...)

	fmt.Println("Part 2: ", finalAnswer, " steps to reach end")
	return finalAnswer
}

func gcd(a, b uint) uint {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func lcm(a, b uint) uint {
	return a / gcd(a, b) * b
}
func lcmAll(a uint, bs ...uint) uint {
	result := a
	for _, b := range bs {
		result = lcm(result, b)
	}

	return result
}

func main() {
	movement, entries := parseFile("input.txt")

	//part1(movement, "AAA", "ZZZ", entries)
	part2(movement, entries)
}
