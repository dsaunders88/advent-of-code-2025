package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/dsaunders88/advent-of-code-2025/utils/data"
)

const INPUT_FILE string = "days/09/input.txt"

type Coord struct {
	X float64
	Y float64
}

type Line struct {
	p1, p2  Coord
	xLength float64 // length between x vals of two connecting points
}

type Polygon []Coord

func main() {
	start := time.Now()
	fmt.Println("  ˳")
	fmt.Println("  ** ---------------------")
	fmt.Println(" / °\\ Day 09 Solutions ~~")
	fmt.Println("/° ・\\ -------------------")

	var total1, total2 int

	// hint: coords are already in "clockwise order"
	input := data.ReadAsLines(INPUT_FILE)

	var coords Polygon
	coordLines := []Line{}

	for i, s := range input {
		parts := strings.Split(s, ",")
		x, _ := strconv.ParseFloat(parts[0], 64)
		y, _ := strconv.ParseFloat(parts[1], 64)
		c := Coord{x, y}
		coords = append(coords, c)

		// find the x lengths between connecting vertices
		var xLength float64
		var l Line
		if i != 0 {
			prev := coords[i-1]
			xLength = math.Abs(prev.X - c.X)
			l = Line{p1: prev, p2: c}
		}
		if i == len(input)-1 {
			first := coords[0]
			xLength = math.Abs(c.X - first.X)
			l = Line{p1: c, p2: first}
		}
		l.xLength = xLength
		coordLines = append(coordLines, l)
	}

	// sort longest x lengths
	slices.SortFunc(coordLines, func(a, b Line) int {
		return cmp.Compare(b.xLength, a.xLength)
	})

	for _, c1 := range coords {
		for _, c2 := range coords {
			// part 1
			area := int(findArea(c1, c2))
			total1 = max(total1, area)

			// part 2: use visual of svg image as guide to find massive gap - longest x lengths
			// check above or below the Y of two largest x lines
			if (c1.Y <= coordLines[0].p1.Y && c2.Y <= coordLines[0].p1.Y) || (c1.Y >= coordLines[1].p1.Y && c2.Y >= coordLines[1].p1.Y) {
				// test if values are in polygon - reduce rect by .5 so points don't fall on edges
				var testC1, testC2 Coord
				if c1.X > c2.X {
					testC1.X = c1.X - 0.5
					testC2.X = c2.X + 0.5
				} else {
					testC1.X = c1.X + 0.5
					testC2.X = c2.X - 0.5
				}
				if c1.Y > c2.Y {
					testC1.Y = c1.Y - 0.5
					testC2.Y = c2.Y + 0.5
				} else {
					testC1.Y = c1.Y + 0.5
					testC2.Y = c2.Y - 0.5
				}
				testC3 := Coord{testC1.X, testC2.Y}
				testC4 := Coord{testC2.X, testC1.Y}

				// ray cast test to see if c3 & c4 are in polygon
				if coords.pointIn(testC3) && coords.pointIn(testC4) {
					area2 := int(findArea(c1, c2))
					total2 = max(total2, area2)
				}
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
	fmt.Println(elapsed)
}

func findArea(c1, c2 Coord) float64 {
	var xLen, yLen float64
	xLen = math.Abs(c2.X - c1.X)
	yLen = math.Abs(c2.Y - c1.Y)
	return (xLen + 1) * (yLen + 1)
}

// ray casting point-in-polygon algorithm from the internet
func (p Polygon) pointIn(point Coord) bool {
	intersections := 0
	numPoints := len(p)

	for i := range numPoints {
		p1 := p[i]
		p2 := p[(i+1)%numPoints]

		if (p1.Y <= point.Y && point.Y < p2.Y) || (p2.Y <= point.Y && point.Y < p1.Y) {
			if point.X < (p2.X-p1.X)*(point.Y-p1.Y)/(p2.Y-p1.Y)+p1.X {
				intersections++
			}
		}
	}

	return intersections%2 == 1
}
