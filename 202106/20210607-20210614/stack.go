package _0210607_20210614

import (
	"plan"
	"strings"
)

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效
func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	template := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]byte, 0)
	val := []byte(s)
	for i := 0; i < len(val); i++ {
		if _, ok := template[val[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != stack[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, val[i])
		}
	}
}

//输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理
func reverseWords(s string) string {
	fields := strings.Split(s, " ")
	var res []string
	for i := 0; i < len(fields); i++ {
		str := strings.TrimSpace(fields[i])
		if len(str) > 0 {
			res = append(res, fields[i])
		}
	}
	return strings.Join(res, " ")
}

//给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
func levelOrderBottom(root *plan.TreeNode) [][]int {
	res := make([][]int, 0)
	process(root, &res, 0)
	left, right := 0, len(res)-1
	for left < right {
		temp := res[left]
		res[left] = res[right]
		res[right] = temp
		left++
		right--
	}
	return res
}

func process(root *plan.TreeNode, res *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*res) > level {
		(*res)[level] = append((*res)[level], root.Val)
	} else {
		*res = append(*res, []int{root.Val})
	}
	process(root.Left, res, level+1)
	process(root.Right, res, level+1)
}

//todo
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度
func lengthOfLongestSubstring(s string) int {
	cnt := len(s)
	if cnt == 0 {
		return 0
	}

	tempMap := make(map[uint8]int)
	l, r, maxLen := 0, 1, 1
	tempMap[s[l]] = l
	for l < r && r < cnt {
		val, ok := tempMap[s[r]]
		if ok && val >= l {
			l = val + 1
		}

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}

		tempMap[s[r]] = r
		r++
	}

	return maxLen
}
