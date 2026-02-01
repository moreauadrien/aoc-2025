package days

import (
	"strconv"
	"strings"
)

type Grid struct {
	w       int
	h       int
	content [][]byte
}

func (g Grid) inBound(x int, y int) bool {
	return (x >= 0) && (x < g.w) && (y >= 0) && (y < g.h)
}

func (d Days) Day04(input string) (string, string) {
	grid := day04_parse(input)
	sol_part1 := day04_part1(grid)
	sol_part2 := day04_part2(grid)

	return strconv.Itoa(sol_part1), strconv.Itoa(sol_part2)
}

func day04_parse(input string) Grid {
	lines := strings.Split(input, "\n")
	h := len(lines)
	w := len(lines[0])

	content := make([][]byte, 0, w)
	for _ = range w {
		content = append(content, make([]byte, h))
	}

	for y, l := range lines {
		for x, c := range []byte(l) {
			content[x][y] = c
		}
	}

	return Grid{w, h, content}
}

var ADJ_8 [8][2]int = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func day04_part1(grid Grid) int {
	count := 0

	for x := range grid.w {
		for y := range grid.h {
			if grid.content[x][y] == '@' {
				if day04_count_adj_paper(grid, x, y) < 4 {
					count += 1
				}
			}
		}
	}

	return count
}

var VISITED_ADJ [4][2]int = [4][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}}

func day04_part2(grid Grid) int {
	count := 0

	q := make([][2]int, 5)
	q[0] = [2]int{0, 0}

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]

		if grid.content[x][y] == '@' {
			if day04_count_adj_paper(grid, x, y) < 4 {
				grid.content[x][y] = '.'
				count += 1

				for _, v_adj := range VISITED_ADJ {
					adj_x := x + v_adj[0]
					adj_y := x + v_adj[1]

					if grid.inBound(adj_x, adj_y) {
						q = append(q, [2]int{adj_x, adj_y})
					}
				}

			}
		}

		next_x := x + 1
		next_y := y

		if next_x >= grid.w {
			next_x = 0
			next_y += 1
		}

		if next_y < grid.h {
			q = append(q, [2]int{next_x, next_y})
		}
	}

	return count
}

func day04_count_adj_paper(grid Grid, x int, y int) int {
	adj_count := 0
	for _, adj := range ADJ_8 {
		adj_x := x + adj[0]
		adj_y := y + adj[1]

		if grid.inBound(adj_x, adj_y) == false {
			continue
		}

		if grid.content[adj_x][adj_y] == '@' {
			adj_count += 1
		}
	}

	return adj_count
}
