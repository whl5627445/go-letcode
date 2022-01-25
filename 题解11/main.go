package main

import (
	"fmt"
)

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] > height[end] {
			high = height[end]
			end--
		} else {
			high = height[start]
			start++
		}
		if max < width*high {
			max = width * high
		}

	}
	return max
}

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(a))
}
