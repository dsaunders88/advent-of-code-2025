package main

import (
	"fmt"
	"regexp"

	"github.com/dsaunders88/advent-of-code-2025/utils/data"
)

const INPUT_FILE string = "days/04/input.txt"

func main() {
	fmt.Println("  ˳")
	fmt.Println("  ** ---------------------")
	fmt.Println(" / °\\ Day 04 Solutions ~~")
	fmt.Println("/° ・\\ -------------------")

	var total1, total2 int

	input := data.ReadAsLines(INPUT_FILE)
	numRows := len(input)
	locations := make([][]int, numRows)
	// init with 1 so while loop runs the first time
	totalRemovedPerRound := []int{1}

	// create grid of initial location indexes
	for i, row := range input {
		pattern := regexp.MustCompile(`@`)
		found := pattern.FindAllStringIndex(row, -1)

		// flatten 2-slice arrays of found patterns
		for _, f := range found {
			locations[i] = append(locations[i], f[0])
		}
	}

	// run as long as last round did not remove 0
	for totalRemovedPerRound[len(totalRemovedPerRound)-1] > 0 {
		var totalRemovedThisRound int
		var newLocations [][]int

		// run through known locations and check surrounding coordinates
		for i, row := range locations {
			var removedFromRow int
			var newRow []int
			rowLength := len(row)

			if rowLength > 0 {
				for j, loc := range row {
					var totalAdjacent int
					// check current row
					// right boundary
					if j < rowLength-1 && row[j+1] == loc+1 {
						totalAdjacent++
					}
					// left boundary
					if j > 0 && row[j-1] == loc-1 {
						totalAdjacent++
					}
					// check row above
					if i > 0 {
						// is slices.Contains() or for loop more performant?
						for _, val := range locations[i-1] {
							if val == loc || val == loc-1 || val == loc+1 {
								totalAdjacent++
							}
						}
					}
					// check row below
					if i < numRows-1 {
						for _, val := range locations[i+1] {
							if val == loc || val == loc-1 || val == loc+1 {
								totalAdjacent++
							}
						}
					}
					if totalAdjacent < 4 {
						removedFromRow++
					} else {
						// add to new row with loc removed
						newRow = append(newRow, loc)
					}
				}
			}

			totalRemovedThisRound += removedFromRow
			newLocations = append(newLocations, newRow)
		}
		// replace locations with new locations
		locations = newLocations
		totalRemovedPerRound = append(totalRemovedPerRound, totalRemovedThisRound)
	}

	// skip first position for totals
	total1 = totalRemovedPerRound[1]
	for i, v := range totalRemovedPerRound {
		if i != 0 {
			total2 += v
		}
	}
	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}
