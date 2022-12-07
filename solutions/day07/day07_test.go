package day07

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, a2, err := Factory().Solve(strings.NewReader(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`))

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if a1 != 95437 {
		t.Errorf("Expected a1 to be 95437, got %d", a1)
	}

	if a2 != 24933642 {
		t.Errorf("Expected a2 to be 24933642, got %d", a2)
	}
}
