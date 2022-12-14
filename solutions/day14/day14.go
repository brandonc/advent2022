// Package day14 simulates sand falling in 2-dimensional space. It can fall down or
// to the left or right diagonal. Sand can come to rest if it can't move in any of these
// directions when blocked by other sand or rocks (the input) and after it does, more sand
// begins to fall.
//
// The approach is a simulation that ends when sand falls past the threshold (part 1) or when the
// room fills up and 'blocks' the appearance of new sand (part 2).
package day14

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/brandonc/advent2022/internal/ui"
	"github.com/brandonc/advent2022/solutions/solution"
)

type day14 struct{}

type position struct {
	y, x int
}

type entity byte

type scan struct {
	entities  map[position]entity
	boundsX1  int
	boundsX2  int
	boundsY   int
	sand      *position
	sandCount int
}

const (
	rock    entity = '#'
	sand    entity = 'o'
	originX        = 500
	originY        = 0
)

// Factory must exist for codegen
func Factory() solution.Solver {
	return day14{}
}

func (s *scan) simulate() {
	sandDeltas := []position{{1, 0}, {1, -1}, {1, 1}}

	for {
		if s.sand == nil {
			p := position{originY, originX}
			if _, blocked := s.entities[p]; blocked {
				return
			}
			s.sand = &p
			s.sandCount += 1
		}

	tick:
		for {
			// abyss!
			if s.sand.y > s.boundsY {
				s.sandCount -= 1
				return
			}

			// moving
			for _, d := range sandDeltas {
				if _, ok := s.entities[position{s.sand.y + d.y, s.sand.x + d.x}]; !ok {
					s.sand.y += d.y
					s.sand.x += d.x
					continue tick
				}
			}

			// at rest
			s.entities[position{s.sand.y, s.sand.x}] = sand
			s.sand = nil
			break
		}
	}
}

func newScan(reader io.Reader) *scan {
	scanner := bufio.NewScanner(reader)
	result := scan{
		entities: make(map[position]entity),
		boundsX1: math.MaxInt,
	}

	for scanner.Scan() {
		var start *position = nil
		coords := strings.Split(scanner.Text(), " -> ")

		for _, pair := range coords {
			xy := strings.Split(pair, ",")
			r := position{}
			x, err := strconv.Atoi(xy[0])
			ui.Assert(err == nil, fmt.Sprintf("expected integer, got %s", xy[0]))
			y, err := strconv.Atoi(xy[1])
			ui.Assert(err == nil, fmt.Sprintf("expected integer, got %s", xy[0]))
			r.x = x
			r.y = y

			if start != nil {
				result.drawLine(*start, r, rock)
			}
			start = &r

			if r.y > result.boundsY {
				result.boundsY = r.y
			}
			if r.x < result.boundsX1 {
				result.boundsX1 = r.x
			}
			if r.x > result.boundsX2 {
				result.boundsX2 = r.x
			}
		}
	}

	return &result
}

func (d day14) Solve(reader io.Reader) (any, any, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return 0, 0, fmt.Errorf("could not read input: %w", err)
	}

	part1 := newScan(strings.NewReader(string(data)))
	part1.simulate()

	part2 := newScan(strings.NewReader(string(data)))
	// Draw floor on Y boundary + 2, double the size of vertical height
	part2.drawLine(position{part2.boundsY + 2, part2.boundsX1 - part2.boundsY}, position{part2.boundsY + 2, part2.boundsX2 + part2.boundsY}, rock)
	part2.boundsY += 2
	part2.simulate()

	return part1.sandCount, part2.sandCount, nil
}
