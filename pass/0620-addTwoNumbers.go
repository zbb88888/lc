package main

import (
	"fmt"
)

// define singly-linked list
type ListNode struct {
	Val    int
	Next *ListNode
}

// format arr to listNode
func format(arr []int) *ListNode {
       if arr == nil {
         return nil
       }
	var head, priv *ListNode
	for i := range arr {
                // println("format: v", arr[i])
		now := &ListNode{Val: arr[i]}
                // init 0
                if head == nil {
                   head = now
                   continue
                }
                
                // init 1
                if priv == nil {
                  head.Next = now
                  priv = head.Next
                  continue
                }
           
                // init n
		priv.Next = now
                priv = priv.Next
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := []int{}
        var l1End, l2End, addOne bool
        l1Next := l1
        l2Next := l2
	for {
                tmp := -1 // not count
                if l1Next == nil {
                    l1End = true
                }else {
                    if tmp == -1 {
                       tmp = 0
                    }
                    tmp += l1Next.Val
		    println("tmp, + l1Next v", tmp, l1Next.Val)
                    l1Next = l1Next.Next
                }
                if l2Next == nil {
                    l2End = true
                }else {
                    if tmp == -1 {
                       tmp = 0
                    }
                    tmp += l2Next.Val
		    println("tmp, + l2Next v", tmp, l2Next.Val)
                    l2Next = l2Next.Next
                }
                if addOne {
                    if tmp == -1 {
                       tmp = 0
                    }
                    tmp += 1
                }
           
                if tmp == -1 {
                   break
                }
		if tmp < 10 {
		  result = append(result, tmp)
                  addOne = false
		} else {
		  result = append(result, tmp-10)
                  addOne = true
                }
                if l1End && l2End {
		  println("end tmp: ", tmp)
                  if addOne {
		      result = append(result, 1)
                  }
                  break
                }
	}
	fmt.Println("result: ", result)
	ln := format(result)
	return ln
}

func main() {
	// arr1 := []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
	// arr2 := []int{5,6,4}
        arr1 := []int{2,4,3}
        arr2 := []int{5,6,4}
        // 其实不需要求和，只需要考虑是否进位即可 
	ln1 := format(arr1)
	ln2 := format(arr2)
	addTwoNumbers(ln1, ln2)
}
