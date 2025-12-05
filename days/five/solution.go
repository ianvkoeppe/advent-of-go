package five

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type FreshRange struct {
	start, end int
}

func partOne(input string) (int64, error) {
	freshRanges, ingredients, err := parse(input)
	if err != nil {
		return 0, err
	}

	mergedRanges := mergeAndSort(freshRanges)

	var fresh int64 = 0
	for _, ingredient := range ingredients {
		if isFresh(mergedRanges, ingredient) {
			fresh++
		}
	}

	return fresh, nil
}

func partTwo(input string) (int64, error) {
	freshRanges, _, err := parse(input)
	if err != nil {
		return 0, err
	}

	mergedRanges := mergeAndSort(freshRanges)

	var fresh int64 = 0
	for _, r := range mergedRanges {
		fresh += int64(r.end) - int64(r.start) + 1
	}

	return fresh, nil
}

func parse(input string) ([]FreshRange, []int, error) {
	parts := strings.Split(input, "\n\n")

	textRanges := strings.Split(parts[0], "\n")
	ranges := make([]FreshRange, len(textRanges))
	for r, textRange := range textRanges {
		rangeParts := strings.Split(textRange, "-")
		start, err := strconv.Atoi(rangeParts[0])
		if err != nil {
			return nil, nil, err
		}

		end, err := strconv.Atoi(rangeParts[1])
		if err != nil {
			return nil, nil, err
		}

		ranges[r] = FreshRange{start, end}
	}

	textIngredients := strings.Split(parts[1], "\n")
	ingredients := make([]int, len(textIngredients))
	for i, textIngredient := range textIngredients {
		ingredient, err := strconv.Atoi(textIngredient)
		if err != nil {
			return nil, nil, err
		}
		ingredients[i] = ingredient
	}

	return ranges, ingredients, nil
}

func mergeAndSort(ranges []FreshRange) []FreshRange {
	sortRanges(ranges)

	merged := make([]FreshRange, 1)
	merged[0] = ranges[0]
	for r := 1; r < len(ranges); r++ {
		if merged[len(merged)-1].end >= ranges[r].start {
			merged[len(merged)-1].end = int(math.Max(float64(merged[len(merged)-1].end), float64(ranges[r].end)))
		} else {
			merged = append(merged, ranges[r])
		}
	}
	return merged
}

func sortRanges(ranges []FreshRange) {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})
}

func isFresh(ranges []FreshRange, ingredient int) bool {
	s, e := 0, len(ranges)-1
	for s <= e {
		mid := s + (e-s)/2
		r := ranges[mid]

		if r.start <= ingredient && ingredient <= r.end {
			return true
		} else if ingredient >= r.start {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}

	return false
}
