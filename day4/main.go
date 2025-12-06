package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("./" +
		"input.txt")
	defer func() {
		file.Close()
	}()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	totalFoundObjects := 0
	newFoundObjects := true
	for newFoundObjects {
		var tempFoundObject int
		var newLines []string
		for i, line := range lines {
			temp := 0
			var newLine string

			if i == 0 {
				temp, newLine = processVectorPosition("", lines[i+1], line)
			} else if i == len(lines)-1 {
				temp, newLine = processVectorPosition(lines[i-1], "", line)
			} else {
				temp, newLine = processVectorPosition(lines[i-1], lines[i+1], line)
			}
			tempFoundObject += temp
			newLines = append(newLines, newLine)
		}

		if tempFoundObject == 0 {
			newFoundObjects = false
		}

		lines = newLines
		totalFoundObjects += tempFoundObject
	}

	fmt.Println(totalFoundObjects)
}

func processVectorPosition(above, below, line string) (int, string) {
	xpos := 0
	l := len(line)
	accessOnLine := 0
	var positionsToRemove []int
	for i, char := range line {

		if string(char) != "@" {
			continue
		}
		xpos = i
		finder := 0
		if above != "" {
			if xpos != 0 {
				if string(above[xpos-1]) == "@" {
					finder++
				}
			}

			if xpos != l-1 {
				if string(above[xpos+1]) == "@" {
					finder++
				}
			}

			if string(above[xpos]) == "@" {
				finder++
			}
		}

		if below != "" {
			if xpos != 0 {
				if string(below[xpos-1]) == "@" {
					finder++
				}
			}

			if xpos != (l - 1) {
				if string(below[xpos+1]) == "@" {
					finder++
				}
			}

			if string(below[xpos]) == "@" {
				finder++
			}
		}

		if xpos != 0 {
			if string(line[xpos-1]) == "@" {
				finder++
			}
		}

		if xpos != l-1 {
			if string(line[xpos+1]) == "@" {
				finder++
			}
		}

		if finder < 4 {
			accessOnLine++
			positionsToRemove = append(positionsToRemove, i)
		}
	}

	for _, position := range positionsToRemove {
		lineBytes := []byte(line)
		lineBytes[position] = '.'
		line = string(lineBytes)
	}

	return accessOnLine, line
}
