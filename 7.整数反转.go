/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	var i int
	for ; x != 0; x /= 10 {
		i = i*10 + (x % 10)
		if int(int32(i)) != i {
			return 0
		}
	}

	return i
}

// @lc code=end
