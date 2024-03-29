package queue

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQListPush(t *testing.T) {
	qList := &QList{
		Items: new(list.List),
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
	}

	val := qList.Pop()
	assert.Equal(t, val.(string), "test1")
	val = qList.Pop()
	assert.Equal(t, val.(string), "test2")
}

func BenchmarkQListPush(b *testing.B) {
	qList := &QList{
		Items: new(list.List),
	}

	for i := 0; i < b.N; i++ {
		qList.Push(fmt.Sprintf("Item-%d", i))
	}
}

func BenchmarkQListPop(b *testing.B) {
	tmpList := new(list.List)
	for i := 0; i < 10000; i++ {
		tmpList.PushBack(fmt.Sprintf("test-%d", i))
	}

	qList := &QList{
		Items: tmpList,
	}

	for i := 0; i < b.N; i++ {
		qList.Pop()
	}
}
