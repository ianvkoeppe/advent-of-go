package one

import (
	"advent-of-go/pkg/io"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	assert(t, partOne, io.ReadExample, 0)
}

func TestPartOne(t *testing.T) {
	assert(t, partOne, io.ReadProblem, 0)
}

func TestPartTwoExample(t *testing.T) {
	assert(t, partTwo, io.ReadExample, 0)
}

func TestPartTwo(t *testing.T) {
	assert(t, partTwo, io.ReadProblem, 0)
}

func assert(t *testing.T, solver func(string) (int, error), reader func() (string, error), expected int) {
	input, err := reader()
	if err != nil {
		t.Fatal(err)
	}

	output, err := solver(input)
	if err != nil {
		t.Fatal(err)
	}

	if output != expected {
		t.Errorf("Output: `%d` | Expected: `%d`", output, expected)
	}
}
