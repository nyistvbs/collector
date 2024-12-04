package queue

import (
	"container/list"
	"fmt"
)

type Queue struct {
	items *list.List
}

func New() *Queue {
	return &Queue{items: list.New()}
}

// 入队操作
func (q *Queue) Enqueue(item interface{}) {
	q.items.PushBack(item)
}

// 出队操作
func (q *Queue) Dequeue() (interface{}, error) {
	if q.items.Len() == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	// 从队列头部删除元素
	front := q.items.Front()
	q.items.Remove(front)

	return front.Value, nil
}

// 获取队列的长度
func (q *Queue) Length() int {
	return q.items.Len()
}

// 查看队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.items.Len() == 0
}
