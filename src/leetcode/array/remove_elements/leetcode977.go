package remove_elements

// SortedSquares
/**
 * 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
 */
func SortedSquares(nums []int) []int {
	ans := make([]int, len(nums))

	left, right := 0, len(nums)-1
	index := len(nums) - 1
	for index >= 0 {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]
		if leftSquare > rightSquare {
			ans[index] = leftSquare
			left++
		} else {
			ans[index] = rightSquare
			right--
		}
		index--
	}
	return ans
}

// MergeSorting
/**
 * 归并排序

 * 总体思路：找到正负数的分界线neg，则num[0] 到 num[neg]单调递减，nums[neg] 到 nums[len(nums)]单调递增，即可排序
 */
func MergeSorting(nums []int) []int {

	length := len(nums)
	lastNegIndex := -1
	for i := 0; i < length && nums[i] < 0; i++ {
		lastNegIndex = i
	}
	ans := make([]int, 0, length)
	for i, j := lastNegIndex, lastNegIndex+1; i >= 0 || j < length; {
		if i < 0 {
			ans = append(ans, nums[j]*nums[j])
			j++
		} else if j == length {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] < nums[j]*nums[j] {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else {
			ans = append(ans, nums[j]*nums[j])
			j++
		}
	}

	return ans
}
