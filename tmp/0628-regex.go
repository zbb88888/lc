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
	var xStar bool // using in x* case
	var pointStar bool // using in .* case
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
			// 1. p 中的单个字符 x 必须要匹配到
			// 1.1 如果 s 存在 x 能和 p 中的 x 匹配上，那就是匹配逻辑
			// 1.2 否则就是匹配不上的逻辑，x 后必须存在一个 x*，然后跳过
			if i+1 < plen {
				if string(p[i+1]) == "*" {
					xStar = true
				}
			}
			// x or x*
			// must get x or skip x*
			must = spv
			if must == ssv {
				if !xStar {
					// 只能匹配一个
					pi ++
					si ++
					continue
				}

				if xStar && preferS && preferLen > 0 {
					// hold x pos of x*
					// 目前一直指向 x 的位置
					for si < lens {
						// 继续贪心匹配更多同样的 x
						ssv := string([s[si]])
						if must == ssv {
							fmt.Println("x* keep match more x:", "si", si, "ssv", ssv, preferLen)
							si ++
							preferLen --
							if preferLen > 0 {
								// 还能继续贪心
								continue
							} else {
								// 直接跳到 x* 后 * 所在位置
								pi += 2
								continue
							}
						} else {
							// 把 * 当作匹配 0 个处理
							pi += 2
							continue
						}
					}
				}

				// s 更长，放弃贪心
				if xStar && !preferS && preferLen > 0 {
					// skip x*
					fmt.Println("skip x*", "pi", pi, "spv", spv, preferLen)
					pi ++
					continue
				}
			}


			if must != ssv {
				// 当前字符不匹配
				pi ++
				if !xStar {
					// 由于不是 x* 所以无法跳过
					// can not match x
					fmt.Println("end not match: ", "pi", pi, "spv", spv, "si", si, "ssv", ssv)
					return false
				}
				if xStar {
					// 把 x* 整个部分当作匹配 0 个处理
					pi++
					fmt.Println("skip x*:", "pi", pi, "spv", spv, "nextStar", xStar)
					continue
				}
			}
		}
		// 下次从这里开始 review 注释
		if sv == "." {
			// 2. 如果只是一个 . 必须匹配一个字符
			// 如果匹配不到，那么继续看是否 .*，这样可以跳过
			if pi+1 < plen {
				nextSV := string(p[pi+1])
				if nextSV == "*" {
					// hold . pos of .*
					pointStar = true
				} else {
					// just .
					pi ++
					if si+1 < slen {
						si ++
					} else {
						fmt.PrintLn("false: point not match")
						return false
					}
				}
			}
			if pointStar { 
				if preferS && preferLen > 0 {
					// .* choose to match how many x
					for si < lens {
						ssv := string([s[si]])
						if must == ssv {
							fmt.Println("keep match .*:", "si", si, "ssv", ssv, preferLen)
							si ++
							preferLen --
							if preferLen > 0 {
								continue
							} else {
								// hold .* until
								pi += 2
								break
							}
						}
						break
					}
				}
			}
			if pointStar && !preferS && preferLen > 0 {
				// skip x*
				fmt.Println("skip x*", "pi", pi, "spv", spv, preferLen)
				pi ++
				continue
			}
			continue
		}
	}
	return false

}
