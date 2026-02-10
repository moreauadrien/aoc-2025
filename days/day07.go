package days

import (
	"slices"
	"strconv"
	"strings"

	ds "github.com/moreauadrien/aoc-2025/datastructures"
)

func (d Days) Day07(input string) (string, string) {
	basePos, splitters := day07_parse(input)

	sol_part1 := day07_part1(basePos, splitters)
	sol_part2 := day07_part2(basePos, splitters)

	return strconv.Itoa(sol_part1), strconv.Itoa(sol_part2)
}

func day07_parse(input string) (int, [][]int) {
	lines := strings.Split(input, "\n")

	basePos := strings.Index(lines[0], "S")

	if basePos == -1 {
		panic("Char 'S' not found on first line")
	}

	splitters := make([][]int, 0, len(lines)-1)

	for i := 1; i < len(lines); i++ {
		s := []int{}

		for j, c := range lines[i] {
			if c == '^' {
				s = append(s, j)
			}
		}

		splitters = append(splitters, s)
	}

	return basePos, splitters
}

func day07_part1(basePos int, splitters [][]int) int {
	splitCount := 0

	beams := ds.NewSet[int]()
	beams.Add(basePos)

	for _, l := range splitters {
		nextBeams := ds.NewSet[int]()

		for b := range beams.All() {
			if slices.Contains(l, b) {
				nextBeams.Add(b - 1)
				nextBeams.Add(b + 1)
				splitCount += 1
			} else {
				nextBeams.Add(b)
			}
		}

		beams = nextBeams
	}

	return splitCount
}

func day07_part2(basePos int, splitters [][]int) int {
	beams := map[int]int{basePos: 1}

	for _, l := range splitters {
		nextBeams := map[int]int{}

		for b, w := range beams {
			if slices.Contains(l, b) {
				v, ok := nextBeams[b-1]
				if ok {
					nextBeams[b-1] = v + w
				} else {
					nextBeams[b-1] = w
				}

				v, ok = nextBeams[b+1]
				if ok {
					nextBeams[b+1] = v + w
				} else {
					nextBeams[b+1] = w
				}
			} else {
				v, ok := nextBeams[b]
				if ok {
					nextBeams[b] = v + w
				} else {
					nextBeams[b] = w
				}
			}
		}

		beams = nextBeams
	}

	res := 0

	for _, v := range beams {
		res += v
	}

	return res
}
