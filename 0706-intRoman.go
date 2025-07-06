package main

import "fmt"

func main() {
	input := 3749
	fmt.Println("input: ", input)
	ret := intToRoman(input)
	fmt.Println("ret: ", ret)
}

// I	1
// V	5
// X	10
// L	50
// C	100
// D	500
// M	1000

func intToRoman(num int) string {
	const (
		M1000 = 1000
		D500  = 500
		C100  = 100
		L50   = 50
		X10   = 10
		V5    = 5
		I1    = 1
	)
	const (
		// 减法形式
		// 4 (IV)，9 (IX)，40 (XL)，90 (XC)，400 (CD) 和 900 (CM)
		CM900 = 900
		CD400 = 400
		XC90  = 90
		XL40  = 40
		IX9   = 9
		IV4   = 4
	)

	// 感觉只是简单的进制转换
	// 从大到小依次整除
	ret := ""
	yu := num
	zheng := 0
	for {
		zheng = 0
		if yu >= M1000 {
			yu = yu % M1000
			zheng = num / M1000
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "M"
			}
			fmt.Println("M: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		}
		if yu >= CM900 {
			yu = yu % CM900
			zheng = num / CM900
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "CM"
			}
			fmt.Println("CM: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		}
		if yu >= D500 {
			yu = yu % D500
			zheng = num / D500
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "D"
			}
			fmt.Println("D: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		}
		if yu >= CD400 {
			yu = yu % CD400
			zheng = num / CD400
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "CD"
			}
			fmt.Println("CD: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		}
		if yu >= C100 {
			yu = yu % C100
			zheng = num / C100
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "C"
			}
			fmt.Println("C: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		}
		if yu >= XC90 {
			yu = yu % XC90
			zheng = num / XC90
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "XC"
			}
			fmt.Println("XC: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		}
		if yu >= L50 {
			yu = yu % L50
			zheng = num / L50
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "L"
			}
			fmt.Println("L: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		}
		if yu >= XL40 {
			yu = yu % XL40
			zheng = num / XL40
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "XL"
			}
			fmt.Println("XL: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		}
		if yu >= X10 {
			yu = yu % X10
			zheng = num / X10
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "X"
			}
			fmt.Println("X: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		}
		if yu >= IX9 {
			yu = yu % IX9
			zheng = num / IX9
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "IX"
			}
			fmt.Println("IX: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		}
		if yu >= V5 {
			yu = yu % V5
			zheng = num / V5
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "V"
			}
			fmt.Println("V: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		}
		if yu == IV4 {
			yu = yu % IV4
			zheng = num / IV4
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "IV"
			}
			fmt.Println("IV: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		}
		if yu >= I1 {
			yu = yu % I1
			zheng = num / I1
			num = yu
			for i := 1; i <= zheng; i++ {
				ret += "I"
			}
			fmt.Println("I: yu:", yu, "zheng", zheng, "ret:", ret)
			continue
		} else {
			fmt.Println("end:", ret)
			break
		}
	}
	return ret
}
