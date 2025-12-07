package seven

import (
	"advent-of-go/pkg/aoc"
	"strings"
)

type Pos struct {
	x, y int
}

func partOne(input string) (int64, error) {
	manifold, start := parse(input)

	splits := int64(0)
	positions := aoc.NewSet[Pos]()
	positions.Add(start)
	for !positions.Empty() {
		for _, pos := range positions.Elements() {
			if manifold[pos.y][pos.x] == '^' {
				positions.Add(Pos{pos.x - 1, pos.y + 1})
				positions.Add(Pos{pos.x + 1, pos.y + 1})
				splits++
			} else if pos.y+1 < len(manifold) {
				positions.Add(Pos{pos.x, pos.y + 1})
			}
			positions.Remove(pos)
		}
	}

	return splits, nil
}

func partTwo(input string) (int64, error) {
	manifold, start := parse(input)

	splits := int64(0)
	positions := map[Pos]int64{start: 1}
	for len(positions) > 0 {
		nextPositions := map[Pos]int64{}
		for pos, paths := range positions {
			if manifold[pos.y][pos.x] == '^' {
				nextPositions[Pos{pos.x - 1, pos.y + 1}] += paths
				nextPositions[Pos{pos.x + 1, pos.y + 1}] += paths
			} else if pos.y+1 < len(manifold) {
				nextPositions[Pos{pos.x, pos.y + 1}] += paths
			} else {
				splits += paths
			}
		}
		positions = nextPositions
	}

	return splits, nil
}

func parse(input string) ([][]rune, Pos) {
	start := Pos{}

	lines := strings.Split(input, "\n")
	manifold := make([][]rune, len(lines))
	for y, line := range lines {
		manifold[y] = make([]rune, len(line))
		for x, char := range line {
			manifold[y][x] = char

			if char == 'S' {
				start.x = x
				start.y = y
			}
		}
	}

	return manifold, start
}
