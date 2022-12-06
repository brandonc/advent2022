package day06

import (
	"strings"
	"testing"
)

type testCase struct {
	sample  string
	answer1 int
	answer2 int
}

func TestSampleInput(t *testing.T) {
	cases := []testCase{
		{sample: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", answer1: 7, answer2: 19},
		{sample: "bvwbjplbgvbhsrlpgdmjqwftvncz", answer1: 5, answer2: 23},
		{sample: "nppdvjthqldpwncqszvftbrmjlhg", answer1: 6, answer2: 23},
		{sample: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", answer1: 10, answer2: 29},
		{sample: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", answer1: 11, answer2: 26},
	}

	solver := Factory()
	for _, c := range cases {
		a1, a2, err := solver.Solve(strings.NewReader(c.sample))
		if err != nil {
			t.Fatalf("Expected no error, got %s", err)
		}

		if a1 != c.answer1 {
			t.Errorf("Expected sample %s, part 1 to be %d, got %d", c.sample, c.answer1, a1)
		}

		if a2 != c.answer2 {
			t.Errorf("Expected sample %s, part 2 to be %d, got %d", c.sample, c.answer2, a2)
		}
	}
}
