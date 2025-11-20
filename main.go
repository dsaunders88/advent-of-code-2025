package main

import (
	"log"

	"github.com/dsaunders88/advent-of-code-2025/utils"
)

func main() {
	// TODO: get current day from CLI flag

	err := utils.ScaffoldDayTemplates()
	if err != nil {
		log.Fatal(err)
	}
}
