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

func Study02() {
	fmt.Printf("\"modify 01 on master\": %v\n", "modify 01 on master")

	fmt.Printf("\"modify 02 on master\": %v\n", "modify 02 on master")
}
