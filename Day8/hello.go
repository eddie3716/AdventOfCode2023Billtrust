package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Counter struct {
	StartKey       string
	LastCheckedKey string
	Steps          uint
	Index          int
}

type Entry struct {
	Left  string
	Right string
}

func maxValueOfCounters(a []Counter) uint {
	var max uint = 0
	for i := 1; i < len(a); i++ {
		if a[i].Steps > a[max].Steps {
			max = uint(i)
		}
	}
	return a[max].Steps
}

func countersAreEqual(a []Counter) bool {
	for i := 1; i < len(a); i++ {
		if a[i].Steps != a[0].Steps {
			return false
		}
	}
	return true
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
	//fmt.Println(movement, " ", entries, " ", len(entries))

	return movement, entries
}

func part1(movement string, currentKey string, terminatingValue string, entries map[string]Entry) (string, uint) {
	var steps uint = 0

restart:
	for i := 0; i < len(movement) && currentKey != terminatingValue; i++ {
		entry := entries[currentKey]
		currentDirection := movement[i]
		//oldKey := currentKey
		if currentDirection == 'L' {
			currentKey = entry.Left
		} else if currentDirection == 'R' {
			currentKey = entry.Right
		} else {
			panic("bad direction")
		}
		steps += 1
		//fmt.Println("Current Steps ", startKey, ": ", stepsToReachZ, " ", " ", oldKey, " ", currentKey, " ", strconv.Itoa(uint(currentDirection)), " ", entry, " ", i)
	}
	if !strings.HasSuffix(currentKey, terminatingValue) {
		goto restart
	}
	//197135 is too high...5111 is too low
	//fmt.Println("Part 1: ", stepsToReachZ, " steps to reach Z")
	return currentKey, steps
}

func part2(movement string, entries map[string]Entry) uint {
	counters := []Counter{}
	keysThatEndWithA := []string{}
	for key := range entries {
		if strings.HasSuffix(key, "A") {
			keysThatEndWithA = append(keysThatEndWithA, key)
		}
	}

	for index, currentKey := range keysThatEndWithA {
		counters = append(counters, Counter{currentKey, currentKey, 0, index})
	}

	// fmt.Println(keysThatEndWithA)
	// fmt.Println(entries)
	var maxSteps uint = 1
tryagain:

	finishedCounters := make(chan *Counter, len(counters))
	wg := sync.WaitGroup{}
	for _, counter := range counters {
		if counter.Steps == maxSteps {
			fmt.Println("Skipping ", counter, " because it has already reached max steps")
			continue
		} else if counter.Steps > maxSteps {
			panic("counter steps is greater than max steps")
		}
		wg.Add(1)
		go func(c *Counter, ch chan *Counter) {
			//fmt.Println("Starting ", c)
			lastKey, steps := part1(movement, c.LastCheckedKey, "Z", entries)
			c.LastCheckedKey = lastKey
			c.Steps += steps
			//fmt.Println("Finished ", c)

			ch <- c
			defer wg.Done()
		}(&counter, finishedCounters)
	}
	//fmt.Println("Waiting for all to finish")
	wg.Wait()
	close(finishedCounters)

	for {
		counter, ok := <-finishedCounters
		if !ok {
			break
		}
		updateCounter := counters[counter.Index]
		updateCounter.Steps = counter.Steps
		updateCounter.LastCheckedKey = counter.LastCheckedKey
		counters[counter.Index] = updateCounter
	}

	if !countersAreEqual(counters) {
		maxSteps = maxValueOfCounters(counters)
		fmt.Println("Max steps so far: ", maxSteps, counters)

		goto tryagain
	}

	fmt.Println("Part 2: ", counters[0].Steps, " steps to reach end")
	return counters[0].Steps
}
func main() {
	movement, entries := parseFile("testinput.txt")

	//part1(movement, "AAA", "ZZZ", entries)
	part2(movement, entries)
}
