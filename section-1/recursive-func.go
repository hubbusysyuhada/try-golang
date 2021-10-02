// recursive function paling mudah/umum digunakan untuk menyelesaikan faktorial
package main

import "fmt"

func main() {
	loop := loopFactorial(5)
	nonLoop := recursive(5)
	fmt.Println(loop, "<<<<")
	fmt.Println(nonLoop, "<<<<")
}

// factorial with loop
func loopFactorial(num int) int {
	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result
}

// factorial with recursive function
func recursive(num int) int {
	fmt.Println(num, "<<<< num from recursive")
	if num == 1 {
		return 1
	} else {
		return num * recursive(num-1)
	}
}
