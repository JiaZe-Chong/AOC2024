package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseStones(r io.Reader) map[int]int {
	stones := make(map[int]int)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	for _, e := range strings.Split(scanner.Text(), " ") {
		num, _ := strconv.Atoi(e)
		stones[num]++
	}

	return stones
}

func part1(stones map[int]int) int {
	numBlinks := 25

	for i := 0; i < numBlinks; i++ {
		nextStones := make(map[int]int)

		for stone, num := range stones {
			if stone == 0 {
				nextStones[1] += num
			} else if int(math.Log10(float64(stone)))%2 == 1 {
				str := strconv.Itoa(stone)
				left, _ := strconv.Atoi(str[:len(str)/2])
				right, _ := strconv.Atoi(str[len(str)/2:])
				nextStones[left] += num
				nextStones[right] += num
			} else {
				nextStones[stone*2024] += num
			}
		}

		stones = nextStones
	}

	ans := 0

	for _, v := range stones {
		ans += v
	}

	return ans
}

func part2(stones map[int]int) int {
	numBlinks := 75

	for i := 0; i < numBlinks; i++ {
		nextStones := make(map[int]int)

		for stone, num := range stones {
			if stone == 0 {
				nextStones[1] += num
			} else if int(math.Log10(float64(stone)))%2 == 1 {
				str := strconv.Itoa(stone)
				left, _ := strconv.Atoi(str[:len(str)/2])
				right, _ := strconv.Atoi(str[len(str)/2:])
				nextStones[left] += num
				nextStones[right] += num
			} else {
				nextStones[stone*2024] += num
			}
		}

		stones = nextStones
	}

	ans := 0

	for _, v := range stones {
		ans += v
	}

	return ans
}

func main() {
	filename := "text.txt"
	f, _ := os.Open(filename)
	stones := parseStones(f)

	start1 := time.Now()
	ans1 := part1(stones)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	ans2 := part2(stones)
	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1 ans: %d time: %s\n", ans1, elapsed1)
	fmt.Printf("Part 2 ans: %d time: %s\n", ans2, elapsed2)
}
