package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Coord struct {
	x int
	y int
}

type Coord2 struct {
	start Coord
	end   Coord
}

func (coord Coord) valid(m, n int) bool {
	return coord.x >= 0 && coord.x < m && coord.y >= 0 && coord.y < n
}

func parseGrid(r io.Reader) ([][]int, []Coord) {
	var ret1 [][]int
	var ret2 []Coord

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var cur []int
		for ind, ele := range scanner.Text() {
			num := int(ele - '0')
			if num == 0 {
				var coord Coord
				coord.x, coord.y = ind, len(ret1)
				ret2 = append(ret2, coord)
			}
			cur = append(cur, int(ele-'0'))
		}
		ret1 = append(ret1, cur)
	}

	return ret1, ret2
}

func part1(grid [][]int, starts []Coord) int {
	m := len(grid)
	n := len(grid[0])
	diffX := []int{0, 0, 1, -1}
	diffY := []int{1, -1, 0, 0}

	set := make(map[Coord2]bool)
	for _, ele := range starts {
		var coord2 Coord2
		coord2.start, coord2.end = ele, ele
		set[coord2] = true
	}

	for cur := 1; cur <= 9; cur++ {
		nextSet := make(map[Coord2]bool)

		for coord := range set {
			for i := 0; i < 4; i++ {
				var nextCoord Coord
				nextCoord.x = coord.end.x + diffX[i]
				nextCoord.y = coord.end.y + diffY[i]
				if nextCoord.valid(m, n) && grid[nextCoord.y][nextCoord.x] == cur {
					var coord2 Coord2
					coord2.start, coord2.end = coord.start, nextCoord
					nextSet[coord2] = true
				}
			}
		}

		set = nextSet
	}

	return len(set)
}

func part2(grid [][]int, starts []Coord) int {
	m := len(grid)
	n := len(grid[0])
	diffX := []int{0, 0, 1, -1}
	diffY := []int{1, -1, 0, 0}

	set := make(map[Coord]int)
	for _, ele := range starts {
		set[ele] = 1
	}

	for cur := 1; cur <= 9; cur++ {
		nextSet := make(map[Coord]int)

		for coord, score := range set {
			for i := 0; i < 4; i++ {
				var nextCoord Coord
				nextCoord.x = coord.x + diffX[i]
				nextCoord.y = coord.y + diffY[i]
				if nextCoord.valid(m, n) && grid[nextCoord.y][nextCoord.x] == cur {
					nextSet[nextCoord] += score
				}
			}
		}

		set = nextSet
	}

	ans := 0

	for _, score := range set {
		ans += score
	}

	return ans
}

func main() {
	filename := "text.txt"
	file, _ := os.Open(filename)
	grid, starts := parseGrid(file)

	start1 := time.Now()
	ans1 := part1(grid, starts)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	ans2 := part2(grid, starts)
	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1 ans: %d time: %s\n", ans1, elapsed1)
	fmt.Printf("Part 2 ans: %d time: %s\n", ans2, elapsed2)
}
