package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Coord struct {
	x int
	y int
}

func (coord Coord) valid(m, n int) bool {
	return coord.x >= 0 && coord.x < m && coord.y >= 0 && coord.y < m
}

func parseGrid(r io.Reader) (map[rune][]Coord, int, int) {
	ret := make(map[rune][]Coord)

	scanner := bufio.NewScanner(r)
	y := 0
	maxX := 0
	for scanner.Scan() {
		maxX = len(scanner.Text())
		for x, ele := range scanner.Text() {
			if ele != '.' {
				var coord Coord
				coord.x, coord.y = x, y
				ret[ele] = append(ret[ele], coord)
			}
		}
		y++
	}

	return ret, maxX, y
}

func part1(grid map[rune][]Coord, m, n int) int {
	set := make(map[Coord]bool)

	for _, coords := range grid {
		for i, base := range coords {
			for j, arm := range coords {
				if i == j {
					continue
				}
				diffX := arm.x - base.x
				diffY := arm.y - base.y
				var antinode Coord
				antinode.x, antinode.y = arm.x+diffX, arm.y+diffY
				if antinode.valid(m, n) {
					set[antinode] = true
				}
			}
		}
	}

	return len(set)
}

func part2(grid map[rune][]Coord, m, n int) int {
	set := make(map[Coord]bool)

	for _, coords := range grid {
		for i, base := range coords {
			for j, arm := range coords {
				if i == j {
					continue
				}
				diffX := arm.x - base.x
				diffY := arm.y - base.y
				antinode := arm
				for antinode.valid(m, n) {
					set[antinode] = true
					antinode.x, antinode.y = antinode.x+diffX, antinode.y+diffY
				}
			}
		}
	}

	return len(set)
}

func main() {
	filename := "text.txt"
	file, _ := os.Open(filename)

	grid, m, n := parseGrid(file)

	ans1 := part1(grid, m, n)
	ans2 := part2(grid, m, n)

	fmt.Printf("Part 1: %d\n", ans1)
	fmt.Printf("Part 2: %d\n", ans2)
}
