// Package day05 accepts a drawing of some stacks of crates and some instructions for moving
// the crates from one stack to another.
//
// In the first part, the crates are moved one at a time
// In the second part, the crates can be moved several at a time
//
// My solution was to use an array of stack data structures that can push/pop N items
package day05

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/brandonc/advent2022/internal/ds"
	"github.com/brandonc/advent2022/internal/ui"
	"github.com/brandonc/advent2022/solutions/solution"
)

type day05 struct {
	stacksP1 []ds.Stack
	stacksP2 []ds.Stack
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day05{}
}

var instructionPattern = regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

// topmostCrates peeks at the top crate of each stack and returns a string
// comprised of each crate letter. This is the puzzle answer
func topmostCrates(s []ds.Stack) string {
	var a strings.Builder
	for _, s := range s {
		a.WriteString(s.Peek())
	}
	return a.String()
}

func (d day05) Solve(reader io.Reader) (any, any, error) {
	// First, read the drawing of the stacks of crates
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), " 1") {
			// This is the end of the stacks input. Read the extra empty line and move to the
			// instructions input
			scanner.Scan()
			break
		}

		if d.stacksP1 == nil {
			// Initialize stack array (the number of chars in the line / 4)
			d.stacksP1 = make([]ds.Stack, len(scanner.Text())/4+1)
			d.stacksP2 = make([]ds.Stack, len(scanner.Text())/4+1)
		}

		ui.Assert((len(scanner.Text())+1)%4 == 0, "Unexpected drawing line: "+scanner.Text())

		for start := 0; start < len(scanner.Text()); start += 4 {
			// Each group of 4 characters is a crate, but could be blank if the stack isn't this tall
			stackIndex := start / 4
			letter := scanner.Text()[start+1]

			if letter == ' ' {
				// No crate at this height
				continue
			}

			// Unshift the crate letter in the stack because the input is processed from the top down
			d.stacksP1[stackIndex].Unshift(string(letter))
			d.stacksP2[stackIndex].Unshift(string(letter))

			ui.Debugf("Unshifted %c onto stack %d", letter, stackIndex)
		}
	}

	// Read the instructions and manipulate the stacks as directed. Example: move 2 from 8 to 1
	for scanner.Scan() {
		matches := instructionPattern.FindStringSubmatch(scanner.Text())
		if matches == nil {
			break
		}

		// Regex ensures that digits appeared in submatches
		count, _ := strconv.Atoi(matches[1])
		source, _ := strconv.Atoi(matches[2])
		dest, _ := strconv.Atoi(matches[3])

		for i := 0; i < count; i++ {
			d.stacksP1[dest-1].Push(d.stacksP1[source-1].Pop())
		}
		d.stacksP2[dest-1].PushN(d.stacksP2[source-1].PopN(count))
	}

	return topmostCrates(d.stacksP1), topmostCrates(d.stacksP2), nil
}
