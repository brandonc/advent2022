// Package day11 is a simulation of monkeys tossing items. It's too dumb to describe fully
// but basically it's about managing very large numbers by exploiting the fact that
// comparing them is only done with prime numbers.
package day11

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/brandonc/advent2022/solutions/solution"
)

type day11 struct{}

type op func(int) int
type curry = func(int) op

var add curry = func(x int) op {
	return func(y int) int {
		return x + y
	}
}
var mult curry = func(x int) op {
	return func(y int) int {
		return x * y
	}
}
var square op = func(x int) int {
	return x * x
}

type monkey struct {
	items         []int
	operation     op
	test          int
	ifTrueMonkey  int
	ifFalseMonkey int
	inspected     int
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day11{}
}

func (m *monkey) inspectNext() (int, bool) {
	if len(m.items) == 0 {
		return 0, false
	}

	result, items := m.items[0], m.items[1:]
	m.items = items

	result = m.operation(result)

	m.inspected += 1
	return result, true
}

func solve(monkeys []*monkey, rounds, relief int) int {
	// All test numbers are prime, so you can multiply them together n
	// and modulo each item by that number to manage its size.
	modulusProduct := 1
	for _, m := range monkeys {
		modulusProduct *= m.test
	}

	for round := 0; round < rounds; round++ {
		for _, monkey := range monkeys {
			item, hasItem := monkey.inspectNext()
			for hasItem {
				item %= modulusProduct
				item /= relief

				if item%monkey.test == 0 {
					monkeys[monkey.ifTrueMonkey].items = append(monkeys[monkey.ifTrueMonkey].items, item)
				} else {
					monkeys[monkey.ifFalseMonkey].items = append(monkeys[monkey.ifFalseMonkey].items, item)
				}
				item, hasItem = monkey.inspectNext()
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected < monkeys[j].inspected
	})

	return monkeys[len(monkeys)-1].inspected * monkeys[len(monkeys)-2].inspected
}

func (d day11) Solve(reader io.Reader) (any, any, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return 0, 0, fmt.Errorf("could not read input: %w", err)
	}

	return solve(parse(strings.NewReader(string(data))), 20, 3),
		solve(parse(strings.NewReader(string(data))), 10_000, 1),
		nil
}
