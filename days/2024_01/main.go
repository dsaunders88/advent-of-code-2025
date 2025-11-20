package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/dsaunders88/advent-of-code-2025/utils/data"
)

const INPUT_FILE string = "days/2024_01/input.txt"

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
	// fmt.Printf("%s\n", text)
	l1, l2 := parseText(text)
	slices.Sort(l1)
	slices.Sort(l2)

	var sum float64
	for i, v := range l1 {
		diff := math.Abs(float64(v - l2[i]))
		sum += diff
	}
	return int(sum)
}

func part2(input string) int {
	text := data.ReadAsString(input)
	l1, l2 := parseText(text)

	var sum int
	for _, v1 := range l1 {
		found := 0
		for _, v2 := range l2 {
			if v1 == v2 {
				found++
			}
		}
		sum += found * v1
	}

	return sum
}

func parseText(text string) (list1, list2 []int) {
	for row := range strings.SplitSeq(text, "\n") {
		v := strings.Split(row, "   ")
		i1, _ := strconv.Atoi(v[0])
		i2, _ := strconv.Atoi(v[1])
		list1 = append(list1, i1)
		list2 = append(list2, i2)
	}

	return list1, list2
}
