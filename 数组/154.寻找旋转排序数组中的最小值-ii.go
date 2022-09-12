/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */
package main

func main() {
	findMin([]int{2, 2, 2, 0, 1})
}

// @lc code=start
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for high > low {
		tmp := low + (high-low)/2
		if nums[tmp] < nums[high] {
			high = tmp
		} else if nums[tmp] > nums[high] {
			low = tmp + 1
		} else {
			high--
		}
	}
	return nums[low]
}

// @lc code=end
