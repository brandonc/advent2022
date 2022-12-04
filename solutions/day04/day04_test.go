package day04

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, a2, err := Factory().Solve(strings.NewReader(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`))

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if a1 != 2 {
		t.Errorf("Expected a1 to be 2, got %d", a1)
	}

	if a2 != 4 {
		t.Errorf("Expected a2 to be 4, got %d", a1)
	}
}
