package main

import (
	"fmt"
	"io"
	"os"

	"github.com/brandonc/advent2022/internal/commands"
	"github.com/brandonc/advent2022/internal/ui"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}

	var input io.Reader
	var err error

	if len(os.Args) > 2 {
		input, err = os.Open(os.Args[2])
		if ui.Die(err) {
			os.Exit(1)
		}
	} else {
		input = os.Stdin
	}

	solutionFactory, ok := commands.SolutionCommands[os.Args[1]]
	if !ok {
		printUsage()
	}

	answer1, answer2, err := solutionFactory().Solve(input)

	if ui.Die(err) {
		os.Exit(1)
	}

	ui.Answer(answer1, answer2)

	ui.Die(err)
	os.Exit(0)
}

func printUsage() {
	fmt.Println("Usage:", os.Args[0], "<day> [input]")
	os.Exit(127)
}
