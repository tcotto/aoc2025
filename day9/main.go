package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Position struct {
	x, y int
}

type GreenTileRange struct {
	a, b   Position
	length int
}

type RedTilePair struct {
	a, b Position
	area int
}

func (p Position) Area(otherPos Position) int {
	length := 0
	length = p.x - otherPos.x
	if length < 0 {
		length *= -1
	}

	breadth := p.y - otherPos.y
	if breadth < 0 {
		breadth *= -1
	}

	return (length + 1) * (breadth + 1)
}

func CreateGreenTileRange(x, y Position) GreenTileRange {
	var gtr GreenTileRange

	// Ensure a has smaller coordinates than b
	if x.x == y.x {
		// Vertical line
		if x.y < y.y {
			gtr = GreenTileRange{x, y, 0}
		} else {
			gtr = GreenTileRange{y, x, 0}
		}
	} else {
		// Horizontal line
		if x.x < y.x {
			gtr = GreenTileRange{x, y, 0}
		} else {
			gtr = GreenTileRange{y, x, 0}
		}
	}

	gtr.length = (gtr.b.x - gtr.a.x) + (gtr.b.y - gtr.a.y)
	return gtr
}

func CreateRedTilePair(x, y Position) RedTilePair {
	// Create a bounding box with a = top-left (min x, min y), b = bottom-right (max x, max y)
	minX := min(x.x, y.x)
	minY := min(x.y, y.y)
	maxX := max(x.x, y.x)
	maxY := max(x.y, y.y)

	rtp := RedTilePair{Position{minX, minY}, Position{maxX, maxY}, 0}
	rtp.area = rtp.a.Area(rtp.b)
	return rtp
}

func (p Position) Compare(otherPos Position) Position {
	if p.y == otherPos.y {
		if p.x < otherPos.x {
			return p
		}

		return otherPos
	}

	if p.x == otherPos.x {
		if p.y < otherPos.x {
			return p
		}

		return otherPos
	}

	if p.x < otherPos.x {
		return p
	}

	return otherPos
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	var nodes []Position
	yIndex := 0
	for scanner.Scan() {
		line := scanner.Text()

		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		nodes = append(nodes, Position{x, y})

		yIndex++
	}

	slices.SortFunc(nodes, func(a, b Position) int {
		if a.y == b.y {
			return a.x - b.x
		}
		return a.y - b.y
	})

	var areas []RedTilePair

	var greenTileRanges []GreenTileRange
	for i := 0; i < len(nodes); i++ {
		nodeCheck := nodes[i]
		for j := i + 1; j < len(nodes); j++ {
			if nodeCheck.x == nodes[j].x || nodeCheck.y == nodes[j].y {
				greenTileRanges = append(greenTileRanges, CreateGreenTileRange(nodeCheck, nodes[j]))
			}

			areas = append(areas, CreateRedTilePair(nodeCheck, nodes[j]))
		}
	}

	slices.SortFunc(areas, func(a, b RedTilePair) int {
		return b.area - a.area
	})
	slices.SortFunc(greenTileRanges, func(a, b GreenTileRange) int {
		return b.length - a.length
	})

	foundArea := 0
	for _, rtp := range areas {
		intersects := false
		for _, gtr := range greenTileRanges {
			if (gtr.a.x < rtp.b.x) && (gtr.a.y < rtp.b.y) && (gtr.b.x > rtp.a.x) && (gtr.b.y > rtp.a.y) {
				intersects = true
				break
			}
		}

		if !intersects {
			foundArea = rtp.area
			break
		}
	}
	fmt.Println(areas)
	fmt.Println(greenTileRanges)
	fmt.Println(foundArea)
}
