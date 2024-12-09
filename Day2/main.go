package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func abs(i int) int {
	return max(i, -i)
}

func readFile(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)
	var reports [][]int

	for scanner.Scan() {
		var report []int
		for _, element := range strings.Split(scanner.Text(), " ") {
			level, err := strconv.Atoi(element)
			if err != nil {
				return reports, err
			}
			report = append(report, level)
		}
		reports = append(reports, report)
	}

	return reports, nil
}

func isSafe(report []int) bool {
	last_diff := 0

	for ind, level := range report {
		if ind == 0 {
			continue
		}

		curr_diff := level - report[ind-1]
		if curr_diff*last_diff < 0 {
			return false
		}
		if abs(curr_diff) < 1 || abs(curr_diff) > 3 {
			return false
		}

		last_diff = curr_diff
	}

	return true
}

func part1(reports [][]int) int {
	ans := 0

	for _, report := range reports {
		if isSafe(report) {
			ans++
		}
	}

	return ans
}

func part2(reports [][]int) int {
	ans := 0

	for _, report := range reports {
		if isSafe(report) {
			ans++
			continue
		}

		for ind := range report {
			temp := make([]int, len(report))
			copy(temp, report)
			temp = append(temp[:ind], temp[ind+1:]...)
			if isSafe(temp) {
				ans++
				break
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

	reports, err := readFile(file)
	if err != nil {
		panic(err)
	}

	start1 := time.Now()
	ans1 := part1(reports)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	ans2 := part2(reports)
	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1 ans: %d time: %s\n", ans1, elapsed1)
	fmt.Printf("Part 2 ans: %d time: %s\n", ans2, elapsed2)
}
