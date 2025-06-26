package main


import "fmt"

func convert(s string, numRows int) string {

// 这道理需要将 s 转换为一个符合 |/| 字的的二维数组
// 但实际上不需要用二维数组来缓存，只需要计算出每一个字符的二维数组的下标
// 需要一个一位数组来缓存每一行的字符串，最后再将该数组逐行拼接


  var result [numRow]string

  // 每个字符的i，j 存在一个数学上的对应关系
  // lenWin = 2 * numRow -1 是一个完整的模式段: 处理窗口
  // 每一个窗口中的 i 都从 0 开始，
  // i <= numRow: i 即是 行号，numRow 的部分是一列
  // i > numRow: lenWin - i 即是行号
  for i := range(s){

  }



}


func main(){

fmt.Println(s)
fmt.Println(numRows)
return convert(s, numRows)

}
