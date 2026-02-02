package days

import (
	"slices"
	"strconv"
	"strings"
)

func (d Days) Day05(input string) (string, string) {
	ranges, ids := day05_parse(input)

	sol_part1 := day05_part1(ranges, ids)
	sol_part2 := day05_part2(ranges)

	return strconv.Itoa(sol_part1), strconv.Itoa(sol_part2)
}

func day05_parse(input string) ([]Range, []int) {
	parts := strings.Split(input, "\n\n")

	ranges_lines := strings.Split(parts[0], "\n")
	ranges := make([]Range, 0, len(ranges_lines))

	for _, l := range ranges_lines {
		se := strings.Split(l, "-")

		start, err := strconv.Atoi(se[0])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(se[1])
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, Range{start, end + 1})
	}

	id_lines := strings.Split(parts[1], "\n")
	ids := make([]int, 0, len(id_lines))

	for _, l := range id_lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}

		ids = append(ids, n)
	}

	return ranges, ids
}

func day05_part1(ranges []Range, ids []int) int {
	count_fresh := 0
	for _, r := range ranges {
		i := 0
		for i < len(ids) {
			if r.contains(ids[i]) {
				count_fresh += 1
				ids = slices.Delete(ids, i, i+1)
			} else {
				i += 1
			}
		}
	}

	return count_fresh
}

func day05_part2(ranges []Range) int {
	var s Set = ranges[0]

	for _, r := range ranges {
		s = s.join(r)
	}

	return s.len()
}

type Set interface {
	len() int
	join(Set) Set
	contains(int) bool
}

type Range struct {
	start int
	end   int
}

type Union struct {
	elems []Range
}

func (u Union) len() int {
	sum := 0

	for _, e := range u.elems {
		sum += e.len()
	}

	return sum
}

func (u Union) join(s Set) Set {
	if r1, isRange := s.(Range); isRange {
		for i, r2 := range u.elems {
			res := r1.join(r2)
			if _, isRange := res.(Range); isRange {
				u.elems = slices.Delete(u.elems, i, i+1)
				return u.join(res)
			}
		}

		u.elems = append(u.elems, r1)
		return u
	} else {
		u2 := s.(Union)

		for _, r := range u2.elems {
			u.join(r)
		}

		return u
	}
}

func (u Union) contains(n int) bool {
	for _, r := range u.elems {
		if r.contains(n) {
			return true
		}
	}

	return false
}

func (r Range) len() int {
	return r.end - r.start
}

func (r Range) contains(n int) bool {
	return r.start <= n && n < r.end
}

func (r1 Range) join(s Set) Set {
	if _, isUnion := s.(Union); isUnion {
		return s.join(r1)
	}

	r2 := s.(Range)

	if r2.start < r1.start {
		r1, r2 = r2, r1
	}

	if r1.end < r2.start {
		return Union{[]Range{r1, r2}}
	}

	if r1.end < r2.end {
		return Range{r1.start, r2.end}
	}

	return Range{r1.start, r1.end}
}
