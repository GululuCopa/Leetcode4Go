package remove_elements

// moveZeroes
/**
 * 283
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

 * 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

 */
// 两次遍历
func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
			right++
		} else {
			right++
		}
	}
	for left < len(nums) {
		nums[left] = 0
		left++
	}
}

func anotherMoveZeros(nums []int) {
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
