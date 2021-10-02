// empty interface same as "any" in typescript

package main

import "fmt"

func main() {
	fmt.Println(eh(1))
	fmt.Println(eh("duh"))
}

func eh(a interface{}) interface{} {
	// this func will receive and return anything
	if a == 1 {
		return fmt.Sprint("number ", a)
	} else {
		return fmt.Sprint("string ", a)
	}
}
