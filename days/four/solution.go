package four

import (
	"strings"
)

type Indices struct {
	x, y int
}

var adjacent = []Indices{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}

func partOne(input string) (int64, error) {
	grid := parse(input)
	grid, removed := removeAccessibleRolls(grid)
	return removed, nil
}

func partTwo(input string) (int64, error) {
	grid := parse(input)

	var removed int64
	for removedLast := int64(1); removedLast > 0; {
		grid, removedLast = removeAccessibleRolls(grid)
		removed += removedLast
	}

	return removed, nil
}

func parse(input string) [][]rune {
	lines := strings.Split(input, "\n")
	columns := make([][]rune, len(lines))

	for l, line := range lines {
		row := make([]rune, len(line))
		for r, char := range line {
			row[r] = char
		}
		columns[l] = row
	}

	return columns
}

func removeAccessibleRolls(grid [][]rune) ([][]rune, int64) {
	removed := make([]Indices, 0)

	for y, col := range grid {
		for x, v := range col {
			if v == '@' {
				adjacentRolls := 0
				for _, indices := range findAdjacentIndices(grid, x, y) {
					if grid[indices.y][indices.x] == '@' {
						adjacentRolls++
					}
				}

				if adjacentRolls < 4 {
					removed = append(removed, Indices{x, y})
				}
			}
		}
	}

	for _, indices := range removed {
		grid[indices.y][indices.x] = '.'
	}

	return grid, int64(len(removed))
}

func findAdjacentIndices(grid [][]rune, x, y int) []Indices {
	neighbors := make([]Indices, 0)
	for _, adj := range adjacent {
		if y+adj.y >= 0 && y+adj.y < len(grid) && x+adj.x >= 0 && x+adj.x < len(grid[y+adj.y]) {
			neighbors = append(neighbors, Indices{x + adj.x, y + adj.y})
		}
	}
	return neighbors
}
