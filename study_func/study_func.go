package study_func

import "fmt"

func StudyDefer() {
	fmt.Printf("\"start\": %v\n", "start")       // 1
	defer fmt.Printf("\"step1\": %v\n", "step1") // 5
	defer fmt.Printf("\"step2\": %v\n", "step2") // 4
	defer fmt.Printf("\"step3\": %v\n", "step3") // 3
	fmt.Printf("\"end\": %v\n", "end")           // 2
}
