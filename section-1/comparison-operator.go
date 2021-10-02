package main

import "fmt"

func main() {
	// same as javascript (< <= > >= ! ==) but the is equal just use 2 "= "instead of 3
	compare1 := "haha" == "hehe"     // false
	compare2 := 2 > 0                // true
	not_compare1 := !compare1        // true
	not_compare2 := !compare2        // false
	compare3 := compare1 && compare2 // false
	compare4 := compare1 || compare2 // true
	fmt.Println(compare1)
	fmt.Println(compare2)
	fmt.Println(compare3)
	fmt.Println(compare4)
	fmt.Println(not_compare1)
	fmt.Println(not_compare2)
}
