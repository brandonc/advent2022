// Package day07 solves a problem of reverse engineering some terminal commands and their
// output into a heirarchy of directories and the size of their files and subfiles.
//
// The approach is to track the directories that are traveled to and the sum of the files within
// that directory. The total size of the directory can be calculated by recursive descent.
//
// The solution to the first part is the sum of the file sizes among the nodes that contain less
// than or equal to 100_000.
//
// The solution to the second part is the smallest directory to delete that is greater than
// the space needed to install an update.
package day07

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/brandonc/advent2022/internal/ds"
	"github.com/brandonc/advent2022/internal/ui"
	"github.com/brandonc/advent2022/solutions/solution"
)

type day07 struct {
	root    *ds.TreeNode
	current *ds.TreeNode
}

const (
	updateSize = 30_000_000
	deviceSize = 70_000_000
)

// Factory must exist for codegen
func Factory() solution.Solver {
	result := day07{
		root: ds.NewTree(0),
	}
	result.current = result.root
	return result
}

// totalValue recursively sums the value of all nodes
func totalValue(n *ds.TreeNode) int {
	cv := n.Value.(int)
	for _, child := range n.Children {
		cv += totalValue(child)
	}
	return cv
}

func sumNodesLessThan(n *ds.TreeNode, max int) int {
	val := 0
	if totalValue(n) <= max {
		val = totalValue(n)
	}
	for _, c := range n.Children {
		val += sumNodesLessThan(c, max)
	}
	return val
}

func smallestDirectoryGreaterThan(min int, node *ds.TreeNode, candidate int) int {
	currentTotal := totalValue(node)
	if currentTotal >= min && currentTotal < candidate {
		for _, c := range node.Children {
			maybe := smallestDirectoryGreaterThan(min, c, currentTotal)
			if maybe != -1 && maybe < currentTotal {
				currentTotal = maybe
			}
		}
		return currentTotal
	}
	return -1
}

func (d *day07) processCommand(command string) {
	switch {
	case command == "cd ..":
		// navigate to parent directory
		d.current = d.current.Parent
	case command == "cd /":
		// navigate to root directory
		for d.current.Parent != nil {
			d.current = d.current.Parent
		}
	case strings.HasPrefix(command, "cd "):
		// navigate to named directory
		childDir := d.current.GetChild(command[3:])
		ui.Assert(childDir != nil, fmt.Sprintf("child directory %s not found", command[3:]))
		d.current = childDir
	case command == "ls":
		// no navigation needed
	default:
		panic("unexpected command: " + command)
	}
}

func (d day07) Solve(reader io.Reader) (interface{}, interface{}, error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$") {
			// This is processed as command input, which can change the current directory
			d.processCommand(line[2:])
			continue
		}

		// Otherwise it is command output:
		// A directory with a name that we will begin tracking
		if strings.HasPrefix(line, "dir") {
			d.current.AddChild(line[4:], 0)
			continue
		}
		// Or a file with a size that we will add to the current directory
		fileRaw := strings.Split(line, " ")
		ui.Assert(len(fileRaw) == 2, fmt.Sprintf("expected line %s to have two fields", line))
		size, err := strconv.Atoi(fileRaw[0])
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse size of %s: %w", line, err)
		}

		previous := d.current.Value.(int)
		d.current.Value = previous + size
	}

	needed := updateSize - (deviceSize - totalValue(d.root))
	return sumNodesLessThan(d.root, 100_000), smallestDirectoryGreaterThan(needed, d.root, math.MaxInt), nil
}
