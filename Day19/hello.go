package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PartRange struct {
	max int64
	min int64
}

type Part struct {
	x int64
	m int64
	a int64
	s int64
}

func greaterThan(a, b int64) bool {
	return a > b
}

func lessThan(a, b int64) bool {
	return a < b
}

func nop(a, b int64) bool {
	return true
}

const (
	GreaterThan = ">"
	LessThan    = "<"
	Result      = ""
)

type Node struct {
	function    func(a, b int64) bool
	operator    string
	operandName string
	operand     int64
	next        string
}

func getFunction(nodeType string) func(a, b int64) bool {
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

func copy(a, b *map[string]int) {
	for k, v := range *b {
		(*a)[k] = v
	}
}

func maxInt(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func getCombinations(partRanges *map[string]PartRange, next string, workflows *map[string]*Workflow) int64 {

	if next == "A" {
		var combos int64 = 1
		for _, partRange := range *partRanges {
			combos *= partRange.max - partRange.min + 1
		}
		return combos
	} else if next == "R" {
		return 0
	}

	workflow, present := (*workflows)[next]

	if !present {
		panic("workflow not found: " + next)
	}

	var combos int64 = 0
	for i := 0; i < len(*workflow.nodes)-1; i++ {
		node := (*workflow.nodes)[i]
		min := (*partRanges)[node.operandName].min
		max := (*partRanges)[node.operandName].max

		var truePartRange, falsePartRange PartRange

		if node.operator == LessThan {
			truePartRange.min = min
			truePartRange.max = minInt(node.operand-1, max)
			falsePartRange.min = maxInt(node.operand, min)
			falsePartRange.max = max
		} else if node.operator == GreaterThan {
			truePartRange.min = maxInt(node.operand+1, min)
			truePartRange.max = max
			falsePartRange.min = min
			falsePartRange.max = minInt(node.operand, max)
		} else {
			panic("unknown operator: " + node.operator)
		}

		if truePartRange.min <= truePartRange.max {
			copy := make(map[string]PartRange)
			for k, v := range *partRanges {
				copy[k] = v
			}
			copy[node.operandName] = truePartRange
			combos += getCombinations(&copy, node.next, workflows)
		}

		if falsePartRange.min <= falsePartRange.max {
			(*partRanges)[node.operandName] = falsePartRange
		} else {
			break
		}
	}

	combos += getCombinations(partRanges, (*(*workflow).nodes)[len((*(*workflow).nodes))-1].next, workflows)
	return combos
}

func (workflow *Workflow) evaluate(part *Part, workflows *map[string]*Workflow) string {
	for _, node := range *workflow.nodes {

		var operand1 int64 = -1
		switch node.operandName {
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

		if node.function(operand1, node.operand) {
			if nextNode, present := (*workflows)[node.next]; present {
				return nextNode.evaluate(part, workflows)
			} else {
				return node.next
			}
		}
	}
	panic("should not be here")
}

type Workflow struct {
	nodes *[]Node
}

func part1(parts *[]Part, workflows *map[string]*Workflow) int64 {
	var answer int64 = 0
	inWorkFlow, _ := (*workflows)["in"]

	for _, part := range *parts {
		if result := inWorkFlow.evaluate(&part, workflows); result == "A" {
			answer += part.x + part.m + part.a + part.s
		}
	}
	return answer
}

func part2(workflows *map[string]*Workflow) int64 {

	partRanges := map[string]PartRange{"a": {min: 1, max: 4000}, "x": {min: 1, max: 4000}, "m": {min: 1, max: 4000}, "s": {min: 1, max: 4000}}
	answer := getCombinations(&partRanges, "in", workflows)

	return answer
}

func main() {
	parts, workflows := parseFile("input.txt")

	fmt.Println("answer part 1:", part1(parts, workflows))

	fmt.Println("answer part 2:", part2(workflows))
}

func parseWorkFlow(workflowString string) *Workflow {
	workflow := Workflow{nodes: &[]Node{}}
	nodestrings := strings.Split(workflowString, ",")
	for _, nodestring := range nodestrings {
		if indexOfColon := strings.Index(nodestring, ":"); indexOfColon > 0 {
			node := Node{}
			node.operandName = nodestring[0:1]
			node.operator = nodestring[1:2]
			node.operand, _ = strconv.ParseInt(nodestring[2:indexOfColon], 10, 64)
			node.function = getFunction(nodestring[1:2])
			node.next = nodestring[indexOfColon+1:]
			*workflow.nodes = append(*workflow.nodes, node)
		} else {
			node := Node{}
			node.operandName = Result
			node.operand = 0
			node.function = nop
			node.next = nodestring
			*workflow.nodes = append(*workflow.nodes, node)
		}
	}
	return &workflow
}

func parseFile(fileName string) (*[]Part, *map[string]*Workflow) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	workFlows := make(map[string]*Workflow)
	parts := []Part{}

	checkWorkflows := true
	for scanner.Scan() {
		line := scanner.Text()
		if checkWorkflows && line == "" {
			checkWorkflows = false
			continue
		}
		if checkWorkflows {
			workflowKey := line[0:strings.Index(line, "{")]

			workflow := parseWorkFlow(line[strings.Index(line, "{")+1 : strings.Index(line, "}")])

			workFlows[workflowKey] = workflow

		} else {
			part := Part{}
			partStrings := strings.Split(line[1:len(line)-1], ",")
			for _, partString := range partStrings {
				switch partString[0] {
				case 'x':
					part.x, _ = strconv.ParseInt(partString[2:], 10, 64)
					break
				case 'm':
					part.m, _ = strconv.ParseInt(partString[2:], 10, 64)
					break
				case 'a':
					part.a, _ = strconv.ParseInt(partString[2:], 10, 64)
					break
				case 's':
					part.s, _ = strconv.ParseInt(partString[2:], 10, 64)
					break
				default:
					panic("unknown part type: " + partString[0:1])
				}
			}
			parts = append(parts, part)
		}
	}

	return &parts, &workFlows
}
