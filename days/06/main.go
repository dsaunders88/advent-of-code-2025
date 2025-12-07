package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dsaunders88/advent-of-code-2025/utils/data"
)

const INPUT_FILE string = "days/06/input.txt"

type Set struct {
	StartIdx int
	Operator string
	Length   int
}

func main() {
	fmt.Println("  ˳")
	fmt.Println("  ** ---------------------")
	fmt.Println(" / °\\ Day 06 Solutions ~~")
	fmt.Println("/° ・\\ -------------------")

	var total1, total2 int

	input := data.ReadAsLines(INPUT_FILE)

	// regex to capture operator with trailing whitespace
	// use to calculate set lengths for part 2
	re := regexp.MustCompile(`(\*|\+)(\s+)`)
	// last row of input = operators
	opRow := re.FindAllString(input[len(input)-1], -1)
	var operators []string
	var sets []Set

	// get operators and set lengths
	for i, s := range opRow {
		set := Set{Length: len(s)}

		for _, char := range s {
			c := string(char)
			if c == "*" || c == "+" {
				set.Operator = c
				operators = append(operators, c)
			}
		}

		// set all but last set on right length to 1 less
		// to reflect single space between the sets
		if i < len(opRow)-1 {
			set.Length = len(s) - 1
		}

		// add start index of each set
		if i > 0 {
			prevSet := sets[i-1]
			set.StartIdx = prevSet.StartIdx + prevSet.Length + 1
		} else {
			set.StartIdx = 0
		}

		sets = append(sets, set)
	}

	// part 1 - key: set index, val: running total for set
	setTotals := make(map[int]int)
	// part 2 - key: col index, val: vertical column combined string values
	vertTotals := make(map[int]string)

	// todo: combine part 1 and 2 in one loop?
	// part 1
	for i, row := range input {
		if i < len(input)-1 {
			// all non whitespace groups
			re := regexp.MustCompile(`[^\s]+`)
			vals := re.FindAllString(row, -1)

			// range through cols
			for j, v := range vals {
				colVal, _ := strconv.Atoi(v)

				if i == 0 {
					// set inital map values
					setTotals[j] = colVal
				} else {
					if prevVal, ok := setTotals[j]; ok {
						var res int
						op := operators[j]
						if op == "*" {
							res = prevVal * colVal
						} else {
							res = prevVal + colVal
						}
						setTotals[j] = res
					}
				}
			}

			// part 2: make vertical column totals
			for j, v := range row {
				n := string(v)
				if i == 0 {
					// set initial map vals on first index
					vertTotals[j] = n
				} else {
					if prevVal, ok := vertTotals[j]; ok {
						vertTotals[j] = fmt.Sprintf("%s%s", prevVal, n)
					}
				}
			}
		}
	}

	for _, v := range setTotals {
		total1 += v
	}

	// part 2
	for _, set := range sets {
		// init as 1 for multiply so not multiplying by zero
		var setTotal int
		if set.Operator == "*" {
			setTotal = 1
		}

		// loop through indexes of the set size
		for j := set.StartIdx; j < set.StartIdx+set.Length; j++ {
			// get stored values for each column
			if val, ok := vertTotals[j]; ok {
				// trim whitespace around values
				n, _ := strconv.Atoi(strings.Trim(val, " "))

				if set.Operator == "*" {
					setTotal = setTotal * n
				} else {
					setTotal = setTotal + n
				}
			}
		}

		total2 += setTotal
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}
