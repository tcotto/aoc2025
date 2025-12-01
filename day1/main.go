package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputs := readInput()
	numberOfTimesPassedZeros := 0
	startingPoint := 50

	for _, input := range inputs {
		var temp int
		if input.noOfTurns > 100 {
			numberOfTimesPassedZeros = numberOfTimesPassedZeros + (input.noOfTurns / 100)
			input.noOfTurns = input.noOfTurns % 100
		}
		if input.turnDirection == "R" {
			temp = modIt(startingPoint+input.noOfTurns, 100)
			if temp < startingPoint {
				numberOfTimesPassedZeros++
			}
		} else {
			temp = modIt(startingPoint-input.noOfTurns, 100)
			if temp > startingPoint {
				numberOfTimesPassedZeros++
			}
		}

		startingPoint = temp
	}

	fmt.Println(numberOfTimesPassedZeros)
}

func modIt(x, y int) int {
	return ((x % y) + y) % y
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
