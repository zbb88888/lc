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
// 如果要求不排序的话，这里就已经是最佳方案
// 如果要求排序的话，先排序，然后再找两个数，等 nums[i] > 0 就可以结束了

func threeSum(nums []int) [][]int {
	// 用于存储 i 的 v：为 j + k
	var ret [][]int
	i, j, k := 0, 0, 0
	sum := 0
	ijkV := make(map[string]bool)
	var l2Builder, l3Builder strings.Builder
	l2Builder.Grow(10)
	l3Builder.Grow(10)
	var key string
	var l2Got []int
	l := len(nums)
	for i = 0; i < l; i++ {
		sum = 0 - nums[i]
		for j = i + 1; j < l; j++ {
			// 这里优先排序, 更容易匹配
			l2Got = []int{nums[i], nums[j]}
			slices.Sort(l2Got)
			l2Builder.Reset()
			l2Builder.WriteString(strconv.Itoa(l2Got[0]))
			l2Builder.WriteString(" ")
			l2Builder.WriteString(strconv.Itoa(l2Got[1]))
			if _, ok := ijkV[l2Builder.String()]; ok {
				// fmt.Println("skip:", key)
				continue
			}
			l3v := sum - nums[j]
			for k = j + 1; k < l; k++ {
				if l3v == nums[k] {
					tmp := append(l2Got, nums[k])
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
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return ret
}
