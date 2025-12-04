package data

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

// Read the contents of an input file from the working directory as a string.
// TODO/improve: `os.Getwd()` will get the working directory from wherever this
// is called, so have to be careful with the `filename` passed in (i.e., it will
// be different for `main.go` and `main_test.go` since tests are run from the
// local directory).
func ReadAsString(filename string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(path.Join(wd, filename))
	if err != nil {
		log.Fatal(err)
	}
	str := string(data)
	// trim new lines from end of input file
	return strings.TrimRight(str, "\n")
}

// Read contents of an input file as array of strings for each line.
func ReadAsLines(filename string) (lines []string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// open new file
	data, err := os.Open(path.Join(wd, filename))
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	// implement new scanner from buffered i/o & scan each line
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		lines = append(lines, strings.TrimRight(scanner.Text(), "\n"))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
