package queue

import "sync"

type QArray struct {
	Items []interface{}
	L     *sync.Mutex
}

func (q *QArray) Push(n interface{}) {
	if q.L != nil {
		q.L.Lock()
		defer q.L.Unlock()
	}
	q.Items = append(q.Items, n)
}

func (q *QArray) Pop() (n interface{}) {
	if q.Len() > 0 {
		if q.L != nil {
			q.L.Lock()
			defer q.L.Unlock()
		}
		n = q.Items[0]
		q.Items[0] = nil
		q.Items = q.Items[1:]
	}

	return
}

func (q *QArray) Len() int {
	if q.L != nil {
		q.L.Lock()
		defer q.L.Unlock()
	}
	return len(q.Items)
}
