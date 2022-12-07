package ds

import "testing"

func TestTree(t *testing.T) {
	tree := NewTree(0)
	a2 := tree.AddChild("a2", 0)
	a2.AddChild("a2_1", 0)

	if a2.GetChild("a2_1") == nil {
		t.Fatal("expected GetChild to not return nil")
	}

	if a2.GetChild("a2_1").Name != "a2_1" {
		t.Errorf("expected GetChild name to return a2_1, got %s", a2.GetChild("a2_1").Name)
	}
}
