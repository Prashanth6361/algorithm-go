package main

import "fmt"

type Stack struct {
	words []string
}

func (s *Stack) Push(i string) {
	s.words = append(s.words, i)
}

func (s *Stack) Pop() string {
	n := len(s.words)
	if n == 0 {
		return "Empty data"
	}
	top := s.words[n-1]
	s.words = s.words[:n-1]
	return top
}

func (s *Stack) Peek() string {
	n := len(s.words)
	temp := s.words[n-1]
	return temp
}

func main() {
	a := Stack{}
	a.Push("Google")
	a.Push("language")
	a.Push("Golang")
	fmt.Println(a)
	fmt.Println(a.Pop())
	fmt.Println(a.Peek())
	fmt.Println(a)
}
