package aoc

import "testing"

func Assert(t *testing.T, solver func(string) (int64, error), reader func() (string, error), expected int64) {
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
