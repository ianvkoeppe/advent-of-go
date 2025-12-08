package eight

import (
	"advent-of-go/pkg/aoc"
	"math"
	"sort"
	"strconv"
	"strings"
)

type JunctionBox struct {
	x, y, z int
}

func (jb JunctionBox) distanceBetween(other JunctionBox) float64 {
	return math.Sqrt(math.Pow(float64(jb.x-other.x), 2) + math.Pow(float64(jb.y-other.y), 2) + math.Pow(float64(jb.z-other.z), 2))
}

type Connection struct {
	from JunctionBox
	to   JunctionBox
	dist float64
}

func newConnection(from JunctionBox, to JunctionBox) Connection {
	return Connection{from: from, to: to, dist: from.distanceBetween(to)}
}

type Circuits struct {
	circuitsById                map[int]*aoc.Set[JunctionBox]
	circuitsByBox               map[JunctionBox]int
	nextId                      int
	firstToFullyCompleteCircuit Connection
}

func newCircuits() *Circuits {
	return &Circuits{circuitsById: make(map[int]*aoc.Set[JunctionBox]), circuitsByBox: make(map[JunctionBox]int), nextId: 0}
}

func (circuits *Circuits) connect(connection Connection) {
	fromGroup, fromOk := circuits.circuitsByBox[connection.from]
	toGroup, toOk := circuits.circuitsByBox[connection.to]

	if fromOk && toOk && fromGroup != toGroup {
		for _, jb := range circuits.circuitsById[fromGroup].Elements() {
			circuits.circuitsById[toGroup].Add(jb)
			circuits.circuitsByBox[jb] = toGroup
		}
		delete(circuits.circuitsById, fromGroup)
	} else if fromOk && !toOk {
		circuits.circuitsById[fromGroup].Add(connection.to)
		circuits.circuitsByBox[connection.to] = fromGroup
		circuits.firstToFullyCompleteCircuit = connection
	} else if toOk && !fromOk {
		circuits.circuitsById[toGroup].Add(connection.from)
		circuits.circuitsByBox[connection.from] = toGroup
		circuits.firstToFullyCompleteCircuit = connection
	} else if !fromOk {
		circuits.circuitsById[circuits.nextId] = aoc.NewSet[JunctionBox]()
		circuits.circuitsById[circuits.nextId].Add(connection.from)
		circuits.circuitsById[circuits.nextId].Add(connection.to)
		circuits.circuitsByBox[connection.from] = circuits.nextId
		circuits.circuitsByBox[connection.to] = circuits.nextId
		circuits.nextId++
	}
}

func partOne(input string, iterations int) (int64, error) {
	boxes, err := parse(input)
	if err != nil {
		return 0, err
	}

	connections := findConnectionsByDistance(boxes)

	circuits := newCircuits()
	for g := 0; g < iterations; g++ {
		shortest := connections[g]
		circuits.connect(shortest)
	}

	sizes := make([]int, len(circuits.circuitsById))
	for _, group := range circuits.circuitsById {
		sizes = append(sizes, group.Size())
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return int64(sizes[0] * sizes[1] * sizes[2]), nil
}

func partTwo(input string) (int64, error) {
	boxes, err := parse(input)
	if err != nil {
		return 0, err
	}

	connections := findConnectionsByDistance(boxes)
	circuits := newCircuits()
	for g := 0; g < len(connections); g++ {
		shortest := connections[g]
		circuits.connect(shortest)
	}

	return int64(circuits.firstToFullyCompleteCircuit.from.x * circuits.firstToFullyCompleteCircuit.to.x), nil
}

func parse(input string) ([]JunctionBox, error) {
	lines := strings.Split(input, "\n")
	boxes := make([]JunctionBox, len(lines))

	for i, line := range lines {
		coordinates := strings.Split(line, ",")
		x, err := strconv.Atoi(coordinates[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(coordinates[1])
		if err != nil {
			return nil, err
		}
		z, err := strconv.Atoi(coordinates[2])
		if err != nil {
			return nil, err
		}

		boxes[i] = JunctionBox{x, y, z}
	}

	return boxes, nil
}

func findConnectionsByDistance(boxes []JunctionBox) []Connection {
	connections := make([]Connection, ((len(boxes)-1)*(len(boxes)))/2)

	b := 0
	for f := 0; f < len(boxes); f++ {
		for t := f + 1; t < len(boxes); t++ {
			if f != t {
				connections[b] = newConnection(boxes[f], boxes[t])
				b++
			}
		}
	}
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].dist < connections[j].dist
	})

	return connections
}
