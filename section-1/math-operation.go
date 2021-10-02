package main

import "fmt"

func main() {
	// same as javascript (+ -  * / %)
	modulo := 6 % 4
	fmt.Println(modulo)
	modulo += 1
	fmt.Println(modulo)
	modulo++ // or modulo--
	fmt.Println(modulo)
}
