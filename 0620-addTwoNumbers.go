package main

import "math"

// define singly-linked list
type ListNode struct {
  V int
  Next *ListNode
}

// format arr to listNode
func format(arr []int) (*ListNode, int) {
// maxPos := len(arr) -1  
var head, prv *ListNode
result := 0
for i := range arr {
    v := arr[i]
    println("v: ", v)
    result += v * int(math.Pow(10, float64(i)))
    println("result: ", result)
    tmp := ListNode{V:v}
    if i == 0 {
      head = &tmp
      prv = &tmp
      continue
    }
    prv.Next = &tmp
  }
  println("format end, sum: ", result)
  return head, result
}

func main(){
 
  l1 := []int{2, 4, 3}
  l2 := []int{5, 6, 4}
  
  // fmt.Println("l1: %v", l1)
  // fmt.Println("l2: %v", l2)

  println("format l1")
  _, sum1 := format(l1)
  println("format l2")
  _, sum2 := format(l2)

  sum := sum1 + sum2
  println("end sum: ", sum)

}
