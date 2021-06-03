package plan

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ConnectTreeNode struct {
	Val   int
	Left  *ConnectTreeNode
	Right *ConnectTreeNode
	Next  *ConnectTreeNode
}
