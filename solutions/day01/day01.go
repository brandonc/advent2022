package day01

import (
	"io"
	"sort"

	"github.com/brandonc/advent2022/internal/input"
	"github.com/brandonc/advent2022/internal/maths"
	"github.com/brandonc/advent2022/internal/ui"
	"github.com/brandonc/advent2022/solutions/solution"
)

type day01 struct{}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day01{}
}

func (d day01) Solve(reader io.Reader) (interface{}, interface{}, error) {
	scanner := input.NewIntScanner(reader)

	elfCalories := make([]int, 0)
	var calories int
	for scanner.Scan() {
		if scanner.Text() == "" {
			elfCalories = append(elfCalories, calories)
			calories = 0
			continue
		}

		calories += scanner.Int()
	}

	elfCalories = append(elfCalories, calories)
	sort.Ints(elfCalories)
	len := len(elfCalories)

	ui.Debugf("There are %d elves", len)

	if len < 3 {
		panic("assertion failed!")
	}

	return elfCalories[len-1], maths.SumSlice(elfCalories[len-3 : len]), nil
}
