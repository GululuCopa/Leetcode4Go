package smallest_subarray

// minSubArrayLen
/**
 * medium
 * 给定一个含有 n 个正整数的数组和一个正整数 target 。

 * 找出该数组中满足其总和大于等于 target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
 */
func minSubArrayLen(target int, nums []int) int {
	left, right, sum := 0, 1, 0
	subLength, result := 0, len(nums)+1
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			subLength = right - left + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if result == len(nums)+1 {
		return 0
	} else {
		return result
	}
}

func anotherSolution(target int, nums []int) int {
	left, right, sum := 0, 0, 0
	minLen, result := 0, len(nums)+1
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			minLen = right - left + 1
			if minLen < result {
				result = minLen
			}

			sum -= nums[left]
			left++
		}
		right++
	}

	if result == len(nums)+1 {
		return 0
	} else {
		return result
	}
}
