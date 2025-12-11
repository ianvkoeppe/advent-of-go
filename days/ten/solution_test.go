package ten

import (
	"advent-of-go/pkg/aoc"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	aoc.Assert(t, partOne, aoc.ReadExample, 7)
}

func TestPartOne(t *testing.T) {
	aoc.Assert(t, partOne, aoc.ReadProblem, 535)
}

func TestPartTwoExample(t *testing.T) {
	aoc.Assert(t, partTwo, aoc.ReadExample, 33)
}

func TestPartTwo(t *testing.T) {
	aoc.Assert(t, partTwo, aoc.ReadProblem, 21021)
}
