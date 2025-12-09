package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type JBox struct {
	x, y, z float64
}

func (box JBox) String() string {
	return fmt.Sprintf("%g,%g,%g", box.x, box.y, box.z)
}

func (box JBox) computeDistance(otherBox JBox) float64 {
	result := math.Sqrt(math.Pow(box.x-otherBox.x, 2) + math.Pow(box.y-otherBox.y, 2) + math.Pow(box.z-otherBox.z, 2))
	if result < 0 {
		result = result * -1
	}

	return result
}

type Graph struct {
	nodes    []JBox
	vertices []Vertex
}

func (g *Graph) computeVertices() {
	g.vertices = []Vertex{}
	seen := make(map[string]bool)
	for i, node := range g.nodes {
		for j := 0; j < len(g.nodes)-1; j++ {
			if i == j {
				continue
			}

			seenKey := node.String() + "-" + g.nodes[j].String()
			if seen[seenKey] {
				continue
			}

			weight := node.computeDistance(g.nodes[j])
			g.vertices = append(g.vertices, Vertex{node, g.nodes[j], weight})
			seenKey = g.nodes[j].String() + "-" + node.String()
			seen[seenKey] = true
		}
	}

	slices.SortFunc(g.vertices, func(a, b Vertex) int {
		if a.weight > b.weight {
			return 1
		} else if a.weight < b.weight {
			return -1
		}

		return 0
	})
}

type Vertex struct {
	nodeA, nodeB JBox
	weight       float64
}

type Circuit []map[string]bool

func (c Circuit) doesExist(box JBox) (int, bool) {
	for i, val := range c {
		_, ok := val[box.String()]
		if ok {
			return i, true
		}
	}
	return 0, false
}

func main() {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	var nodes []JBox
	for scanner.Scan() {
		line := scanner.Text()
		splitLines := strings.Split(line, ",")

		x, _ := strconv.ParseFloat(splitLines[0], 64)
		y, _ := strconv.ParseFloat(splitLines[1], 64)
		z, _ := strconv.ParseFloat(splitLines[2], 64)

		nodes = append(nodes, JBox{x, y, z})
	}

	jboxGraph := Graph{nodes, []Vertex{}}
	jboxGraph.computeVertices()

	circuit := Circuit{}
	foundCircuits := 0
	var circuitClosed Vertex

	for _, vertex := range jboxGraph.vertices {
		a, ok := circuit.doesExist(vertex.nodeA)
		b, ok2 := circuit.doesExist(vertex.nodeB)

		if len(circuit) != 0 && len(circuit[0]) == len(nodes) {
			break
		}

		if ok && ok2 {
			if a != b {
				for key, value := range circuit[b] {
					circuit[a][key] = value
				}

				// Remove circuit[b] after merging
				circuit = append(circuit[:b], circuit[b+1:]...)
				foundCircuits++
				circuitClosed = vertex
				continue
			}
			continue
		}

		if ok {
			circuit[a][vertex.nodeB.String()] = true
			foundCircuits++
			circuitClosed = vertex
			continue
		}

		if ok2 {
			circuit[b][vertex.nodeA.String()] = true
			foundCircuits++
			circuitClosed = vertex
			continue
		}

		circuit = append(circuit, map[string]bool{
			vertex.nodeA.String(): true,
			vertex.nodeB.String(): true,
		})
		circuitClosed = vertex
		foundCircuits++

	}

	slices.SortFunc(circuit, func(a, b map[string]bool) int {
		return len(b) - len(a)
	})

	fmt.Println(int(circuitClosed.nodeA.x) * int(circuitClosed.nodeB.x))
}
