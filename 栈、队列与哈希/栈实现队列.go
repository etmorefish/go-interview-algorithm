package test_queue

import (
	"fmt"
	"testing"
)

// 结构体，包含两个一维切片
type GoQueue struct {
	stack []int
	back  []int
}

// 初始化，
func NewQueue() GoQueue {
	return GoQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}

func (q *GoQueue) Push(x int) {
	q.stack = append(q.stack, x)
}

func (q *GoQueue) Pop() int {
	if len(q.back) == 0 {
		for len(q.stack) != 0 {
			val := q.stack[len(q.stack)-1]
			q.stack = q.stack[:len(q.stack)-1] //切片，更新栈
			q.back = append(q.back, val)
		}
	}
	val := q.back[len(q.back)-1]
	q.back = q.back[:len(q.back)-1]
	return val
}

func (q *GoQueue) Empty() bool {
	return len(q.back) == 0 && len(q.stack) == 0
}

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	for !q.Empty() {
		fmt.Println(q.Pop())
	}
}
