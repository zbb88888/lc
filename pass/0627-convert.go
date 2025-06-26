package main


import "fmt"

func convert(s string, numRows int) string {
  lens := len(s)
  if numRows < 2  {
    return s
  }

// 这道理需要将 s 转换为一个符合 |/| 字的的二维数组
// 但实际上不需要用二维数组来缓存，只需要计算出每一个字符的二维数组的下标
// 需要一个一位数组来缓存每一行的字符串，最后再将该数组逐行拼接

  result := make([]string, numRows)

  // 每个字符的i，j 存在一个数学上的对应关系
  // lenWin = 2 * numRow - 2 是一个完整的模式段: 处理窗口
  // 一个窗口即是要处理的最小问题
  // 每一个窗口中的 i 都从 0 开始，
  // i <= numRow: i 即是 行号，numRow 的部分是一列
  // i > numRow: lenWin - i 即是行号

  lenWin := 2*numRows-2
  win := ""
  for i:=0; i < lens; i+=lenWin{
      // fmt.Println("handle from i", i)
      if (i+lenWin) < lens {
        win = s[i:i+lenWin]
      } else {
        win = s[i:]
      }
      for j:=range win {
        // hang set
        v := string(win[j])
        if j < numRows {
          hang:=j
          result[hang] += v
          // fmt.Println("hang: ", hang, "add", v)
        } else {
          hang:=lenWin-j
          result[hang] += v
          // fmt.Println("hang: ", hang, "add", v)
        }
      }
  }

  res := ""
  for hang, v := range result {
    fmt.Println("hang: ", hang, v)
    res += v
  }
  fmt.Println("res: ", res)
  return res
}


func main(){

s := "PAYPALISHIRING"
numRows := 4
// fmt.Println(s)
// fmt.Println(numRows)

convert(s, numRows)

}
