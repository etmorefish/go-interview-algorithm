package main

import "fmt"

func QuickSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	low := make([]int, 0, 0)
	high := []int{}
	mid := []int{}
	flag := nums[0]
	mid = append(mid, flag)
	for i := 1; i < len(nums); i++ {
		if nums[i] < flag {
			low = append(low, nums[i])
		} else if nums[i] > flag {
			high = append(high, nums[i])
		} else {
			mid = append(mid, nums[i])
		}
	}
	low, high = QuickSort(low), QuickSort(high)
	return append(append(low, mid...), high...)
}
func main() {
	nums := []int{2, 9, 5, 3, 8, 1}
	fmt.Println(QuickSort(nums))
	fmt.Println(QuickSort(nums)[len(nums)-1])

}
