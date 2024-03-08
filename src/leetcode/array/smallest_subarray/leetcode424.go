package smallest_subarray

// CharacterReplacement
/**
 * medium
 * 给你一个字符串 s 和一个整数 k 。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。该操作最多可执行 k 次。
 * 在执行上述操作后，返回 包含相同字母的最长子字符串的长度。
 */
func CharacterReplacement(s string, k int) int {

	left, right := 0, 0
	count, maxLength := 0, 0

	for right < len(s) {
		if s[left] != s[right] {
			count++
		}
		for count > k && left <= right {
			if s[left] == s[right] {
				count--
			} else {
				left++
			}
		}
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
		right++
	}
	return maxLength
}
