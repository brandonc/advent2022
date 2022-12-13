// Package day06 is the problem of finding the first n unique bytes in a string of bytes.
//
// The solution I chose was to read one character at a time into a circular buffer of fixed size n
// and check for uniqueness of all characters using a brute force comparison that checks each
// byte against all others once
package day06

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/brandonc/advent2022/internal/ui"
	"github.com/brandonc/advent2022/solutions/solution"
)

type day06 struct {
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day06{}
}

// isStartOfMarker returns true if all the bytes in buf are unique
// this could be implemented by adding all bytes into a set and checking the length
func (d day06) isStartOfMarker(buf []byte) bool {
	for x := 0; x < len(buf); x++ {
		for y := x + 1; y < len(buf); y++ {
			if x != y && buf[x] == buf[y] {
				return false
			}
		}
	}
	return true
}

func (d day06) findMarker(data []byte, size int) (int, error) {
	reader := bytes.NewReader(data)
	buf := make([]byte, size)

	read, err := reader.Read(buf)
	if err != nil {
		return 0, fmt.Errorf("could not read input: %w", err)
	}
	ui.Assert(read == size, fmt.Sprintf("expected to read at least %d bytes, got %d", size, read))

	begin := -1
	for i := 0; true; i++ {
		if d.isStartOfMarker(buf) {
			ui.Debugf("found start of marker at %d", 4+i)
			begin = size + i
			break
		}

		// Advance buffer by 1
		copy(buf[0:], buf[1:])

		// Read a byte into the end of the buffer
		read, err := reader.Read(buf[size-1:])
		if read < 1 {
			break
		}
		if err != nil {
			return 0, fmt.Errorf("could not read input: %w", err)
		}
	}

	if begin == -1 {
		return 0, errors.New("marker not found")
	}

	return begin, nil
}

func (d day06) Solve(reader io.Reader) (any, any, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return 0, 0, fmt.Errorf("could not copy input to buffer: %w", err)
	}

	part1, err := d.findMarker(data, 4)
	if err != nil {
		return 0, 0, err
	}

	part2, err := d.findMarker(data, 14)
	if err != nil {
		return 0, 0, err
	}

	return part1, part2, nil
}
