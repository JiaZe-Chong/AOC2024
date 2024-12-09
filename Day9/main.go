package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func parseString(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return scanner.Text()
}

func part1(str string) int {
	var blocks []int
	isFile := true
	id := 0

	for _, ele := range str {
		if isFile {
			for i := 0; i < int(ele-'0'); i++ {
				blocks = append(blocks, id)
			}
			id++
		} else {
			for i := 0; i < int(ele-'0'); i++ {
				blocks = append(blocks, -1)
			}
		}

		isFile = !isFile
	}

	left := 0
	right := len(blocks) - 1

	for left < right {
		if blocks[left] != -1 {
			left++
		} else if blocks[right] == -1 {
			right--
		} else {
			blocks[left], blocks[right] = blocks[right], blocks[left]
			left++
			right--
		}
	}

	ans := 0
	for ind, ele := range blocks {
		if ele == -1 {
			break
		}
		ans += ind * ele
	}
	return ans
}

type FreeSpace struct {
	ind int
	len int
}

func part2(str string) int {
	var blocks []int
	var spaces []FreeSpace
	isFile := true
	id := 0

	for _, ele := range str {
		if isFile {
			for i := 0; i < int(ele-'0'); i++ {
				blocks = append(blocks, id)
			}
			id++
		} else {
			var space FreeSpace
			space.ind = len(blocks)
			space.len = int(ele - '0')
			spaces = append(spaces, space)
			for i := 0; i < int(ele-'0'); i++ {
				blocks = append(blocks, -1)
			}
		}

		isFile = !isFile
	}

	set := make(map[int]bool)
	set[-1] = true

	for lastInd, block := range slices.Backward(blocks) {

		if set[block] {
			continue
		}

		firstInd := slices.Index(blocks, block)
		len := lastInd - firstInd + 1

		replaced := false
		for ind, space := range spaces {
			if firstInd <= space.ind {
				break
			}

			if len <= space.len {
				for i := 0; i < len; i++ {
					blocks[firstInd+i], blocks[space.ind+i] = blocks[space.ind+i], blocks[firstInd+i]
				}
				replaced = true
				space.len -= len
				space.ind += len
				spaces[ind] = space
				break
			}
		}
		if !replaced {
			set[block] = true
		}
	}

	ans := 0
	for ind, ele := range blocks {
		if ele == -1 {
			continue
		}
		ans += ind * ele
	}
	return ans
}

func main() {
	filename := "text.txt"
	file, _ := os.Open(filename)
	str := parseString(file)

	ans1 := part1(str)
	ans2 := part2(str)

	fmt.Printf("Part 1: %d\n", ans1)
	fmt.Printf("Part 2: %d\n", ans2)
}
