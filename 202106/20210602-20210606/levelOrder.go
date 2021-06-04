package _0210602_20210606

import "plan"

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）
func levelOrder(root *plan.TreeNode) [][]int {
	result := make([][]int, 0)
	dfs(root, &result, 0)
	return result
}
func dfs(root *plan.TreeNode, result *[][]int, level int) {
	if root != nil {
		if level == len(*result) {
			*result = append(*result, []int{})
		}
		(*result)[level] = append((*result)[level], root.Val)
		dfs(root.Left, result, level+1)
		dfs(root.Right, result, level+1)
	}
}

//翻转二叉树
func invertTree(root *plan.TreeNode) *plan.TreeNode {
	if root == nil {
		return root
	}
	temp := root.Left.Val
	root.Left.Val = root.Right.Val
	root.Right.Val = temp
	invertTree(root.Left)
	invertTree(root.Right)
}

//填充二叉树节点的右侧指针
func connect(root *plan.ConnectTreeNode) *plan.ConnectTreeNode {
	if root == nil {
		return root
	}
	connectTwoNodes(root.Left, root.Right)
	return root
}
func connectTwoNodes(node1 *plan.ConnectTreeNode, node2 *plan.ConnectTreeNode) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectTwoNodes(node1.Left, node1.Right)
	connectTwoNodes(node2.Left, node2.Right)
	connectTwoNodes(node1.Right, node2.Left)
}

//将二叉树展开为链表
func flatten(root *plan.TreeNode) *plan.TreeNode {
	if root == nil {
		return root
	}
	left := flatten(root.Left)
	right := flatten(root.Right)
	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
	return root
}

//给定一个不含重复元素的数组，构造最大二叉树
func constructMaximumBinaryTree(nums []int) *plan.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max, index := nums[0], -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	node := &plan.TreeNode{
		Val:   index,
		Left:  nil,
		Right: nil,
	}
	node.Left = constructMaximumBinaryTree(nums[:index])
	node.Right = constructMaximumBinaryTree(nums[index+1:])
	return node
}
