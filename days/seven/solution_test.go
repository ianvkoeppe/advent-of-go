package seven

import (
	"advent-of-go/pkg/aoc"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	aoc.Assert(t, partOne, aoc.ReadExample, 21)
}

func TestPartOne(t *testing.T) {
	aoc.Assert(t, partOne, aoc.ReadProblem, 1628)
}

func TestPartTwoExample(t *testing.T) {
	aoc.Assert(t, partTwo, aoc.ReadExample, 40)
}

func TestPartTwo(t *testing.T) {
	aoc.Assert(t, partTwo, aoc.ReadProblem, 27055852018812)
}
