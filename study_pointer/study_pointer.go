package study_pointer

import "fmt"

func StudyArr() {
	//arr := make([]int, 2, 3)
	// arr := []int{1, 2, 3}
	// fmt.Printf("arr = %v\n", arr)
	var arr [3]int = [3]int{1, 2, 3}
	var slice []int = arr[0:2]
	fmt.Println(slice) // [1, 2]
}

func StudyArrPointer() {
	var arr [3]int = [3]int{1, 2, 3}
	var ptr *int = &arr[0]
	fmt.Printf("ptr: %v\n", ptr)
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("&arr: %v\n", &arr)
	fmt.Printf("&arr[0]: %p\n", &arr[0])
}

// 切片指针
func StudySlicePointer() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	slice := arr[0:3]
	fmt.Printf("slice: %v\n", slice)
	// "%p"可以打印指针中的内容即变量地址
	fmt.Printf("arr: %p\n", &arr) // 0xc000014240
	// 切片本质上是一个指向数组的指针
	fmt.Printf("slice: %p\n", slice) // 0xc000014240
	// 指向切片的指针：2级指针
	var slice_ptr *[]int = &slice
	fmt.Printf("slice_ptr: %p\n", slice_ptr)  // 0xc00000c030
	fmt.Printf("slice_ptr: %p\n", *slice_ptr) // 0xc000014240
	(*slice_ptr)[1] = 22
	fmt.Printf("arr: %v\n", arr)     // arr: [1 22 3 4 5]
	fmt.Printf("slice: %v\n", slice) // slice: [1 22 3]
}

type Student struct {
	name string
	age  int
}

// 结构体指针
func StudyStructPointer() {
	// 方式1
	var student Student = Student{name: "sorrycc", age: 22}
	var stu_ptr *Student = &student
	fmt.Printf("student: %v\n", student) // student: {sorrycc 22}
	// (*stu_ptr).name = "codeify"
	stu_ptr.name = "codeify123"          // 使用语法糖，结果和上面一样
	fmt.Printf("student: %v\n", student) // student: {codeify123 22}
	// 方式2
	stu_ptr2 := new(Student) // new会返回底层创建的结构体变量的地址
	// 使用语法糖
	stu_ptr2.name = "关羽"
	stu_ptr2.age = 18
	fmt.Printf("*stu_ptr2: %v\n", *stu_ptr2) // *stu_ptr2: {关羽 18}
}
