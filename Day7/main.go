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

type Line struct {
	target uint64
	nums   []uint64
}

type Lines []Line

type Op func(uint64, uint64) uint64

func parseLines(r io.Reader) Lines {
	var lines Lines
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var line Line

		str := scanner.Text()
		ind := strings.Index(str, ": ")

		line.target, _ = strconv.ParseUint(str[:ind], 10, 64)
		for _, ele := range strings.Split(str[ind+2:], " ") {
			num, _ := strconv.ParseUint(ele, 10, 64)
			line.nums = append(line.nums, num)
		}

		lines = append(lines, line)
	}

	return lines
}

func (line Line) check(ops []Op) bool {
	reachableSet := make(map[uint64]bool)

	reachableSet[line.nums[0]] = true

	for ind, ele := range line.nums {
		if ind == 0 {
			continue
		}

		newSet := make(map[uint64]bool)
		for num := range reachableSet {
			for _, op := range ops {
				nextNum := op(num, ele)
				if nextNum <= line.target {
					newSet[nextNum] = true
				}
			}
		}
		reachableSet = newSet
	}

	return reachableSet[line.target]
}

func part1(lines Lines) uint64 {
	ops := []Op{
		func(u1, u2 uint64) uint64 { return u1 + u2 },
		func(u1, u2 uint64) uint64 { return u1 * u2 },
	}

	var ans uint64
	ans = 0

	for _, line := range lines {
		if line.check(ops) {
			ans += line.target
		}
	}

	return ans
}

func part2(lines Lines) uint64 {
	ops := []Op{
		func(u1, u2 uint64) uint64 { return u1 + u2 },
		func(u1, u2 uint64) uint64 { return u1 * u2 },
		func(u1, u2 uint64) uint64 {
			ans, _ := strconv.ParseUint(fmt.Sprintf("%d%d", u1, u2), 10, 64)
			return ans
		},
	}

	var ans uint64
	ans = 0

	for _, line := range lines {
		if line.check(ops) {
			ans += line.target
		}
	}

	return ans
}

func main() {
	filename := "text.txt"
	file, _ := os.Open(filename)
	lines := parseLines(file)

	start1 := time.Now()
	ans1 := part1(lines)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	ans2 := part2(lines)
	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1 ans: %d time: %s\n", ans1, elapsed1)
	fmt.Printf("Part 2 ans: %d time: %s\n", ans2, elapsed2)
}
