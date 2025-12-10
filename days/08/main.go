package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/dsaunders88/advent-of-code-2025/utils/data"
)

const INPUT_FILE string = "days/08/input.txt"

type Vert struct {
	X int
	Y int
	Z int
}

type Edge struct {
	Vert1  int // index of verts in map
	Vert2  int
	Length int
}

func main() {
	fmt.Println("  ˳")
	fmt.Println("  ** ---------------------")
	fmt.Println(" / °\\ Day 08 Solutions ~~")
	fmt.Println("/° ・\\ -------------------")

	var total1, total2 int

	input := data.ReadAsLines(INPUT_FILE)

	junctions := make(map[int]Vert)
	connections := 0
	limit := 1000

	// create vertices from input
	for i, row := range input {
		var v Vert
		for j, part := range strings.Split(row, ",") {
			l, _ := strconv.Atoi(part)
			switch j {
			case 0:
				v.X = l
			case 1:
				v.Y = l
			default:
				v.Z = l
			}
		}
		junctions[i] = v
	}

	// find all edges between vertices
	var edges []Edge
	for i, v1 := range junctions {
		for j, v2 := range junctions {
			if i != j {
				d := dist(v2, v1)
				edges = append(edges, Edge{i, j, d})
			}
		}
	}

	// sort edges by length small->large
	slices.SortFunc(edges, func(a, b Edge) int {
		return cmp.Compare(a.Length, b.Length)
	})

	// part 2: all circuits
	var circuits [][]int
	// part 1: circuits up to limit
	var circuitsPart1 [][]int
	var lastConnection Edge

	// only loop over every other until `limit` connections found
	// this is because edges list has both directions (duplicate values)
	for i := 0; i < len(edges)-1; i += 2 {
		edge := edges[i]

		if i == 0 {
			circuits = append(circuits, []int{edge.Vert1, edge.Vert2})
		} else {
			// what if 1 & 2 are in 2 different sets?
			// track both set indexes
			var v1Exist, v2Exist bool
			var set1Idx, set2Idx int
			for j, set := range circuits {
				if v1Exist == false {
					if slices.Contains(set, edge.Vert1) {
						v1Exist = true
						set1Idx = j
					}
				}
				if v2Exist == false {
					if slices.Contains(set, edge.Vert2) {
						v2Exist = true
						set2Idx = j
					}
				}
			}

			switch {
			case v1Exist && v2Exist:
				// connections exist in different circuits, need to merge
				if set1Idx != set2Idx {
					toMerge := circuits[set2Idx]
					circuits[set1Idx] = append(circuits[set1Idx], toMerge...)
					// delete toMerge
					circuits = append(circuits[:set2Idx], circuits[set2Idx+1:]...)
				}
			case v1Exist:
				circuits[set1Idx] = append(circuits[set1Idx], edge.Vert2)
			case v2Exist:
				circuits[set2Idx] = append(circuits[set2Idx], edge.Vert1)
			default:
				circuits = append(circuits, []int{edge.Vert1, edge.Vert2})
			}
		}

		lastConnection.Vert1 = edge.Vert1
		lastConnection.Vert2 = edge.Vert2

		// nb: every attempt is a connection for part 1 rules, even if it does nothing
		connections++
		if connections == limit {
			circuitsPart1 = append(circuitsPart1, circuits...)
		}
		// part 2: break when circuits has one row with all original vertex indices
		if i != 0 && len(circuits[0]) == len(input) {
			break
		}
	}

	// sort circuits for part 1
	slices.SortFunc(circuitsPart1, func(a, b []int) int {
		return cmp.Compare(len(b), len(a))
	})

	// total part 1
	for i := range 3 {
		circ := circuitsPart1[i]
		if i == 0 {
			total1 = 1
		}

		total1 *= len(circ)
	}

	// total part 2
	last1, ok1 := junctions[lastConnection.Vert1]
	last2, ok2 := junctions[lastConnection.Vert2]

	if ok1 && ok2 {
		total2 = last1.X * last2.X
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}

// distance between vertices in 3d space
func dist(v1, v2 Vert) int {
	x := v2.X - v1.X
	y := v2.Y - v1.Y
	z := v2.Z - v1.Z

	// don't actually need sq. root
	return (x * x) + (y * y) + (z * z)
}
