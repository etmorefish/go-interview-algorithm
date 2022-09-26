/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 思路
外层循环：指针 i 遍历数组。
内层循环：用双指针，去寻找满足三数之和 == 0 的元素

先排序的意义
便于跳过重复元素，如果当前元素和前一个元素相同，跳过。

双指针的移动时，避免出现重复解
找到一个解后，左右指针同时向内收缩，为了避免指向重复的元素，需要：

左指针在保证left < right的前提下，一直右移，直到指向不重复的元素
右指针在保证left < right的前提下，一直左移，直到指向不重复的元素
小优化
排序后，如果外层遍历的数已经大于0，则另外两个数一定大于0，sum不会等于0，直接break。

*/

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 排序
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ { // 外层遍历
		n1 := nums[i]
		if n1 > 0 {
			break // 如果已经爆0，不用做了，break
		}
		if i > 0 && n1 == nums[i-1] {
			continue // 遍历到重复的数，跳过
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++ // 直到指向不一样的数
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++ // 三数和小于0，则左指针右移
			} else {
				r-- // 三数和大于0，则右指针左移
			}
		}
	}
	return res
}

// @lc code=end

