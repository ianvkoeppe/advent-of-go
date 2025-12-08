package eight

import (
	"advent-of-go/pkg/aoc"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	aoc.Assert(t, func(s string) (int64, error) {
		return partOne(s, 10)
	}, aoc.ReadExample, 40)
}

func TestPartOne(t *testing.T) {
	aoc.Assert(t, func(s string) (int64, error) {
		return partOne(s, 1000)
	}, aoc.ReadProblem, 175500)
}

func TestPartTwoExample(t *testing.T) {
	aoc.Assert(t, partTwo, aoc.ReadExample, 25272)
}

func TestPartTwo(t *testing.T) {
	aoc.Assert(t, partTwo, aoc.ReadProblem, 6934702555)
}
