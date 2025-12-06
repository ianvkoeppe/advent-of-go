package six

import (
	"advent-of-go/pkg/aoc"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	aoc.Assert(t, partOne, aoc.ReadExample, 4277556)
}

func TestPartOne(t *testing.T) {
	aoc.Assert(t, partOne, aoc.ReadProblem, 6371789547734)
}

func TestPartTwoExample(t *testing.T) {
	aoc.Assert(t, partTwo, aoc.ReadExample, 3263827)
}

func TestPartTwo(t *testing.T) {
	aoc.Assert(t, partTwo, aoc.ReadProblem, 11419862653216)
}
