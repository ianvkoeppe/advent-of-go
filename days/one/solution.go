package one

import (
	"math"
	"strconv"
	"strings"
)

type Turn struct {
	direction rune
	ticks     int
}

func (t Turn) DirectionalTicks() int {
	if t.direction == 'L' {
		return -t.ticks % 100
	}
	return t.ticks % 100
}

func (t Turn) Rotations() int {
	return int(math.Abs(float64(t.ticks / 100)))
}

func partOne(input string) (int64, error) {
	turns, err := parse(input)
	if err != nil {
		return 0, err
	}

	pos := 50
	var zeros int64 = 0
	for _, turn := range turns {
		pos = (pos + turn.DirectionalTicks() + 100) % 100
		if pos == 0 {
			zeros++
		}
	}

	return zeros, nil
}

func partTwo(input string) (int64, error) {
	turns, err := parse(input)
	if err != nil {
		return 0, err
	}

	pos := 50
	var zeros int64 = 0
	for _, turn := range turns {
		zeros += int64(turn.Rotations())
		if pos != 0 && ((pos+turn.DirectionalTicks()) <= 0 || (pos+turn.DirectionalTicks()) >= 100) {
			zeros++
		}
		pos = (pos + turn.DirectionalTicks() + 100) % 100
	}

	return zeros, nil
}

func parse(input string) ([]Turn, error) {
	lines := strings.Split(input, "\n")
	turns := make([]Turn, len(lines))

	for _, line := range lines {
		ticks, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		turns = append(turns, Turn{direction: rune(line[0]), ticks: ticks})
	}

	return turns, nil
}
