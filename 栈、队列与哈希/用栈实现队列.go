package main

import "fmt"

// 可以通过切片来模拟出队列。 只需要完成简单的 push，pop，再判断是否为空即可！

type Queue struct {
	stack []int
	back  []int
}

func NewQueue() Queue {
	return Queue{
		stack: []int{},
		back:  []int{},
	}
}

func (q *Queue) Push(val int) {
	q.stack = append(q.stack, val)
}

func (q *Queue) Pop() int {
	if len(q.back) == 0 {
		for len(q.stack) > 0 {
			val := q.stack[len(q.stack)-1]
			q.stack = q.stack[:len(q.stack)-1]
			q.back = append(q.back, val)
		}
	}
	val := q.back[len(q.back)-1]
	q.back = q.back[:len(q.back)-1]
	return val
}

func (q *Queue) IsEmpty() bool {
	return len(q.stack) == 0 && len(q.back) == 0
}

func main() {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	for !q.IsEmpty() {
		fmt.Println(q.Pop())
	}
}
