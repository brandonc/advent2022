package day02

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, a2, err := day02{}.Solve(strings.NewReader(`A Y
B X
C Z`))

	if err != nil {
		t.Fatal(err)
	}

	if a1 != 15 {
		t.Errorf("expected 15, got %d", a1)
	}

	if a2 != 12 {
		t.Errorf("expected 12, go %d", a2)
	}
}
