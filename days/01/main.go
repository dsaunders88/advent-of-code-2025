package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dsaunders88/advent-of-code-2025/utils/data"
)

const INPUT_FILE string = "days/01/input.txt"

func main() {
	fmt.Println("  ˳")
	fmt.Println("  ** ---------------------")
	fmt.Println(" / °\\ Day 01 Solutions ~~")
	fmt.Println("/° ・\\ -------------------")

	var current, total1, total2 int
	// starting/current location of the dial
	current = 50

	input := data.ReadAsString(INPUT_FILE)
	for row := range strings.SplitSeq(input, "\n") {
		// clicks = total movements left/right
		// distance = total clicks to move (remainder from clicks % 100)
		// sum = use to calculate new current position of dial
		var clicks, distance, sum int
		// revolutions = full rotation (100) back to current
		var revolutions float64
		// for each combination value, split direction from value
		// return positive int for right, negative for left
		val, found := strings.CutPrefix(row, "R")
		if found {
			clicks, _ = strconv.Atoi(val)
		} else {
			clicks, _ = strconv.Atoi(row[1:])
			clicks = clicks * -1
		}

		// calculate total revolutions
		revolutions = math.Floor(float64(math.Abs(float64(clicks))) / 100)
		if revolutions >= 1 {
			total2 += int(revolutions)
		}

		// calculate total distance to move & new current
		distance = clicks % 100
		sum = current + distance

		switch {
		case sum == 0, sum == 100:
			total1 += 1
			total2 += 1
			current = 0
		case sum < 0:
			if current != 0 {
				total2 += 1
			}
			current = 100 + sum
		case sum > 100:
			if current != 0 {
				total2 += 1
			}
			current = sum - 100
		default:
			current = sum
		}
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}
