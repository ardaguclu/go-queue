package queue

import (
	"container/list"
	"sync"
)

type QList struct {
	Items *list.List
	L     *sync.Mutex
}

func (l *QList) Push(n interface{}) {
	if l.L != nil {
		l.L.Lock()
		defer l.L.Unlock()
	}
	l.Items.PushBack(n)
}

func (l *QList) Pop() (n interface{}) {
	if l.Len() > 0 {
		if l.L != nil {
			l.L.Lock()
			defer l.L.Unlock()
		}
		e := l.Items.Front()
		l.Items.Remove(e)
		n = e.Value
	}

	return
}

func (l *QList) Len() int {
	if l.L != nil {
		l.L.Lock()
		defer l.L.Unlock()
	}
	return l.Items.Len()
}
