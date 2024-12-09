package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func parseRules(r io.Reader) map[int][]int {
	m := make(map[int][]int)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "|")
		left, _ := strconv.Atoi(nums[0])
		right, _ := strconv.Atoi(nums[1])
		m[right] = append(m[right], left)
	}

	return m
}

func parseUpdates(r io.Reader) [][]int {
	var ret [][]int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), ",")
		var updates []int
		for _, ele := range nums {
			num, _ := strconv.Atoi(ele)
			updates = append(updates, num)
		}
		ret = append(ret, updates)
	}

	return ret
}

func checkCorrect(rules map[int][]int, update []int) bool {
	m := make(map[int]bool)
	correct := true

	for _, ele := range update {
		if m[ele] {
			correct = false
			break
		} else {
			for _, num := range rules[ele] {
				m[num] = true
			}
		}
	}

	return correct
}

func part1(rules map[int][]int, updates [][]int) int {
	ans := 0

	for _, update := range updates {
		if checkCorrect(rules, update) {
			ans += update[len(update)/2]
		}
	}

	return ans
}

type Page struct {
	Val          int
	CannotBefore []int
}

type ByPage []Page

func (a ByPage) Len() int           { return len(a) }
func (a ByPage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPage) Less(i, j int) bool { return slices.Contains(a[i].CannotBefore, a[j].Val) }

func fixUpdate(rules map[int][]int, update []int) []int {
	var fixed_update ByPage
	for _, ele := range update {
		var page Page
		page.Val = ele
		page.CannotBefore = rules[ele]
		fixed_update = append(fixed_update, page)
	}

	sort.Sort(fixed_update)

	var ret []int
	for _, ele := range fixed_update {
		ret = append(ret, ele.Val)
	}

	return ret
}

func part2(rules map[int][]int, updates [][]int) int {
	ans := 0

	for _, update := range updates {
		if !checkCorrect(rules, update) {
			fixed_update := fixUpdate(rules, update)
			ans += fixed_update[len(fixed_update)/2]
		}
	}

	return ans
}

func main() {
	rules_filename := "rules.txt"
	updates_filename := "updates.txt"

	rules_file, _ := os.Open(rules_filename)
	updates_file, _ := os.Open(updates_filename)

	rules := parseRules(rules_file)
	updates := parseUpdates(updates_file)

	start1 := time.Now()
	ans1 := part1(rules, updates)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	ans2 := part2(rules, updates)
	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1 ans: %d time: %s\n", ans1, elapsed1)
	fmt.Printf("Part 2 ans: %d time: %s\n", ans2, elapsed2)
}
