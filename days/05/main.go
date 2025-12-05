package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/dsaunders88/advent-of-code-2025/utils/data"
)

const INPUT_FILE string = "days/05/input.txt"

type IDRange struct {
	Lower  int
	Upper  int
	Length int
}

func main() {
	fmt.Println("  ˳")
	fmt.Println("  ** ---------------------")
	fmt.Println(" / °\\ Day 05 Solutions ~~")
	fmt.Println("/° ・\\ -------------------")

	var total1, total2 int

	input := data.ReadAsString(INPUT_FILE)

	var sets []IDRange
	var pastBreak bool        // empty line break in db file
	ids := make(map[int]bool) // ids to check p1

	// create sets and ids from db file
	for s := range strings.SplitSeq(input, "\n") {
		if !pastBreak {
			if s == "" {
				pastBreak = true
				continue
			}
			// ranges
			var set []int
			for v := range strings.SplitSeq(s, "-") {
				n, _ := strconv.Atoi(v)
				set = append(set, n)
			}
			idRange := IDRange{set[0], set[1], set[1] - set[0] + 1}
			sets = append(sets, idRange)
		} else {
			id, _ := strconv.Atoi(s)
			ids[id] = false
		}
	}

	// sort sets first before comparing - thanks reddit
	slices.SortFunc(sets, func(a, b IDRange) int {
		if n := cmp.Compare(a.Lower, b.Lower); n != 0 {
			return n
		}
		// if lower are equal, order by upper
		return cmp.Compare(a.Upper, b.Upper)
	})

	var totalLength int
	var maxBound int

	for i, set := range sets {
		var addedLength int

		if i != 0 {
			prevSet := sets[i-1]
			// check for any overlap
			if prevSet.Upper >= set.Lower {
				// check if not full overlap
				if !(set.Upper <= maxBound) {
					// add remaining from partial overlap
					remaining := set.Upper - prevSet.Upper
					addedLength = remaining
				}
			} else {
				addedLength = set.Length
			}
		} else {
			addedLength = set.Length
		}

		totalLength += addedLength

		if set.Upper > maxBound {
			maxBound = set.Upper
		}

	part1:
		for id := range ids {
			if id >= set.Lower && id <= set.Upper {
				total1++
				delete(ids, id)
			}
			continue part1
		}
	}

	total2 = totalLength

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}
