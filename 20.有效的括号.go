/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	var stack []rune = make([]rune, 0, len(s)/2)

	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
			continue
		case ')':
			if len(stack) == 0 || '(' != stack[len(stack)-1] {
				return false
			}
		case ']':
			if len(stack) == 0 || '[' != stack[len(stack)-1] {
				return false
			}
		case '}':
			if len(stack) == 0 || '{' != stack[len(stack)-1] {
				return false
			}
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

// @lc code=end
