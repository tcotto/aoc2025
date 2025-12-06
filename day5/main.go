package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer func() {
		file.Close()
	}()

	var ranges []R
	var values []int
	switchParse := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			switchParse = true
		} else if !switchParse {
			tmpSplit := strings.Split(line, "-")
			low, _ := strconv.Atoi(tmpSplit[0])
			high, _ := strconv.Atoi(tmpSplit[1])
			ranges = append(ranges, R{low, high})
		} else {
			tmp, _ := strconv.Atoi(line)
			values = append(values, tmp)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Sort ranges by lowerBound
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].lowerBound < ranges[j].lowerBound
	})

	noChange := false
	var newRanges []R

	for !noChange {
		newRanges = sortRanges(ranges)

		if len(newRanges) == len(ranges) {
			noChange = true
		} else {
			ranges = newRanges
		}
	}

	freshMap := map[int]bool{}
	ingredientCounter := 0

	for _, ing := range values {
		_, ok := freshMap[ing]
		if !ok {
			for _, r := range newRanges {
				if r.contains(ing) {
					freshMap[ing] = true
					ingredientCounter++
					break
				}
			}
		}
	}

	totalIngredients := 0
	for _, x := range newRanges {
		totalIngredients += x.values()
	}

	fmt.Println(totalIngredients)
	fmt.Println(newRanges)
	fmt.Println(ingredientCounter)
}

type R struct {
	lowerBound, upperBound int
}

func (r R) contains(val int) bool {
	if val >= r.lowerBound && val <= r.upperBound {
		return true
	}

	return false
}

func (r R) values() int {
	return (r.upperBound - r.lowerBound) + 1
}

func sortRanges(ranges []R) []R {
	var newRanges []R
	for i := 0; i <= len(ranges)-1; i++ {
		r := ranges[i]
		if i == len(ranges)-1 {
			newRanges = append(newRanges, r)
			continue
		}

		if r.upperBound >= ranges[i+1].lowerBound {
			if r.upperBound >= ranges[i+1].upperBound {
				newRanges = append(newRanges, R{r.lowerBound, r.upperBound})
			} else {
				newRanges = append(newRanges, R{r.lowerBound, ranges[i+1].upperBound})
			}
			i++
			continue
		}

		//if r.upperBound > ranges[i+1].lowerBound && r.upperBound <= ranges[i+1].upperBound {
		//	if r.lowerBound < ranges[i+1].lowerBound {
		//		newRanges = append(newRanges, R{r.lowerBound, ranges[i+1].upperBound})
		//	} else {
		//		newRanges = append(newRanges, r)
		//	}
		//	i++
		//	continue
		//}
		newRanges = append(newRanges, r)
	}

	return newRanges
}
