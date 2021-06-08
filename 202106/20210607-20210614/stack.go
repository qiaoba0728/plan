package _0210607_20210614

import "strings"

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
