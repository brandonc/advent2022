// Package day04 finds if two ranges overlap or contain one another given the set format
// like 2-4,3-6
package day04

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2022/internal/ui"
	"github.com/brandonc/advent2022/solutions/solution"
)

type day04 struct{}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day04{}
}

type boundary struct {
	lower int
	upper int
}

func (b boundary) contains(other boundary) bool {
	return b.lower <= other.lower && b.upper >= other.upper
}

func (b boundary) overlaps(other boundary) bool {
	return !(b.upper < other.lower || b.lower > other.upper)
}

func parseRanges(rangesRaw string) (boundary, boundary) {
	split := strings.Split(rangesRaw, ",")
	ui.Assert(len(split) == 2, "expected two ranges")

	parseRange := func(r string) boundary {
		boundsRaw := strings.Split(r, "-")
		ui.Assert(len(split) == 2, "expected two bounds within range")

		l, err := strconv.Atoi(boundsRaw[0])
		ui.Assert(err == nil, "expected integer in bounds")
		u, err := strconv.Atoi(boundsRaw[1])
		ui.Assert(err == nil, "expected integer in bounds")

		return boundary{
			upper: u,
			lower: l,
		}
	}
	return parseRange(split[0]), parseRange(split[1])
}

func (d day04) Solve(reader io.Reader) (interface{}, interface{}, error) {
	scanner := bufio.NewScanner(reader)

	var countContains, countOverlaps int
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		rangeFirst, rangeSecond := parseRanges(scanner.Text())
		if rangeFirst.contains(rangeSecond) || rangeSecond.contains(rangeFirst) {
			countContains++
		}
		if rangeFirst.overlaps(rangeSecond) {
			countOverlaps++
		}
	}
	return countContains, countOverlaps, nil
}
