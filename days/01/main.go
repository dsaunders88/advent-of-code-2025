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
	fmt.Println("Part 1:", part1(INPUT_FILE))
	fmt.Println("Part 2:", part2(INPUT_FILE))
}

func part1(input string) int {
	text := data.ReadAsString(input)
	list := parseText(text)

	// current will always be between 0-99
	var current, total int
	current = 50

	for _, v := range list {
		// get the remaining distance < 100 after all rotations
		distance := v % 100
		sum := current + distance

		switch {
		case sum == 0, sum == 100:
			total += 1
			current = 0
		case sum < 0:
			current = 100 + sum
		case sum > 100:
			current = sum - 100
		default:
			current = sum
		}
	}

	return total
}

func part2(input string) int {
	text := data.ReadAsString(input)
	list := parseText(text)

	var current, total int
	current = 50

	for _, v := range list {
		// total number of full rotations past 0
		// (make sure to get absolute of v)
		rotations := math.Floor(float64(math.Abs(float64(v))) / 100)
		if rotations >= 1 {
			total += int(rotations)
		}
		
		distance := v % 100
		sum := current + distance

		// modified version of above, add to total if passing 0,
		// not starting on 0
		switch {
		case sum == 0, sum == 100:
			total += 1
			current = 0
		case sum < 0:
			if current != 0 {
				total += 1
			}
			current = 100 + sum
		case sum > 100:
			if current != 0 {
				total += 1
			}
			current = sum - 100
		default:
			current = sum
		}
	}

	return total
}

func parseText(text string) (list []int) {
	for row := range strings.SplitSeq(text, "\n") {
		// for each combination value, split direction from value
		// return positive int for right, negative for left
		val, foundR := strings.CutPrefix(row, "R")
		if foundR {
			i, _ := strconv.Atoi(val)
			list = append(list, i)
		} else {
			i, _ := strconv.Atoi(row[1:])
			list = append(list, i * -1)
		}
	}

	return list
}
