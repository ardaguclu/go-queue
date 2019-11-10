package queue

import (
	"container/list"
	"sync"
)

// CQList is container/list based thread safe queue implementation for concurrent access.
type CQList struct {
	Items *list.List
	L     *sync.Mutex
}

// NewCQList initializes and returns new CQList.
func NewCQList() *CQList {
	return &CQList{
		Items: new(list.List),
		L:     new(sync.Mutex),
	}
}

// Push adds new item tail of the queue.
func (l *CQList) Push(n interface{}) {
	l.L.Lock()
	defer l.L.Unlock()
	l.Items.PushBack(n)
}

// Pop returns first item from the queue.
func (l *CQList) Pop() (n interface{}) {
	if l.Len() > 0 {
		l.L.Lock()
		defer l.L.Unlock()
		e := l.Items.Front()
		l.Items.Remove(e)
		n = e.Value
	}

	return
}

// Len returns the item length of queue.
func (l *CQList) Len() int {
	l.L.Lock()
	defer l.L.Unlock()
	return l.Items.Len()
}
