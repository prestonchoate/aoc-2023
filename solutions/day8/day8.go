package day8

import (
	"fmt"
	"strings"
)

type pathNode struct {
	value         string
	leftNodeName  string
	rightNodeName string
}

func createPathNode(val string) *pathNode {
	node := pathNode{
		value: strings.TrimSpace(val),
	}

	return &node
}

type instruction struct {
	steps     []string
	iterator  int
	stepCount int
}

func createInstruction(steps string) *instruction {
	i := instruction{}
	for _, step := range steps {
		i.steps = append(i.steps, string(step))
	}

	return &i
}

func (i *instruction) getNextStep() string {
	if i.iterator == len(i.steps) {
		i.iterator = 0
	}
	step := i.steps[i.iterator]
	i.iterator++
	i.stepCount++
	return step
}

var nodes map[string]pathNode

func SolveP1(input string) int {
	nodes = make(map[string]pathNode)
	parts := strings.Split(input, "\n\n")
	instructions := createInstruction(parts[0])
	nodes := map[string]pathNode{}
	var curNode pathNode

	for _, line := range strings.Split(parts[1], "\n") {
		node := parseNode(line)
		nodes[node.value] = node
	}

	curNode = nodes["AAA"]
	for curNode.value != "ZZZ" {
		if curNode.leftNodeName == curNode.value && curNode.rightNodeName == curNode.value {
			fmt.Printf("Found a cyclical loop in node %v\n", curNode.value)
			return -1
		}

		fmt.Printf("Checking current node: %+v\n", curNode)
		nextStep := instructions.getNextStep()
		fmt.Printf("Going %v\n\n", nextStep)
		if curNode.value == "QNF" {
			fmt.Printf("Got to the right node but we are going %v\n", nextStep)
		}
		if nextStep == "L" {
			//fmt.Printf("Trying to get node %v from map\n", curNode.leftNodeName)
			foundNode, ok := nodes[curNode.leftNodeName]
			if !ok {
				fmt.Println("Something went wrong")
				break
			}
			curNode = foundNode
		} else {
			//fmt.Printf("Trying to get node %v from map\n", curNode.rightNodeName)
			foundNode, ok := nodes[curNode.rightNodeName]
			if !ok {
				fmt.Println("Something went wrong")
				break
			}
			curNode = foundNode
		}
	}
	return instructions.stepCount
}

func SolveP2(input string) int {
	nodes = make(map[string]pathNode)
	parts := strings.Split(input, "\n\n")
	instructions := createInstruction(parts[0])
	nodes := map[string]pathNode{}
	starters := []string{}
	cycleLens := []int{}

	for _, line := range strings.Split(parts[1], "\n") {
		node := parseNode(line)
		nodes[node.value] = node
		if strings.HasSuffix(node.value, "A") {
			starters = append(starters, node.value)
		}
	}

	fmt.Printf("Starter nodes: %+v\n", starters)

	for _, name := range starters {
		curNode := nodes[name]
		cycleLen := []int{}
		instructions.stepCount = 0
		instructions.iterator = 0
		firstZ := ""

		for true {
			for instructions.stepCount == 0 || !strings.HasSuffix(curNode.value, "Z") {
				nextStep := instructions.getNextStep()
				if nextStep == "L" {
					//fmt.Printf("Trying to get node %v from map\n", curNode.leftNodeName)
					foundNode, ok := nodes[curNode.leftNodeName]
					if !ok {
						fmt.Println("Something went wrong")
						break
					}
					curNode = foundNode
				} else {
					//fmt.Printf("Trying to get node %v from map\n", curNode.rightNodeName)
					foundNode, ok := nodes[curNode.rightNodeName]
					if !ok {
						fmt.Println("Something went wrong")
						break
					}
					curNode = foundNode
				}
			}

			cycleLen = append(cycleLen, instructions.stepCount)

			if firstZ == "" {
				firstZ = curNode.value
				instructions.stepCount = 0
			} else if curNode.value == firstZ {
				break
			}
		}

		cycleLens = append(cycleLens, cycleLen...)
	}

	fmt.Printf("%+v\n", cycleLens)

	lcm := cycleLens[0]

	for _, num := range cycleLens {
		lcm = lcm * num / gcd(lcm, num)
	}

	return lcm
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	tmp := a
	a = b
	b = tmp % a
	return gcd(a, b)
}

func parseNode(line string) pathNode {
	pieces := strings.Split(line, "=")
	node := createPathNode(pieces[0])
	pathStrs := strings.Split(pieces[1], ", ")
	for i := 0; i < len(pathStrs); i++ {
		str := strings.ReplaceAll(pathStrs[i], "(", "")
		str = strings.ReplaceAll(str, ")", "")
		str = strings.TrimSpace(str)
		if i == 0 {
			node.leftNodeName = str
		} else {
			node.rightNodeName = str
		}
	}

	return *node
}
