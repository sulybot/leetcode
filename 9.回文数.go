/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 || x != 0 && x%10 == 0 {
		return false
	}

	var i int
	for ; i < x; x /= 10 {
		i = i*10 + (x % 10)
	}

	return i == x || i/10 == x
}

// @lc code=end
