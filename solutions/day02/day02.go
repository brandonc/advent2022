package day02

import (
	"bufio"
	"io"

	"github.com/brandonc/advent2022/solutions/solution"
)

type day02 struct{}

const (
	rock     = 1
	paper    = 2
	scissors = 3
)

// Factory must exist for codegen
func Factory() solution.Solver {
	return day02{}
}

type game struct {
	opp byte
	me  byte
}

func (g game) oppChoice() int {
	switch g.opp {
	case 'A':
		return rock
	case 'B':
		return paper
	case 'C':
		return scissors
	default:
		panic("unexpected opponent choice: " + string(g.opp))
	}
}

func (g game) meChoice() int {
	switch g.me {
	case 'X':
		return rock
	case 'Y':
		return paper
	case 'Z':
		return scissors
	default:
		panic("unexpected me choice: " + string(g.me))
	}
}

func (g game) opponentWins() bool {
	return (g.oppChoice() == paper && g.meChoice() == rock) ||
		(g.oppChoice() == scissors && g.meChoice() == paper) ||
		(g.oppChoice() == rock && g.meChoice() == scissors)
}

func (g game) scoreFirst() int {
	if g.oppChoice() == g.meChoice() {
		// Draw
		return 3 + g.meChoice()
	} else if g.opponentWins() {
		// Loss
		return 0 + g.meChoice()
	}
	return 6 + g.meChoice()
}

func (g game) toWin() int {
	switch g.oppChoice() {
	case rock:
		return paper
	case paper:
		return scissors
	case scissors:
		return rock
	default:
		panic("this can't happen")
	}
}

func (g game) toLose() int {
	switch g.oppChoice() {
	case rock:
		return scissors
	case paper:
		return rock
	case scissors:
		return paper
	default:
		panic("this can't happen")
	}
}

func (g game) scoreSecond() int {
	switch g.me {
	case 'X':
		// Lose
		return 0 + g.toLose()
	case 'Y':
		// Draw
		return 3 + g.oppChoice()
	case 'Z':
		// Win
		return 6 + g.toWin()
	default:
		panic("unexpected me choice: " + string(g.me))
	}
}

func (d day02) Solve(reader io.Reader) (interface{}, interface{}, error) {
	scanner := bufio.NewScanner(reader)

	guideScore := 0
	clarifiedScore := 0
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) != 3 {
			panic("unexpected input: " + input)
		}

		round := game{
			opp: input[0],
			me:  input[2],
		}

		guideScore += round.scoreFirst()
		clarifiedScore += round.scoreSecond()
	}
	return guideScore, clarifiedScore, nil
}
