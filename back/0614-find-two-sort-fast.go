package main

func main() {
	nums1 := []int{1,3}
	nums2 := []int{2}
	findMedianSortedArrays(nums1, nums2)

}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    len1 := len(nums1)
    len2 := len(nums2)

    total := len1 + len2
    
    passByTotal := 0
    holder1 := make(chan int)
    holder2 := make(chan int)
    
    iGoFirst := false
    if nums1[0] <= nums2[0] {
      iGoFirst = true
    }

    result1 := 0
    result2 := 0
    go goBy(nums1, holder1, holder2, "nums1", total, iGoFirst, &passByTotal, &result1, &result2)
    goBy(nums2, holder2, holder1, "nums2", total, !iGoFirst, &passByTotal, &result1, &result2)
    result := float64(result1 + result2) / 2
    println(result)
    return result

}

func goBy(nums []int, my chan <- int, she <-chan int, role string, total int, iGoFirst bool, passByTotal, result1, result2 *int){
  // mid pos
  mid := total/2
  mid1 := mid
  mid2 := 0
  if total % 2 == 0 {
      mid2 = mid -1
      *passByTotal ++
  } 

  if iGoFirst {
      // init go, let she know
      my <- nums[0]
  }

  for _,v := range nums {
	  sheIsHold := <- she
	  if *passByTotal == mid1 {
              *result1 = v
	  }
	  if *passByTotal == mid2 {
              *result2 = v
	  }
	  if v <= sheIsHold {
             // i go
             continue
	  } else {
             // she go
	     my <- v
	  }
  }
}


