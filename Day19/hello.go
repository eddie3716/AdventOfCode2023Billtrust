package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type PartRange struct {
	xMin int
	xMax int
	mMin int
	mMax int
	aMin int
	aMax int
	sMin int
	sMax int
}

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
	GreaterThan = ">"
	LessThan    = "<"
	Result      = ""
)

type Node struct {
	function    func(a, b int) bool
	operator    string
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

func (workflow *Workflow) getRanges(partRange *PartRange, workflows *orderedmap.OrderedMap[string, *Workflow]) *[]PartRange {

	partRanges := []PartRange{}
	operandType := ""
	for _, node := range *workflow.nodes {
		operand1Min := -1
		operand1Max := -1
		switch node.operandType {
		case "a":
			operandType = "a"
			operand1Min = partRange.aMin
			operand1Max = partRange.aMax
			break
		case "x":
			operandType = "x"
			operand1Min = partRange.xMin
			operand1Max = partRange.xMax
			break
		case "m":
			operandType = "m"
			operand1Min = partRange.mMin
			operand1Max = partRange.mMax
			break
		case "s":
			operandType = "s"
			operand1Min = partRange.sMin
			operand1Max = partRange.sMax
		}

		minFits := node.function(operand1Min, node.operand2)
		maxFits := node.function(operand1Max, node.operand2)

		kindOfFits := true
		switch node.operator {
		case GreaterThan:
			if !maxFits {
				kindOfFits = false
				break
			}
			if !minFits {
				operand1Min = node.operand2 + 1
			}
			break
		case LessThan:
			if !maxFits {
				operand1Max = node.operand2 - 1
			}
			if !minFits {
				kindOfFits = false
			}
		}

		if !kindOfFits {
			continue
		}

		partRangeToSend := &PartRange{}
		partRangeToSend.aMax = partRange.aMax
		partRangeToSend.aMin = partRange.aMin
		partRangeToSend.xMax = partRange.xMax
		partRangeToSend.xMin = partRange.xMin
		partRangeToSend.mMax = partRange.mMax
		partRangeToSend.mMin = partRange.mMin
		partRangeToSend.sMax = partRange.sMax
		partRangeToSend.sMin = partRange.sMin

		if operandType != "" {
			switch operandType {
			case "a":
				partRangeToSend.aMax = operand1Max
				partRangeToSend.aMin = operand1Min
				break
			case "x":
				partRangeToSend.xMax = operand1Max
				partRangeToSend.xMin = operand1Min
				break
			case "m":
				partRangeToSend.mMax = operand1Max
				partRangeToSend.mMin = operand1Min
				break
			case "s":
				partRangeToSend.sMax = operand1Max
				partRangeToSend.sMin = operand1Min
			}
		}

		if nextNode, present := workflows.Get(node.next); present {
			partRanges = append(partRanges, *nextNode.getRanges(partRangeToSend, workflows)...)
		} else if node.next == "A" {
			partRanges = append(partRanges, *partRangeToSend)
		}
	}
	return &partRanges
}

func (workflow *Workflow) evaluate(part *Part, workflows *orderedmap.OrderedMap[string, *Workflow]) string {
	for _, node := range *workflow.nodes {

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
	panic("should not be here")
}

type Workflow struct {
	nodes *[]Node
}

func part1(parts *[]Part, workflows *orderedmap.OrderedMap[string, *Workflow]) int {
	answer := 0
	inWorkFlow, _ := workflows.Get("in")

	for _, part := range *parts {
		if result := inWorkFlow.evaluate(&part, workflows); result == "A" {
			answer += part.x + part.m + part.a + part.s
		}
	}
	return answer
}

func part2(workflows *orderedmap.OrderedMap[string, *Workflow]) int {
	answer := 0
	inWorkFlow, _ := workflows.Get("in")

	partRange := PartRange{xMin: 1, xMax: 4000, mMin: 1, mMax: 4000, aMin: 1, aMax: 4000, sMin: 1, sMax: 4000}

	partRanges := inWorkFlow.getRanges(&partRange, workflows)

	for _, partRange := range *partRanges {
		fmt.Println(partRange)
		answer += (partRange.xMax - partRange.xMin) * (partRange.mMax - partRange.mMin) * (partRange.aMax - partRange.aMin) * (partRange.sMax - partRange.sMin)
	}

	return answer
}

func main() {
	parts, workflows := parseFile("testinput.txt")

	fmt.Println("answer part 1:", part1(parts, workflows))

	fmt.Println("answer part 2:", part2(workflows))
}

func parseWorkFlow(workflowString string) *Workflow {
	workflow := Workflow{nodes: &[]Node{}}
	nodestrings := strings.Split(workflowString, ",")
	for _, nodestring := range nodestrings {
		if indexOfColon := strings.Index(nodestring, ":"); indexOfColon > 0 {
			node := Node{}
			node.operandType = nodestring[0:1]
			node.operator = nodestring[1:2]
			node.operand2, _ = strconv.Atoi(nodestring[2:indexOfColon])
			node.function = getFunction(nodestring[1:2])
			node.next = nodestring[indexOfColon+1:]
			*workflow.nodes = append(*workflow.nodes, node)
		} else {
			node := Node{}
			node.operandType = Result
			node.operand2 = 0
			node.function = nop
			node.next = nodestring
			*workflow.nodes = append(*workflow.nodes, node)
		}
	}
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
			continue
		}
		if checkWorkflows {
			workflowKey := line[0:strings.Index(line, "{")]

			workflow := parseWorkFlow(line[strings.Index(line, "{")+1 : strings.Index(line, "}")])

			workFlows.Set(workflowKey, workflow)

		} else {
			part := Part{}
			partStrings := strings.Split(line[1:len(line)-1], ",")
			for _, partString := range partStrings {
				switch partString[0] {
				case 'x':
					part.x, _ = strconv.Atoi(partString[2:])
					break
				case 'm':
					part.m, _ = strconv.Atoi(partString[2:])
					break
				case 'a':
					part.a, _ = strconv.Atoi(partString[2:])
					break
				case 's':
					part.s, _ = strconv.Atoi(partString[2:])
					break
				default:
					panic("unknown part type: " + partString[0:1])
				}
			}
			parts = append(parts, part)
		}
	}

	return &parts, workFlows
}
