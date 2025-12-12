package twelve

import (
	"strconv"
	"strings"
)

type Tree struct {
	area     int
	presents []int
}

func partOne(input string) (int64, error) {
	trees, err := parse(input)
	if err != nil {
		return 0, err
	}

	valid := 0
	for _, tree := range trees {
		sum := 0
		for _, p := range tree.presents {
			sum += p
		}

		if tree.area >= sum*9 {
			valid++
		}
	}

	return int64(valid), nil
}

func parse(input string) ([]Tree, error) {
	sections := strings.Split(input, "\n\n")

	treeText := strings.Split(sections[len(sections)-1], "\n")
	trees := make([]Tree, len(treeText))
	for t, tree := range treeText {
		parts := strings.Split(tree, ": ")
		dimensions := strings.Split(parts[0], "x")
		width, err := strconv.Atoi(dimensions[0])
		if err != nil {
			return nil, err
		}

		length, err := strconv.Atoi(dimensions[1])
		if err != nil {
			return nil, err
		}

		presentCountText := strings.Split(parts[1], " ")
		presentCount := make([]int, len(presentCountText))
		for p, countText := range presentCountText {
			presentCount[p], err = strconv.Atoi(countText)
			if err != nil {
				return nil, err
			}
		}

		trees[t] = Tree{width * length, presentCount}
	}

	return trees, nil
}
