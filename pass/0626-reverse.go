package main

import (
"fmt"
"math"
)


func reverse(x int) int {
   
  // 1. 标记符号
  // 2. 整除10之后再*10累加
  if -10 < x && x < 10 {
    return x
  }

  res := 0
  tmp := 0
  for {
  
  // 余数
  tmp = x % 10
  // println("tmp: ", tmp)
  // 整数
  x = x / 10
  // println("x: ", x)

  // 余数变整数
  res = res*10 + tmp*10
  // println("tmp res: ", res)
  
  if -10 < x && x < 10 {
    // 整数变余数
    res += x
    if res > math.MaxInt32 || res < math.MinInt32 {
      println("end 0: ")
      return 0
    } 
    println("end res: ", res)
    return res
  }
  }

}


func main() {

input := 1534236469

fmt.Println("input: ", input)

reverse(input)


}

