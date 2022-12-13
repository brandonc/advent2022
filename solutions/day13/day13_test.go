package day13

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, a2, err := Factory().Solve(strings.NewReader(`[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`))

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if a1 != 13 {
		t.Errorf("Expected a1 to be 13, got %d", a1)
	}

	if a2 != 140 {
		t.Errorf("Expected a2 to be 140, got %d", a2)
	}
}
