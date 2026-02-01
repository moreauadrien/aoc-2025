package days

import (
	"strconv"
	"strings"
)

func (d Days) Day02(input string) (string, string) {
	moves := parse(input)
	sol_part1 := part1(moves)
	sol_part2 := part2(moves)

	return strconv.Itoa(sol_part1), strconv.Itoa(sol_part2)
}

func parse(input string) []int64 {
	lines := strings.Split(input, "\n")
	moves := make([]int64, 0, len(lines))

	for _, l := range lines {
		left := l[0] == 'L'
		val, err := strconv.ParseInt(l[1:], 10, 0)
		if err != nil {
			panic(err)
		}

		if left {
			val = -val
		}

		moves = append(moves, val)
	}

	return moves
}

func part1(moves []int64) int {
	var dial int64 = 50
	count := 0

	for _, m := range moves {
		dial = (dial + m) % 100
		if dial == 0 {
			count += 1
		}
	}

	return count
}

func part2(moves []int64) int {
	var dial int64 = 50
	count := 0
	for _, m := range moves {
		var step int64 = 1
		if m < 0 {
			step = -1
		}
		for i := int64(0); i != m; i += step {
			dial = ((dial+step)%100 + 100) % 100
			if dial == 0 {
				count++
			}
		}
	}
	return count
}
