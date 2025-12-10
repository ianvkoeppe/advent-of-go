package nine

import (
	"advent-of-go/pkg/aoc"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	aoc.Assert(t, partOne, aoc.ReadExample, 50)
}

func TestPartOne(t *testing.T) {
	aoc.Assert(t, partOne, aoc.ReadProblem, 4725826296)
}

func TestPartTwoExample(t *testing.T) {
	aoc.Assert(t, partTwo, aoc.ReadExample, 24)
}

func TestPartTwo(t *testing.T) {
	aoc.Assert(t, partTwo, aoc.ReadProblem, 1637556834)
}
