/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 反转链表，返回反转后的头和尾，便于拼接链表
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	cur := head
	for pre != tail {
		next := cur.Next //暂存后继节点 cur.next
		cur.Next = pre   //修改 Next 引用指向
		pre = cur        //pre 暂存 cur
		cur = next       //cur 访问下一节点
	}
	return tail, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		nex := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return dummy.Next
}

// @lc code=end

