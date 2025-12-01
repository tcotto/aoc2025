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
	totalZeros := 0

	for _, i := range input {
		zeros := 0
		dialNode, zeros = turnNodes(i.turnDirection, i.noOfTurns, 0, dialNode, true)
		totalZeros = totalZeros + zeros
	}

	fmt.Println(totalZeros)

}

func turnNodes(direction string, numberOfTurns int, numberOfZeros int, node *DoubleLLNode, newPass bool) (*DoubleLLNode, int) {
	newNumberOfZeros := numberOfZeros
	if node.value == 0 && newPass == false {
		newNumberOfZeros = newNumberOfZeros + 1
	}

	if numberOfTurns != 0 {
		if direction == "R" {
			node, newNumberOfZeros = turnNodes("R", numberOfTurns-1, newNumberOfZeros, node.next, false)
		} else {
			node, newNumberOfZeros = turnNodes("L", numberOfTurns-1, newNumberOfZeros, node.previous, false)
		}
	}

	return node, newNumberOfZeros
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
