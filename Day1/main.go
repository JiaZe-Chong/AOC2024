package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"time"
)

func readInts(r io.Reader) ([]int, []int, error) {
	scanner := bufio.NewScanner(r)
	var left []int
	var right []int

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return left, right, err
		}

		if len(right) >= len(left) {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	return left, right, nil
}

func part1(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	ans := 0

	for ind, element := range left {
		diff := element - right[ind]
		if diff > 0 {
			ans += diff
		} else {
			ans += -diff
		}
	}

	return ans
}

func part2(left []int, right []int) int {
	m := make(map[int]int)

	for _, element := range right {
		m[element]++
	}

	ans := 0

	for _, element := range left {
		ans += m[element] * element
	}

	return ans
}

func main() {
	filename := "text.txt"
	file, er := os.Open(filename)
	if er != nil {
		panic(er)
	}

	left, right, err := readInts(file)
	if err != nil {
		panic(err)
	}

	start1 := time.Now()
	ans1 := part1(left, right)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	ans2 := part2(left, right)
	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1 ans: %d took: %s\n", ans1, elapsed1)
	fmt.Printf("Part 2 ans: %d took: %s\n", ans2, elapsed2)
}
