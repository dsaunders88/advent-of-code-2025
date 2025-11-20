package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_FILE string = "input-test.txt"

func TestPart1(t *testing.T) {
	expected := 11
	res := part1(TEST_FILE)
	assert.Equal(t, expected, res)
}

func TestPart2(t *testing.T) {
	expected := 31
	res := part2(TEST_FILE)
	assert.Equal(t, expected, res)
}
