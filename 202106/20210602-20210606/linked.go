package _0210602_20210606

import "plan"

//反转链表
func reverseNode(root *plan.Linked) *plan.Linked {
	if root == nil {
		return root
	}
	node := reverseNode(root.Next)
	root.Next.Next = root
	root.Next = nil
	return node
}

//回文链表判断
func isPalindrome(root *plan.Linked) bool {
	if root == nil {
		return false
	}
	fast, slow := root, root
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast == nil {
		slow = slow.Next
	}
	right := reverseNode(slow)
	left := root
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
