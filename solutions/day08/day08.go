// Package day08 is a 2D grid lookaround problem, consisting of a two dimensional array of
// integers representing the heights of trees. The algorithms involve examining the heights of
// the surrounding trees in four directions (north, south, east, west)
//
// In part 1, you have to count the number of visible trees. A visible tree is one that is not
// surrounded on all 4 sides with a taller or equal height tree.
//
// In part 2, you have to score each tree based on how many other trees can be seen from that
// location. The score is the number of visible trees in all 4 directions multiplied together.
//
// The solution is tedious, and involves storing all integers in a 2-dimensional array and
// looping in the correct direction for each directional "look". I simplified things a bit by
// defining a function to look in each direction and invoke a callback for each tree height
// found there.
package day08

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/brandonc/advent2022/internal/ui"
	"github.com/brandonc/advent2022/solutions/solution"
)

type day08 struct{}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day08{}
}

// traversal defines a function that can traverses a direction beginning at startX, startY
// and calls a given function for each tree height found in that direction. Return false
// from the callback to break.
type traversal func(startX, startY int, grid [][]int, cb func(h int) bool)

// north implements traversing north
var north traversal = func(startX, startY int, grid [][]int, cb func(h int) bool) {
	for yy := startY - 1; yy >= 0; yy-- {
		if !cb(grid[yy][startX]) {
			break
		}
	}
}

// south implements traversing south
var south traversal = func(startX, startY int, grid [][]int, cb func(h int) bool) {
	for yy := startY + 1; yy < len(grid); yy++ {
		if !cb(grid[yy][startX]) {
			break
		}
	}
}

// east implements traversing east
var east traversal = func(startX, startY int, grid [][]int, cb func(h int) bool) {
	for xx := startX + 1; xx < len(grid); xx++ {
		if !cb(grid[startY][xx]) {
			break
		}
	}
}

// west implements traversing west
var west traversal = func(startX, startY int, grid [][]int, cb func(h int) bool) {
	for xx := startX - 1; xx >= 0; xx-- {
		if !cb(grid[startY][xx]) {
			break
		}
	}
}

// tallest returns the tallest tree in a traversal
func tallest(x, y int, grid [][]int, look traversal) int {
	max := 0
	look(x, y, grid, func(h int) bool {
		if h > max {
			max = h
		}
		return true
	})
	return max
}

// visible returns the number of trees visible in a traversal
func visible(x, y int, grid [][]int, look traversal) int {
	height := grid[y][x]
	score := 0
	look(x, y, grid, func(h int) bool {
		score += 1
		return h < height
	})
	return score
}

func countVisible(grid [][]int) int {
	result := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			height := grid[y][x]

			if y == 0 || y == len(grid)-1 || x == 0 || x == len(grid[y])-1 {
				result += 1
				continue
			}

			if tallest(x, y, grid, north) < height ||
				tallest(x, y, grid, south) < height ||
				tallest(x, y, grid, east) < height ||
				tallest(x, y, grid, west) < height {
				result += 1
				continue
			}
		}
	}
	return result
}

func bestScenicScore(grid [][]int) int {
	result := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			score := visible(x, y, grid, north) *
				visible(x, y, grid, south) *
				visible(x, y, grid, east) *
				visible(x, y, grid, west)

			if score > result {
				result = score
			}
		}
	}
	return result
}

func (d day08) Solve(reader io.Reader) (any, any, error) {
	scanner := bufio.NewScanner(reader)
	ui.Assert(scanner.Scan(), "no input")

	lenX := len(scanner.Text())
	grid := make([][]int, 0, 16)
	for {
		row := make([]int, lenX)
		for i, c := range scanner.Bytes() {
			h, err := strconv.Atoi(string(c))
			if err != nil {
				return 0, 0, fmt.Errorf("failed to convert digit to integer: %w", err)
			}
			row[i] = h
		}
		grid = append(grid, row)

		if !scanner.Scan() {
			break
		}
	}

	return countVisible(grid), bestScenicScore(grid), nil
}
