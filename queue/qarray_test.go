package queue

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQArrayPush(t *testing.T) {
	qArr := &QArray{
		Items: nil,
		L:     nil,
	}

	for i := 0; i < 100; i++ {
		qArr.Push(fmt.Sprintf("Item-%d", i))
		assert.Equal(t, i+1, len(qArr.Items))
	}
}

func TestQArrayPushWithLock(t *testing.T) {
	qArr := &QArray{
		Items: nil,
		L:     new(sync.Mutex),
	}

	for i := 0; i < 100; i++ {
		qArr.Push(fmt.Sprintf("Item-%d", i))
		assert.Equal(t, i+1, len(qArr.Items))
	}
}

func TestQArrayPop(t *testing.T) {
	tmpItems := make([]interface{}, 0, 2)
	tmpItems = append(tmpItems, "test1")
	tmpItems = append(tmpItems, "test2")
	qArr := &QArray{
		Items: tmpItems,
		L:     nil,
	}

	val := qArr.Pop()
	assert.Equal(t, val.(string), "test1")
	val = qArr.Pop()
	assert.Equal(t, val.(string), "test2")
}

func TestQArrayPopWithLock(t *testing.T) {
	tmpItems := make([]interface{}, 0, 2)
	tmpItems = append(tmpItems, "test1")
	tmpItems = append(tmpItems, "test2")
	qArr := &QArray{
		Items: tmpItems,
		L:     new(sync.Mutex),
	}

	val := qArr.Pop()
	assert.Equal(t, val.(string), "test1")
	val = qArr.Pop()
	assert.Equal(t, val.(string), "test2")
}

func BenchmarkQArrayPush(b *testing.B) {
	qArr := &QArray{
		Items: nil,
		L:     nil,
	}

	for i := 0; i < b.N; i++ {
		qArr.Push(fmt.Sprintf("Item-%d", i))
	}
}

func BenchmarkQArrayPushWithLock(b *testing.B) {
	qArr := &QArray{
		Items: nil,
		L:     new(sync.Mutex),
	}

	for i := 0; i < b.N; i++ {
		qArr.Push(fmt.Sprintf("Item-%d", i))
	}
}

func BenchmarkQArrayPop(b *testing.B) {
	tmpItems := []interface{}{"test1", "test2", "test3", "test4", "test5", "test6", "test7", "test8", "test9", "test10"}
	qArr := &QArray{
		Items: tmpItems,
		L:     nil,
	}

	for i := 0; i < b.N; i++ {
		qArr.Pop()
	}
}

func BenchmarkQArrayPopWithLock(b *testing.B) {
	tmpItems := []interface{}{"test1", "test2", "test3", "test4", "test5", "test6", "test7", "test8", "test9", "test10"}
	qArr := &QArray{
		Items: tmpItems,
		L:     new(sync.Mutex),
	}

	for i := 0; i < b.N; i++ {
		qArr.Pop()
	}
}
