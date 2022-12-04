// Package day03 finds common characters within a string representing the contents
// of a bag. Each item is represented by the bytes a-z or A-Z. Each line represents a single bag.
//
// The first answer is derived by finding the items that are common in the first and second half
// of each line. Each line/bag should have an even number of bytes/items.
//
// The second answer is derived by finding the item (byte) common between each group of 3
// bags (lines).
//
// The answers are equal to the sum of all scores of the common items, where
// a=1, b=2...A=27, B=28, etc.
//
// The solution for the second answer works line by line, tracking the contents of each group's
// items in an array of 3 maps (one for each bag). After 3 lines are cataloged, we then look
// for each possible item present in all 3. There are 52 possible items: a-z + A-Z.
package day03

import (
	"bufio"
	"io"

	"github.com/brandonc/advent2022/internal/ui"
	"github.com/brandonc/advent2022/solutions/solution"
)

type item byte

type day03 struct {
	score        int
	groupScore   int
	groupCatalog []map[item]int // Length 3 of histogram contents of each group
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day03{
		groupCatalog: []map[item]int{
			make(map[item]int),
			make(map[item]int),
			make(map[item]int),
		},
	}
}

func (d *day03) resetCatalog() {
	for elf := range d.groupCatalog {
		for x := range d.groupCatalog[elf] {
			d.groupCatalog[elf][x] = 0
		}
	}
}

// scoreCatalog is a helper for determining which item is common among the group and adding it
// to the group score
func (d *day03) scoreCatalog() {
	for b := item('a'); b <= 'z'; b++ {
		if d.groupCatalog[0][b] > 0 && d.groupCatalog[1][b] > 0 && d.groupCatalog[2][b] > 0 {
			d.groupScore += scoreItem(b)
			return
		}
	}
	for b := item('A'); b <= 'Z'; b++ {
		if d.groupCatalog[0][b] > 0 && d.groupCatalog[1][b] > 0 && d.groupCatalog[2][b] > 0 {
			d.groupScore += scoreItem(b)
			return
		}
	}
}

func scoreItem(a item) int {
	// Helpers for returning the point value for an item
	// a = 1 point, b = 2 ... A = 27, B = 28, etc
	if a >= 'a' && a <= 'z' {
		return int(a - 'a' + 1)
	}
	if a >= 'A' && a <= 'Z' {
		return int(a - 'A' + 1 + 26)
	}
	ui.Assert(false, "item not found in scoring key")
	return 0 // never
}

func (d day03) Solve(input io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(input)

	var index = 0
	for scanner.Scan() {
		var items = scanner.Text()
		var elfIndex = index % 3 // Index within group: 0, 1, or 2

		// For each group, score the previous group and reset to prepare for the next group
		if elfIndex == 0 && index != 0 {
			d.scoreCatalog()
			d.resetCatalog()
		}

		alreadyScored := false
		compartmentAContents := make(map[byte]int)

		// Catalog first compartment
		for x := 0; x < len(items)/2; x++ {
			compartmentAContents[items[x]]++
			d.groupCatalog[elfIndex][item(items[x])]++
		}

		// Catalog second compartment
		for y := len(items) / 2; y < len(items); y++ {
			d.groupCatalog[elfIndex][item(items[y])]++

			if !alreadyScored && compartmentAContents[items[y]] > 0 {
				d.score += scoreItem(item(items[y]))
				alreadyScored = true
			}
		}
		index++
		ui.Assert(alreadyScored, "Did not find duplicate item")
	}

	ui.Assert(index%3 == 0, "Not an even number of groups in the input")

	// Score final group
	d.scoreCatalog()

	return d.score, d.groupScore, nil
}
