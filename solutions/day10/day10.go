// Package day10 describes a set of computer instructions that add a number to a register.
// Each instruction requires some number of cycles to complete.
//
// Part 1 is sum of the register multiplied by the cycle number at 6 intervals
// Part 2 is a rendering of a sprite in a horizontal buffer defined by the register
//
// The approach is to perform CRT operations during every cycle, but store the last CPU
// instruction until enough cycles have passed to execute it.
//
// The solution is unique because the CRT 'image' is written to a text file 'image.txt'
package day10

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/brandonc/advent2022/internal/maths"
	"github.com/brandonc/advent2022/internal/ui"
	"github.com/brandonc/advent2022/solutions/solution"
)

type day10 struct {
	x                 int
	cycle             int
	instructionCycles int
	finishedExecuting func()
	crtRowBuffer      []byte
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day10{
		x:                 1,
		cycle:             0,
		instructionCycles: 1,
		finishedExecuting: noop(),
		crtRowBuffer:      make([]byte, crtColumns),
	}
}

const crtColumns = 40

func noop() func() {
	return func() {}
}

func addx(n string, register *int) func() {
	return func() {
		arg, err := strconv.Atoi(n)
		if err != nil {
			panic(fmt.Sprintf("expected addx <number>, got %s", n))
		}
		*register += arg
	}
}

func (d day10) crtPosition() int {
	return (d.cycle - 1) % crtColumns
}

// spriteIsVisible determines if the crtPosition is timed to the x register, which is the
// position of the center of a 3-pixel sprite
func (d day10) spriteIsVisible() bool {
	return d.crtPosition() == d.x-1 || d.crtPosition() == d.x || d.crtPosition() == d.x+1
}

// beginExecuting interprets the next instruction for future execution
func (d *day10) beginExecuting(instr string) {
	if strings.HasPrefix(instr, "addx") {
		d.instructionCycles = 2
		d.finishedExecuting = addx(instr[5:], &d.x)
	} else if strings.HasPrefix(instr, "noop") {
		d.instructionCycles = 1
		d.finishedExecuting = noop()
	}
}

func (d day10) Solve(reader io.Reader) (interface{}, interface{}, error) {
	signalStrengths := make([]int, 0, 5)
	scanner := bufio.NewScanner(reader)
	image, err := os.Create("image.txt")
	if err != nil {
		return 0, 0, fmt.Errorf("could not write to image.txt: %w", err)
	}
	defer image.Close()

	for {
		d.cycle += 1

		d.instructionCycles -= 1
		if d.instructionCycles == 0 {
			// Instruction is complete. Execute it and get the next.
			d.finishedExecuting()
			if !scanner.Scan() {
				break
			}
			d.beginExecuting(scanner.Text())
		}

		// At these cycle intervals, record the cycle * x, which is the signal strength
		if d.cycle == 20 || (d.cycle+20)%40 == 0 {
			ui.Debugf("at cycle %d, the signal strength is %d * %d = %d", d.cycle, d.x, d.cycle, d.x*d.cycle)
			signalStrengths = append(signalStrengths, d.x*d.cycle)
		}

		var pixel byte = '.'
		if d.spriteIsVisible() {
			pixel = '#'
		}
		d.crtRowBuffer[d.crtPosition()] = pixel

		// At the end of the CRT row, write the buffer to the output file
		if d.crtPosition() == crtColumns-1 {
			image.Write(d.crtRowBuffer)
			image.WriteString("\n")
		}
	}

	return fmt.Sprint(maths.SumSlice(signalStrengths)), "see image.txt", nil
}
