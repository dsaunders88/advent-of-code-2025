package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dsaunders88/advent-of-code-2025/utils"
	"github.com/urfave/cli/v3"
)

var dayArgument = &cli.StringArg{
	Name: "day",
}

func main() {
	cmd := &cli.Command{
		Name:                   "aoc",
		Usage:                  "Tool to scaffold files and solve puzzles for Advent of Code.",
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:      "scaffold",
				Aliases:   []string{"f"},
				Usage:     "Generate directory and templates for a day's puzzle. Use the 'day' argument to specify the name of the directory (should be a two digit number with a leading 0 if applicable, i.e., '01')",
				Arguments: []cli.Argument{dayArgument},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					day := cmd.StringArg("day")
					if day == "" {
						log.Fatal("Need to specify a day")
					}
					fmt.Printf("Scaffolding directory and templates for day '%s'\n", cmd.StringArg("day"))
					err := utils.ScaffoldDayTemplates(day)
					if err != nil {
						log.Fatal(err)
					}

					return nil
				},
			},
			// {
			// 	Name:      "solve",
			// 	Aliases:   []string{"s"},
			// 	Usage:     "Run the `main.go` file in a given day's directory to return a solution from an input file.",
			// 	Arguments: []cli.Argument{dayArgument},
			// 	Action: func(ctx context.Context, cmd *cli.Command) error {
			// 		return nil
			// 	},
			// },
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
