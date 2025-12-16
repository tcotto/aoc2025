package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Nodes map[string][]string

var memoCache map[string]int

func main() {
	memoCache = make(map[string]int)
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()
	nodes := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		nodeName := strings.TrimSpace(parts[0])
		connectionsStr := strings.TrimSpace(parts[1])

		// Split connections by spaces
		connections := strings.Fields(connectionsStr)
		nodes[nodeName] = connections
	}

	nodes["out"] = []string{}

	fmt.Println(part1("you", nodes))
	fmt.Println(part2("svr", nodes, false, false))

}

func part1(node string, nodes Nodes) int {
	if node == "out" {
		return 1
	} else {
		sum := 0
		for _, x := range nodes[node] {
			sum += part1(x, nodes)
		}

		return sum
	}
}

func part2(node string, nodes Nodes, vistedDac, visitedFft bool) int {
	cacheKey := fmt.Sprintf("%s%t%t", node, vistedDac, visitedFft)
	_, ok := memoCache[cacheKey]
	if ok {
		return memoCache[cacheKey]
	}
	if node == "out" {
		if vistedDac && visitedFft {
			return 1
		}
		return 0
	} else {
		sum := 0
		for _, x := range nodes[node] {
			if x == "fft" {
				tmpSeenFft = true
			}

			if x == "dac" {
				tmpSeenDac = true
			}

			sum += part2(x, nodes, tmpSeenDac, tmpSeenFft)
		}
		memoCache[cacheKey] = sum
		return sum
	}
}
