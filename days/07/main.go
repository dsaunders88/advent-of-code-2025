package main

import (
	"fmt"
	"regexp"

	"github.com/dsaunders88/advent-of-code-2025/utils/data"
)

const INPUT_FILE string = "days/07/input.txt"

func main() {
	fmt.Println("  ˳")
	fmt.Println("  ** ---------------------")
	fmt.Println(" / °\\ Day 07 Solutions ~~")
	fmt.Println("/° ・\\ -------------------")

	var total1, total2 int

	input := data.ReadAsLines(INPUT_FILE)

	re := regexp.MustCompile(`S`)
	start := re.FindStringIndex(input[0])
	// initialize with start S
	// keys = indexes of current beams, updates as loop through rows
	currentBeams := map[int]bool{start[0]: true}
	// part 2: indexes of total beams current paths
	// key: idx, val: number of branches / beam overlaps
	totalBeams := map[int]int{start[0]: 1}

	// only loop over the rows that will have splitters
	// assume one row of blank space between
	for i := 2; i < len(input)-1; i += 2 {
		row := input[i]
		re := regexp.MustCompile(`\^`)
		splits := re.FindAllStringIndex(row, -1)

		for _, sp := range splits {
			// part 1
			// if splitter occupies pos of a current beam, split and remove idx
			if _, ok := currentBeams[sp[0]]; ok {
				total1++
				delete(currentBeams, sp[0])
				// add new current beams to left and right of the current splitter
				currentBeams[sp[0]-1] = true
				currentBeams[sp[1]] = true
			}

			// part 2
			// same as 1, but carry over overlap values on indexes
			if existing, ok := totalBeams[sp[0]]; ok {
				delete(totalBeams, sp[0])
				// if this would overwrite, create additional branch
				if v, ok := totalBeams[sp[0]-1]; ok {
					totalBeams[sp[0]-1] = existing + v
				} else {
					totalBeams[sp[0]-1] = existing
				}
				if v, ok := totalBeams[sp[1]]; ok {
					totalBeams[sp[1]] = existing + v
				} else {
					totalBeams[sp[1]] = existing
				}
			}
		}
	}

	for _, b := range totalBeams {
		total2 += b
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}
