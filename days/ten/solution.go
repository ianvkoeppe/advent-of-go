package ten

import (
	"advent-of-go/pkg/aoc"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Machine struct {
	lights  []int
	buttons [][]int
	jolts   []int

	ops      map[string][]int
	patterns map[string][][]int
	cache    map[string]int
}

func (m *Machine) findMinPressesForLights() int64 {
	m.ops = make(map[string][]int)
	m.patterns = make(map[string][][]int)

	for mask := 0; mask < (1 << len(m.buttons)); mask++ {
		pressed := make([]int, len(m.buttons))
		for b := 0; b < len(m.buttons); b++ {
			if mask&(1<<b) != 0 {
				pressed[b] = 1
			}
		}

		jolt := make([]int, len(m.jolts))
		for i, p := range pressed {
			if p == 1 {
				for _, l := range m.buttons[i] {
					jolt[l] += 1
				}
			}
		}

		lights := make([]int, len(m.jolts))
		for i, x := range jolt {
			lights[i] = x % 2
		}

		m.ops[fmt.Sprint(pressed)] = jolt

		keyLights := fmt.Sprint(lights)
		m.patterns[keyLights] = append(m.patterns[keyLights], pressed)
	}

	minPresses := math.MaxInt
	for _, p := range m.patterns[fmt.Sprint(m.lights)] {
		if s := aoc.Sum(p); s < minPresses {
			minPresses = s
		}
	}

	return int64(minPresses)
}

func (m *Machine) findMinPressesToAchieveJolts() int64 {
	m.findMinPressesForLights() // Prepopulate `ops` and `patterns`.
	m.cache = make(map[string]int)
	return int64(m.findMinPressesToAchieve(m.jolts))
}

func (m *Machine) findMinPressesToAchieve(jolts []int) int {
	key := fmt.Sprint(jolts)
	if v, ok := m.cache[key]; ok {
		return v
	}

	allZero := true
	for _, x := range jolts {
		if x < 0 {
			return 1 << 30
		}
		if x != 0 {
			allZero = false
		}
	}
	if allZero {
		return 0
	}

	lights := make([]int, len(m.lights))
	for i, x := range jolts {
		lights[i] = x % 2
	}
	keyLights := fmt.Sprint(lights)

	minPresses := 1 << 30
	for _, pressed := range m.patterns[keyLights] {
		diff := m.ops[fmt.Sprint(pressed)]

		newTarget := make([]int, len(m.lights))
		for j := range jolts {
			newTarget[j] = (jolts[j] - diff[j]) / 2
		}

		minPresses = min(minPresses, aoc.Sum(pressed)+2*m.findMinPressesToAchieve(newTarget))
	}

	m.cache[key] = minPresses
	return minPresses
}

func partOne(input string) (int64, error) {
	return findMinPresses(input, func(m Machine) int64 { return m.findMinPressesForLights() })
}

func partTwo(input string) (int64, error) {
	return findMinPresses(input, func(m Machine) int64 { return m.findMinPressesToAchieveJolts() })
}

func parse(input string) ([]Machine, error) {
	lines := strings.Split(input, "\n")

	machines := make([]Machine, len(lines))
	for l, line := range lines {
		parts := strings.Split(line, " ")

		lights := make([]int, len(parts[0])-2)
		for l, light := range parts[0][1 : len(parts[0])-1] {
			if light == '#' {
				lights[l] = 1
			} else {
				lights[l] = 0
			}
		}
		machines[l].lights = lights

		buttonTexts := parts[1 : len(parts)-1]
		buttons := make([][]int, len(buttonTexts))
		for w, buttonText := range buttonTexts {
			values := strings.Split(buttonText[1:len(buttonText)-1], ",")
			buttons[w] = make([]int, len(values))
			for v, valueText := range values {
				value, err := strconv.Atoi(valueText)
				if err != nil {
					return nil, err
				}

				buttons[w][v] = value
			}
		}
		machines[l].buttons = buttons

		joltTexts := strings.Split(parts[len(parts)-1][1:len(parts[len(parts)-1])-1], ",")
		jolts := make([]int, len(joltTexts))
		for j, joltText := range joltTexts {
			joltage, err := strconv.Atoi(joltText)
			if err != nil {
				return nil, err
			}
			jolts[j] = joltage
		}
		machines[l].jolts = jolts
	}

	return machines, nil
}

func findMinPresses(input string, objective func(m Machine) int64) (int64, error) {
	machines, err := parse(input)
	if err != nil {
		return 0, nil
	}

	presses := int64(0)
	for _, m := range machines {
		presses += objective(m)
	}
	return presses, nil
}
