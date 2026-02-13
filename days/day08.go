package days

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	u "github.com/moreauadrien/aoc-2025/utils"
)

func (d Days) Day08(input string) (string, string) {
	boxes := day08_parse(input)
	sol_part1 := day08_part1(boxes, 1000)
	sol_part2 := day08_part2(boxes)

	return strconv.Itoa(sol_part1), strconv.Itoa(sol_part2)
}

type Box struct {
	x uint32
	y uint32
	z uint32
}

func (b1 Box) square_dist(b2 Box) uint64 {
	dx := int64(b1.x) - int64(b2.x)
	dy := int64(b1.y) - int64(b2.y)
	dz := int64(b1.z) - int64(b2.z)

	return uint64(dx*dx + dy*dy + dz*dz)
}

//func (b Box) compact() uint64 {
//	// coords are in the range of [0; 10000] (17 bits)
//	return uint64(((b.x) << 40) | ((b.y) << 20) | (b.z))
//}

func day08_parse(input string) []Box {
	lines := strings.Split(input, "\n")

	boxes := make([]Box, 0, len(lines))

	for _, l := range lines {
		coords := strings.Split(l, ",")

		x, err := strconv.Atoi(coords[0])
		u.Check(err)

		y, err := strconv.Atoi(coords[1])
		u.Check(err)

		z, err := strconv.Atoi(coords[2])
		u.Check(err)

		boxes = append(boxes, Box{uint32(x), uint32(y), uint32(z)})
	}

	return boxes
}

type Dist struct {
	dist uint64
	b0   int
	b1   int
}

func day08_part1(boxes []Box, nb_conn int) int {
	circuits := [][]int{}
	circuits_lookup := map[int]int{}
	n := len(boxes)
	l := n * (n - 1) / 2
	distances := make([]Dist, 0, l)

	var i int
	var j int

	for i = 0; i < n; i++ {
		b0 := boxes[i]
		for j = i + 1; j < n; j++ {
			b1 := boxes[j]
			distances = append(distances, Dist{b0.square_dist(b1), i, j})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist < distances[j].dist
	})

	for i, _ := range boxes {
		circuits = append(circuits, []int{i})
		circuits_lookup[i] = i
	}

	for i, d := range distances {
		if i == nb_conn {
			break
		}

		c0 := circuits_lookup[d.b0]
		c1 := circuits_lookup[d.b1]

		if c0 == c1 {
			continue
		}

		circuits[c0] = slices.Concat(circuits[c0], circuits[c1])

		for _, b := range circuits[c1] {
			circuits_lookup[b] = c0
		}

		circuits[c1] = nil
	}

	cLen := make([]int, 0, len(circuits))
	for _, c := range circuits {
		cLen = append(cLen, len(c))
	}

	slices.Sort(cLen)

	res := 1

	fmt.Println(cLen)

	for i := len(cLen) - 3; i < len(cLen); i++ {
		res *= cLen[i]
	}

	return res
}

func day08_part2(boxes []Box) int {
	circuits := [][]int{}
	circuits_lookup := map[int]int{}
	n := len(boxes)
	l := n * (n - 1) / 2
	distances := make([]Dist, 0, l)

	var i int
	var j int

	for i = 0; i < n; i++ {
		b0 := boxes[i]
		for j = i + 1; j < n; j++ {
			b1 := boxes[j]
			distances = append(distances, Dist{b0.square_dist(b1), i, j})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist < distances[j].dist
	})

	for i, _ := range boxes {
		circuits = append(circuits, []int{i})
		circuits_lookup[i] = i
	}

	for _, d := range distances {
		c0 := circuits_lookup[d.b0]
		c1 := circuits_lookup[d.b1]

		if c0 == c1 {
			continue
		}

		circuits[c0] = slices.Concat(circuits[c0], circuits[c1])

		if len(circuits[c0]) == len(boxes) {
			return int(boxes[d.b0].x) * int(boxes[d.b1].x)
		}

		for _, b := range circuits[c1] {
			circuits_lookup[b] = c0
		}

		circuits[c1] = nil
	}

	return -1
}
