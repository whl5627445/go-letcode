package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 题目可以理解成寻找两个数组合并之后的中位数， 因为两个数组长度固定，那么中位数的位置也是已知的，所以可以更简化为寻找第k个最小值，官方的这部分解释很好理解
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		// 当第一个数组已经取值完毕，那么第k项必定在第二个数组中，可以直接取值
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		// 同上
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		// 这里是当k的位置在两个数组中已经缩小到2个数， 取最小的那个
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		// 下面的部分对官方答案进行了简化， 比较好理解一点， 大概意思就是谁小把谁踢掉，踢到第k个为止
		data1 := nums1[min(index1, len(nums1)-1)]
		data2 := nums2[min(index2, len(nums2)-1)]
		if data1 <= data2 {
			index1 += 1
		} else {
			index2 += 1
		}
		k -= 1
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}
	fmt.Println(findMedianSortedArrays(a, b))
}
