package main

import (
  "fmt"
  "math"
  "strconv"
)


func myAtoi(s string) int {

// 去除空格
// 判断 + -
// 跳过 0
// 读取
// 直到结尾（非数字）

numbers := ""
for i, v := range s {
  
  sv := string(v)
  if sv == " " {
    if numbers != "" {
      fmt.Println("end at: ", i, sv)
      break
    }
    continue
  }
  

  if sv == "+" || sv == "-" {
    if numbers != "" {
      fmt.Println("end at: ", i, sv)
      break
    }
    numbers += sv
  }

  // get 0-9
  if "0" <= sv && sv <= "9" {
    numbers += sv
    continue
  }

  // end
  if sv == "." {
    fmt.Println("end at: ", i, sv)
    break
  }

  if "a" <= sv && sv <= "z" {
    fmt.Println("end at: ", i, sv)
    break
  }

}

fmt.Println("numbers: ", numbers)
result, err := strconv.Atoi(numbers)
if err != nil {
  fmt.Println("err: ", err)
}

fmt.Println("result: ", result)
if result < math.MinInt32 {
  return math.MinInt32
}

if result > math.MaxInt32 {
  return math.MaxInt32
}

return result

}



func main(){


// 空格：读入字符串并丢弃无用的前导空格（" "）
// 符号：检查下一个字符（假设还未到字符末尾）为 '-' 还是 '+'。如果两者都不存在，则假定结果为正。
// 转换：通过跳过前置零来读取该整数，直到遇到非数字字符或到达字符串的结尾。如果没有读取数字，则结果为0。
// 舍入：如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数

// 使其保持在这个范围内。具体来说，小于 −2^31 的整数应该被舍入为 −2^31 ，大于 2^31 − 1 的整数应该被舍入为 2^31 − 1 。
// 返回整数作为最终结果。

s := "1337c0d3"

myAtoi(s)

}
