package day15

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	d := Factory().(day15)
	d.part1Row = 10
	d.part2Max = 20

	a1, a2, err := d.Solve(strings.NewReader(`Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`))

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if a1 != 26 {
		t.Errorf("Expected a1 to be 26, got %d", a1)
	}

	if a2 != 56000011 {
		t.Errorf("Expected a2 to be 56000011, got %d", a2)
	}
}
