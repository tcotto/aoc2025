package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	var linkedList []*DoubleLLNode
	firstNode := DoubleLLNode{0, nil, nil}
	lastNode := DoubleLLNode{99, nil, nil}
	lastNode.next = &firstNode
	firstNode.previous = &lastNode

	linkedList = append(linkedList, &firstNode)
	for i := 1; i <= 98; i++ {
		if i == 1 {
			node := DoubleLLNode{i, &firstNode, nil}
			firstNode.next = &node
			linkedList = append(linkedList, &node)
		} else if i == 98 {
			node := DoubleLLNode{i, linkedList[i-1], &lastNode}
			lastNode.previous = &node
			linkedList[i-1].next = &node
			linkedList = append(linkedList, &node)
		} else {
			node := DoubleLLNode{i, linkedList[i-1], nil}
			linkedList[i-1].next = &node
			linkedList = append(linkedList, &node)
		}
	}
	linkedList[98].next = &lastNode
	linkedList = append(linkedList, &lastNode)

	dialNode := linkedList[50]
	zeros := 0

	for _, i := range input {
		dialNode = turnNodes(i.turnDirection, i.noOfTurns, dialNode)
		if dialNode.value == 0 {
			zeros = zeros + 1
		}
	}

	fmt.Println(zeros)

}

func turnNodes(direction string, numberOfTurns int, node *DoubleLLNode) *DoubleLLNode {
	if numberOfTurns != 0 {
		if direction == "R" {
			node = turnNodes("R", numberOfTurns-1, node.next)
		} else {
			node = turnNodes("L", numberOfTurns-1, node.previous)
		}
	}

	return node
}

func readInput() []TextInput {
	file, _ := os.Open("./input.txt")
	defer func() {
		file.Close()
	}()
	var directions []TextInput
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		directionRune := string(line[0])
		turnNumber, _ := strconv.Atoi(line[1:])

		directions = append(directions, TextInput{directionRune, turnNumber})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return directions
}

type TextInput struct {
	turnDirection string
	noOfTurns     int
}

type DoubleLLNode struct {
	value    int
	previous *DoubleLLNode
	next     *DoubleLLNode
}
