package main

import (
	"fmt"
	// "math"
	// "regexp"
	// "slices"
	// "strconv"
	// "strings"
	"github.com/dsaunders88/advent-of-code-2025/utils/data"
)

const INPUT_FILE string = "days/{{.Name}}/input-test.txt"

func main() {
	fmt.Println("  ˳")
	fmt.Println("  ** ---------------------")
	fmt.Println(" / °\\ Day {{.Name}} Solutions ~~")
	fmt.Println("/° ・\\ -------------------")

	var total1, total2 int

	input := data.ReadAsString(INPUT_FILE)
	fmt.Printf("%v\n", input)

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}
