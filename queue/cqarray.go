package queue

import "sync"

// CQArray is slice based thread safe queue implementation for concurrent access.
type CQArray struct {
	Items []interface{}
	L     *sync.Mutex
}

// NewCQArray initializes and returns new CQArray.
func NewCQArray() *CQArray {
	return &CQArray{
		Items: make([]interface{}, 0, 1024),
		L:     new(sync.Mutex),
	}
}

// Push adds new item tail of the queue.
func (q *CQArray) Push(n interface{}) {
	q.L.Lock()
	defer q.L.Unlock()
	q.Items = append(q.Items, n)
}

// Pop returns first item from the queue.
func (q *CQArray) Pop() (n interface{}) {
	if q.Len() > 0 {
		q.L.Lock()
		defer q.L.Unlock()
		n = q.Items[0]
		q.Items[0] = nil
		q.Items = q.Items[1:]
	}

	return
}

// Len returns the item length of queue.
func (q *CQArray) Len() int {
	q.L.Lock()
	defer q.L.Unlock()
	return len(q.Items)
}
