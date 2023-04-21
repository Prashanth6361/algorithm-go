package main

import (
	"ds/queue"
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
	a := queue.New(5)
	a.Enqueue(1)
	a.Enqueue(2)
	a.Enqueue(3)
	log.Println(a)
	v1, e1 := a.Dequeue()
	errorCheck(e1)
	log.Println(v1)
	log.Println("Length of queue:", a.Length())
	log.Println(a)
}
