package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"time"
)

type Coord struct {
	x int
	y int
}

func parseGrid(r io.Reader) [][]rune {
	var grid [][]rune

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	return grid
}

func toRegions(grid [][]rune) map[string][]Coord {
	diffX := []int{0, 0, -1, 1}
	diffY := []int{-1, 1, 0, 0}
	diff := 4

	numRegion := make(map[rune]int)
	regions := make(map[string][]Coord)

	for y, row := range grid {
		for x, ch := range row {
			if ch != '.' {

				str := string(ch) + strconv.Itoa(numRegion[ch])
				numRegion[ch]++

				var stack []Coord
				coord := Coord{x: x, y: y}
				stack = append(stack, coord)
				regions[str] = append(regions[str], coord)

				for len(stack) != 0 {
					coord := stack[len(stack)-1]
					stack = stack[:len(stack)-1]

					grid[coord.y][coord.x] = '.'

					for i := 0; i < diff; i++ {
						newCoord := Coord{x: coord.x + diffX[i], y: coord.y + diffY[i]}
						if newCoord.x < 0 || newCoord.x >= len(grid[0]) || newCoord.y < 0 || newCoord.y >= len(grid) {
							continue
						}
						if grid[newCoord.y][newCoord.x] != ch {
							continue
						}
						if slices.Contains(regions[str], newCoord) {
							continue
						}

						regions[str] = append(regions[str], newCoord)
						stack = append(stack, newCoord)
					}

				}

			}
		}
	}
	return regions
}

func calculatePerimeter(region []Coord) int {
	perimeter := 0

	diffX := []int{0, 0, -1, 1}
	diffY := []int{-1, 1, 0, 0}
	diff := 4

	for _, coord := range region {
		for i := 0; i < diff; i++ {
			newCoord := Coord{x: coord.x + diffX[i], y: coord.y + diffY[i]}
			if !slices.Contains(region, newCoord) {
				perimeter++
			}
		}
	}

	return perimeter
}

func part1(regions map[string][]Coord) int {
	score := 0
	for _, region := range regions {
		score += calculatePerimeter(region) * len(region)
	}
	return score
}

type Perimeter struct {
	coord Coord
	diffX int
	diffY int
}

func calculateSides(region []Coord) int {
	diffX := []int{0, 0, -1, 1}
	diffY := []int{-1, 1, 0, 0}
	diff := 4

	var perimeters []Perimeter

	for _, coord := range region {
		for i := 0; i < diff; i++ {
			newCoord := Coord{x: coord.x + diffX[i], y: coord.y + diffY[i]}
			if !slices.Contains(region, newCoord) {
				perimeters = append(perimeters, Perimeter{coord: coord, diffX: diffX[i], diffY: diffY[i]})
			}
		}
	}

	set := make(map[Perimeter]bool)
	sides := 0

	diffSide := []int{-1, 1}

	for _, perimeter := range perimeters {
		if !set[perimeter] {
			sides++
			set[perimeter] = true
			if perimeter.diffX == 0 {
				for _, d := range diffSide {
					perimeter.coord.x = perimeter.coord.x + d
					for slices.Contains(perimeters, perimeter) {
						set[perimeter] = true
						perimeter.coord.x = perimeter.coord.x + d
					}
				}
			} else {
				for _, d := range diffSide {
					perimeter.coord.y = perimeter.coord.y + d
					for slices.Contains(perimeters, perimeter) {
						set[perimeter] = true
						perimeter.coord.y = perimeter.coord.y + d
					}
				}
			}
		}
	}

	return sides
}

func part2(regions map[string][]Coord) int {
	score := 0
	for _, region := range regions {
		score += calculateSides(region) * len(region)
	}
	return score
}

func main() {
	filename := "text.txt"
	f, _ := os.Open(filename)
	grid := parseGrid(f)

	start := time.Now()
	regions := toRegions(grid)
	elapsed := time.Since(start)
	fmt.Printf("Region calculation time: %s\n", elapsed)

	start1 := time.Now()
	ans1 := part1(regions)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	ans2 := part2(regions)
	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1 ans: %d time: %s\n", ans1, elapsed1)
	fmt.Printf("Part 2 ans: %d time: %s\n", ans2, elapsed2)
}
