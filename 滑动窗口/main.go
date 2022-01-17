package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	s_map := map[string]int{}
	left := 0
	right := 0
	ans := 0
	for index, _ := range s {
		if _, ok := s_map[string(s[index])]; !ok {
			s_map[string(s[index])] = index
		} else {
			if s_map[string(s[index])] >= left {
				left = s_map[string(s[index])] + 1
			}
			s_map[string(s[index])] = index
		}
		l := s[left : index+1]
		if len(l) > ans {
			ans = len(l)
		}
		if right+1 < len(s) {
			right += 1
		}

	}
	// return ans
	fmt.Println(ans)
}
