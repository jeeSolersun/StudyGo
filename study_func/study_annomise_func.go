package study_func

import "fmt"

func StudyAnnomiseFunc() {
	f1 := func(num1 int32, num2 int32) int32 {
		return num1 + num2
	}
	i := f1(1, 21)
	fmt.Printf("i: %v\n", i)
}
