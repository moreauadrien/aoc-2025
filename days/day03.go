package days

import (
	"strconv"
	"strings"
)

func (d Days) Day03(input string) (string, string) {
	banks := day03_parse(input)
	sol_part1 := day03(banks, 2)
	sol_part2 := day03(banks, 12)

	return strconv.Itoa(sol_part1), strconv.Itoa(sol_part2)
}

type bank []int8

func day03_parse(input string) []bank {
	lines := strings.Split(input, "\n")

	banks := make([]bank, 0, len(lines))

	for _, l := range lines {
		b := make(bank, 0, len(l))
		for _, c := range []byte(l) {
			b = append(b, int8(c-'0'))
		}

		banks = append(banks, b)
	}

	return banks
}

func day03(banks []bank, n_turn int) int {
	res := 0

	for _, bank := range banks {
		n_bat := len(bank)
		digits := make([]int8, n_turn)

		for i, bat := range bank {
			rem_bat := n_bat - i

			for j, d := range digits {
				rem_turn := n_turn - j
				if bat > d && (rem_bat >= rem_turn) {
					digits[j] = bat

					for k := j + 1; k < n_turn; k++ {
						digits[k] = 0
					}

					break
				}
			}
		}

		n := 0

		for _, d := range digits {
			n = n*10 + int(d)
		}

		res += n
	}

	return res
}
