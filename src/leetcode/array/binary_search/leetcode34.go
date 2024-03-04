package binary_search

// SearchRange
/**
 * 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

 * 如果数组中不存在目标值 target，返回 [-1, -1]。

 * 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
 */
func SearchRange(nums []int, target int) []int {
	left := queryShowTime(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{1, -1}
	}
	right := queryShowTime(nums, target+1) - 1
	return []int{left, right}
}

func queryShowTime(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

// AnotherSolve
/**
 * 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

 * 如果数组中不存在目标值 target，返回 [-1, -1]。

 * 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
 */
func AnotherSolve(nums []int, target int) []int {
	leftBorder, isLeftExist := getLeftBorder(nums, target)
	rightBorder, isRightExist := getRightBorder(nums, target)

	if !(isLeftExist && isRightExist) {
		return []int{-1, -1}
	}
	return []int{leftBorder, rightBorder}
}

func getRightBorder(nums []int, target int) (int, bool) {
	left, right := 0, len(nums)
	isRightExist := false
	for left < right {
		middle := left + (right-left)>>1
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] == target {
			isRightExist = true
			left = middle + 1
		} else {
			right = middle
		}
	}
	return right, isRightExist
}

func getLeftBorder(nums []int, target int) (int, bool) {
	left, right := 0, len(nums)
	isLeftExist := false
	for left < right {
		middle := left + (right-left)>>1
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] == target {
			isLeftExist = true
			right = middle
		} else {
			right = middle
		}
	}
	return left, isLeftExist
}
