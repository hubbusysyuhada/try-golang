package main

import "fmt"

func main() {
	variadic(1, 2, 3, 4, 5, 6, 7)
	customNum := multipleAndSumAll(2, 2, 3, 4)
	fmt.Println(customNum, "<<< custom num")

	// inserting a slice into varags in variadic function
	slice := []int{1, 2, 3, 4}
	multipleAndSumAll(2, slice...)
}

func variadic(num ...int) {
	// ciri dari variadic function adalah parameternya memiliki "..."
	// disebut dengan varargs (variable arguments) hanya dapat terletak di argumen terakhir
	// varargs akan diconvert menjadi tipe array/slice
	for _, v := range num {
		fmt.Println(v, "<<< v")
	}
}

func multipleAndSumAll(num1 int, slice ...int) int {
	result := 0
	for _, value := range slice {
		result += num1 * value
	}
	return result
}
