package main

import (
	"fmt"
)

func main() {
	// 打印表头
	fmt.Printf("%-8s %-8s %-8s %-12s %-8s\n", "十进制", "十六进制", "ASCII", "Unicode编码点", "Unicode字符")
	fmt.Println("----------------------------------------------------------------")

	// 输出标准ASCII字符表 (0-127)
	for i := 0; i <= 127; i++ {
		// 控制字符通常不可打印，用点号表示
		var asciiChar string
		if i >= 32 && i <= 126 { // 可打印字符范围
			asciiChar = string(rune(i))
		} else {
			asciiChar = "."
		}

		// 获取Unicode编码点和字符
		r := rune(i)
		unicodePoint := fmt.Sprintf("U+%04X", i)
		
		// 构建Unicode字符的字符串表示
		var unicodeChar string
		if i >= 32 && i <= 126 {
			unicodeChar = string(r)
		} else {
			unicodeChar = "."
		}

		// 打印行
		fmt.Printf("%-8d 0x%-6X %-8s %-12s %-8s\n", i, i, asciiChar, unicodePoint, unicodeChar)

		// 每16个字符后添加分隔线
		if (i+1)%16 == 0 {
			fmt.Println("----------------------------------------------------------------")
		}
	}
}
