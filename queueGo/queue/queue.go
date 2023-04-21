package queue

import "errors"

type Queue struct {
	size  int
	value []interface{}
}

func New(size int) *Queue {
	return &Queue{
		size:  size,
		value: make([]interface{}, 0),
	}
}

func (a *Queue) Enqueue(v interface{}) error {
	l := a.Length()
	if l == a.size {
		return errors.New("Full queue")
	}
	a.value = append(a.value, v)
	return nil
}

func (a *Queue) Dequeue() (interface{}, error) {
	n := a.Length()
	if n == 0 {
		return nil, errors.New("Empty queue")
	}
	temp := a.value[0]
	a.value = a.value[1:]
	return temp, nil
}

func (a *Queue) Length() int {
	return len(a.value)
}
