/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */
package main

func main() {
	findMin([]int{3, 4, 5, 1, 2})
}

// @lc code=start
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for high > low {
		tmp := low + (high-low)/2
		if nums[tmp] < nums[high] {
			high = tmp
		} else {
			low = tmp + 1
		}
	}
	return nums[low]
}

// @lc code=end
