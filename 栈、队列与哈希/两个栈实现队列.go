package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AppendTail(3)
	fmt.Println(obj.DeleteHead())

	obj.AppendTail(2)
	obj.AppendTail(1)
	fmt.Println(obj.instack)
	fmt.Println(obj.outstack)
}

type CQueue struct {
	instack, outstack []int
}

func Constructor() CQueue {
	return CQueue{
		instack:  []int{},
		outstack: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.instack = append(this.instack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outstack) == 0 {
		if len(this.instack) == 0 {
			return -1
		}
		for len(this.instack) > 0 {
			index := len(this.instack) - 1
			this.outstack = append(this.outstack, this.instack[index])
			this.instack = this.instack[:index]
		}
	}
	index := len(this.outstack) - 1
	val := this.outstack[index]
	this.outstack = this.outstack[:index]
	return val
}
