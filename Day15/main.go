package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Cell int8

const (
	Empty    Cell = iota
	Box      Cell = iota
	Wall     Cell = iota
	Robot    Cell = iota
	BoxLeft  Cell = iota
	BoxRight Cell = iota
)

type Move struct {
	x int
	y int
}

var (
	Up    Move = Move{x: 0, y: -1}
	Down  Move = Move{x: 0, y: 1}
	Left  Move = Move{x: -1, y: 0}
	Right Move = Move{x: 1, y: 0}
)

type Coord struct {
	x int
	y int
}

func (coord Coord) nextCoord(move Move) Coord {
	return Coord{x: coord.x + move.x, y: coord.y + move.y}
}

func parseGrid(r io.Reader) (grid [][]Cell, coord Coord) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var row []Cell
		for ind, ele := range scanner.Text() {
			switch ele {
			case '#':
				row = append(row, Wall)
				break
			case 'O':
				row = append(row, Box)
				break
			case '@':
				coord.x, coord.y = ind, len(grid)
				row = append(row, Robot)
				break
			case '.':
				row = append(row, Empty)
				break
			default:
				panic("Unknown cell")
			}
		}
		grid = append(grid, row)
	}

	return
}

func get2WGrid(grid [][]Cell) (grid2W [][]Cell, coord Coord) {
	for _, row := range grid {
		var row2W []Cell
		for _, ele := range row {
			switch ele {
			case Wall:
				row2W = append(row2W, Wall)
				row2W = append(row2W, Wall)
				break
			case Box:
				row2W = append(row2W, BoxLeft)
				row2W = append(row2W, BoxRight)
				break
			case Robot:
				coord.x, coord.y = len(row2W), len(grid2W)
				row2W = append(row2W, Robot)
				row2W = append(row2W, Empty)
				break
			case Empty:
				row2W = append(row2W, Empty)
				row2W = append(row2W, Empty)
				break
			default:
				panic("Unknown cell")
			}
		}
		grid2W = append(grid2W, row2W)
	}

	return
}

func parseMoves(r io.Reader) (moves []Move) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		for _, ele := range scanner.Text() {
			switch ele {
			case '^':
				moves = append(moves, Up)
				break
			case 'v':
				moves = append(moves, Down)
				break
			case '<':
				moves = append(moves, Left)
				break
			case '>':
				moves = append(moves, Right)
				break
			default:
				panic("Unknown move")
			}
		}
	}

	return
}

func print(grid [][]Cell) {
	for _, row := range grid {
		for _, ele := range row {
			switch ele {
			case Robot:
				fmt.Print("@")
				break
			case Wall:
				fmt.Print("#")
				break
			case Box:
				fmt.Print("O")
				break
			case Empty:
				fmt.Print(" ")
				break
			case BoxLeft:
				fmt.Print("[")
				break
			case BoxRight:
				fmt.Print("]")
				break
			}
		}
		fmt.Println()
	}
}

func part1(moves []Move, grid [][]Cell, coord Coord) int {
	for _, move := range moves {
		if grid[coord.y][coord.x] != Robot {
			panic("Coord and Robot mismatch")
		}

		nextCoord := coord.nextCoord(move)

		switch grid[nextCoord.y][nextCoord.x] {
		case Empty:
			grid[coord.y][coord.x] = Empty
			grid[nextCoord.y][nextCoord.x] = Robot
			coord = nextCoord
			break
		case Wall:
			break
		case Box:
			swapCoord := nextCoord.nextCoord(move)
			for grid[swapCoord.y][swapCoord.x] == Box {
				swapCoord = swapCoord.nextCoord(move)
			}

			if grid[swapCoord.y][swapCoord.x] == Empty {
				grid[swapCoord.y][swapCoord.x] = Box
				grid[coord.y][coord.x] = Empty
				grid[nextCoord.y][nextCoord.x] = Robot
				coord = nextCoord
			} else if grid[swapCoord.y][swapCoord.x] != Wall {
				panic("Cell not recognized")
			}

			break
		}
	}

	score := 0
	for j, row := range grid {
		for i, ele := range row {
			if ele == Box {
				score += i + 100*j
			}
		}
	}
	return score
}

func checkShift(move Move, grid [][]Cell, coord Coord) bool {
	newCoord := coord.nextCoord(move)

	if grid[newCoord.y][newCoord.x] == Wall {
		return false
	}

	if grid[newCoord.y][newCoord.x] == Empty {
		return true
	}

	if grid[newCoord.y][newCoord.x] != BoxLeft && grid[newCoord.y][newCoord.x] != BoxRight {
		panic("Error in check shift unknown cell")
	}

	if move == Left || move == Right {
		return checkShift(move, grid, newCoord)
	}

	var otherHalf Coord
	if grid[newCoord.y][newCoord.x] == BoxLeft {
		otherHalf = Coord{x: newCoord.x + 1, y: newCoord.y}
	} else if grid[newCoord.y][newCoord.x] == BoxRight {
		otherHalf = Coord{x: newCoord.x - 1, y: newCoord.y}
	} else {
		panic("Error in check shift unknown cell")
	}

	return checkShift(move, grid, newCoord) && checkShift(move, grid, otherHalf)
}

func shift(move Move, grid [][]Cell, coord Coord) [][]Cell {
	newCoord := coord.nextCoord(move)

	if grid[newCoord.y][newCoord.x] == Wall {
		return grid
	}

	if grid[newCoord.y][newCoord.x] == Empty {
		grid[newCoord.y][newCoord.x], grid[coord.y][coord.x] = grid[coord.y][coord.x], grid[newCoord.y][newCoord.x]
		return grid
	}

	if grid[newCoord.y][newCoord.x] != BoxLeft && grid[newCoord.y][newCoord.x] != BoxRight {
		panic("Error in shift unknown cell")
	}

	if move == Left || move == Right {
		grid = shift(move, grid, newCoord)
		grid[newCoord.y][newCoord.x], grid[coord.y][coord.x] = grid[coord.y][coord.x], grid[newCoord.y][newCoord.x]
		return grid
	}

	var otherHalf Coord
	if grid[newCoord.y][newCoord.x] == BoxLeft {
		otherHalf = Coord{x: newCoord.x + 1, y: newCoord.y}
	} else if grid[newCoord.y][newCoord.x] == BoxRight {
		otherHalf = Coord{x: newCoord.x - 1, y: newCoord.y}
	} else {
		panic("Error in shift unknown cell")
	}

	shift(move, grid, newCoord)
	shift(move, grid, otherHalf)

	grid[newCoord.y][newCoord.x], grid[coord.y][coord.x] = grid[coord.y][coord.x], grid[newCoord.y][newCoord.x]

	return grid
}

func part2(moves []Move, grid [][]Cell, coord Coord) int {
	for _, move := range moves {
		if grid[coord.y][coord.x] != Robot {
			panic("Coord and Robot mismatch")
		}

		nextCoord := coord.nextCoord(move)

		switch grid[nextCoord.y][nextCoord.x] {
		case Empty:
			grid[coord.y][coord.x] = Empty
			grid[nextCoord.y][nextCoord.x] = Robot
			coord = nextCoord
			break
		case Wall:
			break
		case BoxRight:
			if checkShift(move, grid, coord) {
				grid = shift(move, grid, coord)
				coord = nextCoord
			}
			break
		case BoxLeft:
			if checkShift(move, grid, coord) {
				grid = shift(move, grid, coord)
				coord = nextCoord
			}
			break
		}
	}

	score := 0
	for j, row := range grid {
		for i, ele := range row {
			if ele == BoxLeft {
				score += i + 100*j
			}
		}
	}
	return score
}

func main() {
	gridFilename := "grid.txt"
	movesFilename := "moves.txt"

	gridFile, _ := os.Open(gridFilename)
	movesFile, _ := os.Open(movesFilename)

	grid, coord := parseGrid(gridFile)
	grid2, coord2 := get2WGrid(grid)
	moves := parseMoves(movesFile)

	start1 := time.Now()
	ans1 := part1(moves, grid, coord)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	ans2 := part2(moves, grid2, coord2)
	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1 ans: %d time: %s\n", ans1, elapsed1)
	fmt.Printf("Part 2 ans: %d time: %s\n", ans2, elapsed2)
}
