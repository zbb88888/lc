package main

import (
	"fmt"
)

// define singly-linked list
type ListNode struct {
	V    int
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
		now := &ListNode{V: arr[i]}
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
        v := n.V
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
                v = n.V
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
	result := []int{}
	res := sum
	tmp := 0

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
	arr1 := []int{2, 4, 3}
	arr2 := []int{5, 6, 4}
	l1 := format(arr1)
	l2 := format(arr2)
	addTwoNumbers(l1, l2)
}
