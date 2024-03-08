package main

import (
	"demo1/src/leetcode/array/smallest_subarray"
	"fmt"
)

func main() {
	longest := smallest_subarray.CharacterReplacement("AABABBA", 1)
	fmt.Println(longest)
}
