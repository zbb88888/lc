package main
import "fmt"

func main() {
	
	arr1 := []int {1, 2, 3, 4}
	result := append1(arr1, 5, 6, 7)
	fmt.Println(result)
}

func append1(arr []int, i ...int ) []int {
	// 不使用 append，实现 append(arr, i)
	
	lenArr := len(arr)
	lenI := len(i)
	maxLen := lenArr + lenI
	
	println(lenArr, lenI, maxLen)
	
	result := make([]int, maxLen)
	for i1, v := range arr {
		result[i1] = v
	}
	
	// result[lenArr:] = i[:]
	for i2, v := range i {
		result[lenArr+i2] = v
	}
	
	return result
}