package remove_elements

// BackspaceCompare
/**
 * 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。

 * 注意：如果对空文本输入退格字符，文本继续为空。
 */
func BackspaceCompare(s string, t string) bool {
	return buildStr(s) == buildStr(t)
}

func buildStr(str string) string {
	var slices []byte
	left, right := 0, 0
	for right < len(str) {
		if str[right] == '#' {
			if left == 0 {
				right++
				continue
			} else {
				left--
				slices = slices[0:left]
			}
		} else {
			slices = append(slices, str[right])
			left++
		}
		right++
	}

	return string(slices[0:left])
}
