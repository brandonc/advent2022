package day12

import (
	"bufio"
	"io"
	"math"
	"sort"

	"github.com/brandonc/advent2022/solutions/solution"
)

type day12 struct{}

type navigation struct {
	heightMap [][]byte
	start     position
	goal      position
}

type position struct {
	y, x int
}

type heap map[position]int

func (h *heap) pop() position {
	var cheapest position
	var cheapestCost int = math.MaxInt
	for p, cost := range *h {
		if cost < cheapestCost {
			cheapest = p
			cheapestCost = cost
		}
	}
	delete(*h, cheapest)
	return cheapest
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day12{}
}

// aStar implements the A* pathfinding algorithm
func (n navigation) aStar() int {
	frontier := make(heap)
	cost := make(map[position]int)

	frontier[n.start] = 0
	cost[n.start] = 0

	for len(frontier) > 0 {
		pos := frontier.pop()

		if pos == n.goal {
			break
		}

		neighbors := n.neighbors(pos)
		for _, neighbor := range neighbors {
			newCost := cost[pos] + 1
			previousCost, visited := cost[neighbor]
			if !visited || newCost < previousCost {
				cost[neighbor] = newCost
				frontier[neighbor] = newCost
			}
		}
	}

	return cost[n.goal]
}

func (n navigation) neighbors(p position) []position {
	result := make([]position, 0, 4)
	maxHeight := n.heightMap[p.y][p.x] + 1

	// North
	if p.y > 0 && n.heightMap[p.y-1][p.x] <= maxHeight {
		result = append(result, position{p.y - 1, p.x})
	}

	// South
	if p.y < len(n.heightMap)-1 && n.heightMap[p.y+1][p.x] <= maxHeight {
		result = append(result, position{p.y + 1, p.x})
	}

	// West
	if p.x > 0 && n.heightMap[p.y][p.x-1] <= maxHeight {
		result = append(result, position{p.y, p.x - 1})
	}

	// East
	if p.x < len(n.heightMap[0])-1 && n.heightMap[p.y][p.x+1] <= maxHeight {
		result = append(result, position{p.y, p.x + 1})
	}

	return result
}

func (d day12) Solve(reader io.Reader) (interface{}, interface{}, error) {
	scanner := bufio.NewScanner(reader)
	nav := navigation{
		heightMap: make([][]byte, 0, 8),
	}
	y := 0
	lowestPoints := make([]position, 0, 8)
	for scanner.Scan() {
		row := make([]byte, len(scanner.Text()))
		for x, c := range scanner.Bytes() {
			row[x] = c

			// The start and end locations are at a and z height respectively
			if c == 'S' {
				nav.start = position{y, x}
				row[x] = 'a'
			} else if c == 'E' {
				nav.goal = position{y, x}
				row[x] = 'z'
			}

			// Store all the lowest points to evaluate part 2
			if row[x] == 'a' {
				lowestPoints = append(lowestPoints, position{y, x})
			}
		}
		nav.heightMap = append(nav.heightMap, row)
		y += 1
	}

	part1 := nav.aStar()

	pathsFromLowest := make([]int, 0, len(lowestPoints))
	for _, o := range lowestPoints {
		nav.start = o
		cost := nav.aStar()
		if cost > 0 {
			pathsFromLowest = append(pathsFromLowest, cost)
		}
	}

	sort.Ints(pathsFromLowest)
	part2 := pathsFromLowest[0]
	return part1, part2, nil
}
