package binary_search

// MySqrt
/**
 * easy
 * 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

 * 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

 */
func MySqrt(x int) int {
	left, right := 0, x
	var result int

	for left <= right {
		middle := left + (right-left)>>1
		if middle*middle > x {
			right = middle - 1
		} else if middle*middle == x {
			result = middle
			break
		} else {
			left = middle + 1
			result = middle
		}
	}
	return result
}
