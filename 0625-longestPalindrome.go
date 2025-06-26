package main
import (
  "fmt"
)


func longestPalindrome(s string) string {
  max := 0
  result := ""
  // 初始化滑动窗口
  lens := len(s)
  for window:=1; window <= lens; window ++ {
    for i:=0; i <= lens-window; i ++ {
      // 缓存窗口大小，判断
      long := s[i:i+window]
      yes := check(long)
      fmt.Println("check: matched ?", yes, long)
      if yes && len(long) > max {
          result = long 
          println("find longest: ", result)
      }
    }
  }
  return result
}

func check(s string)  bool {
  // 判断当前是否依旧符合回文
  maxIndex := len(s)-1
  mid := maxIndex/2
  for i:=0; i <= mid; i++ {
    if s[i] == s[maxIndex-i] {
      continue
    }else { 
      return false
    }
  }
   return true
}

func main(){

// 这是一个滑动窗口题
// 窗口大小 从 1 到 len(s)
// 每次取窗口大小的list，然后判断是否回文，然后步进一，继续
// 穷举的核心就是： 缓存 判断 缓存 继续
s := "aacabdkacaa"
println("input: ", s)

result := longestPalindrome(s)

println(result)

}

