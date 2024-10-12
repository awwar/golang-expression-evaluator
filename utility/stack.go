package utility

import "fmt"

type Stack[T any] struct {
	items []*T
}

func (s *Stack[T]) Push(data *T) {
	s.items = append(s.items, data)
}

func (s *Stack[T]) Pop() {
	if s.IsEmpty() {
		return
	}

	s.items = s.items[:len(s.items)-1]
}

func (s *Stack[T]) Top() (*T, error) {
	var result *T

	if s.IsEmpty() {
		return result, fmt.Errorf("stack is empty")
	}

	result = s.items[len(s.items)-1]

	return result, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
