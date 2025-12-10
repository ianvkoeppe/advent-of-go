package nine

import (
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Segment struct {
	start, end Point
}

func (s Segment) intersects(segment Segment) bool {
	o1 := orient(s.start, s.end, segment.start)
	o2 := orient(s.start, s.end, segment.end)
	o3 := orient(segment.start, segment.end, s.start)
	o4 := orient(segment.start, segment.end, s.end)
	return o1 != 0 && o2 != 0 && o3 != 0 && o4 != 0 && o1 != o2 && o3 != o4
}

func orient(a, b, c Point) int {
	cross := (b.x-a.x)*(c.y-a.y) - (b.y-a.y)*(c.x-a.x)

	if cross > 0 {
		return 1
	}
	if cross < 0 {
		return -1
	}
	return 0
}

func newSegment(start, end Point) Segment {
	if start.x == end.x {
		if start.y > end.y {
			start.y, end.y = end.y, start.y
		}
	} else if start.y == end.y {
		if start.x > end.x {
			start.x, end.x = end.x, start.x
		}
	}

	return Segment{start, end}
}

func (s Segment) contains(p Point) bool {
	if s.start.x == s.end.x {
		return p.x == s.start.x && s.start.y <= p.y && p.y <= s.end.y
	} else {
		return p.y == s.start.y && s.start.x <= p.x && p.x <= s.end.x
	}
}

type Rectangle struct {
	minX, maxX, minY, maxY int
}

func newRectangle(f, s Point) Rectangle {
	if f.x > s.x {
		f.x, s.x = s.x, f.x
	}
	if f.y > s.y {
		f.y, s.y = s.y, f.y
	}
	return Rectangle{f.x, s.x, f.y, s.y}
}

func (r Rectangle) corners() []Point {
	return []Point{{r.minX, r.minY}, {r.minX, r.maxY}, {r.maxX, r.minY}, {r.maxX, r.maxY}}
}

func (r Rectangle) area() float64 {
	return float64((r.maxY - r.minY + 1) * (r.maxX - r.minX + 1))
}

type Polygon struct {
	edges []Segment
}

func newPolygon(points []Point) Polygon {
	edges := make([]Segment, 0)
	for t := range points {
		o := t + 1
		if o == len(points) {
			o = 0
		}
		edges = append(edges, newSegment(points[t], points[o]))
	}
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].start.x == edges[j].start.x {
			return edges[i].start.y < edges[j].start.y
		}
		return edges[i].start.x > edges[j].start.x // Descending for ray casting.
	})

	return Polygon{edges}
}

func (p *Polygon) containsRectangle(rectangle Rectangle) bool {
	for _, corner := range rectangle.corners() {
		if !p.containsPoint(corner) {
			return false
		}
	}

	segments := []Segment{
		{Point{rectangle.minX, rectangle.minY}, Point{rectangle.maxX, rectangle.minY}},
		{Point{rectangle.minX, rectangle.maxY}, Point{rectangle.maxX, rectangle.maxY}},
		{Point{rectangle.minX, rectangle.minY}, Point{rectangle.minX, rectangle.maxY}},
		{Point{rectangle.maxX, rectangle.minY}, Point{rectangle.maxX, rectangle.maxY}},
	}
	for _, edge := range p.edges {
		for _, segment := range segments {
			if segment.intersects(edge) {
				return false
			}
		}
	}

	return true
}

func (p *Polygon) containsPoint(point Point) bool {
	intersections := 0
	for _, seg := range p.edges {
		if seg.contains(point) {
			return true
		}

		if seg.start.x <= point.x {
			break
		}

		if seg.start.y > point.y != (seg.end.y > point.y) {
			t := (point.y - seg.start.y) / (seg.end.y - seg.start.y)
			if seg.start.x+t*(seg.end.x-seg.start.x) > point.x {
				intersections++
			}
		}
	}
	return intersections%2 == 1
}

func partOne(input string) (int64, error) {
	points, err := parse(input)
	if err != nil {
		return 0, err
	}

	return int64(findCandidateRectangles(points)[0].area()), nil
}

func partTwo(input string) (int64, error) {
	points, err := parse(input)
	if err != nil {
		return 0, err
	}

	polygon := newPolygon(points)
	for _, rectangle := range findCandidateRectangles(points) {
		if polygon.containsRectangle(rectangle) {
			return int64(rectangle.area()), nil
		}
	}

	return int64(0), nil
}

func parse(input string) ([]Point, error) {
	lines := strings.Split(input, "\n")
	tiles := make([]Point, len(lines))
	for l, line := range lines {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		tiles[l].x = x

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		tiles[l].y = y
	}
	return tiles, nil
}

func findCandidateRectangles(points []Point) []Rectangle {
	candidates := make([]Rectangle, len(points)*(len(points)-1)/2)
	for t := range points {
		for o := t + 1; o < len(points); o++ {
			candidates[t*(len(points)-1)-(t*(t-1))/2+(o-t-1)] = newRectangle(points[t], points[o])
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].area() > candidates[j].area()
	})

	return candidates
}
