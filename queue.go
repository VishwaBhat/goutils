package main

// Queue int variant of Queue implementation
type Queue struct {
	items []interface{}
}

// NewQueue creates new Queue
func NewQueue() *Queue {
	return &Queue{}
}

// Len provides length of the queue
func (s *Queue) Len() int {
	return len(s.items)
}

// Remove removes the element from top of the Queue
func (s *Queue) Remove() error {
	if s.Len() == 0 {
		return ErrNoElements
	}
	s.items = s.items[1:]
	return nil
}

// Add adds new element to the top of the Queue
func (s *Queue) Add(item interface{}) {
	s.items = append(s.items, item)
}

// Peek returns the top element of the Queue without removing it from the Queue
func (s *Queue) Peek() (interface{}, error) {
	if s.Len() == 0 {
		return nil, ErrNoElements
	}
	firstEl := s.items[0]
	s.items = s.items[1:]
	return firstEl, nil
}
