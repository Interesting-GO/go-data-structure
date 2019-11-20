package queue

import (
	"sync"
)

type queueItem struct {
	Next *queueItem  // 她的下一级
	Data interface{} // 本节的数据
}

type queue struct {
	sync.Mutex     // 队列操作时加锁保证数据安全
	size       int // 定义队列中数据的长度

	head, tail *queueItem // 队列的头和尾巴
}

// 数据入队
func (q *queue) Append(data interface{}) {
	item := new(queueItem)
	q.Lock()
	defer q.Unlock()
	if q.head == nil {
		q.head = item
	}
	q.size++
	end := q.tail
	if q.tail != nil {
		end.Next = item
	}
	q.tail = item
	item.Data = data
}

// 队列是否为空
func (q *queue) Empty() bool {
	q.Lock()
	defer q.Unlock()
	if q.head == nil {
		return true
	}
	return false
}

// 返回队列长度
func (q *queue) Len() int {
	q.Lock()
	defer q.Unlock()
	return q.size
}

// 出列
func (q *queue) Next() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()
	if q.head == nil {
		return nil, false
	}
	q.size--
	item := q.head
	q.head = item.Next
	if q.tail == item {
		q.tail = nil
	}
	item.Next = nil
	return item.Data, true
}
