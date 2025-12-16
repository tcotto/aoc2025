package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	desiredState string
	buttons      [][]int
}

func (m *Machine) ComputeState() int {
	currentState := "."
	for i := 1; i < len(m.desiredState); i++ {
		currentState = currentState + "."
	}
	
}

func main() {
	file, _ := os.Open("./smolinput.txt")

	scanner := bufio.NewScanner(file)
	var machines []Machine
	for scanner.Scan() {
		line := scanner.Text()

		machine := Machine{}
		// Extract content between brackets []
		re := regexp.MustCompile(`\[([^\]]*)\]`)
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			content := matches[1]
			machine.desiredState = matches[1] // ".##."
			fmt.Println("Brackets content:", content)
		}

		// Extract all content inside parentheses ()
		reParen := regexp.MustCompile(`\(([^\)]*)\)`)
		parenMatches := reParen.FindAllStringSubmatch(line, -1)
		for _, match := range parenMatches {
			if len(match) > 1 {
				fmt.Println(match[1])
				buttonStrings := strings.Split(match[1], ",")
				var buttonNumbers []int
				for _, bs := range buttonStrings {
					x, _ := strconv.Atoi(bs)
					buttonNumbers = append(buttonNumbers, x)
				}
				machine.buttons = append(machine.buttons, buttonNumbers)
			}
		}

		// Extract all content inside curly braces {}
		reBrace := regexp.MustCompile(`\{([^\}]*)\}`)
		braceMatches := reBrace.FindAllStringSubmatch(line, -1)
		for _, match := range braceMatches {
			if len(match) > 1 {
				fmt.Println("Curly braces content:", match[1])
			}
		}

		machines = append(machines, machine)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(machines)
}
