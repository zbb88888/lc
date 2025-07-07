package main

import (
	"fmt"
	"slices"
)

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("input:", input)
	ret := threeSum(input)
	fmt.Println("ret: ", ret)
}

// 简化思路
// 先减去一个值，然后找两个数是否能够凑够剩下的和

func threeSum(nums []int) [][]int {
	// 用于存储 i 的 v：为 j + k
	var ret [][]int
	// 用于存储 i 对应的 j k 的位置
	iNeed := make(map[int]int)
	for i, v := range nums {
		iNeed[i] = 0 - v
	}

	i, j, k := 0, 0, 0
	v, vv, sum := 0, 0, 0
	for i = 0; i < len(iNeed); i++ {
		// fmt.Println("i", i)
		sum = iNeed[i]
		for j = i + 1; j < len(nums); j++ {
			// fmt.Println("j", j)
			v = nums[j]
			for k = j + 1; k < len(nums); k++ {
				// k 指针跑最快
				vv = nums[k]
				// fmt.Println("ijk", i, j, k)
				if sum == v+vv {
					got := []int{0 - iNeed[i], v, vv}
					fmt.Println("got:", got)
					ret = append(ret, got)
				}
			}
		}
	}
	// 这个比对去重太慢了，比如对 [0 -1 1] 排序的处理建立一种字符串
	// 然后用map key 维护该字符串作为key，即可即时去重
	var ret1, sorts [][]int
	for _, v := range ret {
		duplicate := false
		tmp := slices.Clone(v)
		slices.Sort(tmp)
		for _, vv := range sorts {
			if slices.Equal(vv, tmp) {
				duplicate = true
				break
			}
		}
		if !duplicate {
			sorts = append(sorts, tmp)
			ret1 = append(ret1, v)
		}
	}
	return ret1
}
