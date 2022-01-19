package main

import (
	"fmt"
	"math"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	groupLen := numRows*2 - 2
	groupNum := int(math.Ceil(float64(len(s)) / float64(groupLen)))
	var ansString []byte

	for i := 0; i < numRows; i++ {
		// 计算第 i 行字符串
		for j := 0; j < groupNum; j++ {
			// 计算第 j 组字符串
			charIndex := groupLen*j + i
			if charIndex >= len(s) {
				continue
			}
			ansString = append(ansString, s[charIndex])
			if i != 0 && i != numRows-1 {
				charIndex = groupLen*(j+1) - i
				if charIndex < len(s) {
					ansString = append(ansString, s[charIndex])
				}
			}
		}

	}
	return string(ansString)

}

func main() {
	s := "PAYPALISHIRING"
	numRows := 4
	fmt.Println(convert(s, numRows))
}
