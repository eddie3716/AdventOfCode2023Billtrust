package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sets := parseFile("input.txt")

	//part1(sets)

	part2(sets)
}

func parseFile(fileName string) [][]rune {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	contraption := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		contraption = append(contraption, []rune(line))
	}

	return contraption
}

func printContraption(contraption [][]rune) {
	for _, line := range contraption {
		fmt.Println(string(line))
	}
	fmt.Println()
}

const (
	North = iota*10 + 10
	South
	East
	West
)

type Path struct {
	startX int
	startY int
	North  *Path
	South  *Path
	East   *Path
	West   *Path
}

func (p Path) GetKey() string {
	return fmt.Sprintf("%d,%d", p.startX, p.startY)
}

func (p Path) Visit(visited map[string]map[int]bool, direction int) bool {
	key := p.GetKey()
	if didVisit, _ := visited[key][direction]; !didVisit {
		visited[key][direction] = true
		return true
	}
	return false
}

func goNorth(visited map[string]map[int]bool, contraption [][]rune, startX int, startY int) *Path {
	if startX < 0 || startX >= len(contraption[0]) || startY < 0 || startY >= len(contraption) {
		return nil
	}
	path := Path{startX, startY, nil, nil, nil, nil}
	if !path.Visit(visited, North) {
		return nil
	}

	if contraption[startY][startX] == '.' || contraption[startY][startX] == '|' {
		path.North = goNorth(visited, contraption, startX, startY-1)
	} else if contraption[startY][startX] == '-' {
		path.West = goWest(visited, contraption, startX-1, startY)
		path.East = goEast(visited, contraption, startX+1, startY)
	} else if contraption[startY][startX] == '\\' {
		path.West = goWest(visited, contraption, startX-1, startY)
	} else if contraption[startY][startX] == '/' {
		path.East = goEast(visited, contraption, startX+1, startY)
	} else {
		fmt.Println("Found a letter: ", contraption[startY][startX])
		panic("Found a letter")
	}
	return &path
}

func goSouth(visited map[string]map[int]bool, contraption [][]rune, startX int, startY int) *Path {
	if startX < 0 || startX >= len(contraption[0]) || startY < 0 || startY >= len(contraption) {
		return nil
	}
	path := Path{startX, startY, nil, nil, nil, nil}

	if !path.Visit(visited, South) {
		return nil
	}

	if contraption[startY][startX] == '.' || contraption[startY][startX] == '|' {
		path.South = goSouth(visited, contraption, startX, startY+1)
	} else if contraption[startY][startX] == '-' {
		path.West = goWest(visited, contraption, startX-1, startY)
		path.East = goEast(visited, contraption, startX+1, startY)
	} else if contraption[startY][startX] == '\\' {
		path.East = goEast(visited, contraption, startX+1, startY)
	} else if contraption[startY][startX] == '/' {
		path.West = goWest(visited, contraption, startX-1, startY)
	} else {
		fmt.Println("Found a letter: ", contraption[startY][startX])
		panic("Found a letter")
	}
	return &path
}

func goWest(visited map[string]map[int]bool, contraption [][]rune, startX int, startY int) *Path {
	if startX < 0 || startX >= len(contraption[0]) || startY < 0 || startY >= len(contraption) {
		return nil
	}
	path := Path{startX, startY, nil, nil, nil, nil}
	if !path.Visit(visited, West) {
		return nil
	}
	if contraption[startY][startX] == '.' || contraption[startY][startX] == '-' {
		path.West = goWest(visited, contraption, startX-1, startY)
	} else if contraption[startY][startX] == '|' {
		path.North = goNorth(visited, contraption, startX, startY-1)
		path.South = goSouth(visited, contraption, startX, startY+1)
	} else if contraption[startY][startX] == '\\' {
		path.North = goNorth(visited, contraption, startX, startY-1)
	} else if contraption[startY][startX] == '/' {
		path.South = goSouth(visited, contraption, startX, startY+1)
	} else {
		fmt.Println("Found a letter: ", contraption[startY][startX])
		panic("Found a letter")
	}
	return &path
}

func goEast(visited map[string]map[int]bool, contraption [][]rune, startX int, startY int) *Path {
	if startX < 0 || startX >= len(contraption[0]) || startY < 0 || startY >= len(contraption) {
		return nil
	}
	path := Path{startX, startY, nil, nil, nil, nil}

	if !path.Visit(visited, East) {
		return nil
	}

	if contraption[startY][startX] == '.' || contraption[startY][startX] == '-' {
		path.East = goEast(visited, contraption, startX+1, startY)
	} else if contraption[startY][startX] == '|' {
		path.North = goNorth(visited, contraption, startX, startY-1)
		path.South = goSouth(visited, contraption, startX, startY+1)
	} else if contraption[startY][startX] == '\\' {
		path.South = goSouth(visited, contraption, startX, startY+1)
	} else if contraption[startY][startX] == '/' {
		path.North = goNorth(visited, contraption, startX, startY-1)
	} else {
		fmt.Println("Found a letter: ", contraption[startY][startX])
		panic("Found a letter")
	}
	return &path
}

func part1(contraption [][]rune) {

	printContraption(contraption)
	answer := 0

	visited := make(map[string]map[int]bool)
	for i := 0; i < len(contraption); i++ {
		for j := 0; j < len(contraption[i]); j++ {
			key := fmt.Sprintf("%d,%d", i, j)
			visited[key] = make(map[int]bool)
			visited[key][East] = false
			visited[key][West] = false
			visited[key][North] = false
			visited[key][South] = false
		}
	}

	path := goEast(visited, contraption, 0, 0)
	fmt.Println(visited)

	printContraption(contraption)
	for i := 0; i < len(contraption); i++ {
		for j := 0; j < len(contraption[i]); j++ {
			key := fmt.Sprintf("%d,%d", j, i)
			visitCount := 0
			visitDirection := 0
			for key, value := range visited[key] {
				if value {
					visitCount++
					visitDirection = key
				}
			}
			if visitCount == 1 && contraption[i][j] == '.' {
				switch visitDirection {
				case North:
					contraption[i][j] = '^'
					break
				case South:
					contraption[i][j] = 'v'
					break
				case East:
					contraption[i][j] = '>'
					break
				case West:
					contraption[i][j] = '<'
					break
				default:
					break
				}
			} else if visitCount > 1 && contraption[i][j] == '.' {
				contraption[i][j] = rune(visitCount + '0')
			}
		}
	}

	printContraption(contraption)
	for path != nil {
		fmt.Println(*path)
		path = path.East
	}

	for _, value := range visited {
		for _, visited := range value {
			if visited {
				answer++
				break
			}
		}
	}

	fmt.Println("Part 1: ", answer)
}

func part2(contraption [][]rune) {

	printContraption(contraption)
	answer := 0

	visited := make(map[string]map[int]bool)
	for i := 0; i < len(contraption); i++ {
		for j := 0; j < len(contraption[i]); j++ {
			key := fmt.Sprintf("%d,%d", i, j)
			visited[key] = make(map[int]bool)
			visited[key][East] = false
			visited[key][West] = false
			visited[key][North] = false
			visited[key][South] = false
		}
	}

	for i := 0; i < len(contraption[0]); i++ {
		goSouth(visited, contraption, i, 0)
		answer = Max(answer, countVisits(visited))
		resetCache(visited)
	}

	for i := 0; i < len(contraption[0]); i++ {
		goNorth(visited, contraption, i, len(contraption)-1)
		answer = Max(answer, countVisits(visited))
		resetCache(visited)
	}

	for i := 0; i < len(contraption); i++ {
		goEast(visited, contraption, 0, i)
		answer = Max(answer, countVisits(visited))
		resetCache(visited)
	}

	for i := 0; i < len(contraption); i++ {
		goWest(visited, contraption, len(contraption[0])-1, i)
		answer = Max(answer, countVisits(visited))
		resetCache(visited)
	}

	fmt.Println("Part 2: ", answer)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func countVisits(visited map[string]map[int]bool) int {
	count := 0
	for _, value := range visited {
		for _, visited := range value {
			if visited {
				count++
				break
			}
		}
	}
	return count
}

func resetCache(visited map[string]map[int]bool) {
	for _, value := range visited {
		for key := range value {
			value[key] = false
		}
	}
}
