package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	bytes, _ := io.ReadAll(file)
	input := string(bytes)

	presents := make(map[int]int)
	inputSplit := strings.Split(strings.TrimSpace(input), "\n\n")
	result := 0

	for i := 0; i <= 5; i++ {
		presentSize := 0
		present := inputSplit[i]
		for _, char := range present {
			if string(char) == "#" {
				presentSize++
			}
		}
		presents[i] = presentSize
	}

	trees := strings.Split(inputSplit[6], "\n")

	fmt.Println(trees)

	for _, tree := range trees {
		treeSplit := strings.Split(tree, ":")
		treeDimensions := strings.Split(treeSplit[0], "x")
		treeLength, _ := strconv.Atoi(treeDimensions[0])
		treeBreadth, _ := strconv.Atoi(treeDimensions[1])
		treeArea := treeLength * treeBreadth

		requiredPresents := strings.Split(strings.Trim(treeSplit[1], " "), " ")
		totalPresentArea := 0

		for i, presentNumberString := range requiredPresents {
			presentNumber, _ := strconv.Atoi(presentNumberString)
			totalPresentArea += presents[i] * presentNumber
		}

		if totalPresentArea <= treeArea {
			result++
		}

		fmt.Println(result)
	}
}
