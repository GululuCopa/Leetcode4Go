package smallest_subarray

// LongestOnes
/**
 * medium
 * 给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。
 */
func LongestOnes(nums []int, k int) int {
	zeroCount := 0
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		}
	}
	if k >= zeroCount {
		return len(nums)
	}

	left, right := 0, 0
	t, maxLength := 0, -1

	for right < len(nums) {
		if nums[right] == 0 {
			t++
		}
		for t > k && left <= right {
			if nums[left] == 0 {
				t--
			}
			left++
		}

		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
		right++

	}
	return maxLength
}

// LongestOnes2
/**
 * 使用数组前缀和的方式求解，维护两个变量，分别代表left前缀和、right前缀和
 * 缘由：前缀和数组一定单调递增（数据为0或1）
 */
func LongestOnes2(nums []int, k int) (ans int) {
	var left, lsum, rsum int
	for right, v := range nums {
		rsum += 1 - v
		for lsum < rsum-k {
			lsum += 1 - nums[left]
			left++
		}
		ans = maxInteger(ans, right-left+1)
	}
	return
}

func maxInteger(a, b int) int {
	if a > b {
		return a
	}
	return b
}
