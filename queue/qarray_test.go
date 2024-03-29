package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQArrayPush(t *testing.T) {
	qArr := &QArray{
		Items: nil,
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
	}

	val := qArr.Pop()
	assert.Equal(t, val.(string), "test1")
	val = qArr.Pop()
	assert.Equal(t, val.(string), "test2")
}

func BenchmarkQArrayPush(b *testing.B) {
	qArr := &QArray{
		Items: nil,
	}

	for i := 0; i < b.N; i++ {
		qArr.Push(fmt.Sprintf("Item-%d", i))
	}
}

func BenchmarkQArrayPop(b *testing.B) {
	tmpItems := make([]interface{}, 0, 1024)
	for i := 0; i < 10000; i++ {
		tmpItems = append(tmpItems, fmt.Sprintf("test-%d", i))
	}
	qArr := &QArray{
		Items: tmpItems,
	}

	for i := 0; i < b.N; i++ {
		qArr.Pop()
	}
}
