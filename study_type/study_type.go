package study_type

import "fmt"

func Study() {
	var num int32 = 22
	fmt.Printf("num: %v\n", num)
	var num2 int64 = int64(num) // 使用type()函数进行转换
	fmt.Printf("num2: %v\n", num2)
}

func Study02() {
	fmt.Printf("\"modify 01 on dev\": %v\n", "modify 01 on dev")

	fmt.Printf("\"modify 02 on dev\": %v\n", "modify 02 on dev")

}
