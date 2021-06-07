package _0210607_20210614

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
