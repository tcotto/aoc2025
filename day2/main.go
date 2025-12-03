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
	inputString := string(bytes)

	split := strings.Split(inputString, ",")
	counter := 0

	for _, item := range split {
		ids := strings.Split(item, "-")
		firstId := ids[0]
		lastId := ids[1]

		firstIdAsNum, _ := strconv.Atoi(firstId)
		lastIdAsNum, _ := strconv.Atoi(lastId)

		for i := firstIdAsNum; i <= lastIdAsNum; i++ {
			idAsString := strconv.Itoa(i)
			lengthOfid := len(idAsString)
			for j := lengthOfid - 1; j > 0; j-- {
				var foundPattern string
				match := false
				if lengthOfid%j == 0 {
					match = true
					oldIndex := 0
					foundPattern = idAsString[:j]
					for k := j; k <= lengthOfid; k += j {
						if idAsString[oldIndex:k] != foundPattern {
							match = false
							break
						}
						oldIndex = k
					}
				}

				if match == true {
					counter += i
					break
				}
			}
		}

	}

	fmt.Println(counter)
}
