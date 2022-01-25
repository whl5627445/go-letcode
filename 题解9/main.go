package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	// a := strconv.Itoa(x)
	// len_a := len(a)
	// for i := 0; i < len_a; i++ {
	// 	if len_a == 1 {
	// 		return true
	// 	}
	// 	if a[i] != a[len_a-i-1] {
	// 		return false
	// 	}
	// 	if i >= len_a-i-1 {
	// 		break
	// 	}
	// }
	// return true
	if x < 0 {
		return false
	}
	a := x
	b := 0
	for x != 0 {
		b = b*10 + x%10
		x = x / 10
	}
	return a == b
}

func main() {
	fmt.Println(isPalindrome(1000021))
}
