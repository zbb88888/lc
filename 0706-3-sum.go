package main

import (
	"fmt"
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
	for i, sum = range iNeed {
		fmt.Println("i", i)
		for j, v = range nums {
			fmt.Println("j", j)
			// j 比 i 快
			if j <= i {
				continue
			}
			for k, vv = range nums {
				fmt.Println("ijk", i, j, k)
				// k 指针跑最快
				if k <= i {
					continue
				}
				if k <= j {
					continue
				}
				if sum == v+vv {
					got := []int{0 - iNeed[i], v, vv}
					fmt.Println("got:", got)
					ret = append(ret, got)
				}
			}
		}
	}
	return ret
}
