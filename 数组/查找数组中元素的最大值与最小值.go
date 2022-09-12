package main

import "fmt"

func FindSimple(nums []int) (max, min int) {
	max = nums[0]
	min = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return max, min
}

func main() {
	nums := []int{2, 9, 5, 3, 8, 1}
	fmt.Println(FindSimple(nums))
}
