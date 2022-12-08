package day08

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, a2, err := Factory().Solve(strings.NewReader(`30373
25512
65332
33549
35390`))

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if a1 != 21 {
		t.Errorf("Expected a1 to be 21, got %d", a1)
	}

	if a2 != 8 {
		t.Errorf("Expected a2 to be 8, got %d", a2)
	}
}
