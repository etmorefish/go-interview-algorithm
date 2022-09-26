/*
  - @lc app=leetcode.cn id=20 lang=golang
    *
  - [20] 有效的括号
    有效括号字符串的长度，一定是偶数！

右括号前面，必须是相对应的左括号，才能抵消！
右括号前面，不是对应的左括号，那么该字符串，一定不是有效的括号！
*/
package main

import "fmt"

func main() {
	s := "{[()]}"
	isValid(s)
}

// @lc code=start
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	m := map[byte]byte{'}': '{', ')': '(', ']': '['}
	stack := []byte{}
	for i := 0; i < n; i++ {
		fmt.Println(s[i])
		if m[s[i]] > 0 { // 匹配到对应m中的key
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1] // 出栈
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// @lc code=end
