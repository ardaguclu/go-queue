# Go Queue Implementations

There is no queue implementation in the standard library. 
That's why, package queue includes two different types of queue implementations, 
which are slice based/list based and their benchmark results are shown below. 

Bearing mind that, neither container/list nor slice operations are thread safe. 
Thread safety supported versions are also implemented separately.

____

### USAGE

For slice based standard queue implementation;

    qArr := queue.NewQArray()
    qArr.Push("QArrayItem1")
    qArr.Push("QArrayItem2")
    qArr.Pop()
   
For slice based thread safe queue implementation;

    cqArr := queue.NewCQArray()
	cqArr.Push("CQArrayItem1")
	cqArr.Push("CQArrayItem2")
	cqArr.Pop()

For container/list based standard queue implementation;

    qList := queue.NewQList()
	qList.Push("QListItem1")
	qList.Push("QListItem2")
	qList.Pop()

For container/list based thread safe queue implementation;

    cqList := queue.NewCQList()
	cqList.Push("CQListItem1")
	cqList.Push("CQListItem2")
	cqList.Pop()

### BENCHMARK RESULTS

As shown benchmark results below; slice based queue implementation is faster than container/list based queue implementation. 
This case is acceptable for not only standard implementation but also thread safe implementation which is used mutex.

    BenchmarkCQArrayPush-8           3435589               330 ns/op
    BenchmarkCQListPush-8            3279596               345 ns/op
    BenchmarkQArrayPush-8            4522604               276 ns/op
    BenchmarkQListPush-8             4408953               340 ns/op

Slice based standard queue implementation is faster than container/list based. 
However, in this case clearly seen that there is a huge latency come from mutex usage for concurrency support. 
If the priority performance and no need for concurrency there is huge gain for standard queue implementations.

    BenchmarkCQArrayPop-8           23455449                47.9 ns/op
    BenchmarkCQListPop-8            24149019                47.2 ns/op
    BenchmarkQArrayPop-8            1000000000               0.538 ns/op
    BenchmarkQListPop-8             383273991                3.08 ns/op
