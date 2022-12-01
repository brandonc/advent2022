package input

import (
	"bufio"
	"io"
	"strconv"

	"github.com/brandonc/advent2022/internal/ui"
)

type IntScanner struct {
	*bufio.Scanner
}

func NewIntScanner(reader io.Reader) IntScanner {
	return IntScanner{
		bufio.NewScanner(reader),
	}
}

func (i IntScanner) Int() int {
	item, err := strconv.Atoi(i.Text())
	if ui.Die(err) {
		panic("Expected int")
	}

	return item
}
