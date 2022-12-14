package day14

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, a2, err := Factory().Solve(strings.NewReader(`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`))

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if a1 != 24 {
		t.Errorf("Expected a1 to be 24, got %d", a1)
	}

	if a2 != 93 {
		t.Errorf("Expected a2 to be 93, got %d", a1)
	}
}
