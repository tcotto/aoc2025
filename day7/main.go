package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var seen2 map[string]int

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()
	var input [][]string
	seen := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var lineInput []string
		for _, char := range line {
			tmp := string(char)
			lineInput = append(lineInput, tmp)
		}

		input = append(input, lineInput)
	}
	var result map[string]bool
	seen2 = make(map[string]int)
	resultDay2 := 0
	for i, char := range input[0] {
		if char == "S" {
			result = part1(input, i, 0, seen)
			resultDay2 = part2(input, i, 0)
		}
	}

	fmt.Println(len(result))

	for k, _ := range result {
		x := strings.Split(k, "-")
		rowI, _ := strconv.Atoi(x[0])
		colI, _ := strconv.Atoi(x[1])
		fmt.Printf("Value at row-col %s: %s\n", k, input[rowI][colI])
	}

	fmt.Println(resultDay2)
}

func part1(input [][]string, colI, rowI int, seen map[string]bool) map[string]bool {
	if rowI == len(input)-1 {
		return seen
	}

	var leftSplit, rightSplit map[string]bool
	for i := rowI; i < len(input); i++ {
		if input[i][colI] == "^" {
			key := fmt.Sprintf("%d-%d", i, colI)
			_, ok := seen[key]
			if ok {
				return seen
			}
			leftSplit = part1(input, colI-1, i+1, seen)
			rightSplit = part1(input, colI+1, i+1, seen)
			seen[key] = true
			break
		}
	}

	for k, v := range leftSplit {
		seen[k] = v
	}
	for k, v := range rightSplit {
		seen[k] = v
	}

	return seen
}

func part2(input [][]string, colI, rowI int) int {
	if rowI == len(input)-1 {
		return 1
	}

	var leftSplit, rightSplit int
	for i := rowI; i < len(input); i++ {
		if i == len(input)-1 {
			return 1
		}
		if input[i][colI] == "^" {
			key := fmt.Sprintf("%d-%d", i, colI)
			_, ok := seen2[key]
			if ok {
				return seen2[key]
			}
			leftSplit = part2(input, colI-1, i+1)
			rightSplit = part2(input, colI+1, i+1)
			seen2[key] = leftSplit + rightSplit
			break
		}
	}

	return leftSplit + rightSplit
}
