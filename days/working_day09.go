package days

/*
package days

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	u "github.com/moreauadrien/aoc-2025/utils"
)

type Point struct {
	x, y int
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

type SegmentDB struct {
	horizontal_segments []Segment
	vertical_segments   []Segment
	sorted              bool
}

func (db *SegmentDB) Insert(s Segment) error {
	switch {
	case s.isVertical():
		db.vertical_segments = append(db.vertical_segments, s)
		db.sorted = false
	case s.isHorizontal():
		db.horizontal_segments = append(db.horizontal_segments, s)
		db.sorted = false
	default:
		return errors.New("segments should be either vertical or horizontal")
	}
	return nil
}

func (db *SegmentDB) sort() {
	if !db.sorted {
		sort.Slice(db.horizontal_segments, func(i, j int) bool {
			return db.horizontal_segments[i].start.y < db.horizontal_segments[j].start.y
		})
		sort.Slice(db.vertical_segments, func(i, j int) bool {
			return db.vertical_segments[i].start.x < db.vertical_segments[j].start.x
		})
		db.sorted = true
	}
}

// CountIntersections compte les intersections entre le segment s (qui doit être un bord du rectangle)
// et les segments de la boucle, en ignorant les points rouges.
func (db *SegmentDB) CountIntersections(s Segment, redSet map[Point]bool) (int, error) {
	db.sort()

	if s.isVertical() {
		// s est vertical : on cherche les segments horizontaux qui intersectent
		searchUniverse := db.horizontal_segments
		if len(searchUniverse) == 0 {
			return 0, nil
		}
		x := s.start.x
		yMin := u.IntMin(s.start.y, s.end.y)
		yMax := u.IntMax(s.start.y, s.end.y)

		// segments horizontaux dont la coordonnée y (start.y) est dans [yMin, yMax]
		i := sort.Search(len(searchUniverse), func(k int) bool {
			return searchUniverse[k].start.y >= yMin
		})
		j := sort.Search(len(searchUniverse), func(k int) bool {
			return searchUniverse[k].start.y > yMax
		}) - 1

		if i > j || i < 0 || j >= len(searchUniverse) {
			return 0, nil
		}

		count := 0
		for _, seg := range searchUniverse[i : j+1] {
			xMin := u.IntMin(seg.start.x, seg.end.x)
			xMax := u.IntMax(seg.start.x, seg.end.x)
			if xMin <= x && x <= xMax {
				p := Point{x, seg.start.y}
				if !redSet[p] {
					count++
				}
			}
		}
		return count, nil

	} else if s.isHorizontal() {
		// s est horizontal : on cherche les segments verticaux
		searchUniverse := db.vertical_segments
		if len(searchUniverse) == 0 {
			return 0, nil
		}
		y := s.start.y
		xMin := u.IntMin(s.start.x, s.end.x)
		xMax := u.IntMax(s.start.x, s.end.x)

		i := sort.Search(len(searchUniverse), func(k int) bool {
			return searchUniverse[k].start.x >= xMin
		})
		j := sort.Search(len(searchUniverse), func(k int) bool {
			return searchUniverse[k].start.x > xMax
		}) - 1

		if i > j || i < 0 || j >= len(searchUniverse) {
			return 0, nil
		}

		count := 0
		for _, seg := range searchUniverse[i : j+1] {
			yMinSeg := u.IntMin(seg.start.y, seg.end.y)
			yMaxSeg := u.IntMax(seg.start.y, seg.end.y)
			if yMinSeg <= y && y <= yMaxSeg {
				p := Point{seg.start.x, y}
				if !redSet[p] {
					count++
				}
			}
		}
		return count, nil
	}
	return -1, errors.New("segment must be vertical or horizontal")
}

// isInside détermine si un point est à l'intérieur ou sur le bord de la région verte définie par la boucle.
func isInside(p Point, pts []Point, db *SegmentDB) bool {
	// Vérifier si le point est sur un segment de la boucle
	for _, seg := range db.horizontal_segments {
		if seg.start.y == p.y && p.x >= u.IntMin(seg.start.x, seg.end.x) && p.x <= u.IntMax(seg.start.x, seg.end.x) {
			return true
		}
	}
	for _, seg := range db.vertical_segments {
		if seg.start.x == p.x && p.y >= u.IntMin(seg.start.y, seg.end.y) && p.y <= u.IntMax(seg.start.y, seg.end.y) {
			return true
		}
	}

	// Ray casting (demi-droite horizontale vers la droite)
	inside := false
	n := len(pts)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		xi, yi := pts[i].x, pts[i].y
		xj, yj := pts[j].x, pts[j].y
		if yi == yj {
			continue // segment horizontal ignoré
		}
		if (yi > p.y) != (yj > p.y) {
			// Calcul de l'intersection en x
			x := float64(xi) + float64(p.y-yi)*float64(xj-xi)/float64(yj-yi)
			if x >= float64(p.x) {
				inside = !inside
			}
		}
	}
	return inside
}

func day09_parse(input string) []Point {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	pts := make([]Point, 0, len(lines))
	for _, l := range lines {
		parts := strings.Split(l, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		pts = append(pts, Point{x, y})
	}
	return pts
}

func day09_part1(pts []Point) int {
	maxArea := 0
	l := len(pts)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			area := (1 + u.IntAbs(pts[i].x-pts[j].x)) * (1 + u.IntAbs(pts[i].y-pts[j].y))
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func day09_part2(pts []Point) int {
	maxArea := 0
	l := len(pts)

	// Construire la base de segments de la boucle
	db := SegmentDB{}
	for i := 0; i < l; i++ {
		a := pts[i]
		b := pts[(i+1)%l]
		db.Insert(newSegment(a.x, a.y, b.x, b.y))
	}

	// Ensemble des points rouges
	redSet := make(map[Point]bool)
	for _, p := range pts {
		redSet[p] = true
	}

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			a := pts[i]
			b := pts[j]
			leftX := u.IntMin(a.x, b.x)
			rightX := u.IntMax(a.x, b.x)
			topY := u.IntMin(a.y, b.y)
			botY := u.IntMax(a.y, b.y)

			// Cas du point unique
			if leftX == rightX && topY == botY {
				if 1 > maxArea {
					maxArea = 1
				}
				continue
			}

			// Vérifier les deux autres coins (ceux qui ne sont pas a ou b)
			otherCorners := []Point{
				{leftX, topY},
				{leftX, botY},
				{rightX, topY},
				{rightX, botY},
			}
			valid := true
			for _, p := range otherCorners {
				if p == a || p == b {
					continue
				}
				if !isInside(p, pts, &db) {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			// Vérifier les quatre bords (sans les coins)
			segs := make([]Segment, 0, 4)
			if topY+1 <= botY-1 {
				segs = append(segs, newSegment(leftX, topY+1, leftX, botY-1))
				segs = append(segs, newSegment(rightX, topY+1, rightX, botY-1))
			}
			if leftX+1 <= rightX-1 {
				segs = append(segs, newSegment(leftX+1, topY, rightX-1, topY))
				segs = append(segs, newSegment(leftX+1, botY, rightX-1, botY))
			}
			for _, s := range segs {
				c, err := db.CountIntersections(s, redSet)
				if err != nil {
					valid = false
					break
				}
				if c > 0 {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			area := (rightX - leftX + 1) * (botY - topY + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func (d Days) Day09(input string) (string, string) {
	pts := day09_parse(input)
	sol_part1 := day09_part1(pts)
	sol_part2 := day09_part2(pts)

	return strconv.Itoa(sol_part1), strconv.Itoa(sol_part2)
}
*/
