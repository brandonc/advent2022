package day01

import (
	"os"
	"sort"

	"github.com/brandonc/advent2022/internal/input"
	"github.com/brandonc/advent2022/internal/maths"
	"github.com/brandonc/advent2022/internal/ui"

	"github.com/mitchellh/cli"
)

type day01 struct{}

func Command() (cli.Command, error) {
	return day01{}, nil
}

func (d day01) Help() string {
	return "Day 1 solution"
}

func (d day01) Synopsis() string {
	return "Day 1 solution"
}

func (d day01) Run(args []string) int {
	fs, err := os.Open(args[0])
	if ui.Die(err) {
		return 1
	}

	scanner := input.NewIntScanner(fs)

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

	ui.AnswerInt(elfCalories[len-1], "Most calories")

	if len < 3 {
		panic("assertion failed!")
	}

	ui.AnswerInt(maths.SumSlice(elfCalories[len-3:len]), "Top 3 sum")

	return 0
}
