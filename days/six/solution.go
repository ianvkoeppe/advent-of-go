package six

import (
	"regexp"
	"strconv"
	"strings"
)

type Problem struct {
	nums []int
	op   rune
}

var whiteSpace = regexp.MustCompile("\\s+")

var ops = map[rune]func(int64, int64) int64{
	'+': func(a, b int64) int64 { return a + b },
	'*': func(a, b int64) int64 { return a * b },
}

func partOne(input string) (int64, error) {
	lines := strings.Split(input, "\n")

	problems := make([]Problem, len(whiteSpace.Split(strings.TrimSpace(lines[0]), -1)))
	for p := range problems {
		problems[p].nums = make([]int, len(lines)-1)
	}

	for l, line := range lines[:len(lines)-1] {
		for n, numText := range whiteSpace.Split(strings.TrimSpace(line), -1) {
			num, err := strconv.Atoi(numText)
			if err != nil {
				return 0, err
			}

			problems[n].nums[l] = num
		}
	}

	for o, opText := range whiteSpace.Split(strings.TrimSpace(lines[len(lines)-1]), -1) {
		problems[o].op = rune(opText[0])
	}

	return solveAll(problems), nil
}

func partTwo(input string) (int64, error) {
	lines := strings.Split(input, "\n")
	problems := make([]Problem, 1)

	lastOp := ' '
	for x := 0; x < len(lines[0]); x++ {
		if problems[len(problems)-1].nums == nil {
			problems[len(problems)-1].nums = make([]int, 0)
		}

		numText := "0"
		for y := 0; y < len(lines)-1; y++ {
			if lines[y][x] != ' ' {
				numText += string(lines[y][x])
			}
		}
		num, err := strconv.Atoi(numText)
		if err != nil {
			return 0, err
		}

		if num == 0 {
			problems[len(problems)-1].op = lastOp
			problems = append(problems, Problem{})
		} else {
			problems[len(problems)-1].nums = append(problems[len(problems)-1].nums, num)
			if lines[len(lines)-1][x] != ' ' {
				lastOp = rune(lines[len(lines)-1][x])
			}
		}
	}
	problems[len(problems)-1].op = lastOp

	return solveAll(problems), nil
}

func solveAll(problems []Problem) int64 {
	grandTotal := int64(0)
	for _, problem := range problems {
		total := int64(problem.nums[0])
		for _, num := range problem.nums[1:] {
			total = ops[problem.op](total, int64(num))
		}
		grandTotal += total
	}
	return grandTotal
}
