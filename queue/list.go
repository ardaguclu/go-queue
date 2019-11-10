package queue

import (
	"container/list"
	"sync"
)

type List struct {
	Items *list.List
	L     *sync.Mutex
}

func (l *List) Push(n interface{}) {
	if l.L != nil {
		l.L.Lock()
		defer l.L.Unlock()
	}
	l.Items.PushBack(n)
}

func (l *List) Pop() (n interface{}) {
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

func (l *List) Len() int {
	if l.L != nil {
		l.L.Lock()
		defer l.L.Unlock()
	}
	return l.Items.Len()
}
