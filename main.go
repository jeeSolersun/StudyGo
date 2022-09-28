package main

import (
	"fmt"
	"studygo/study_type"
)

func main() {
	//Hello()
	//StudyIf()
	// TestSwitch()
	//StudyFor()
	// study_func.StudyDefer()
	// study_scanner.Study()
	// study_pointer.StudyArr()
	study_type.Study()

	// study_pointer.StudyPointer()
	// study_pointer.StudySlicePointer()
	// study_pointer.StudyStructPointer()
}

func StudyFor() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				goto END
			}
		}
	}
END:
	fmt.Printf("\"END\": %v\n", "goto END...")
}

func TestSwitch() {
	var num int32 = 22
	switch num {
	case 11:
		fmt.Printf("num: %v\n", num)
	case 22, 23:
		fmt.Printf("22, 23: %v\n", "num是22或者23")
	}
}

func Hello() {
	/* fmt.Println("hello go")
	count := 3
	for i := 0; i < count; i++ {
		fmt.Printf("i = %d\n", i)
	} */
}

func StudyIf() {
	// 分支结构
	var a = 22
	if a < 22 {
		fmt.Printf("a的值小于 22\n")
	} else if a == 22 {
		fmt.Printf("a的值等于 22\n")
	} else {
		fmt.Printf("a的值大于 22\n")
	}
}
