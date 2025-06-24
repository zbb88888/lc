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
func format(arr []int) (*ListNode) {
       if arr == nil {
         return nil
       }
	var head, priv *ListNode
	for i := range arr {
                println("format: v", arr[i])
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

func sumListNode(l *ListNode) int {
        if l == nil {
          return 0
        }
	n := l
        v := n.Val
	result := 0
	decimal := 1
	for {
                println("sum v: ", v)
		result += decimal * v
                // next loop + next v
		n = n.Next
                if n == nil {
                break
                }
                v = n.Val
		decimal *= 10
	}
        println("end sum result: ", result)
	return result
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum1 := sumListNode(l1)
	sum2 := sumListNode(l2)
	sum := sum1 + sum2
	println("end sum: ", sum)
        // 0

	result := []int{}
	tmp := 0
        if sum < 10 {
	  result = append(result, sum)
	  ln := format(result)
	  return ln
        }

	res := sum
	for {
		tmp = res % 10
		res = res / 10
		println("res tmp", res, tmp)
		result = append(result, tmp)
		if res < 10 {
		  result = append(result, res)
		  break
		}
	}
	fmt.Println("result: ", result)
	ln := format(result)
	return ln
}

func main() {
	arr1 := []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
	arr2 := []int{5,6,4}
	l1 := format(arr1)
	l2 := format(arr2)
	addTwoNumbers(l1, l2)
}
