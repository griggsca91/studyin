package stack

import (
	"fmt"
	"strings"
)

type Stack[T any] interface {
	Pop() (T, error)
	Add(T)
}

type ArrayStack[T any] struct {
	data []T
	top  int
	size int
}

func NewArrayStack[T any](size int) ArrayStack[T] {
	return ArrayStack[T]{
		data: make([]T, size),
		top:  -1,
	}
}

func (s *ArrayStack[T]) Pop() (T, error) {
	if s.top == -1 {
		return *new(T), fmt.Errorf("nothing to pop")
	}

	value := s.data[s.top]
	s.top--

	return value, nil
}

func (s *ArrayStack[T]) Add(value T) {
	s.top++
	if s.top >= len(s.data) {
		s.data = append(s.data, value)
	} else {
		s.data[s.top] = value
	}
}

func (s ArrayStack[T]) String() string {
	var strs []string
	for _, value := range s.data[0 : s.top+1] {
		strs = append(strs, fmt.Sprintf("%v", value))
	}

	return fmt.Sprintf(
		"len = %v cap = %v top=%v [%s]",
		len(s.data),
		cap(s.data),
		s.top,
		strings.Join(strs, ", "),
	)
}
