# Advent of Code 2025 in Go

My [Advent of Code](https://adventofcode.com/2025) solutions for 2025 in Go.

## CLI Commands

Run the following commands by installing this module with `go install` or by running `go run main.go [global options] [command [command options]]`.

| Command | Alias | Description |
| ------- | ----- | ----------- |
| `scaffold` | `f` | Generate directory and templates for a day's puzzle. Creates `main.go`, `main_test.go`, `input.txt`, and `input-test.txt` files. |
| `solve` | `s` | Run the `main.go` file in a day's directory to return a solution from an input file. |

> [!NOTE]
> Input filepaths assume that these commands are run from the project root. Tests on the other hand should be run from the day directory.