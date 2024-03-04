package binary_search

/**
 *	给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
 *  写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
 */

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	index := -1
	for i, num := range nums {
		if num == target {
			index = i
		}
	}
	return index

}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := left + (right-left)/2
		if target > nums[middle] {
			left = middle + 1
		} else if target < nums[middle] {
			right = middle - 1
		} else {
			return middle
		}

	}
	return -1
}
