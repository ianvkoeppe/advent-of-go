package three

import (
	"strings"
)

type Bank struct {
	batteries []int
}

func partOne(input string) (int64, error) {
	return findMaxJoltage(input, 2), nil
}

func partTwo(input string) (int64, error) {
	return findMaxJoltage(input, 12), nil
}

func parse(input string) []Bank {
	lines := strings.Split(input, "\n")
	banks := make([]Bank, len(lines))

	for l, line := range lines {
		batteries := make([]int, len(line))
		for b, battery := range line {
			batteries[b] = int(battery) - '0'
		}
		banks[l] = Bank{batteries}
	}

	return banks
}

func findMaxJoltage(input string, digits int) int64 {
	var total int64 = 0

	for _, bank := range parse(input) {
		var joltage int64 = 0
		needed := digits

		for start, end := 0, len(bank.batteries)-1; needed > 0; {
			m, max := findMaxWithin(bank, start, end)
			if len(bank.batteries)-m >= needed {
				joltage = joltage*10 + int64(max)
				start = m + 1
				end = len(bank.batteries) - 1
				needed--
			} else {
				end = m - 1
			}
		}

		total += joltage
	}

	return total
}

func findMaxWithin(bank Bank, start int, end int) (int, int) {
	m, max := 0, 0
	for b := start; b <= end; b++ {
		if bank.batteries[b] > max {
			max = bank.batteries[b]
			m = b
		}
	}
	return m, bank.batteries[m]
}
