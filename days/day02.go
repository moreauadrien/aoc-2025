package days

import (
	"fmt"
	"strconv"
	"strings"
)

type idRange struct {
	start int
	end   int
}

func (d Days) Day02(input string) (string, string) {
	ranges := day02_parse(input)
	sol_part1 := day02(ranges, isValidId_part1)
	sol_part2 := day02(ranges, isValidId_part2)

	return strconv.Itoa(sol_part1), strconv.Itoa(sol_part2)
}

func day02_parse(input string) []idRange {
	str_ranges := strings.Split(input, ",")
	ranges := make([]idRange, 0, len(str_ranges))

	for _, r := range str_ranges {
		parts := strings.Split(r, "-")

		s, err_s := strconv.Atoi(parts[0])
		e, err_e := strconv.Atoi(parts[1])

		if err_s != nil || err_e != nil {
			panic(fmt.Sprintf("Can't parse the range '%v'", r))
		}

		ranges = append(ranges, idRange{s, e})
	}

	return ranges
}

func isValidId_part1(n int) bool {
	n_str := strconv.Itoa(n)
	l := len(n_str)

	if l%2 != 0 {
		return true
	}

	return n_str[0:l/2] != n_str[l/2:l]
}

func day02(ranges []idRange, isValidId func(int) bool) int {
	count := 0

	for _, r := range ranges {
		for i := r.start; i <= r.end; i += 1 {
			if isValidId(i) == false {
				count += i
			}
		}
	}

	return count
}

func isValidId_part2(n int) bool {
	n_str := strconv.Itoa(n)
	l := len(n_str)

outer:
	for win_size := 1; win_size <= l/2; win_size += 1 {
		if l%win_size != 0 {
			continue
		}

		patt := n_str[0:win_size]

		for w := win_size; w+win_size <= l; w += win_size {
			if n_str[w:w+win_size] != patt {
				continue outer
			}
		}

		return false
	}

	return true
}
