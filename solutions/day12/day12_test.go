package day12

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, a2, err := Factory().Solve(strings.NewReader(`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`))

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if a1 != 31 {
		t.Errorf("Expected a1 to be 31, got %d", a1)
	}

	if a2 != 29 {
		t.Errorf("Expected a2 to be 29, got %d", a1)
	}
}
