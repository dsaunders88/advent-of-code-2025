package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dsaunders88/advent-of-code-2025/utils/data"
)

const INPUT_FILE string = "days/02/input.txt"

func main() {
	fmt.Println("  ˳")
	fmt.Println("  ** ---------------------")
	fmt.Println(" / °\\ Day 02 Solutions ~~")
	fmt.Println("/° ・\\ -------------------")

	var total1, total2 int

	input := data.ReadAsString(INPUT_FILE)
	for set := range strings.SplitSeq(input, ",") {
		var bounds []string
		bounds = strings.Split(set, "-")
		lower, _ := strconv.Atoi(bounds[0])
		upper, _ := strconv.Atoi(bounds[1])

		for i := lower; i <= upper; i++ {
			s := fmt.Sprint(i)
			length := len(s)

			// part 1: check for evens only - first half of string
			// part 2: check for evens, 3s, 5s, and 1xs
			if length%2 == 0 {
				repeatable := s[:len(s)/2]

				switch {
				case strings.Count(s, repeatable) > 1:
					total1 += i
					total2 += i
				case hasRepeats(s, length, 3):
					total2 += i
				case hasRepeats(s, length, 5):
					total2 += i
				}
			} else if length > 1 {
				// check for singles & 3s on odd lengths, skip single digits
				first := s[:1]

				switch {
				case strings.Count(s, first) == length:
					total2 += i
				case hasRepeats(s, length, 3):
					total2 += i
				}
			}

		}
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}

// s = string to check; length = length of string; div = repeatable divison
// count = number of occurrences, ok = if modulo calc is zero & count is length/freq
func hasRepeats(s string, length int, div int) bool {
	if length%div == 0 {
		repeatable := s[:len(s)/div]
		freq := length / div
		if strings.Count(s, repeatable) == length/freq {
			return true
		}
	}

	return false
}
