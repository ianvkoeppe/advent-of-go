package ten

import (
	"advent-of-go/pkg/aoc"
	"sort"
	"strings"
)

type ServerRack struct {
	devices map[string][]string
	counts  map[string]int64
}

func (s *ServerRack) countPaths(current string, mustSee *aoc.Set[string]) int64 {
	key := createKey(current, mustSee)
	if count, ok := s.counts[key]; ok {
		return count
	}

	if current == "out" {
		if mustSee.Size() == 0 {
			return 1
		}
		return 0
	}

	sum := int64(0)
	for _, device := range s.devices[current] {
		nowMustSee := mustSee
		if mustSee.Contains(device) {
			nowMustSee = aoc.NewSet[string](mustSee.Elements()...)
			nowMustSee.Remove(device)
		}

		sum += s.countPaths(device, nowMustSee)
	}
	s.counts[key] = sum
	return s.counts[key]
}

func createKey(prefix string, set *aoc.Set[string]) string {
	elements := set.Elements()
	sort.Strings(elements)
	return prefix + "|" + strings.Join(elements, ",")
}

func partOne(input string) (int64, error) {
	rack := parse(input)
	return rack.countPaths("you", aoc.NewSet[string]()), nil
}

func partTwo(input string) (int64, error) {
	rack := parse(input)
	return rack.countPaths("svr", aoc.NewSet[string]("fft", "dac")), nil
}

func parse(input string) ServerRack {
	lines := strings.Split(input, "\n")

	devices := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, ": ")

		outputs := strings.Split(parts[1], " ")
		devices[parts[0]] = make([]string, len(outputs))
		for p, part := range outputs {
			devices[parts[0]][p] = part
		}
	}

	return ServerRack{devices, make(map[string]int64)}
}
