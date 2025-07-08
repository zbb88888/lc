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
// 如果要求排序的话，先排序，然后再找两个数，等 nums[i] > 0 就可以结束了
// 先减去一个值，然后找两个数是否能够凑够剩下的和
// 如果要求不排序的话，这里就已经是最佳方案

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
	slices.Sort(nums) // 先排序
	// fmt.Println("sorted nums:", nums)
	var newNums []int
	match0 := 0
	last2Same := false
	newNums = append(newNums, nums[0])
	last := 0
	for i = 1; i < l; i++ {
		// 预处理
		// 快速跳过连续重复的值
		// 当重复的值超过两个时，才可以跳过
		// 0 最多连续保留四个
		if nums[i] == 0 {
			if match0 >= 4 {
				continue
			}
			match0++
			newNums = append(newNums, nums[i])
			last++
			continue
		}
		// 其他的值最多连续保留两个
		if newNums[last] != nums[i] {
			newNums = append(newNums, nums[i])
			last++
			last2Same = false
		} else {
			if last2Same {
				continue
			} else {
				newNums = append(newNums, nums[i])
				last++
				last2Same = true
				continue
			}
		}
	}

	// fmt.Println("newNums:", newNums)
	l = len(newNums)
	for i = 0; i < l; i++ {
		if newNums[i] > 0 {
			// fmt.Println("v > 0")
			return ret
		}
		sum = 0 - newNums[i]
		for j = i + 1; j < l; j++ {
			// 这里优先排序, 更容易匹配
			l2Got = []int{newNums[i], newNums[j]}
			slices.Sort(l2Got)
			l2Builder.Reset()
			l2Builder.WriteString(strconv.Itoa(l2Got[0]))
			l2Builder.WriteString(" ")
			l2Builder.WriteString(strconv.Itoa(l2Got[1]))
			if _, ok := ijkV[l2Builder.String()]; ok {
				// fmt.Println("skip:", key)
				continue
			}
			l3v := sum - newNums[j]
			for k = j + 1; k < l; k++ {
				if l3v == newNums[k] {
					tmp := append(l2Got, newNums[k])
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
					ret = append(ret, []int{newNums[i], newNums[j], newNums[k]})
				}
			}
		}
	}
	return ret
}
