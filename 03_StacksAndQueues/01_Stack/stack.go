package stack

import "errors"

type Stack struct {
	array []int
}

func (s *Stack) Empty() bool {
	return len(s.array) == 0
}

func (s *Stack) Push(x int) {
	s.array = append(s.array, x)
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, errors.New("underflow")
	}
	res := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return res, nil
}
