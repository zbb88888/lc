package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	nums := []int{-1000, -1000, -1000}
	target := 10000
	ret := threeSumClosest(nums, target)
	fmt.Println("ret: ", ret)
}

// 先排序
// i j k
// i 从前往后
// j 从 i + 1 开始
// k 从 len-1 开始
// 因为前面的是负值，后面的是正值，可以尽快匹配
// 如果 i > 0, 后面都 > 0  不可能继续匹配到了

func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	slices.Sort(nums)
	// fmt.Println("nums:", nums, "target", target)
	var sum, minSub, res, tmpSub int
	minSub = math.MaxInt
	for i := 0; i < l; i++ {
		j := i + 1
		k := l - 1
		for j < k {
			sum = nums[i] + nums[j] + nums[k]
			// fmt.Println("sum ", sum, "=", nums[i], nums[j], nums[k])
			if sum > target {
				// 正数太大了
				k--
				tmpSub = sum - target
			} else if sum < target {
				// 负数太小了
				j++
				tmpSub = target - sum
			} else {
				return sum
			}
			if tmpSub < minSub {
				minSub = tmpSub
				res = sum
			}
		}
	}
	return res
}
