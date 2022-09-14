package main

import "fmt"

func singleNumbers(nums []int) []int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	x, y, tmp := 0, 0, 1
	for res&tmp == 0 { //与运算相同为1不同为0
		tmp <<= 1 //如果不同一直循环，tmp一直左移，直到第一个不为0的点
	}
	//找到第一个不为0的点，将数组分为两部分，分别异或，得出最终两个值
	for _, v := range nums {
		if v&tmp == 0 {
			x ^= v
		} else {
			y ^= v
		}
	}
	return []int{x, y}
}

func main() {
	nums := []int{1, 2, 10, 4, 1, 4, 3, 3}
	fmt.Println(singleNumbers(nums))
	fmt.Println(uint(0))
}
