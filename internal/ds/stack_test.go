package ds

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push("hello")
	s.Push("world")
	s.Push("this")
	s.Push("is")
	s.Push("homer")

	t.Run("Len returns length", func(t *testing.T) {
		if s.Len() != 5 {
			t.Errorf("Expected length 5, got %d", s.Len())
		}
	})

	t.Run("Peek returns last item", func(t *testing.T) {
		if s.Peek() != "homer" {
			t.Errorf("Expected Peek to return homer, got %s", s.Peek())
		}
	})

	t.Run("Unshift inserts an item into the bottom", func(t *testing.T) {
		s := NewStack()
		s.Push("hello")
		s.Push("world")
		s.Unshift("surprise!")

		var r string
		if r = s.Pop(); r != "world" {
			t.Errorf("Expected r to equal world, got %s", r)
		}
		if r = s.Pop(); r != "hello" {
			t.Errorf("Expected r to equal hello, got %s", r)
		}
		if r = s.Pop(); r != "surprise!" {
			t.Errorf("Expected r to equal surprise, got %s", r)
		}
	})

	t.Run("Pop returns each item in FILO order", func(t *testing.T) {
		var r string
		if r = s.Pop(); r != "homer" {
			t.Errorf("Expected r to equal homer, got %s", r)
		}
		if r = s.Pop(); r != "is" {
			t.Errorf("Expected r to equal is, got %s", r)
		}
		if r = s.Pop(); r != "this" {
			t.Errorf("Expected r to equal this, got %s", r)
		}
		if r = s.Pop(); r != "world" {
			t.Errorf("Expected r to equal world, got %s", r)
		}
		if r = s.Pop(); r != "hello" {
			t.Errorf("Expected r to equal hello, got %s", r)
		}
		if s.Len() != 0 {
			t.Errorf("Expected length 0, got %d", s.Len())
		}
		if s.Peek() != "" {
			t.Errorf("Expected Peek to return empty, got %s", s.Peek())
		}
	})
}

func TestStackPopN(t *testing.T) {
	s := NewStack()
	s.Push("hello")
	s.Push("world")
	s.Push("this")
	s.PushN([]string{"is", "homer"})

	result := s.PopN(3)
	if result[0] != "this" && result[1] != "is" && result[2] != "homer" {
		t.Errorf("Expected PopN to return this is homer, got %s %s %s", result[0], result[1], result[2])
	}
}
