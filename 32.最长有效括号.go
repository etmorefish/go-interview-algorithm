/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */
package main

func main() {
	s := ")()())"
	longestValidParentheses(s)
}

// @lc code=start
func longestValidParentheses(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, right*2)
		} else if left < right {
			//从左往右遍历，首次出现“）”，一定是无效括号，重新计数
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, left*2)
		} else if left > right {
			//从右往左遍历，首次出现“（”，一定是无效括号，重新计数
			left, right = 0, 0
		}
	}
	return maxLength
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

// @lc code=end
