/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i, j := 0, len(nums)-1
	for i < j {
		for i < len(nums) && nums[i] != val {
			i++
		}

		for j > 0 && nums[j] == val {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	if nums[j] == val {
		return j
	} else {
		return j + 1
	}
}

// @lc code=end
