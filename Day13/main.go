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

type Button struct {
	x int
	y int
}

type Machine struct {
	A      Button
	B      Button
	target Button
}

func parseMachines(r io.Reader) []Machine {
	var machines []Machine

	patternA := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	patternB := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	patternPrice := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		str := scanner.Text()
		if patternA.MatchString(str) {
			match := patternA.FindStringSubmatch(str)
			var button Button
			button.x, _ = strconv.Atoi(match[1])
			button.y, _ = strconv.Atoi(match[2])
			machines = append(machines, Machine{A: button})
		} else if patternB.MatchString(str) {
			match := patternB.FindStringSubmatch(str)
			var button Button
			button.x, _ = strconv.Atoi(match[1])
			button.y, _ = strconv.Atoi(match[2])
			machines[len(machines)-1].B = button
		} else if patternPrice.MatchString(str) {
			match := patternPrice.FindStringSubmatch(str)
			var button Button
			button.x, _ = strconv.Atoi(match[1])
			button.y, _ = strconv.Atoi(match[2])
			machines[len(machines)-1].target = button
		}
	}

	return machines
}

func (machine Machine) valid(a, b int) bool {
	return machine.A.x*a+machine.B.x*b == machine.target.x && machine.A.y*a+machine.B.y*b == machine.target.y
}

func part1(machines []Machine) int {
	ans := 0

	for _, machine := range machines {
		det := machine.A.x*machine.B.y - machine.A.y*machine.B.x
		a := (machine.target.x*machine.B.y - machine.target.y*machine.B.x) / det
		b := (machine.A.x*machine.target.y - machine.A.y*machine.target.x) / det

		if machine.valid(a, b) {
			ans += 3*a + b
		}
	}

	return ans
}

func part2(machines []Machine) int {
	ans := 0

	for _, machine := range machines {
		machine.target.x += 10000000000000
		machine.target.y += 10000000000000
		det := machine.A.x*machine.B.y - machine.A.y*machine.B.x
		a := (machine.target.x*machine.B.y - machine.target.y*machine.B.x) / det
		b := (machine.A.x*machine.target.y - machine.A.y*machine.target.x) / det

		if machine.valid(a, b) {
			ans += 3*a + b
		}
	}

	return ans
}

func main() {
	filename := "text.txt"
	f, _ := os.Open(filename)
	machines := parseMachines(f)

	start1 := time.Now()
	ans1 := part1(machines)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	ans2 := part2(machines)
	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1 ans: %d time: %s\n", ans1, elapsed1)
	fmt.Printf("Part 2 ans: %d time: %s\n", ans2, elapsed2)
}
