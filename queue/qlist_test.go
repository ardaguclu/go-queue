package queue

import (
	"container/list"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQListPush(t *testing.T) {
	qList := &QList{
		Items: new(list.List),
		L:     nil,
	}

	for i := 0; i < 100; i++ {
		qList.Push(fmt.Sprintf("Item-%d", i))
		assert.Equal(t, i+1, qList.Len())
	}
}

func TestQListPushWithLock(t *testing.T) {
	qList := &QList{
		Items: new(list.List),
		L:     new(sync.Mutex),
	}

	for i := 0; i < 100; i++ {
		qList.Push(fmt.Sprintf("Item-%d", i))
		assert.Equal(t, i+1, qList.Len())
	}
}

func TestQListPop(t *testing.T) {
	tmpList := new(list.List)
	tmpList.PushBack("test1")
	tmpList.PushBack("test2")
	qList := &QList{
		Items: tmpList,
		L:     nil,
	}

	val := qList.Pop()
	assert.Equal(t, val.(string), "test1")
	val = qList.Pop()
	assert.Equal(t, val.(string), "test2")
}

func TestQListPopWithLock(t *testing.T) {
	tmpList := new(list.List)
	tmpList.PushBack("test1")
	tmpList.PushBack("test2")
	qList := &QList{
		Items: tmpList,
		L:     new(sync.Mutex),
	}

	val := qList.Pop()
	assert.Equal(t, val.(string), "test1")
	val = qList.Pop()
	assert.Equal(t, val.(string), "test2")
}

func BenchmarkQListPush(b *testing.B) {
	qList := &QList{
		Items: nil,
		L:     nil,
	}

	for i := 0; i < b.N; i++ {
		qList.Push(fmt.Sprintf("Item-%d", i))
	}
}

func BenchmarkQListPushWithLock(b *testing.B) {
	qList := &QList{
		Items: nil,
		L:     new(sync.Mutex),
	}

	for i := 0; i < b.N; i++ {
		qList.Push(fmt.Sprintf("Item-%d", i))
	}
}

func BenchmarkQListPop(b *testing.B) {
	tmpList := new(list.List)
	tmpList.PushBack("test1")
	tmpList.PushBack("test2")
	tmpList.PushBack("test3")
	tmpList.PushBack("test4")
	tmpList.PushBack("test5")
	tmpList.PushBack("test6")
	tmpList.PushBack("test7")
	tmpList.PushBack("test8")
	tmpList.PushBack("test9")
	tmpList.PushBack("test10")
	qList := &QList{
		Items: tmpList,
		L:     nil,
	}

	for i := 0; i < b.N; i++ {
		qList.Pop()
	}
}

func BenchmarkQListPopWithLock(b *testing.B) {
	tmpList := new(list.List)
	tmpList.PushBack("test1")
	tmpList.PushBack("test2")
	tmpList.PushBack("test3")
	tmpList.PushBack("test4")
	tmpList.PushBack("test5")
	tmpList.PushBack("test6")
	tmpList.PushBack("test7")
	tmpList.PushBack("test8")
	tmpList.PushBack("test9")
	tmpList.PushBack("test10")
	qList := &QList{
		Items: tmpList,
		L:     new(sync.Mutex),
	}

	for i := 0; i < b.N; i++ {
		qList.Pop()
	}
}
