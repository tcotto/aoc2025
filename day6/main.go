package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	problems := make(map[int][]int)
	newProblems := make(map[int][]string)
	//var newProblemsOperands []string
	total := 0
	total2 := 0

	allBreakIndex := make(map[int]bool)
	var operandList []string

	// get the last line which is the operands
	for i, operand := range lines[len(lines)-1] {
		op := string(operand)
		if i == 0 {
			operandList = append(operandList, op)
			continue
		}

		if op == "*" || op == "/" || op == "+" || op == "-" {
			allBreakIndex[i-1] = true
			operandList = append(operandList, op)
		}
	}

	for i, line := range lines {
		if i == len(lines)-1 {
			break
		}
		temp := ""
		currentindex := 0

		//if i == len(lines)-1 {
		//	for j, char := range line {
		//		if string(char) == " " {
		//			if temp != "" {
		//				total += processProblem(problems[currentindex], temp)
		//				newProblemsOperands = append(newProblemsOperands, temp)
		//				currentindex++
		//			}
		//
		//			temp = ""
		//			continue
		//		}
		//		if j == len(line)-1 && temp == "" {
		//			newProblemsOperands = append(newProblemsOperands, string(char))
		//			total += processProblem(problems[currentindex], string(char))
		//		}
		//		temp = string(char)
		//	}
		//	break
		//}

		for j, char := range line {
			_, ok := allBreakIndex[j]
			if ok {
				_, ok = newProblems[currentindex]
				if ok {
					newProblems[currentindex] = append(newProblems[currentindex], temp)
				} else {
					newProblems[currentindex] = []string{temp}
				}
				temp = ""
				currentindex++
				continue
			}
			temp = temp + string(char)

			if j == len(line)-1 {
				_, ok = newProblems[currentindex]
				if ok {
					newProblems[currentindex] = append(newProblems[currentindex], temp)
				} else {
					newProblems[currentindex] = []string{temp}
				}
			}

			//if j == len(line)-1 {
			//	if temp != "" {
			//		num, _ := strconv.Atoi(temp + string(char))
			//		_, ok := problems[currentindex]
			//		if ok {
			//			problems[currentindex] = append(problems[currentindex], num)
			//			newProblems[currentindex] = append(newProblems[currentindex], temp+string(char))
			//		} else {
			//			problems[currentindex] = []int{num}
			//			newProblems[currentindex] = []string{temp + string(char)}
			//		}
			//	}
			//
			//	break
			//}

			//temp = temp + string(char)
		}
	}

	mapLen := len(newProblems)
	for i := mapLen - 1; i >= 0; i-- {
		total2 += processReverseProblem(newProblems[i], operandList[i])
	}

	fmt.Println(problems)
	fmt.Println(newProblems)
	fmt.Println(total)
	fmt.Println(total2)
}

func processProblem(numbers []int, operand string) int {
	finalSum := numbers[0]

	for i := 1; i <= len(numbers)-1; i++ {
		switch operand {
		case "+":
			finalSum += numbers[i]
		case "-":
			finalSum -= numbers[i]
		case "*":
			finalSum = finalSum * numbers[i]
		case "/":
			finalSum = finalSum / numbers[i]
		}
	}

	return finalSum
}

func processReverseProblem(numbers []string, operand string) int {
	x := make(map[int]string)
	maxLength := 0
	total := 0

	for _, number := range numbers {
		for i, char := range number {
			_, ok := x[i]
			if ok {
				x[i] = x[i] + string(char)
			} else {
				x[i] = string(char)
			}

			if i > maxLength {
				maxLength = i
			}
		}
	}

	tmp, _ := strconv.Atoi(strings.TrimSpace(x[maxLength]))
	total = tmp
	for i := maxLength - 1; i >= 0; i-- {
		if string(x[i][0]) == " " {
			x[i] = x[i][1:]
		}
		tmp2, _ := strconv.Atoi(strings.TrimSpace(x[i]))
		switch operand {
		case "+":
			total += tmp2
		case "-":
			total -= tmp2
		case "*":
			total = total * tmp2
		case "/":
			total = total / tmp2
		}
	}

	return total
}
