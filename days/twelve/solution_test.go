package twelve

import (
	"advent-of-go/pkg/aoc"
	"testing"
)

func TestPartOne(t *testing.T) {
	aoc.Assert(t, partOne, aoc.ReadProblem, 577)
}
