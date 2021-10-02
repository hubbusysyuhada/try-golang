package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 5 || i == 7 {
			continue
		}
		if i == 9 {
			break
		}
		fmt.Println(i, "<<<< i")
	}
}
