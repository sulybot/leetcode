/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if "" == needle {
		return 0
	}

outter:
	for i, hl, nl := 0, len(haystack), len(needle); i < hl-nl+1; i++ {
		for j := 0; j < nl; j++ {
			if needle[j] != haystack[i+j] {
				continue outter
			}
		}

		return i
	}

	return -1
}

// @lc code=end
