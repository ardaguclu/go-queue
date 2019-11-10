package queue

import (
	"container/list"
)

// QList is container/list based standard queue implementation without thread safety for concurrent access.
type QList struct {
	Items *list.List
}

// NewQList initializes and returns new QList.
func NewQList() *QList {
	return &QList{
		Items: new(list.List),
	}
}

// Push adds new item tail of the queue.
func (l *QList) Push(n interface{}) {
	l.Items.PushBack(n)
}

// Pop returns first item from the queue.
func (l *QList) Pop() (n interface{}) {
	if l.Len() > 0 {
		e := l.Items.Front()
		l.Items.Remove(e)
		n = e.Value
	}

	return
}

// Len returns the item length of queue.
func (l *QList) Len() int {
	return l.Items.Len()
}
