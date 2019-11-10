package main

import (
	"container/list"
	"fmt"
	"sync"

	"github.com/ardaguclu/go-queue/queue"
)

func main() {
	qArr := &queue.Array{
		Items: nil,
		L:     new(sync.Mutex),
	}
	qArr.Push("ArrayItem1")
	qArr.Push("ArrayItem2")

	fmt.Println(qArr.Pop())

	qList := &queue.List{
		Items: list.New(),
		L:     new(sync.Mutex),
	}

	qList.Push("ListItem1")
	qList.Push("ListItem2")
	fmt.Println(qList.Pop())
}
