package two

import (
	"math"
	"strconv"
	"strings"
)

type ProductIdRange struct {
	start int64
	end   int64
}

func (r ProductIdRange) StartNumberOfDigits() int {
	return len(strconv.FormatInt(r.start, 10))
}

func (r ProductIdRange) EndNumberOfDigits() int {
	return len(strconv.FormatInt(r.end, 10))
}

func partOne(input string) (int64, error) {
	ranges, err := parse(input)
	if err != nil {
		return 0, err
	}

	var invalid int64 = 0
	for _, r := range ranges {
		invalid += countInvalidProductIds(r, 1)
	}

	return invalid, nil
}

func partTwo(input string) (int64, error) {
	ranges, err := parse(input)
	if err != nil {
		return 0, err
	}

	var invalid int64 = 0
	for _, r := range ranges {
		invalid += countInvalidProductIds(r, r.EndNumberOfDigits())
	}

	return invalid, nil
}

func parse(input string) ([]ProductIdRange, error) {
	lines := strings.Split(input, ",")
	ranges := make([]ProductIdRange, len(lines))

	for _, line := range lines {
		ids := strings.Split(line, "-")
		start, err := strconv.ParseInt(ids[0], 10, 64)
		if err != nil {
			return nil, err
		}

		end, err := strconv.ParseInt(ids[1], 10, 64)
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, ProductIdRange{start, end})
	}

	return ranges, nil
}

func countInvalidProductIds(r ProductIdRange, maxRepeat int) int64 {
	seen := map[int64]bool{}
	var invalid int64 = 0
	for repeats := 1; repeats <= maxRepeat; repeats++ {
		for d := r.StartNumberOfDigits() / (1 + repeats); d <= r.EndNumberOfDigits()/(1+repeats); d++ {
			for i := int(math.Pow10(d - 1)); i < 10*int(math.Pow10(d-1)); i++ {
				repeated := strings.Builder{}
				repeated.WriteString(strconv.Itoa(i))
				for x := 1; x <= repeats; x++ {
					repeated.WriteString(strconv.Itoa(i))
				}
				scaled, _ := strconv.ParseInt(repeated.String(), 10, 64)
				if scaled > r.end {
					break
				}

				if scaled >= r.start && !seen[scaled] {
					invalid += scaled
					seen[scaled] = true
				}
			}
		}
	}
	return invalid
}
