package study_scanner

import "fmt"

func Study() {
	var (
		num1 int32
		num2 int32
	)
	// fmt.Scanf("%d%d", &num1, &num2)
	// fmt.Printf("num1: %v\n", num1)
	// fmt.Printf("num2: %v\n", num2)
	fmt.Scanln(&num1, &num2)
	fmt.Printf("num1: %v\n", num1)
	fmt.Printf("num2: %v\n", num2)
}
