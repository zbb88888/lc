package main

import (
	"fmt"
)

func main() {

	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("input:", input)
	result := maxArea(input)
	println(result)
}

func maxArea(height []int) int {
	// 用双指针
	// 计算两个指针之间的水量
	// 双循环 穷举所有值
	area := 0
	h := 0
	w := 0
	l := len(height)
	maxH := 0
	for i, v := range height {
		// 倒指针移动
		if v <= maxH {
			continue
		}
		for ii := l - 1; ii > i; ii-- {
			vv := height[ii]
			if vv <= maxH {
				continue
			}
			if i == ii {
				continue
			}
			h = min(v, vv)
			if h > maxH {
				maxH = h
				w = ii - i
				tmp := maxH * w
				if area < tmp {
					area = tmp
					// fmt.Println("get bigger: ", area, "i", i, "ii", ii)
				}
			}
		}
	}
	return area
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
