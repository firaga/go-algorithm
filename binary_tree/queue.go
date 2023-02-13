package binary_tree

import (
	"container/list"
	"sync"
)

type Queue struct {
	sync.RWMutex
	L *list.List
}

func NewQueue() *Queue {
	return &Queue{L: list.New()}
}

func (q *Queue) Add(v interface{}) *list.Element {
	q.Lock()
	defer q.Unlock()
	e := q.L.PushBack(v)
	return e
}

func (q *Queue) Get() interface{} {
	q.Lock()
	defer q.Unlock()
	if elem := q.L.Front(); elem != nil {
		item := q.L.Remove(elem)
		return item
	}
	return nil
}
