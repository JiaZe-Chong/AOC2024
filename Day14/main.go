package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Robot struct {
	x  int
	y  int
	dX int
	dY int
}

func (robot Robot) move(seconds, m, n int) Robot {
	robot.x = ((robot.x+robot.dX*seconds)%m + m) % m
	robot.y = ((robot.y+robot.dY*seconds)%n + n) % n
	return robot
}

func parseRobots(r io.Reader) []Robot {
	var robots []Robot

	pattern := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		match := pattern.FindStringSubmatch(scanner.Text())
		var robot Robot
		robot.x, _ = strconv.Atoi(match[1])
		robot.y, _ = strconv.Atoi(match[2])
		robot.dX, _ = strconv.Atoi(match[3])
		robot.dY, _ = strconv.Atoi(match[4])
		robots = append(robots, robot)
	}

	return robots
}

func part1(robots []Robot) int {
	m := 101
	n := 103
	seconds := 100

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for _, robot := range robots {
		temp := robot.move(seconds, m, n)
		if temp.x < m/2 && temp.y < n/2 {
			q1 += 1
		} else if temp.x > m/2 && temp.y < n/2 {
			q2 += 1
		} else if temp.x < m/2 && temp.y > n/2 {
			q3 += 1
		} else if temp.x > m/2 && temp.y > n/2 {
			q4 += 1
		}
	}

	return q1 * q2 * q3 * q4
}

func renderRobots(robots []Robot, ite int) {
	m := 101
	n := 103

	upleft := image.Point{0, 0}
	lowRight := image.Point{m, n}
	img := image.NewNRGBA(image.Rectangle{upleft, lowRight})

	for _, robot := range robots {
		robot := robot.move(ite, m, n)
		img.Set(robot.x, robot.y, color.White)
	}

	f, _ := os.Create(strconv.Itoa(ite) + ".png")
	png.Encode(f, img)
}

func calVariance(robots []Robot) (float64, float64) {
	sumX := 0
	sumY := 0
	for _, robot := range robots {
		sumX += robot.x
		sumY += robot.y
	}

	avgX := float64(sumX) / float64(len(robots))
	avgY := float64(sumY) / float64(len(robots))

	var sumDiffX float64 = 0
	var sumDiffY float64 = 0
	for _, robot := range robots {
		sumDiffX += math.Pow(float64(robot.x)-avgX, 2)
		sumDiffY += math.Pow(float64(robot.y)-avgY, 2)
	}

	return sumDiffX / float64(len(robots)-1), sumDiffY / float64(len(robots)-1)
}

func part2(robots []Robot) int {
	m := 101
	n := 103
	var minVarX float64 = -1
	var minVarY float64 = -1
	minX := -1
	minY := -1

	for i := 1; i <= 103*101; i++ {
		for ind, robot := range robots {
			robots[ind] = robot.move(1, m, n)
		}
		curVarX, curVarY := calVariance(robots)

		if curVarX < minVarX || minVarX == -1 {
			minVarX = curVarX
			minX = i
		}
		if curVarY < minVarY || minVarY == -1 {
			minVarY = curVarY
			minY = i
		}
	}

	t := minX + 51*(minY-minX)*m
	t = ((t % (m * n)) + m*n) % (m * n)
	return t
}

func main() {
	filename := "text.txt"
	f, _ := os.Open(filename)
	robots := parseRobots(f)

	start1 := time.Now()
	ans1 := part1(robots)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	ans2 := part2(robots)
	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1 ans: %d time: %s\n", ans1, elapsed1)
	fmt.Printf("Part 2 ans: %d time: %s\n", ans2, elapsed2)

	renderRobots(robots, ans2)
}
