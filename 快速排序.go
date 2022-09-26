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

func BubbleSort(nums []int, k int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		fmt.Println(i+1, nums)
		k--
		if k == 0 {
			return nums
		}
	}
	return nums
}

func qk(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	l, m, h := []int{}, []int{}, []int{}
	flag := nums[0]
	m = append(m, flag)
	for i := 1; i < len(nums); i++ {
		if nums[i] > flag {
			h = append(h, nums[i])
		} else if nums[i] < flag {
			l = append(l, nums[i])
		} else {
			m = append(m, nums[i])
		}
	}
	l, h = qk(l), qk(h)
	return append(append(l, m...), h...)

}

func main() {
	nums := []int{2, 9, 5, 3, 8, 1}
	// fmt.Println(QuickSort(nums))
	// fmt.Println(QuickSort(nums)[len(nums)-1])
	// fmt.Println(BubbleSort(nums, 3))
	fmt.Println(qk(nums))

}
