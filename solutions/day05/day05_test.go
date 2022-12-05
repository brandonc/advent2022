package day05

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, a2, err := Factory().Solve(strings.NewReader("    [D]    " + `
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`))

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if a1 != "CMZ" {
		t.Errorf("Expected a1 to be \"CMZ\", got %s", a1)
	}

	if a2 != "MCD" {
		t.Errorf("Expected a2 to be \"MCD\", got %s", a2)
	}
}
