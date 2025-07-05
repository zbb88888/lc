package main

import (
	"fmt"
)

func main() {
	input := []int{183, 219, 57, 193, 94, 233, 202, 154, 65, 240, 97, 234, 100, 249, 186, 66, 90, 238, 168, 128, 177, 235, 50, 81, 185, 165, 217, 207, 88, 80, 112, 78, 135, 62, 228, 247, 211}
	fmt.Println("input:", input)
	result := rob(input)
	println("result:", result)
}

func rob(nums []int) int {

	// 问题划分
	// len 为 1 返回 nums[0]
	// len 为 2 返回 较大的那个
	// len >= 3 时，其实有可能出现
	// 其实根据这三个值没有办法确定到底选中间还是左右两个
	// 到底选哪个只有等到第四个值被加上的时候才能确定是选13还是24
	// 这道题使用动态规划： 会同时保存最大的值和次大的值，直到最后才能确定最大的路线是哪一个

	// 理解动态规划：
	// 在这个例子中，在抢劫问题中：列表中的每一个值，都有抢和不抢两种选择，而且当前无法判定是否抢还是不抢，所以只能都抢
	// 然后根据后续的值的相关关系，不断丢弃之前的状态
	// 动态规划其实就是分身：用不同分身应付不同情况，不同分身就是缓存，目的是为了解决当前无法解决的问题，所以只能先把当前的状态存下来，留于后续可以判断的时候再判断
	len := len(nums)
	if len == 0 {
		return 0
	}
	if len == 1 {
		return nums[0]
	}
	if len == 2 {
		return max(nums[0], nums[1])
	}
	less := nums[0]
	more := max(nums[0], nums[1])
	for i := 2; i < len; i++ {
		less, more = more, max(less+nums[i], more)
		println("less:", less, "more: ", more)
	}

	if more >= less {
		// 不到最后，其实无法确认哪个最多
		return more
	}
	return less
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
