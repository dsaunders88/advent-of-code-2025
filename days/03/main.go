package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dsaunders88/advent-of-code-2025/utils/data"
)

const INPUT_FILE string = "days/03/input.txt"

func main() {
	fmt.Println("  ˳")
	fmt.Println("  ** ---------------------")
	fmt.Println(" / °\\ Day 03 Solutions ~~")
	fmt.Println("/° ・\\ -------------------")

	var total1, total2 int

	input := data.ReadAsString(INPUT_FILE)
	for row := range strings.SplitSeq(input, "\n") {
		var joltage [2]int
		// max 12 but use slice here so easier to reform later
		var max []int

		for _, char := range row {
			j, _ := strconv.Atoi(string(char))

			// part 1
			// compare right w/ left, then compare current val with right
			if joltage[1] > joltage[0] {
				joltage[0] = joltage[1]
				joltage[1] = j
			} else if j > joltage[1] {
				joltage[1] = j
			}

			// part 2
			// fill slice, then run check once 12 spots are filled
			max = append(max, j)
			if len(max) > 12 {
				// range through slice of 12 to compare each
				// need to delete 1
				for i, v := range max {
					// stay within slice bounds for indexing
					if i+1 < len(max) {
						// if next value is greater than current
						if max[i+1] > v {
							// delete value and break
							max = append(max[:i], max[i+1:]...)
							break
						}
					} else {
						// look at last two values
						// if last is greater, delete penultimate
						if max[len(max)-1] > max[len(max)-2] {
							max = append(max[:len(max)-2], max[len(max)-1:]...)
						} else {
							// else delete last
							max = max[:len(max)-1]
						}
					}
				}
			}
		}

		join := fmt.Sprintf("%v%v", joltage[0], joltage[1])

		// construct max joltage from slice of 12
		var joinMax string
		for _, n := range max {
			joinMax = fmt.Sprintf("%s%v", joinMax, n)
		}

		v1, _ := strconv.Atoi(join)
		v2, _ := strconv.Atoi(joinMax)
		total1 += v1
		total2 += v2
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}
