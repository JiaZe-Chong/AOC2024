package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"
)

func readFile(r io.Reader) []string {
	var ans []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ans = append(ans, scanner.Text())
	}

	return ans
}

func part1(strings []string) int {
	regex := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	ans := 0

	for _, str := range strings {
		matches := regex.FindAllStringSubmatch(str, -1)

		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			ans += num1 * num2
		}

	}

	return ans
}

func part2(strings []string) int {
	regex := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)")

	ans := 0
	do := true

	for _, str := range strings {
		matches := regex.FindAllStringSubmatch(str, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				do = true
			} else if match[0] == "don't()" {
				do = false
			} else if do {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				ans += num1 * num2
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

	strings := readFile(file)

	start1 := time.Now()
	ans1 := part1(strings)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	ans2 := part2(strings)
	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1 ans: %d time: %s\n", ans1, elapsed1)
	fmt.Printf("Part 2 ans: %d time: %s\n", ans2, elapsed2)
}
