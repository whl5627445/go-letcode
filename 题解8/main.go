package main

import (
	"fmt"
)

func myAtoi(s string) int {
	z := 1
	i := 0
	var a []byte

	for i < len(s) {
		if string(s[i]) == " " {
			i++
		} else {
			break
		}
	}
	if len(s[i:]) == 0 {
		return 0
	}
	if string(s[i]) == "-" || string(s[i]) == "+" {
		if string(s[i]) == "-" {
			z = -1
		}
		i++
	}
	for i < len(s) {
		if 58 > s[i] && s[i] >= 48 {
			a = append(a, s[i])
			i++
		} else {
			break
		}
	}
	n := 0
	for _, b := range a {
		b -= '0'
		n = 10*n + int(b)
		if n*z > 2<<30-1 {
			return 2147483647
		}
		if n*z < -2<<30 {
			return -2147483648
		}
	}
	n = n * z
	return n
}
func main() {
	a := " "
	fmt.Println(myAtoi(a))
}
