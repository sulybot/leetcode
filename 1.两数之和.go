/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		if v, ok := hashMap[target-num]; ok {
			return []int{v, i}
		}

		hashMap[num] = i
	}

	return nil
}

// @lc code=end
