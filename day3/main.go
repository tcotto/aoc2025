package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer func() {
		file.Close()
	}()

	scanner := bufio.NewScanner(file)

	totalPower := 0
	otherTotalPower := 0

	for scanner.Scan() {
		line := scanner.Text()
		maxSum := 0

		maxSum += findMax(line)
		totalPower += maxSum

		tempLine := line
		for len(tempLine) > 12 {
			tempLine = findMin(tempLine)
		}
		tempTotal, _ := strconv.Atoi(tempLine)
		otherTotalPower += tempTotal
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(totalPower)
	fmt.Println(otherTotalPower)
}

func findMax(battery string) int {
	firstMax, _ := strconv.Atoi(string(battery[0]))
	secondMax := 0

	for i := 1; i <= len(battery)-1; i++ {
		temp, _ := strconv.Atoi(string(battery[i]))
		if i == 0 {
			continue
		}

		if i == len(battery)-1 {
			if temp > secondMax {
				secondMax = temp
				break
			}
		}

		if i == len(battery)-2 {
			if temp > firstMax {
				firstMax = temp
				secondMax, _ = strconv.Atoi(string(battery[i+1]))
				break
			}
		}

		if temp > firstMax {
			firstMax = temp
			secondMax = 0
			continue
		}

		if temp > secondMax {
			secondMax = temp
			continue
		}
	}

	return (10 * firstMax) + secondMax
}

func findMin(battery string) string {
	//minValue, _ := strconv.Atoi(string(battery[0]))
	minValueIndex := 0

	for i := 0; i <= len(battery)-1; i++ {
		if i == len(battery)-1 {
			return battery[:i]
		}

		temp, _ := strconv.Atoi(string(battery[i]))
		tempNext, _ := strconv.Atoi(string(battery[i+1]))
		if temp < tempNext {
			//minValue = temp
			minValueIndex = i
			break
		}
	}

	return battery[:minValueIndex] + battery[minValueIndex+1:]
}
