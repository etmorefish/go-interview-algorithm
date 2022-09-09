package main

import (
	"errors"
	"fmt"
)

/*
	实现一个栈的数据结构，功能：压栈、弹栈、取出栈顶元素、判断栈是否为空、元素个数

采用从头结点插入新结点的方法，使用带头结点的链表，可以保证对每个结点的操作都是相同的
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkStack struct {
	head *ListNode
}

func (p *LinkStack) IsEmpty() bool {
	return p.head.Next == nil
}

func (p *LinkStack) Size() int {
	size := 0
	node := p.head.Next
	for node != nil {
		size++
		node = node.Next
	}
	return size
}

func (p *LinkStack) Top() int {
	if p.head.Next != nil {
		return p.head.Next.Val
	}
	panic(errors.New("empty stack!"))
}

func (p *LinkStack) Pop() int {
	tmp := p.head.Next
	if tmp != nil {
		p.head.Next = tmp.Next
		return tmp.Val
	}
	panic(errors.New("empty stack!"))
}

func (p *LinkStack) Push(i int) {
	node := &ListNode{Val: i, Next: p.head.Next}
	p.head.Next = node
}

func (p *LinkStack) Status() []int {
	cur := p.head
	arr := []int{}
	for cur.Next != nil {
		arr = append(arr, cur.Next.Val)
		cur = cur.Next
	}
	return arr
}

func SliceMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("LinkNode init stackstruct start...")
	st := &LinkStack{head: &ListNode{}}
	st.Push(99)
	st.Push(78)
	st.Push(32)
	st.Push(12)
	fmt.Println("stack Top: ", st.Top())
	fmt.Println("stack Size:", st.Size())
	st.Pop()
	fmt.Println("stack Size:", st.Size())
	fmt.Println("stack Pop ok: ", st.Status())

}
func main() {
	SliceMode()

}
