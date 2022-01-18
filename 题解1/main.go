package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	s_map := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		res := target - nums[i]
		if _, ok := s_map[res]; ok {
			return []int{s_map[res], i}
		}
		s_map[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{12, 4, 5, 7, 8, 85, 3, 325, 456, 67, 23, 14, 43, 6}
	res := twoSum(nums, 30)
	fmt.Println(res)
}
