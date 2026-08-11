package days

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	u "github.com/moreauadrien/aoc-2025/utils"
)

func (d Days) Day09(input string) (string, string) {
	pts := day09_parse(input)
	sol_part1 := day09_part1(pts)
	sol_part2 := day09_part2(pts)

	return strconv.Itoa(sol_part1), strconv.Itoa(sol_part2)
}

type Point struct {
	x, y int
}

func day09_parse(input string) []Point {
	lines := strings.Split(input, "\n")

	pts := make([]Point, 0, len(lines))

	for _, l := range lines {
		parts := strings.Split(l, ",")

		x, err := strconv.Atoi(parts[0])
		u.Check(err)

		y, err := strconv.Atoi(parts[1])
		u.Check(err)

		pts = append(pts, Point{x, y})
	}

	return pts
}

func day09_part1(pts []Point) int {
	maxArea := 0
	l := len(pts)

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			ptsA := pts[i]
			ptsB := pts[j]
			t := (1 + u.IntAbs(ptsA.x-ptsB.x)) * (1 + u.IntAbs(ptsA.y-ptsB.y))

			if t > maxArea {
				maxArea = t
			}
		}
	}

	return maxArea
}

type Segment struct {
	start Point
	end   Point
}

func newSegment(startX, startY, endX, endY int) Segment {
	return Segment{start: Point{startX, startY}, end: Point{endX, endY}}
}

func (s Segment) isVertical() bool {
	return s.start.x == s.end.x
}

func (s Segment) isHorizontal() bool {
	return s.start.y == s.end.y
}

type SegmentDS struct {
	horizontal_segments []Segment
	vertical_segments   []Segment
	sorted              bool
}

func (db *SegmentDS) Insert(s Segment) error {

	switch {

	case s.isVertical():
		db.vertical_segments = append(db.vertical_segments, s)
		db.sorted = false

	case s.isHorizontal():
		db.horizontal_segments = append(db.horizontal_segments, s)
		db.sorted = false

	default:
		return errors.New("segments should be either vertical or horizontal to be added to a SegmentDB")

	}

	return nil
}

func (db *SegmentDS) sort() {
	if db.sorted == false {
		sort.Slice(db.horizontal_segments, func(i, j int) bool {
			return db.horizontal_segments[i].start.y < db.horizontal_segments[j].start.y
		})

		sort.Slice(db.vertical_segments, func(i, j int) bool {
			return db.vertical_segments[i].start.x < db.vertical_segments[j].start.x
		})

		db.sorted = true
	}
}

func day09_part2(pts []Point) int {
	return 0
}
