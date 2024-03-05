package smallest_subarray

// FindSubstring
/**
 * hard
 * 给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。

 * s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。

 * 例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
 * 返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。
 */
func FindSubstring(s string, words []string) []int {
	var indexes []int
	var str []byte
	wordSize := 0
	if len(words) == 0 {
		return indexes
	}
	if len(words[0]) == 0 {
		return indexes
	}

	wordsMap := make(map[string]int)
	for _, w := range words {
		wordSize = len(w)
		wordsMap[w]++
	}

	left, right, wordSum := 0, len(words)*wordSize-1, len(words)*wordSize

	for right <= len(s) {
		str = []byte(s[left:right])
		for left <= right-wordSum {
			if checkString(str, wordsMap, wordSize) {
				indexes = append(indexes, left)
			}
			left++
		}
		right++
	}
	return indexes
}

func checkString(str []byte, wordsMap map[string]int, size int) bool {
	wordMap := make(map[string]int)
	for i := 0; i < len(str); i += size {
		word := string(str[i : i+size])
		if _, ok := wordsMap[word]; ok {
			wordMap[word]++
		} else {
			return false
		}
	}

	for k, v := range wordsMap {
		if wordMap[k] < v {
			return false
		}
	}
	return true
}
