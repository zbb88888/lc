package main

import "fmt"

// p 是正则表达式
// p 是外层循环
// 参考 p 进行最小窗口问题划分：
// 1. p 中的单个字符 x 必须要匹配到
// 如果遇到 x 匹配不到，那么继续看是否是 x*， 这样可以跳过
// 2. 如果只是一个 . 必须匹配一个字符
// 如果匹配不到，那么继续看是否 .*，这样可以跳过
// 3. 贪心匹配： 根据当前p和s的剩余差值判断至少需要跳过多少个字符
// 如果 .* 后存在一个字符，还是要优先匹配该字符，如果不匹配（根据剩余差值判断至少需要跳过多少个字符）

func main() {
	s := "aab"
	p := "c*a*b"
	fmt.Println("input: ", s, p)
	yes := isMatch(s, p)
	fmt.Println("yes: ", yes)
}

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	lenS := len(s)
	lenP := len(p)
	fmt.Println("lenS:", lenS, "lenP:", lenP)
	if lenS == 0 || lenP == 0 {
		return false
	}


	// 优先处理通配逻辑
	// 目前仅面向一次性通配的场景
	// found x* 视作 0
	// foud .* 视作无穷
	// 计算通配能匹配到的最大｜最小字符数目
	// 1. 首先只考虑都是通配符的场景
	// 2. 如果 x 只有一个字符，则考虑判断 x* 的情况

	onlyX := "" // 记录唯一的一个非通配字符
	xOnly := false

	// 通配数目范围
	maxAs := 0
	minAs := 0
	for i, v := range p {
		sv := string(v)
		if xOnly && sv == onlyX {
			maxAs = ++
			minAs = ++
			continue
		}
		if "a" <= sv && sv <= "z" {
			if i+1 < plen {
				nextSV := string(p[i+1])
				if nextSV == "*" {
					maxAs ++
					continue
				}
			} else {
				fmt.Println("simple x end1")
			}
			if onlyX == "" {
				onlyX = sv
				xOnly = true
				maxAs ++
				minAs ++
				continue
			}
			if onlyX != sv {
				xOnly := false
				onlyX += sv
				continue
			}
			fmt.Println("simple unknown case1")
		}
		if sv == "." {
			if i+1 < plen {
				nextSV := string(p[i+1])
				if nextSV == "*" {
					maxAs += 20
					minAs += 0
					continue
				}
			} else {
				fmt.Println("simple . end1")
			}
			maxAs += 1
			minAs += 1
			continue
		}
		fmt.Println("simple unknown case2")
	}

	if xOnly && lenS == 1 && onlyX == s && minAs <= lenS{
	   fmt.Println("simple one true")
	   return true
	}
	if onlyX == s && minAs <= lenS {
	   fmt.Println("simple match true")
	   return true
	}

	// 2. 动态贪心：可以来回调整是否贪心
	// 最小窗口：p 一次只能处理一个通配符或者一个字符
	preferS := false
	// s 更长，优先用通配正则消费更多的 s
	preferLen := plen - slen
	// 最多只能通配这么多

	var pi, si, preferLen
	var pv, sv, must string
	var nextStart bool
	// 参考 p 进行最小窗口问题划分：
	// 1. p 中的单个字符 x 必须要匹配到
	// 1.1 如果 s 存在 x 能和 p 中的 x 匹配上，那就是匹配逻辑
	// 1.2 否则就是匹配不上的逻辑，x 后必须存在一个 x*，然后跳过
	// 2. 如果只是一个 . 必须匹配一个字符
	// 如果匹配不到，那么继续看是否 .*，这样可以跳过
	// 3. 贪心匹配： 根据当前 p 和 s 的剩余差值判断至少需要跳过多少个字符
	// 如果 .* 后存在一个字符，还是要优先匹配该字符，如果不匹配（根据剩余差值判断至少需要跳过多少个字符）
	for pi, pv = range p {
		pLeft := p[pi:lenP]
		sLeft := s[si:lenS]
		spv := string(pv)
		ssv := string([s[si]])
		if slen >= plen {
			preferS = true
			preferLen = slen - plen
		} else {
			preferS = false
			preferLen = plen - slen
		}
		fmt.Println("preferS:", preferS, "preferLen:", preferLen)
		if "a" <= spv && spv <= "z" {
			must = spv
			// x or x*
			// must get x or skip x*
			if i+1 < plen {
				if string(p[i+1]) == "*" {
					nextStar = true
				}
			}
			if must == ssv {
				// match x, so pass it
				pi ++
				si ++
				if !nextStar {
					// only match x
					continue
				}
				if nextStar && preferS && preferLen > 0 {
					// choose to match how many x
					for si < lens {
						ssv := string([s[si]])
						if must == ssv {
							fmt.Println("keep match x*:", "si", si, "ssv", ssv, preferLen)
							si ++
							preferLen --
							if preferLen > 0 {
								continue
							} else {
								break
							}
						}
						break
					}
				}
				if nextStar && !preferS && preferLen > 0 {
					// skip x*
					fmt.Println("skip x*", "pi", pi, "spv", spv, preferLen)
					pi ++
					continue
				}
			}
			if must != ssv {
				pi ++
				if !nextStar {
					// can not match x
					fmt.Println("end not match: ", "pi", pi, "spv", spv, "si", si, "ssv", ssv)
					return false
				}
				if nextStar {
					pi++
					fmt.Println("skip x*:", "pi", pi, "spv", spv, "nextStar", nextStar)
					continue
				}
			}
		}
		if sv == "." {
			// . or .*
			if i+1 < plen {
				nextSV := string(p[i+1])
				if nextSV == "*" {
					maxAs += 20
					minAs += 0
					continue
				}
			}
			maxAs += 1
			minAs += 1
			continue
		}
		fmt.Println("dynamic unknown case2")
	}
	return false

}
