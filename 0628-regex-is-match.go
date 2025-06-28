package main
import "fmt"


	// . 可以匹配任意一个字符
	// * 可以匹配0或多个前面的那一个元素
	// 双指针： 
	// 1. 缓存s 的字符
	// 2. 缓存p 中 . 或者* 的位置

	// 最小窗口问题划分： s 的任何位置肯定都是字符
	// 1. p 所在的位置也是字符，需要一一相等
	// 2. p 所在位置是 . 或者 *， 需要 hold 住该位置，然后对比 s
	//  哪个最晚结束哪个在外层循环
	// star 有一个难处理的地方，* 号一旦用了自身也是要进一的，但又要有 hold 在这个状态一直匹配的逻辑
	// 所以星号也要有一个缓存: 能够满足数组先进位，同时又能保持继续和s匹配
	// .* 如果是连接的，需要当一个符号判断，直接进两位

func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)

	fmt.Println("slen", slen, "plen", plen)

	if slen == 0 || plen == 0 {
          return false
	}

	// 优先处理通配逻辑
        // found * 
	// foud . just after *
	// 计算通配能匹配到的最大字符数目
	// 这里只考虑都是通配符的场景
	maxAs := 0
	minAs := 0
	for i, v := range p {
	    sv := string(v)
            if "a" <= sv && sv <= "z" {
	       if i+1 < plen {
                   nextSV := string(p[i+1])
		   if nextSV == "*" {
		      // ignore x*
		      // as 0
	              continue
		   }
	           maxAs = 0
	           minAs = 0
                   break
	       }
	       maxAs = 0
	       minAs = 0
               break

            }
	    if sv == "." {
	       if i+1 < plen {
                   nextSV := string(p[i+1])
		   if nextSV == "*" {
                       maxAs += 20
                       minAs += 0
		       continue
		   }
	       }
               minAs += 1
	       continue
	    }

	}

	// 如果都是通配。直接比较字符长度

	if maxAs > 0 || minAs > 0  {
	  fmt.Println("just simple maxAs", maxAs, "minAs", minAs)
	  if minAs <= slen && slen <= maxAs {
	    fmt.Println("just simple match")
	    return true
          } else {
	    fmt.Println("just simple not")
	    return false
	  } 
	}

	spos := 0
	ppos := 0
	beforeSSV := ""
	findStartAfterX := false // 开始的时候 xx*, x* 为空， 可跳过
	findStarAfterX := false // 然后只能 x* 为空， 可跳过
	holdStar := false    // 如果匹配则一直匹配，如果不匹配则跳过 star（匹配到0个），* 为空，可跳过
        mustGet := ""	     // mustGet 是正则模板中 .* 之后必须要匹配到的字符，如果找不到，就是不匹配
	// 这两个状态是互相独立的 没有优先或者依赖关系
	for spos < slen {
             if spos > 0 {
               beforeSSV = string(s[spos-1])
	     }
	     fmt.Println("loop1 spos ppos hold findStart findStar: ", spos, ppos, holdStar, findStartAfterX, findStarAfterX)
             if spos == slen {
	        break
             }
             if ppos == plen {
	        break
             }
	     ssv := string(s[spos])
             spv := string(p[ppos])
             if mustGet == "." {
	          // .* .*. .*..  以此类， 只需要继续迭代，拿到下一个字符 
                  // 保持现状
		  ppos ++
		  // 至少对应一个字符
		  spos ++
		  continue
	     }

	     if mustGet != "" {
		if ssv == spv {
                    fmt.Println("re start from must get: ", mustGet, "spos", spos)
		    mustGet = ""
		    spos ++
		    continue
		} 
		fmt.Println("must get skip: ", "spos", spos, "ssv", ssv)
		spos ++
		continue
	     }

	     if findStartAfterX {
	       if spv != "*" {
                  ppos ++
		  continue
	       } else {
		 // start with xx* is empty, pass
		 findStartAfterX = false
                 ppos ++
                 continue
	       }
	     }

	     if findStarAfterX {
	       if spv != "*" {
		  fmt.Println("end loop1: false")
		  return false
	       } else {
		 // next with x* is empty, pass
		 findStarAfterX = false
                 ppos ++
                 continue
	       }
	     }
	     // 如果匹配成功则停留在 star 这里一直匹配
             // 如果匹配不成功，则跳过 star
	     // 如果下一个是 *，实际上是两个 * 连在一起了，是异常
	     if beforeSSV != "" && holdStar{
                if beforeSSV == ssv {
		   // 一直匹配
	           fmt.Println("loop1 *: ", spos, ppos)
                   spos ++
		   continue
		} else {
		   // 允许失败: * 匹配到 0  个
                   holdStar = false
		   ppos ++
		}
	     }
             for ppos < plen {
	       fmt.Println("loop2 spos ppos: ", spos, ppos)
	       fmt.Println("holdStar: ", holdStar)
               if spos == slen {
	          break
                }
               if ppos == plen {
	          break
                }
	        ssv := string(s[spos])
	        spv := string(p[ppos])
		fmt.Println("loop2 check: ", "ssv", ssv, "spv", spv, "beforeSSV", beforeSSV)
		// . 直接都前进1
                if spv == "." {
                      ppos ++
		      // 判断下一个是否是 *， 如果是当作一个字符进行逻辑判断
		      if ppos < plen && string(p[ppos]) == "*" {
			 // 经过一个.*，至少再跳过2个字符，贪心
                         spos += 2
			 // .* 可以看作 零个或者 多个
			 // 如果是 .* 结尾，直接结束
                         ppos ++
		         if ppos == plen {
                              fmt.Println(".* just end")
			      return true
		         }
			 // 如果 .* 之后还有字符，依旧需要优先匹配字符
			 // 缓存该字符，直接跳到该字符后，继续
		         if ppos < plen {
			      mustGet = string(p[ppos])
			      fmt.Println("after .* mustGet: ", mustGet)
			      break
			 }
		      } 
		      // 只是 . 必须匹配一个字符
		      spos ++
		      break
		 }
	         // 匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
		 if spv == "*" {
		    // loop l1 一直匹配
		    holdStar = true
		    break
		    // 如果匹配成功则停留在 star 这里一直匹配
		    // 如果匹配不成功，则跳过 star
		    // 如果下一个是 *，实际上是两个 * 连在一起了，是异常
		 }

		 if spv == ssv {
                      ppos ++
		      spos ++ 
		      break
		  } else {
		      // 不相等的时候判断是否是如下组合，则可以跳过
		      // x* xx* 为空
	              // x* means empty, skip empty
		      // l1 need find the first *, 然后开始
		      if spos == 0 {
		        // star after x* xx* 为空
		        findStartAfterX = true
		        ppos ++
		        break
		      } else {
		        // next x* 为空
		        findStarAfterX = true
		        ppos ++
		        break
		      }
		  }

             }
             if spos == slen && ppos == plen{
                    fmt.Println("loop1 check: ", "spos", spos, "ppos", ppos)
	            return true
             }
	}


send := false
pend := false
if spos >= slen {
   send = true
}
if ppos >= plen {
   pend = true
}

fmt.Println("return0", "send", send, "pend", pend, "holdStar", holdStar)
fmt.Println("return0", "spos", spos, "ppos", ppos)

if send && pend{
   fmt.Println("end true")
   return true
}


if !send {
  fmt.Println("end false 1")
  return false
}

pleft := plen - ppos
if pleft == 1 {
  if string(p[ppos]) == "*" {
     return true
  } 
  if string(p[ppos]) == "." {
     fmt.Println("pleft 1 end false 1")
     return false
  }
  fmt.Println("pleft 1 end false 2")
  return false
}

if pleft == 2 && beforeSSV == string(p[ppos+1]) {
  return true
}


// 这里有两个贪心：
// 1. 正则中的贪心： 正则模板尽可能多的匹配字符串，这样可以尽早
// 2. 最后的一个或者连续字符的贪心，尽量匹配晚所有的正则模板

lasts :=  string(s[slen-1])

if pleft >= 2 {

  left := p[ppos+1:]
  fmt.Println("left: ", left)
  if slen == 1 {
    if "a" <= string(left[0]) && string(left[0]) <= "z" {
       fmt.Println("one return false 1")
       return false
    }

    if string(left[0]) == "." && string(left[0]) != "*" {
       fmt.Println("one return false 2")
       return false
    }
  }


  // 由于正则表达式不能以* 号开头，所以都以 x* 格式开头 相当于 空
  // 由于 第一个 x* 视作为空
  // 后续的 x* * 可视作为空
  // 去空后看是否 < = 2

  if len(left) == 1 {
    if string(left[0]) == "."{
         return true
    }
    if string(left[0]) == "*"{
     return true
    }
  }
  cutFrom := 0
  findStar := false
  for i, v := range left {
    if string(v) == "*" {
      cutFrom = i
      findStar = true
    }
  }
  if findStar {
     thePLeft := left[cutFrom+1:]
     endLen := len(thePLeft)
     fmt.Println("thePLeft: ", thePLeft)
     if endLen == 0 {
       // all left is empty
       return true
     }
     if send {
        // just one but matched
	if slen == 1 && spos == 1 {
	   // 只有1个，且已经匹配过
           if endLen > 0 { 
             fmt.Println("end one false 3")
             return false

	   }

	}
        fmt.Println("end false 3")
        if endLen == 1 { 
	    if string(thePLeft[0]) == "."{
               return true
            }
	    if string(thePLeft[0]) == "*"{
               return true
            }
	    if string(thePLeft[0]) == lasts{
               return true
            }
        }
        if endLen == 2 && string(thePLeft[0]) == lasts{
           return true
        }
	return false
     }
  } else {
    fmt.Println("end false 4")
    return false
  }
}



return false


}

// main

func main(){

s := "ab"
p := ".*c"

// 由于正则表达式不能以* 号开头，所以都以 x* 格式开头 相当于 空
// 关于是否贪心的问题，如果P 已经用尽，那么只能选择贪心即用 * 号多匹配一次
// for 循环中是贪心的判断，那就会导致P 还有剩余的情况下导致错误判断：
// 如果P 没有用尽，那必须在去空 x* 之后，还有一个x 或者 x* 或者 . 可以用于匹配
// 由于 第一个 x* 视作为空
// 后续的 x* 视作为空
// *x* 也视作为空
// 所以如果不匹配的时候，需要考虑找到 x* xx* 几个x 都可以，第一个字母必须匹配


// 字符串匹配和通配是两种逻辑

// 所以单独先处理通配的逻辑


fmt.Println("input: ", s, p)

yes := isMatch(s, p)

fmt.Println("yes: ", yes)


}
