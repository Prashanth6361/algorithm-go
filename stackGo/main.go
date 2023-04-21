package main

import (
	"ds/stack"
	"log"
	"os"
)

func errorCheck(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	a := stack.New(5)
	a.Push("Google")
	a.Push("language")
	a.Push("Golang")
	log.Println(a)
	err := a.Push(3)
	errorCheck(err)
	v1, e1 := a.Pop()
	errorCheck(e1)
	log.Println(v1)
	v2, e2 := a.Peek()
	errorCheck(e2)
	log.Println(v2)
	log.Println(a)
}
