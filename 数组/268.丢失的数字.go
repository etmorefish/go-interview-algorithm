/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start

func missingNumber(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return total - sum
}

// @lc code=end

func missingNumber1(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	for i := 0; i <= len(nums); i++ {
		result ^= i
	}
	return result
}

func missingNumber2(nums []int) int {
	sort.Ints(nums)
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return len(nums)
}