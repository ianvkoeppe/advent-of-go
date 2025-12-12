package ten

import (
	"advent-of-go/pkg/aoc"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	aoc.Assert(t, partOne, func() (string, error) { return aoc.Read("example-1") }, 5)
}

func TestPartOne(t *testing.T) {
	aoc.Assert(t, partOne, aoc.ReadProblem, 652)
}

func TestPartTwoExample(t *testing.T) {
	aoc.Assert(t, partTwo, func() (string, error) { return aoc.Read("example-2") }, 2)
}

func TestPartTwo(t *testing.T) {
	aoc.Assert(t, partTwo, aoc.ReadProblem, 362956369749210)
}
