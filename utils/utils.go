package utils

import "math"

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Assert(b bool, msg string) {
	if b == false {
		panic(msg)
	}
}

func IntAbs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func IntMax(n ...int) int {
	m := math.MinInt

	for _, e := range n {
		if e > m {
			m = e
		}
	}

	return m
}

func IntMin(n ...int) int {
	m := math.MaxInt

	for _, e := range n {
		if e < m {
			m = e
		}
	}

	return m
}
