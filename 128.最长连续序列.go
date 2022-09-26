/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */
package main

import "fmt"

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}

	fmt.Println(longestConsecutive(nums))
}

// @lc code=start
func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, nums := range nums {
		numSet[nums] = true
	}
	maxLength := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentnum := num
			currentlen := 1
			for numSet[currentnum+1] {
				currentnum++
				currentlen++
			}
			if currentlen > maxLength {
				maxLength = currentlen
			}
		}
	}
	return maxLength
}

// @lc code=end
