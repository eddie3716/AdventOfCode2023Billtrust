package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type Part struct {
	x int
	m int
	a int
	s int
}

func greaterThan(a, b int) bool {
	return a > b
}

func lessThan(a, b int) bool {
	return a < b
}

func nop(a, b int) bool {
	return true
}

const (
	GreaterThan = "GreaterThan"
	LessThan    = "LessThan"
	Result      = "Result"
)

type Node struct {
	function    func(a, b int) bool
	operandType string
	operand2    int
	next        string
}

func getFunction(nodeType string) func(a, b int) bool {
	switch nodeType {
	case GreaterThan:
		return greaterThan
	case LessThan:
		return lessThan
	case Result:
		return nop
	default:
		panic("unknown node type: " + nodeType)
	}
}

func (workflow *Workflow) evaluate(part *Part, workflows *orderedmap.OrderedMap[string, *Workflow]) string {
	for _, node := range *workflow.node {
		operand1 := -1
		switch node.operandType {
		case "a":
			operand1 = part.a
			break
		case "x":
			operand1 = part.x
			break
		case "m":
			operand1 = part.m
			break
		case "s":
			operand1 = part.s
		}

		if node.function(operand1, node.operand2) {
			if nextNode, present := workflows.Get(node.next); present {
				return nextNode.evaluate(part, workflows)
			} else {
				return node.next
			}
		}
	}

	return "0"
}

type Workflow struct {
	node *[]Node
}

func part1(parts *[]Part, workflows *orderedmap.OrderedMap[string, *Workflow]) int {
	return 0
}

func part2(parts *[]Part, workflows *orderedmap.OrderedMap[string, *Workflow]) int {
	return 0
}

func main() {
	parts, workflows := parseFile("input.txt")

	fmt.Println("answer part 1:", part1(parts, workflows))

	fmt.Println("answer part 2:", part2(parts, workflows))
}

func parseWorkFlow(workflowString string) *Workflow {
	workflow := Workflow{}
	return &workflow
}

func parseFile(fileName string) (*[]Part, *orderedmap.OrderedMap[string, *Workflow]) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	workFlows := orderedmap.New[string, *Workflow]()
	parts := []Part{}

	checkWorkflows := true
	for scanner.Scan() {
		line := scanner.Text()
		if checkWorkflows && line == "" {
			checkWorkflows = false
		}
		if checkWorkflows {
			workflowKey := line[0 : strings.Index(line, "{")-1]

			workflow := parseWorkFlow(line[strings.Index(line, "{")+1 : strings.Index(line, "}")-1])

			workFlows.Set(workflowKey, workflow)

		} else {

		}
	}

	return &parts, workFlows
}
