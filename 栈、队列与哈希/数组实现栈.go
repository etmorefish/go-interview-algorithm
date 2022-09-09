package main

import (
	"errors"
	"fmt"
)

/* 实现一个栈的数据结构，功能：压栈、弹栈、取出栈顶元素、判断栈是否为空、元素个数

可以把数组的首元素当作栈底，同时记录栈中元素的个数size，假设数组首地址是arr，
则压栈的操作其实就是待压栈的元素放到数组arr[size]中，然后执行size++；
同理出栈，其实就是取数组arr[size-1]元素，然后执行size--

*/

type SliceStack struct {
	arr       []int
	stackSize int
}

func (s *SliceStack) IsEmpty() bool {
	return s.stackSize == 0
}

func (s *SliceStack) Size() int {
	return s.stackSize
}

func (s *SliceStack) Top() int {
	if s.IsEmpty() {
		panic(errors.New("empty stack!"))
	}
	return s.arr[s.stackSize-1]
}

func (s *SliceStack) Pop() int {
	if s.stackSize > 0 {
		s.stackSize--
		res := s.arr[s.stackSize]
		s.arr = s.arr[:s.stackSize]
		return res
	}
	panic(errors.New("empty stack!"))
}

func (s *SliceStack) Push(i int) {
	s.arr = append(s.arr, i)
	s.stackSize++
}

func (s *SliceStack) Status() []int {
	return s.arr
}

func SliceMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("slice init stackstruct start...")
	st := &SliceStack{arr: make([]int, 0)}
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
