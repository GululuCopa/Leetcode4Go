package smallest_subarray

// LengthOfLongestSubstring
/**
 * medium
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
 */
func LengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	maxLength := -1
	dataMap := make(map[byte]int)

	for right < len(s) {
		dataMap[s[right]]++
		for checkMap(dataMap) && left <= right {
			dataMap[s[left]]--
			if dataMap[s[left]] == 0 {
				delete(dataMap, s[left])
			}
			left++
		}
		maxLength = maxInteger(maxLength, right-left+1)
		right++
	}
	if maxLength == -1 {
		return 0
	}
	return maxLength
}

func checkMap(dataMap map[byte]int) bool {
	for _, v := range dataMap {
		if v > 1 {
			return true
		}
	}
	return false
}

// LengthOfLongestSubstring2
/**
 * 滑动窗口 map中存储的是字符和字符的下标
 */
func LengthOfLongestSubstring2(s string) int {
	dict := make(map[byte]int)
	ans := 0
	left := -1
	for right, _ := range s {
		if val, ok := dict[s[right]]; ok {
			// 更新左指针
			left = max(left, val)
		}
		// Hash 表中更新为最新的下标
		dict[s[right]] = right
		ans = max(ans, right-left)
	}
	return ans

}
