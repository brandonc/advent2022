// Package day15 describes pairs of sensors and beacons. The input is the location of a sensor
// and the nearest beacon to it. The goal is to extrapolate the distance between each pair into
// a map of the total coverage of each sensor.
//
// Part 1 requires that you indicate how much of a particular row is covered by all sensors
// Part 2 requires that you find a single gap between coverage in a broad range of the map
//
// The approach is to not materialize the coverage in any data structure, but use the distance of
// the nearest beacon to describe a circle around the sensor. A particular row's horizontal
// coverage can be derived as a proportion of the vertical distance from the sensor.
//
// To find the gap for part 2, I decided to re-use the detect function from part 1 but
// reduce all the sensor areas in a row to a union of the ranges of all proximate sensors
// in order to detect a gap.
package day15

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"

	"github.com/brandonc/advent2022/internal/maths"
	"github.com/brandonc/advent2022/internal/ui"
	"github.com/brandonc/advent2022/solutions/solution"
)

type day15 struct {
	part1Row int // 2_000_000 normally
	part2Max int // 4_000_000 normally
}

type position struct {
	y, x int
}

type bounds struct {
	left, right int
}

type sensorPair struct {
	sensor position
	beacon position
}

func (p position) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p position) distance(other position) int {
	return maths.AbsInt(p.y-other.y) + maths.AbsInt(p.x-other.x)
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day15{
		// These are modified for tests
		part1Row: 2_000_000,
		part2Max: 4_000_000,
	}
}

var inputPattern = regexp.MustCompile(`x=([\d-]+), y=([\d-]+)`)

func parseInput(reader io.Reader) []sensorPair {
	scanner := bufio.NewScanner(reader)
	sensors := make([]sensorPair, 0, 32)
	for scanner.Scan() {
		matches := inputPattern.FindAllStringSubmatch(scanner.Text(), -1)

		// Parses two pairs of coords using regex pattern
		sensorX, _ := strconv.Atoi(matches[0][1])
		sensorY, _ := strconv.Atoi(matches[0][2])
		beaconX, _ := strconv.Atoi(matches[1][1])
		beaconY, _ := strconv.Atoi(matches[1][2])

		sensors = append(sensors, sensorPair{
			sensor: position{sensorY, sensorX},
			beacon: position{beaconY, beaconX},
		})
	}
	return sensors
}

func objects(sensors []sensorPair) map[position]struct{} {
	result := make(map[position]struct{}, 32)
	for _, s := range sensors {
		result[s.beacon] = struct{}{}
		result[s.sensor] = struct{}{}
	}
	return result
}

func (d day15) detect(sensors []sensorPair, row int) []bounds {
	ranges := make([]bounds, 0, 8)
	for _, s := range sensors {
		radius := s.sensor.distance(s.beacon)
		// Is target row within range of this beacon?
		if row >= s.sensor.y-radius && row <= s.sensor.y+radius {
			// The sensor distance in the horizontal direction is proportional to the y
			ranges = append(ranges, bounds{
				left:  s.sensor.x - (radius - maths.AbsInt(row-s.sensor.y)),
				right: s.sensor.x + (radius - maths.AbsInt(row-s.sensor.y)),
			})
		}
	}

	// Union all ranges together into as few boundaries as possible. The puzzle indicates
	// there will be only one gap between boundaries, so after the union there should be at most
	// two bounds in the result.
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].left < ranges[j].left
	})

	union := make([]bounds, 1, 2)
	union[0] = ranges[0]
	for index := 1; index < len(ranges); index++ {
		r := ranges[index]
		// Use the beginning of each successive range to see if it can be included in the
		// previously added range.
		if union[len(union)-1].right >= r.left-1 {
			union[len(union)-1].right = maths.Max(union[len(union)-1].right, r.right)
		} else {
			union = append(union, r)
		}
	}

	return union
}

func (d day15) Solve(reader io.Reader) (any, any, error) {
	sensors := parseInput(reader)

	// part 1
	ranges := d.detect(sensors, d.part1Row)
	ui.Assert(len(ranges) == 1, fmt.Sprintf("expected one range in row %d", d.part1Row))
	coverage := ranges[0].right - ranges[0].left + 1

	// Deduct the sensors and beacons from the possible coverage count
	for object := range objects(sensors) {
		if object.y == d.part1Row && object.x >= ranges[0].left && object.x <= ranges[0].right {
			coverage -= 1
		}
	}

	// part 2
	distressBeacon := position{}
	for y := 0; y <= d.part2Max; y++ {
		ranges := d.detect(sensors, y)
		ui.Assert(len(ranges) <= 2, fmt.Sprintf("expected up to two ranges in row %d", y))
		if len(ranges) == 2 {
			distressBeacon.y = y
			distressBeacon.x = ranges[0].right + 1
			break
		}
	}

	return coverage, distressBeacon.x*4_000_000 + distressBeacon.y, nil
}
