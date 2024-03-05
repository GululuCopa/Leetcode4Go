package main

import (
	"demo1/src/leetcode/array/smallest_subarray"
	"fmt"
)

func main() {
	longest := smallest_subarray.FindSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})
	fmt.Println(longest)
}
