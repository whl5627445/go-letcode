package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	length := len(s)
	ans := 0
	cmap := make([]int, 128)
	left := 0
	for right := 0; right < length; right++ {
		index := s[right]
		left = max(left, cmap[index])
		ans = max(ans, right-left+1)
		cmap[index] = right + 1
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ss := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(ss))
}
