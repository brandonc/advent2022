// Package day09 simulates ropes with different numbers of knots moving in 2 dimensional space.
//
// The rope moves one unit at a time, and the knot immediately following moves to follow
// if it is not one space away, either in the x or y direction, or both at once to simulate
// diagonal movement.
//
// The solution is the number of spaces that the trailing knot has visited in space. Part one
// is a rope with two knots, part two has ten, but the approach is the same:
//
// The approach is to implement the knots as a linked list, where the knot only has to
// consider the previous knot's position. If the two are already touching within 1 space,
// it and the remaining knots do not have to move.
package day09

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/brandonc/advent2022/internal/maths"
	"github.com/brandonc/advent2022/solutions/solution"
)

type knot struct {
	x, y int
	next *knot
}

type day09 struct{}

type rope struct {
	head    knot
	visited map[knot]struct{}
}

func (r knot) touching(other knot) bool {
	diffX, diffY := maths.AbsInt(r.x-other.x), maths.AbsInt(r.y-other.y)

	return (diffX == 0 && diffY == 0) || // same spot
		(diffX == 1 && diffY == 1) || // diag
		(diffX == 0 && diffY == 1) || // above or below
		(diffX == 1 && diffY == 0) // left or right
}

func (r *knot) followX(other knot) {
	if other.x > r.x {
		r.x += 1
	} else {
		r.x -= 1
	}
}

func (r *knot) followY(other knot) {
	if other.y > r.y {
		r.y += 1
	} else {
		r.y -= 1
	}
}

func (r *knot) follow(other knot, visited map[knot]struct{}) {
	if r == nil {
		visited[other] = struct{}{}
		return
	}

	if r.touching(other) {
		return
	}

	// Follow logic
	if other.x == r.x {
		// follow up or down
		r.followY(other)
	} else if other.y == r.y {
		// follow left or right
		r.followX(other)
	} else {
		// follow diag
		r.followX(other)
		r.followY(other)
	}

	r.next.follow(*r, visited)
}

func (r *rope) move(dir byte) {
	switch dir {
	case 'U':
		r.head.y += 1
	case 'D':
		r.head.y -= 1
	case 'L':
		r.head.x -= 1
	case 'R':
		r.head.x += 1
	default:
		panic(fmt.Sprintf("unexpected dir %c", dir))
	}

	r.head.next.follow(r.head, r.visited)
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day09{}
}

func (d day09) Solve(reader io.Reader) (any, any, error) {
	scanner := bufio.NewScanner(reader)

	twoKnot := rope{
		head:    knot{0, 0, &knot{0, 0, nil}},
		visited: make(map[knot]struct{}),
	}

	tenKnot := rope{
		head:    knot{0, 0, &knot{0, 0, &knot{0, 0, &knot{0, 0, &knot{0, 0, &knot{0, 0, &knot{0, 0, &knot{0, 0, &knot{0, 0, &knot{0, 0, nil}}}}}}}}}},
		visited: make(map[knot]struct{}),
	}

	// Both tails have visited the initial space
	twoKnot.visited[knot{0, 0, nil}] = struct{}{}
	tenKnot.visited[knot{0, 0, nil}] = struct{}{}

	for scanner.Scan() {
		if len(scanner.Text()) < 3 {
			continue
		}
		dir := scanner.Text()[0]
		times, err := strconv.Atoi(scanner.Text()[2:])
		if err != nil {
			return 0, 0, fmt.Errorf("expected number of moves, got %s", scanner.Text()[2:])
		}

		for i := 0; i < times; i++ {
			twoKnot.move(dir)
			tenKnot.move(dir)
		}
	}
	return len(twoKnot.visited), len(tenKnot.visited), nil
}
