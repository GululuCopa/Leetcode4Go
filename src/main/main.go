package main

import (
	"demo1/src/leetcode/array/smallest_subarray"
	"fmt"
)

func main() {
	longest := smallest_subarray.LongestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 0)
	fmt.Println(longest)
}
