package main

import "fmt"

func main() {
	// multiple ways to declare a variable
	var A string
	A = "A"
	var B string = "B"
	var C = "C"
	D := "D"
	var (
		E        = "E"
		F string = "F"
	)

	fmt.Println(A, B, C, D, E, F)
}
