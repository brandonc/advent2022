package day03

import (
	"strings"
	"testing"
)

func TestSampleIn(t *testing.T) {
	a1, a2, err := Factory().Solve(strings.NewReader(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`))

	if err != nil {
		t.Fatal("An error occurred:", err)
	}

	if a1 != 157 {
		t.Errorf("Expected a1 to be 157 but got %d", a1)
	}

	if a2 != 70 {
		t.Errorf("Expected a2 to be 70 but got %d", a2)
	}
}
