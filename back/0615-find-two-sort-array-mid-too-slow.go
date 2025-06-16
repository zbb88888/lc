package main

func main() {
	nums1 := []int{1,3}
	nums2 := []int{2,7}
	f := findMedianSortedArrays(nums1, nums2)
	println(f)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // get mid
	needTwo := false

    len1 := len(nums1)
    len2 := len(nums2)
	
	if len1 == 0 && len2 == 0 {
		return float64(0)
	}
    total := len1 + len2
	

	
	mid1 := 0
	mid2 := 0
	if total % 2 == 0 {
		needTwo = true
		mid := total/2
		mid1 = mid
		mid2 = mid+1
		println("should find two result at: ", mid1, mid2)
	} else {
		mid := (total/2)+1 // 向下取整，需要加1
		mid1 = mid
		println("should find one result at: ", mid1)
	}
    result1 := 0
    result2 := 0		
	

	onlyOne := false
	var onlyNums []int
	if len1 ==0 {
		onlyNums = nums2
	}
	if len2 ==0 {
		onlyNums = nums1
	}
	for i, v :=  range onlyNums {
		if needTwo {
			if i == mid1-1 {
				result1 = v
			}
			if i == mid2-1 {
				result2 = v
				onlyOne = true
			}
		} else {
			if i == mid1-1 {
				result1 = v
				onlyOne = true
			}
		}
	}
	if onlyOne {
		if needTwo {
			return float64(result1 + result2) / 2
		} 
		return float64(result1)
	}


	nextFrom := 0
	goby := 0 // goby, find one smaller, goby ++
	var v1, v2 int
	
	// for loop: smaller go first, bigger stop
	nums1End := false
    for i1 := 0; i1 < len(nums1); i1++ {
		v1 = nums1[i1]
		v2 = nums2[nextFrom]
		if result1 > 0 {
			// find result2
			if needTwo {
				if result2 > 0 {
					// end
					println("nums1 end, already has the result2", result2)
					break
				}
				if v1 <= v2 {
					// go
					goby ++
					if goby == mid2 {
						// should find the first result2
						result2 = v1
						println("nums1 end, find the result2", result2, "at n1 n2", i1, nextFrom)
						break
					}
					if i1 == (len(nums1)-1){
						nums1End = true
					}
				}
			} else {
				// end
				println("nums1 end", result1)
				break
			}
		} else {
			//  find result1
			if v1 <= v2 { // 每一次比较都必然导致一次 goby ++， 每一个 v 都需要比较两次，其实这是一个单循环
				// go
				goby ++
				if goby == mid1 {
					result1 = v1
					println("nums1 find the result1", result1, "at n1 n2", i1, nextFrom)
				}
				if i1 == (len(nums1)-1){
					nums1End = true
				}
			} 
		}
		
		// nums1 end, just go by nums2		
		if nums1End {
				for i2 := nextFrom; i2 < len(nums2); i2++ {
					v2 = nums2[i2]
					goby ++
					if result1 > 0 {
						println("just nums2 try find result2", i2)
						// find result2
						if needTwo {
							if result2 > 0 {
								println("just nums2 end, has the result2", result2)
								break
							}
							if goby == mid2 {
								// should find the first result2
								result2 = v2
								println("just nums2 end, find the result2", result2, "at n2", i2)
								break
							}
						} else {
							// only find one
							println("just nums2 end", result1)
							break
						}
					} else {
						//  find result1
						if goby == mid1 {
								result1 = v2
								println("just nums2 find the result1", result1, "at n2", i2)
							}
					}
				}
					
			} else {
					for i2 := nextFrom; i2 < len(nums2); i2++ {
						v2 = nums2[i2]
						if result1 > 0 {
							println("nums2 try find result2", i2)
							// find result2
							if needTwo {
								if result2 > 0 {
									println("nums2 end, has the result2", result2)
									break
								}
								if v2 <= v1 {
									// go
									goby ++
									nextFrom = i2
									if goby == mid2 && result2 < 0 {
										// should find the first result2
										result2 = v2
										println("nums2 end, find the result2", result2, "at n2", i2)
										break
									}
								}
							} else {
								// only find one
								println("nums2 end", result1)
								break
							}
						} else {
							//  find result1
							if v2 <= v1  {
								// go
								goby ++
								nextFrom = i2
								if goby == mid1 {
									result1 = v2
									println("nums2 find the result1", result1, "at n2", i2)
								}
							} else {
								goby ++
								if goby == mid1 {
									result1 = v2
									println("nums2 find the result1", result1, "at n2", i2)
								}
								break
							}
						}
					}
			}
	}

    println(result1, result2)
	
	if needTwo {
		return float64(result1 + result2) / 2
	} 
	return  float64(result1)
}