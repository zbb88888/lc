package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := []int{1, 0, -4, 3, -3, 2, 4, -2, 2, 2, 3, -4, -3, -1, -5, -1, 1}
	fmt.Println("input:", input)
	ret := threeSum(input)
	fmt.Println("ret: ", ret)
}

// 简化思路
// 先减去一个值，然后找两个数是否能够凑够剩下的和

func threeSum(nums []int) [][]int {
	// 用于存储 i 的 v：为 j + k
	var ret [][]int
	i, j, k := 0, 0, 0
	v, vv, sum := 0, 0, 0
	ijkV := make(map[string]bool)
	var key string
	var l2Builder, l3Builder strings.Builder
	for i = 0; i < len(nums); i++ {
		sum = 0 - nums[i]
		for j = i + 1; j < len(nums); j++ {
			v = nums[j]
			// 这里优先排序, 更容易匹配
			l2Got := []int{nums[i], v}
			slices.Sort(l2Got)
			l2Builder.Reset()
			l2Builder.WriteString(strconv.Itoa(l2Got[0]))
			l2Builder.WriteString(" ")
			l2Builder.WriteString(strconv.Itoa(l2Got[1]))
			key = l2Builder.String()
			if _, ok := ijkV[key]; ok {
				// fmt.Println("skip:", key)
				continue
			}
			for k = j + 1; k < len(nums); k++ {
				vv = nums[k]
				if sum == v+vv {
					got := []int{nums[i], v, vv}
					tmp := append(l2Got, vv)
					slices.Sort(tmp)
					// 只需要缓存两个值就足够
					l3Builder.Reset()
					l3Builder.WriteString(strconv.Itoa(tmp[0]))
					l3Builder.WriteString(" ")
					l3Builder.WriteString(strconv.Itoa(tmp[1]))
					// 这里是排序的结果
					key = l3Builder.String()
					if _, ok := ijkV[key]; ok {
						continue
					}
					ijkV[key] = true
					// fmt.Println("cache:", key)
					ret = append(ret, got)
				}
			}
		}
	}
	return ret
}
