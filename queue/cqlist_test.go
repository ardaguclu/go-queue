package queue

import (
	"container/list"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCQListPush(t *testing.T) {
	cqList := &CQList{
		Items: new(list.List),
		L:     new(sync.Mutex),
	}

	for i := 0; i < 100; i++ {
		cqList.Push(fmt.Sprintf("Item-%d", i))
		assert.Equal(t, i+1, cqList.Len())
	}
}

func TestCQListPop(t *testing.T) {
	tmpList := new(list.List)
	tmpList.PushBack("test1")
	tmpList.PushBack("test2")
	cqList := &CQList{
		Items: tmpList,
		L:     new(sync.Mutex),
	}

	val := cqList.Pop()
	assert.Equal(t, val.(string), "test1")
	val = cqList.Pop()
	assert.Equal(t, val.(string), "test2")
}

func BenchmarkCQListPush(b *testing.B) {
	cqList := &CQList{
		Items: new(list.List),
		L:     new(sync.Mutex),
	}

	for i := 0; i < b.N; i++ {
		cqList.Push(fmt.Sprintf("Item-%d", i))
	}
}

func BenchmarkCQListPop(b *testing.B) {
	tmpList := new(list.List)
	for i := 0; i < 10000; i++ {
		tmpList.PushBack(fmt.Sprintf("test-%d", i))
	}
	cqList := &CQList{
		Items: tmpList,
		L:     new(sync.Mutex),
	}

	for i := 0; i < b.N; i++ {
		cqList.Pop()
	}
}
