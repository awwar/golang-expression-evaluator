package utility

import "fmt"

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]*T, 0)}
}

type Stack[T any] struct {
	items []*T
}

func (s *Stack[T]) Push(data *T) {
	s.items = append(s.items, data)
}

func (s *Stack[T]) PushCopy(data T) {
	s.items = append(s.items, &data)
}

func (s *Stack[T]) Pop() (*T, error) {
	var result *T

	if s.IsEmpty() {
		return result, fmt.Errorf("stack is empty")
	}

	result = s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return result, nil
}

func (s *Stack[T]) IsEmpty() bool {
	if s == nil {
		return true
	}

	return len(s.items) == 0
}

func (s *Stack[T]) Iterate(callback func(T)) {
	for _, item := range s.items {
		callback(*item)
	}
}
