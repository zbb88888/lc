package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // just nil
	needTwo := false
    len1 := len(nums1)
    len2 := len(nums2)
	if len1 == 0 && len2 == 0 {
		return float64(0)
	}
	
	// get mids
    total := len1 + len2
	mid1 := 0
	mid2 := 0
	if total % 2 == 0 {
		needTwo = true
		mid := total/2
		mid1 = mid
		mid2 = mid+1
		println("should find two pos: ", mid1, mid2)
	} else {
		mid1 = (total/2) + 1 // 向下取整导致-1，补回来
		println("should find one pos: ", mid1)
	}
	
	// -1 means not found
    result1 := 0
    result2 := 0

	gotResult1 := false

	var longNums, shortNums []int
	var longLen, shortLen int
	longEnd := false
	shortEnd := false
	
	if len1 <= len2 {
		longNums = nums2
		longLen = len2
		shortLen = len1
		shortNums = nums1
	} else {
		longNums = nums1
		longLen = len1
		shortLen = len2
		shortNums = nums2
	}
	
	fmt.Printf("longNums: %v \n", longNums)
	fmt.Printf("shortNums: %v \n", shortNums)
	
	var goby int
	// case 1: just one list, go until the mid
	if shortLen == 0 {
		goby = 1 // goby, find one smaller, goby ++
		for _, v := range longNums {
			if !gotResult1 {
				if goby == mid1 {
					result1 = v
					println("1. get result1", result1)
					if !needTwo {
						println("1. end result1", result1)
						return float64(result1)
					}
					gotResult1 = true
				}
				goby ++
				continue
			}
			// find result2 after got result1
			if goby == mid2 {
				result2 = v
				println("1. end result1 result2", result1, result2)
				return float64(result1 + result2) / 2
			}
			goby ++
		}
	}
	
	// case 2: two list
	// 每一次比较都必然导致一次 goby ++
	// 如果当前的 v 小于 vv，则 goby v
	// 如果当前 vv 小于 v， 则 goby vv
	// 每一次 goby，都确定一个小值，谁小，步进一
	// 如果一个数组已经 goby 到了最后，则回退到 case 1
	
	
	// 优先使用当前 for range list 的下一个 v 去继续比较：
	// 当前循环步进优先： 直到在当前循环通过步进，找到一个比另一个循环体的 v 更大的值，再进入另一循环体
	
	// end, go by v 的时候，v 的下标已经是数组最后一个下标
	
	vvIndex := 0
	var ii, vv int
	goby = 1
	var longStop, shortStop int // 只剩下一个时，从剩下的那个的停留位置开始
	for i, v := range longNums {
		if longEnd || shortEnd {
			println("longEnd, shortEnd, longStop, shortStop, goby： ", longEnd, shortEnd, longStop, shortStop, goby)
			break
		}
		vv = shortNums[vvIndex]
		if !gotResult1 {
			// find result1
			if v <= vv {
				if goby == mid1 {
					result1 = v
					println("2.1 v is the result1", result1, "at ", i)
					if !needTwo {
						println("2.1 end result1", result1)
						return float64(result1)
					}
					gotResult1 = true
				}
				if i == (longLen -1 ) {
					println("2.1 longEnd")
					longEnd = true
					longStop = i
					shortStop = vvIndex
					break
				}
				goby++
				println("2.1. goby i v vv", i, v, vv)
				i++ // current v is smaller
				continue // prefer: next v compare with current vv
			} else {
				if goby == mid1 {
					result1 = vv
					println("2.2 vv is the result1", result1, "at ", vvIndex)
					if !needTwo {
						println("2.2 end result1", result1)
						return float64(result1)
					}
					gotResult1 = true
				}
				if vvIndex == (shortLen -1 ) {
					println("2.2 shortEnd")
					shortEnd = true
					longStop = i
					shortStop = vvIndex
					break
				}
				goby++
				println("2.2 goby i v vv", i, v, vv)
				vvIndex++ // current vv is smaller
				// goto l2: current v compare with next vv
			}
		} else {
			// find result2 after got result1
			if v <= vv {
				if goby == mid2 {
					result2 = v
					println("2.3 end the result1, result2", result1, result2)
					return float64(result1 + result2) / 2
				}
				if i == (longLen -1 ) {
					println("2.3 longEnd")
					longEnd = true
					longStop = i
					shortStop = vvIndex
					break
				}
				goby++
				println("2.3 goby i v vv", i, v, vv)
				i++ // current v is smaller
				continue // prefer: next v compare with current vv
			} else {
				if goby == mid2 {
					result2 = vv
					println("2.4 end the result1, result2", result1, result2)
					return float64(result1 + result2) / 2
				}
				if vvIndex == (shortLen -1 ) {
					println("2.4 shortEnd")
					shortEnd = true
					longStop = i
					shortStop = vvIndex
					break
				}
				goby++
				println("2.4 goby i v vv", i, v, vv)
				vvIndex++ // current vv is smaller
				// goto l2: current v compare with next vv
			}
		}
		println("l2 loop: i ii v vv", i, vvIndex, v, vv)
		// 4. go to l2 loop, v compare with next vv
		for ii = vvIndex; ii < shortLen; {
			vv = shortNums[vvIndex]
			if !gotResult1 {
				// find result1
				if vv <= v {
					if goby == mid1 {
						result1 = vv
						println("4.1 vv is the result1", result1, "at ", goby)
						if !needTwo {
							println("4.1 end result1", result1)
							return float64(result1)
						}
						gotResult1 = true
					}
					if ii == (shortLen-1) {
						println("4.1 vv shortEnd")
						shortEnd = true
						longStop = i
						shortStop = ii
						break
					}
					goby++
					println("4.1 goby ii v vv", ii, v, vv)
					ii ++ // current vv is smaller
					continue // prefer: next vv compare with current v
				} else {
					if goby == mid1 {
						result1 = v
						println("4.2 v is the result1", result1, "at ", goby)
						if !needTwo {
							println("4.2 end result1", result1)
							return float64(result1)
						}
						gotResult1 = true
					}
					vvIndex = ii 
					if i == (longLen -1) {
						longEnd = true
						longStop = i
						shortStop = ii
						break
					}
					goby++
					println("4.2 goby ii v vv", ii, v, vv)
					i++ // current v is smaller
					break // goto l1: current vv compare with next v
				}
			} else {
				if vv <= v {
					if goby == mid2 {
						result2 = vv
						println("4.3 vv end the result1, result2", result1, result2)
						return float64(result1 + result2) / 2
					}
					if i == (longLen -1 ) {
						println("4.3 longEnd")
						longEnd = true
						longStop = i
						shortStop = ii
						break
					}
					goby ++
					ii ++ // current vv is smaller
					println("4.3 goby ii v vv", ii, v, vv)
					continue // prefer: next vv compare with current v
				} else {
					if goby == mid2 {
						result2 = v
						println("4.4 v end the result1, result2", result1, result2)
						return float64(result1 + result2) / 2
					}
					vvIndex = ii // current vv compare with next v
					if i == (longLen -1) {
						longEnd = true
						longStop = i
						shortStop = ii
						break
					}
					goby++
					println("4.4 goby ii v vv", ii, v, vv)
					i++ // current v is smaller
					break // goto l1: current vv compare with next v
				}
			}
		}
	}
	
	// 5. longEnd, 循环 shortNums 剩下的部分肯定能找到
	if longEnd {
		println("5. just loop short: result1, result2, goby", result1, result2, goby)
		println("5. longStop shortStop", longStop, shortStop)
		for i:=shortStop; i < shortLen; i++ {
			goby ++
			if !gotResult1 && goby == mid1 {
					result1 = longNums[i]
					println("5. get result1", result1)
					if !needTwo {
						println("5. end result1", result1)
						return float64(result1)
					}
					continue
			}
			if goby == mid2 {
				result2 = longNums[i]
				println("5. end result1 result2", result1, result2)
				return float64(result1 + result2) / 2
			}
		}
	}
	
	// 6. shortEnd 已经到了最后, 循环 longNums 剩下的部分肯定能找到
	// 需要记录 longNums 已经走到的位置，并从那个位置继续，不然 goby 了之前的路
	if shortEnd {
		println("6. just loop long: result1, result2, goby", result1, result2, goby)
		println("6. longStop shortStop", longStop, shortStop)
		for i:=longStop; i < longLen; i++ {
			goby ++
			if !gotResult1 && goby == mid1 {
					result1 = longNums[i]
					println("6. get result1", result1)
					if !needTwo {
						println("6. end result1", result1)
						return float64(result1)
					}
					continue
			}
			if goby == mid2 {
				result2 = longNums[i]
				println("6. end result1 result2", result1, result2)
				return float64(result1 + result2) / 2
			}
		}
	}
	return float64(0.0)
}

func main() {
	nums1 := []int{3}
	nums2 := []int{-2,-1}
	res := findMedianSortedArrays(nums1, nums2)
	println(res)
}
