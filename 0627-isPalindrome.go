package main

import (
  "fmt"
)


func isPalindrome(x int) bool {
// 不将整数转为字符串来解决
// 使用进制的方式把这个数反转过来相等则回文


  if x < 0 {
    return false
  }

  if x < 10 {
     return true
  }


  if x == 10 {
    return false
  }



    
}



func main(){

input: = 123

isPalindrome(input)

}
