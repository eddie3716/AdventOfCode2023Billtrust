package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value    int
	Prev     *Node
	Next     *Node
	Ancestor *Node
}

func (node *Node) Reverse() *Node {
	// Reverse the linked list
	var prev *Node = nil
	current := node
	for current != nil {
		next := current.Next
		current.Next = prev
		current.Prev = next
		prev = current
		current = next
	}
	node = prev
	return node
}

func (node *Node) PrintGraph(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("   ")
	}
	for currentNode := node; currentNode != nil; currentNode = currentNode.Next {
		fmt.Printf("%6d", currentNode.Value)
	}
	fmt.Println()
	if node.Ancestor != nil {
		node.Ancestor.PrintGraph(depth + 1)
	}
}

func (currentGeneration *Node) GenerateGraph() {
	startNode := currentGeneration
	for startNode != nil {
		ancestorBuffer := []*Node{}
		for currentNode := startNode; currentNode != nil; currentNode = currentNode.Next {
			if currentNode.Next != nil {
				ancestorValue := currentNode.Next.Value - currentNode.Value
				ancestor := &Node{Value: ancestorValue, Ancestor: nil}
				currentNode.Ancestor = ancestor
				ancestorBuffer = append(ancestorBuffer, ancestor)
			}
		}

		for index, ancestor := range ancestorBuffer {
			if index < len(ancestorBuffer)-1 {
				ancestor.Next = ancestorBuffer[index+1]
				ancestorBuffer[index+1].Prev = ancestor
			}
		}
		startNode = nil
		if len(ancestorBuffer) > 1 {
			startNode = ancestorBuffer[0]
		}
		ancestorBuffer = nil
	}
	if currentGeneration.Ancestor != nil {
		currentGeneration.Ancestor.GenerateGraph()
	}
}

func (node *Node) ComputeNextValue() int {
	ending := node
	for ; ending.Next != nil; ending = ending.Next {
	}

	if ending.Ancestor == nil && (ending.Prev == nil || ending.Prev.Ancestor == nil) {
		return ending.Value
	} else {
		return ending.Value + ending.Prev.Ancestor.ComputeNextValue()
	}
}

func (node *Node) ComputePrevValue() int {
	starting := node

	if starting.Ancestor == nil && starting.Next == nil {
		return starting.Value
	} else {
		return starting.Value - starting.Ancestor.ComputePrevValue()
	}
}

func parseFile(fileName string) []*Node {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	startingNodes := []*Node{}

	var prevNode *Node = nil
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		for _, token := range tokens {
			value, _ := strconv.Atoi(token)
			currentNode := Node{Next: nil, Value: value}
			if prevNode == nil {
				startingNodes = append(startingNodes, &currentNode)
			} else {
				prevNode.Next = &currentNode
				currentNode.Prev = prevNode
			}
			prevNode = &currentNode
		}
		prevNode = nil
	}

	return startingNodes
}

func part1(nodes []*Node) {
	nextGeneration := []int{}
	for _, node := range nodes {
		node.GenerateGraph()
		node.PrintGraph(0)
		nextGeneration = append(nextGeneration, node.ComputeNextValue())
	}
	fmt.Println(nextGeneration)

	sum := 0
	for _, value := range nextGeneration {
		sum += value
	}
	fmt.Println(sum)
}

func part2(nodes []*Node) {
	nextGeneration := []int{}
	for _, node := range nodes {
		node.GenerateGraph()
		node.PrintGraph(0)
		nextGeneration = append(nextGeneration, node.ComputePrevValue())
	}
	fmt.Println(nextGeneration)

	sum := 0
	for _, value := range nextGeneration {
		sum += value
	}
	fmt.Println(sum)
}

func main() {
	currentGeneration := parseFile("input.txt")

	//part1(currentGeneration)
	//part2(currentGeneration)
	for index, node := range currentGeneration {
		newNode := node.Reverse()
		currentGeneration[index] = newNode
	}
	part1(currentGeneration)
}
