package study_type

import "fmt"

func Study() {
	var num int32 = 22
	fmt.Printf("num: %v\n", num)
	var num2 int64 = int64(num) // 使用type()函数进行转换
	fmt.Printf("num2: %v\n", num2)
}
