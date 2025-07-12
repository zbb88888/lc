package main

import (
	"fmt"
	"slices"
)

func main() {
	input := []int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}
	fmt.Println("input:", input)
	ret := threeSum(input)
	fmt.Println("ret: ", ret)
}

// 先排序
// i j k
// i 从前往后
// j 从 i + 1 开始
// k 从 len-1 开始
// 因为前面的是负值，后面的是正值，可以尽快匹配
// 如果 i > 0, 后面都 > 0  不可能继续匹配到了
//

func threeSum(nums []int) [][]int {
	// special
	var res [][]int
	l := len(nums)
	if l < 3 {
		return nil
	}
	slices.Sort(nums)
	fmt.Println("nums:", nums)

	var key string
	ijk := make(map[string]bool)
	for i := 0; i < l; i++ {
		if nums[i] > 0 {
			return res
		}
		j := i + 1
		k := l - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				key = fmt.Sprintf("%d %d %d", nums[i], nums[j], nums[j])
				if _, ok := ijk[key]; ok {
					// dup
					// only move k
					k--
					continue
				}
				res = append(res, []int{nums[i], nums[j], nums[k]})
				ijk[key] = true
				if j < k {
					if nums[i] == nums[i+1] {
						// 同值快进
						j++
					}
					if nums[k] == nums[k-1] {
						// 同值快进
						j--
					}
				}
				j++
				k--
			} else if nums[i]+nums[j]+nums[k] > 0 {
				// 正数太大了
				k--
			} else {
				// 负数太小了
				j++
			}
		}
	}
	return res
}
