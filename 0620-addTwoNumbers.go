package main

// define singly-linked list
type ListNode struct {
  V int
  Next *ListNode
}


func addTwoNumbers(l1, l2 *ListNode) *ListNode {

println("printListNode l1")
printListNode(l1)
println("printListNode l2")
printListNode(l2)

return nil

}

func printListNode(ln *ListNode) {
if ln == nil {
    println("ln is nil")
    return
}
println("ln.V: ", ln.V)
if ln.Next == nil {
    println("ln.Next is nil")
    return
}

for {
      n := ln.Next
      if n != nil {
        println(n.V)
        ln = ln.Next
      } else {
        break
      }
  }
}

// format arr to listNode
func format(arr []int) *ListNode {
maxPos := len(arr) -1  
var head, prv *ListNode
for i := range arr {
    // init node
    v := arr[maxPos-i]
    println(v)
    tmp := ListNode{V:v}
    if i == 0 {
      head = &tmp
      prv = &tmp
      continue
    }
    prv.Next = &tmp
  }
return head
}

func main(){
 
  l1 := []int{2, 4, 3}
  l2 := []int{5, 6, 4}

  ln1 := format(l1)
  ln2 := format(l2)

  ln := addTwoNumbers(ln1, ln2)
  println("printListNode result")
  printListNode(ln)
}
