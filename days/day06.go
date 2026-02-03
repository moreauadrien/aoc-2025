package days

import (
	"strconv"
	"strings"
)

func (d Days) Day06(input string) (string, string) {
	calcs1 := day06_parse(input)
	sol_part1 := day06_part1(calcs1)
	calcs2 := day06_parse_part2(input)
	sol_part2 := day06_part1(calcs2)

	return strconv.Itoa(sol_part1), strconv.Itoa(sol_part2)
}

type Calc struct {
	op       byte
	operands []int
}

func day06_parse(input string) []Calc {
	lines := strings.Split(input, "\n")
	nb_op := len(lines) - 1

	calcs := make([]Calc, 0)

	for i := 0; i < len(lines)-1; i++ {
		l := lines[i]
		j := 0
		for o := range strings.SplitSeq(l, " ") {
			if len(o) == 0 {
				continue
			}

			if i == 0 {
				ops := make([]int, 0, nb_op)

				n, err := strconv.Atoi(o)
				if err != nil {
					panic(err)
				}

				ops = append(ops, n)
				calcs = append(calcs, Calc{' ', ops})
			} else {
				n, err := strconv.Atoi(o)
				if err != nil {
					panic(err)
				}

				calc := calcs[j]
				calc.operands = append(calc.operands, n)

				calcs[j] = calc

				j += 1
			}
		}
	}

	j := 0

	for _, o := range []byte(lines[len(lines)-1]) {
		if o == ' ' {
			continue
		}

		calc := calcs[j]
		calc.op = o

		calcs[j] = calc

		j += 1
	}

	return calcs
}

func day06_part1(calcs []Calc) int {
	sum := 0

	for _, calc := range calcs {
		var res int
		if calc.op == '+' {
			res = 0
		} else if calc.op == '*' {
			res = 1
		}

		for _, o := range calc.operands {
			if calc.op == '+' {
				res += o
			} else if calc.op == '*' {
				res *= o
			}
		}

		sum += res
	}

	return sum
}

func day06_parse_part2(input string) []Calc {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return nil
	}

	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	paddedLines := make([]string, len(lines))
	for i, line := range lines {
		if len(line) < maxWidth {
			paddedLines[i] = line + strings.Repeat(" ", maxWidth-len(line))
		} else {
			paddedLines[i] = line
		}
	}

	calcs := make([]Calc, 0)
	var currentCalc *Calc
	inNumber := false

	for col := maxWidth - 1; col >= 0; col-- {
		columnDigits := ""
		for row := 0; row < len(paddedLines)-1; row++ {
			if col < len(paddedLines[row]) {
				columnDigits += string(paddedLines[row][col])
			} else {
				columnDigits += " "
			}
		}

		var opChar byte = ' '
		if col < len(paddedLines[len(paddedLines)-1]) {
			opChar = paddedLines[len(paddedLines)-1][col]
		}

		hasDigit := false
		for _, c := range columnDigits {
			if c >= '0' && c <= '9' {
				hasDigit = true
				break
			}
		}

		if hasDigit {
			if !inNumber {
				inNumber = true
				if currentCalc == nil {
					currentCalc = &Calc{op: ' ', operands: make([]int, 0)}
				}
			}

			numStr := ""
			for _, c := range columnDigits {
				if c >= '0' && c <= '9' {
					numStr += string(c)
				}
			}

			if len(numStr) > 0 {
				n, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				currentCalc.operands = append(currentCalc.operands, n)
			}

			if opChar == '+' || opChar == '*' {
				currentCalc.op = opChar
			}
		} else {
			if inNumber && currentCalc != nil {
				calcs = append(calcs, *currentCalc)
				currentCalc = nil
				inNumber = false
			}
		}
	}

	if currentCalc != nil && len(currentCalc.operands) > 0 {
		calcs = append(calcs, *currentCalc)
	}

	return calcs
}
