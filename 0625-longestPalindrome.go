package main

func longestPalindrome(s string) string {

// 这个是一个栈的操作思路
// 需要一个切片，存储最长的字符串
// 字符串不在切片中，存入
// 字符串在切片中，判断是否回文
// 清空：
// 1. 从栈底部开始清空不再符合回文的部分
// 2. 回文结束，清空整个回文的部分，判断是否是最长的回文片段并记录
// 继续迭代


long := []string // 缓存最长回文

for v := range s {
  if !contain(long, v) {
    long = append(long, v)
    continue
  }

  
  long = append(long, v)
  // 判断当前是否依旧符合回文

}
    
}

func check(list []string) (list []string, bool) {
  maxIndex := len(list)-1
  mid := maxIndex/2
  goby := 0

  for i:=0; i <= mid; i++ {
    if list[i] == list[maxIndex-1] {
      // 符合回文，继续
      // 如果到最后 mid, 则就是一个完整的回文字符串
      println("match at : ", i, maxIndex-1, "both are", list[i])
      goby = i
      continue
    }
    println("not match at: ", i, maxIndex-1, list[i], "!=", list[maxIndex - 1])
    // 不再符合，则清理
    break
  }

  if goby == mid {
    // 始终符合回文
     return list, true
  }

  // 不再符合，则从底部开始清理
  println("before cut from: ", goby)
  newList := list[goby:]
  println("after cut from: ", goby)
}

func contain(list []string, s string) bool {
  for v := range list{
    if v == s {
      return true
    }
  }
  return false
}


func main(){

s := "babad"

result := longestPalindrome(s)

println(result)

}

