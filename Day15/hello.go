package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func main() {
	sets := parseFile("input.txt")

	part1(sets)

	part2(sets)
}

func hash(token string) int {
	currentValue := 0
	for _, char := range token {
		currentValue += int(char)
		currentValue *= 17
		currentValue %= 256
	}
	//fmt.Println("token: ", token, " hashes to: ", currentValue)
	return currentValue
}

func parseFile(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	tokens := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == ',' || r == ' '
		})
		tokens = append(tokens, fields...)
	}

	return tokens
}

func part1(tokens []string) {

	answer := 0
	for _, token := range tokens {
		answer += hash(token)
	}

	fmt.Println("Part 1: ", answer)
}

func part2(tokens []string) {

	boxes := []orderedmap.OrderedMap[string, int]{}
	for i := 0; i < 256; i++ {
		boxes = append(boxes, *orderedmap.New[string, int]())
	}
	answer := 0
	for _, token := range tokens {
		if strings.Contains(token, "-") {
			key := strings.Split(token, "-")[0]
			hash := hash(key)
			box := boxes[hash]
			//fmt.Println("deleting: ", key, " from box: ", hash)
			box.Delete(key)
		} else if strings.Contains(token, "=") {
			kvp := strings.Split(token, "=")
			key := kvp[0]
			hash := hash(key)
			box := boxes[hash]
			value, err := strconv.Atoi(kvp[1])
			if err != nil {
				panic(err)
			}
			//fmt.Println("setting: ", key, " to: ", value, " in box: ", hash)
			box.Set(key, value)
		} else {
			panic("unknown token: " + token)
		}
	}

	for iBox, box := range boxes {
		boxNumber := iBox + 1
		lenseNumber := 1
		for pair := box.Oldest(); pair != nil; pair = pair.Next() {
			//fmt.Println("box: ", iBox, ", lenseNumber: ", lenseNumber, ", pair: ", pair.Key, " = ", pair.Value)
			answer += pair.Value * lenseNumber * boxNumber
			lenseNumber++
		}
	}

	fmt.Println("Part 2: ", answer)
}
