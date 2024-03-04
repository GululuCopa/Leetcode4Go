package smallest_subarray

// MinWindow
/**
 * 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

 * 注意：

 * 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 * 如果 s 中存在这样的子串，我们保证它是唯一的答案。
 */
func MinWindow(s string, t string) string {
	left, right := 0, 0
	minLength := len(s) + 1
	resultStart := 0

	tStrMap := make(map[string]int)
	for _, tStr := range t {
		tStrMap[string(tStr)]++
	}

	sStrMap := make(map[string]int)

	for right < len(s) {
		_, ok := tStrMap[string(s[right])]
		if len(sStrMap) == 0 && !ok {
			right++
			left++
			continue
		}

		sStrMap[string(s[right])]++
		for check(&sStrMap, &tStrMap) && left <= right {
			if right-left+1 < minLength {
				minLength = right - left + 1
				resultStart = left
			}
			s2 := string(s[left])
			minusOne(&sStrMap, &s2)
			left++
		}
		right++
	}
	if minLength == len(s)+1 {
		return ""
	}
	return s[resultStart : resultStart+minLength]
}

func check(sStrMap *map[string]int, tStrMap *map[string]int) bool {
	for k, v := range *tStrMap {
		if (*sStrMap)[k] < v {
			return false
		}
	}
	return true
}

func minusOne(dataMap *map[string]int, key *string) {
	_, ok := (*dataMap)[(*key)]
	if ok {
		(*dataMap)[(*key)]--
	}
}
