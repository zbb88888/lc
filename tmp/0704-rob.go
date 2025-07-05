package main

import (
	"fmt"
	"slices"
)

func main() {
	input := []int{183, 219, 57, 193, 94, 233, 202, 154, 65, 240, 97, 234, 100, 249, 186, 66, 90, 238, 168, 128, 177, 235, 50, 81, 185, 165, 217, 207, 88, 80, 112, 78, 135, 62, 228, 247, 211}
	fmt.Println("input:", input)
	result := rob(input)
	println("result:", result)
}

// sorts := slices.Clone(nums)
// slices.Sort(sorts)
// fmt.Println("sorts:", sorts)

// 是不是抢劫路线只有从第一家开始和从第二家开始两种选择？ 这是相隔1家的路线选择
// 不相邻问题的最小问题是: 1/1, 1/2, 1/3 or 2/3. 这些情况都被奇偶路线涵盖
// 如果超过3，那就只能从走下述逻辑
// 以上不全面：
// 还有第三种路线：比如当长度为 4 时，只取首尾
// 穷举：最小窗口问题：
// 应该是 3个中选一个，然后选定之后，再从新的三个中选一个
// 三个中最大是左边，意味着中间只能抛弃，最右的可选可不选（从这里开始再找两个，凑够三个继续判断）
// 最大是中间， 如果是中间，则左右都抛弃，直接往下再选三个
// 最大的不可能是右边，因为这样就要先选最左边
// 如果存在三个数：
// 一次判断，要么选最左（左边只要大于中间），要么选中间（中间>左右之和）

func rob(nums []int) int {
	reversed := slices.Clone(nums)
	slices.Reverse(reversed)
	max0 := jiou(nums)
	max1 := rob3(nums)
	max2 := rob3(reversed)
	fmt.Println("max0:", max0, "max1:", max1, "max2", max2)
	if max0 >= max1 && max0 >= max2 {
		return max0
	}
	if max1 >= max0 && max1 >= max2 {
		return max1
	}
	if max2 >= max0 && max2 >= max1 {
		return max2
	}
	fmt.Println("err:", 0)
	return 0
}

func rob3(nums []int) int {
	maxPos := len(nums) - 1
	fmt.Println("nums:", nums, "maxPos:", maxPos)
	for i := 0; i <= maxPos; {
		left := nums[i]
		fmt.Println("i:", i, "v:", left, "picked", nums)
		if i+1 <= maxPos {
			mid := nums[i+1]
			// 左边只要大于中间，选左
			if left >= mid {
				i++
				// 直接把邻居置为0
				fmt.Println("1 set0", i)
				nums[i] = 0
				// 单独处理 0XXXX 的情况
				if maxPos-i == 4 {
					fmt.Println("last4s4:", nums[i+1:])
					last4(nums[i+1:])
					break
				}
				// 直接到邻居的右边，重头判定是否要抢劫
				i++
				continue
			} else {
				if i+2 <= maxPos {
					right := nums[i+2]
					// 中间>=左右之和，选中间
					if mid >= left+right {
						fmt.Println("2 set0", i)
						nums[i] = 0
						// 单独处理 0XXXX 的情况
						if maxPos-i == 4 {
							fmt.Println("last4s1:", nums[i+1:])
							last4(nums[i+1:])
							break
						}
						// 已经选定中间的，把其左右置为0
						i += 2
						// 直接把邻居置为0
						fmt.Println("3 set0", i)
						nums[i] = 0
						// 单独处理 0XXXX 的情况
						if maxPos-i == 4 {
							fmt.Println("last4s2:", nums[i+1:])
							last4(nums[i+1:])
							break
						}
						continue
					} else {
						// 中间 < 左右之和，选左边
						i++
						// 直接把邻居置为0
						fmt.Println("1 set0", i)
						nums[i] = 0
						// 单独处理 0XXXX 的情况
						if maxPos-i == 4 {
							fmt.Println("last4s3:", nums[i+1:])
							last4(nums[i+1:])
							break
						}
						// 直接到邻居的右边，重头判定是否要抢劫
						i++
						continue
					}
				} else {
					// 只剩下当前的左右两个
					// 谁大选哪个
					// 上面 if 已经判定过了 left >= mid
					// 这里在 else 里只能选右边的
					nums[i] = 0
					// 单独处理 0XXXX 的情况
					if maxPos-i == 4 {
						fmt.Println("last4s4:", nums[i+1:])
						last4(nums[i+1:])
						break
					}
					i++
				}
			}
		} else {
			fmt.Println("end")
			break
		}
	}

	fmt.Println("picked:", nums)
	max1 := 0
	for _, v := range nums {
		max1 += v
	}
	fmt.Println("max1:", max1)
	return max1
}
func last4(last4s []int) {
	// 0xxxx last 4
	// 这种情况下确实要多考虑一种组合
	if len(last4s) != 4 {
		return
	}
	shouwei := last4s[0] + last4s[3]
	yisan := last4s[0] + last4s[2]
	ersi := last4s[1] + last4s[3]
	fmt.Println("shouwei:", shouwei, "yisan", yisan, "ersi", ersi)
	if shouwei >= yisan && shouwei >= ersi {
		last4s[1] = 0
		last4s[2] = 0
	}
	if yisan >= shouwei && yisan >= ersi {
		last4s[1] = 0
		last4s[3] = 0
	}
	if ersi >= shouwei && ersi >= yisan {
		last4s[0] = 0
		last4s[2] = 0
	}
}

func jiou(nums []int) int {
	// 返回奇偶数的最大值
	jishuhe := 0
	oushuhe := 0
	for i, v := range nums {
		if i%2 == 0 {
			oushuhe += v
		} else {
			jishuhe += v
		}
	}
	max0 := 0
	// max0 路线 1 的最大值
	if jishuhe >= oushuhe {
		max0 = jishuhe
	} else {
		max0 = oushuhe
	}
	fmt.Println("jishuhe:", jishuhe, "oushuhe", oushuhe)
	return max0
}
