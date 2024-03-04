package binary_search

// isPerfectSquare
/**
 * 367
 * 给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

 * 完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。

 */
func isPerfectSquare(num int) bool {
	left, right := 0, num

	for left <= right {
		middle := left + (right-left)>>1
		if middle*middle > num {
			right = middle - 1
		} else if middle*middle < num {
			left = middle + 1
		} else {
			return true
		}
	}
	return false
}
