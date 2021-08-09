package main

// Stack int variant of Stack implementation
type Stack struct {
	items []interface{}
	top   int
}

// NewStack creates new stack
func NewStack() *Stack {
	return &Stack{
		top: -1,
	}
}

// Pop removes the element from top of the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.top == -1 {
		return nil, ErrNoElements
	}
	el := s.items[s.top]
	s.top--
	s.items = s.items[:(s.top)]
	return el, nil
}

// Push adds new element to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
	s.top++
}

// Peek returns the top element of the stack without removing it from the stack
func (s *Stack) Peek() (interface{}, error) {
	if s.top == -1 {
		return nil, ErrNoElements
	}
	el := s.items[s.top]
	return el, nil
}
