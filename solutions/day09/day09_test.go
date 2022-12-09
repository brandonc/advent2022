package day09

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, a2, err := Factory().Solve(strings.NewReader(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`))

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if a1 != 13 {
		t.Errorf("Expected a1 to be 13, got %d", a1)
	}

	if a2 != 1 {
		t.Errorf("Expected a2 to be 1, got %d", a2)
	}
}
