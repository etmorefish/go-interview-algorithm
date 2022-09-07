/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {return list2}
	if list2 == nil {return list1}
	var dummy *ListNode
	if list1.Val < list2.Val {
		dummy = list1
		dummy.Next = mergeTwoLists(list1.Next, list2)
	} else {
		dummy = list2
		dummy.Next = mergeTwoLists(list1, list2.Next)
	}
	return dummy
}

// @lc code=end

