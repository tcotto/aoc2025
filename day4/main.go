package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer func() {
		file.Close()
	}()

	scanner := bufio.NewScanner(file)
	var lines []string
	foundObjects := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, line := range lines {
		if i == 0 {
			foundObjects += processVectorPosition("", lines[i+1], line)
			continue
		}

		if i == len(lines)-1 {
			foundObjects += processVectorPosition(lines[i-1], "", line)
			break
		}

		foundObjects += processVectorPosition(lines[i-1], lines[i+1], line)
	}

	if foundObjects != 1367 {
		fmt.Errorf("WRONG ANSWER!")
		os.Exit(1)
	}

	fmt.Println(foundObjects)
}

func processVectorPosition(above, below, line string) int {
	xpos := 0
	l := len(line)
	accessOnLine := 0
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
		}
	}

	return accessOnLine
}
