// Package ds provides some simple data structures
package ds

import "fmt"

type Stack struct {
	array []string
}

func (s Stack) Len() int {
	return len(s.array)
}

func (s Stack) Peek() string {
	if len(s.array) == 0 {
		return ""
	}
	return s.array[len(s.array)-1]
}

func NewStack() Stack {
	return Stack{
		array: make([]string, 0),
	}
}

func (s *Stack) Pop() string {
	return s.PopN(1)[0]
}

func (s *Stack) PopN(n int) []string {
	if len(s.array) < n {
		panic(fmt.Sprintf("cannot pop %d items from stack len %d", n, len(s.array)))
	}

	result := s.array[len(s.array)-n:]
	s.array = s.array[:len(s.array)-n] // subject to memory leak
	return result
}

func (s *Stack) Push(item string) {
	s.array = append(s.array, item)
}

func (s *Stack) PushN(items []string) {
	s.array = append(s.array, items...)
}

func (s *Stack) Unshift(item string) {
	if len(s.array) == 0 {
		s.Push(item)
		return
	}

	newArray := make([]string, len(s.array)+1)
	if copied := copy(newArray[1:], s.array); copied != len(s.array) {
		panic(fmt.Sprintf("expected to copy %d items but got %d", len(s.array), copied))
	}
	newArray[0] = item
	s.array = newArray
}
