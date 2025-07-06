package main

import "fmt"

func main() {
	input := []string{"flower", "flow", "flight"}
	fmt.Println("input:", input)
	ret := longestCommonPrefix(input)
	fmt.Println("ret:", ret)
}

func matchLen(a, b string) (string, int) {
	match := 0
	s := ""
	fmt.Println("checking:", a, b)
	small := b
	if len(a) < len(b) {
		small = a
	}
	for i := range small {
		v := a[i]
		vv := b[i]
		if v == vv {
			match++
			s += string(vv)
			continue
		}
		return s, match
	}
	return s, match
}
func longestCommonPrefix(strs []string) string {
	// 双指针
	// 每一个值都和其他所有值比较前缀
	// 性能优化
	// ii >= i+1
	// vv 小于 maxCom 直接跳过
	l := len(strs)
	if l == 0 {
		return ""
	}
	if l < 2 {
		return strs[0]
	}
	var tmpS, v, vv string
	tmpMatch := 0
	v = strs[0]
	vv = strs[1]
	tmpS, tmpMatch = matchLen(v, vv)
	for i := 1; i < l; i++ {
		v = strs[i]
		tmpS, tmpMatch = matchLen(tmpS, v)
		if tmpMatch != 0 {
			fmt.Println("match: ", tmpS)
		} else {
			return ""
		}
	}
	return tmpS
}
