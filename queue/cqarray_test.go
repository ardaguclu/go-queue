package queue

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCQArrayPush(t *testing.T) {
	cqArr := &CQArray{
		Items: nil,
		L:     new(sync.Mutex),
	}

	for i := 0; i < 100; i++ {
		cqArr.Push(fmt.Sprintf("Item-%d", i))
		assert.Equal(t, i+1, len(cqArr.Items))
	}
}

func TestCQArrayPop(t *testing.T) {
	tmpItems := make([]interface{}, 0, 2)
	tmpItems = append(tmpItems, "test1")
	tmpItems = append(tmpItems, "test2")
	cqArr := &CQArray{
		Items: tmpItems,
		L:     new(sync.Mutex),
	}

	val := cqArr.Pop()
	assert.Equal(t, val.(string), "test1")
	val = cqArr.Pop()
	assert.Equal(t, val.(string), "test2")
}

func BenchmarkCQArrayPush(b *testing.B) {
	cqArr := &CQArray{
		Items: nil,
		L:     new(sync.Mutex),
	}

	for i := 0; i < b.N; i++ {
		cqArr.Push(fmt.Sprintf("Item-%d", i))
	}
}

func BenchmarkCQArrayPop(b *testing.B) {
	tmpItems := make([]interface{}, 0, 1024)
	for i := 0; i < 10000; i++ {
		tmpItems = append(tmpItems, fmt.Sprintf("test-%d", i))
	}
	cqArr := &CQArray{
		Items: tmpItems,
		L:     new(sync.Mutex),
	}

	for i := 0; i < b.N; i++ {
		cqArr.Pop()
	}
}
