/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, jin int
	var res = &ListNode {
		Val: 0,
		Next: nil
	}

	o1, o2, r := l1, l2, res

	for {
		sum, jin = sum+jin, 0

		if nil != o1 {
			sum += o1.Val
			o1 = o1.Next
		}

		if nil != o2 {
			sum += o2.Val
			o2 = o2.Next
		}

		r.Val = sum % 10
		jin = sum / 10

		if nil == o1 && nil == o2 {
			break
		}
	}

	return res
}

// @lc code=end
