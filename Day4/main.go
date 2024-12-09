package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func readFile(r io.Reader) []string {
	var ret []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}

func part1(grid []string) int {
	ans := 0

	m := len(grid)
	n := len(grid[0])
	delta_x := []int{0, 0, 1, 1, 1, -1, -1, -1}
	delta_y := []int{1, -1, 0, -1, 1, 0, -1, 1}
	directions := 8
	words := []byte{'X', 'M', 'A', 'S'}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if grid[i][j] == 'X' {
				for dir := 0; dir < directions; dir++ {

					is_word := true
					for k := 0; k < 4; k++ {
						x := i + delta_x[dir]*k
						y := j + delta_y[dir]*k

						if x < 0 || x >= m || y < 0 || y >= n {
							is_word = false
							break
						}

						if grid[x][y] != words[k] {
							is_word = false
							break
						}

					}

					if is_word {
						ans++
					}

				}
			}

		}
	}

	return ans
}

func part2(grid []string) int {
	ans := 0

	m := len(grid)
	n := len(grid[0])
	delta_x := []int{1, 1, -1, -1}
	delta_y := []int{1, -1, 1, -1}
	directions := 4
	words := []byte{'M', 'A', 'S'}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			num_word := 0
			if grid[i][j] == 'A' {
				for dir := 0; dir < directions; dir++ {

					is_word := true
					for k := -1; k < 2; k++ {
						x := i + delta_x[dir]*k
						y := j + delta_y[dir]*k

						if x < 0 || x >= m || y < 0 || y >= n {
							is_word = false
							break
						}

						if grid[x][y] != words[k+1] {
							is_word = false
							break
						}
					}

					if is_word {
						num_word++
					}

				}
			}
			if num_word == 2 {
				ans++
			}
		}
	}

	return ans
}

func main() {
	filename := "text.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	grid := readFile(file)

	start1 := time.Now()
	ans1 := part1(grid)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	ans2 := part2(grid)
	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1 ans: %d time: %s\n", ans1, elapsed1)
	fmt.Printf("Part 2 ans: %d time: %s\n", ans2, elapsed2)

}
