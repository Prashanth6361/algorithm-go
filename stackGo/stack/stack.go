package stack

import "errors"

type Stack struct {
	size  int
	value []interface{}
}

func New(size int) *Stack {
	return &Stack{
		size:  size,
		value: make([]interface{}, 0),
	}
}

func (s *Stack) Push(v interface{}) error {
	l := len(s.value)
	if l == s.size {
		return errors.New("Full stack")
	}
	s.value = append(s.value, v)
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	n := len(s.value)
	if n == 0 {
		return nil, errors.New("Empty stack")
	}
	top := s.value[n-1]
	s.value = s.value[:n-1]
	return top, nil
}

func (s *Stack) Peek() (interface{}, error) {
	n := len(s.value)
	if n == 0 {
		return nil, errors.New("Empty stack")
	}
	temp := s.value[n-1]
	return temp, nil
}
