package stack

import "errors"

type Stack struct {
	array [100]int
	top   int
}

func NewStack() *Stack {
	return &Stack{top: -1}
}

func (s *Stack) Empty() bool {
	return s.top == -1
}

func (s *Stack) Push(x int) {
	s.top = s.top + 1
	s.array[s.top] = x
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, errors.New("underflow")
	}
	res := s.array[s.top]
	s.top = s.top - 1
	return res, nil
}
